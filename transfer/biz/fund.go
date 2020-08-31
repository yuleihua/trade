package biz

import (
	"errors"

	"github.com/yuleihua/trade/transfer/errcode"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/yuleihua/trade/pkg/dbclient"
	"github.com/yuleihua/trade/transfer/dummy"
	"github.com/yuleihua/trade/transfer/model"
	"github.com/yuleihua/trade/transfer/types"
)

func HandleTransferFund(req *types.RequestTransfer, c *model.Customer) (*model.Trade, error) {
	var useBranch *model.Branch
	var useFund *model.Fund

	tradeSequence, err := model.GetTradeSeqByUUID(req.UUID)
	if err != nil {
		log.Errorf("get trade-sequence error, uuid: %s, %v", req.UUID, err)
		return nil, err
	}

	if req.ToAccount != "" {
		toBranch, err := model.GetBranch("account = ?", req.ToAccount)
		if err != nil {
			log.Errorf("get branch error, account: %s, %v", req.ToAccount, err)
			return nil, err
		}
		useBranch = toBranch[0]

		if len(toBranch) == 0 {
			log.Errorf("no branch, %s, %v", req.ToAccount, err)
			return nil, errors.New("branch account is not existed")
		}

		toFund, err := model.GetFundByBranchId(toBranch[0].Id)
		if err != nil {
			log.Errorf("get fund error, %d, %v", toBranch[0].Id, err)
			return nil, err
		}
		useFund = toFund
	} else if req.ToName != "" {
		c, err := model.GetCustomerByName("CN", req.ToName)
		if err != nil {
			log.Errorf("get customer error, account: %s, %v", req.ToAccount, err)
			return nil, err
		}

		toBranch, err := model.GetBranchByCustomerId(c.Id)
		if err != nil {
			log.Errorf("get branch error, account: %s, %v", req.ToAccount, err)
			return nil, err
		}
		useBranch = toBranch[0]

		if len(toBranch) == 0 {
			log.Errorf("no branch, account: %s, %v", req.ToAccount, err)
			return nil, errors.New("branch account is not existed")
		}

		toFund, err := model.GetFundByBranchId(toBranch[0].Id)
		if err != nil {
			log.Errorf("get fund error, account: %s, %v", req.ToAccount, err)
			return nil, err
		}
		useFund = toFund
	}

	fromBranch, err := model.GetBranchByCustomerId(c.Id)
	if err != nil {
		log.Errorf("get from branch error, account: %s, %v", req.ToAccount, err)
		return nil, err
	}
	//todo: get the account. using default branch ID(Suppose a customer has a fund account)

	fromFund, err := model.GetFundByBranchId(fromBranch[0].Id)
	if err != nil {
		log.Errorf("get from fund error, account: %s, %v", req.ToAccount, err)
		return nil, err
	}

	// begin fund op: downstream
	if fromFund.Balance-fromFund.FreezeBalance-req.MoneyAmt <= 0 {
		return nil, errors.New("no enough money balance")
	}

	tx := dbclient.NewDBTransaction(dbclient.DBDefault())
	if err := tx.Begin(); err != nil {
		log.Errorf("new transaction error, account: %s, %v", req.ToAccount, err)
		return nil, err
	}

	transDate, transTime := GetBizDatetime()
	tradeUUID := uuid.New().String()
	transCode := 0

	//step1: calc fee
	feeBalance, err := dummy.CalcFee(req.MoneyAmt)
	if err != nil {
		return nil, err
	}

	feeObj := model.TradeFee{
		TransDate: transDate,
		TransTime: transTime,
		Uuid:      tradeUUID,
		MoneyType: req.MoneyType,
		Amt:       req.MoneyAmt,
		Fee:       feeBalance,
		Remark:    req.Comment,
	}

	if err := model.AddTradeFeeTX(tx, feeObj); err != nil {
		tx.Rollback()
		log.Errorf("add trade-fee error, fee: %v, %v", feeObj, err)
		return nil, err
	}

	// step2: check risk?
	isRisk := false
	if req.MoneyAmt > dummy.GetMoneyLimit(0) {
		riskObj := model.TradeRisk{
			TransDate: transDate,
			TransTime: transTime,
			Level:     5,
			Uuid:      tradeUUID,
			Amt:       req.MoneyAmt,
			Fee:       feeBalance,
			Remark:    "",
		}

		if err := model.AddTradeRiskTX(tx, riskObj); err != nil {
			tx.Rollback()
			log.Errorf("add trade risk error, risk: %v, %v", riskObj, err)
			return nil, err
		}
		isRisk = true
		transCode = errcode.RiskTradeCode
	}

	// step 3: sub balance
	if isRisk {
		if err := model.UpdateFundTXDownFreezeBalanceAll(tx, fromFund.Id, req.MoneyAmt); err != nil {
			tx.Rollback()
			log.Errorf("update fund error, fromFund: %#v, %v", fromFund, err)
			return nil, err
		}
	} else {
		if err := model.UpdateFundTXDownBalance(tx, fromFund.Id, req.MoneyAmt); err != nil {
			tx.Rollback()
			log.Errorf("update fund error, fund: %#v, %v", fromFund, err)
			return nil, err
		}
	}

	// step4: add balance
	if !isRisk {
		upAmt := req.MoneyAmt - feeBalance
		if err := model.UpdateFundTXUpBalance(tx, useFund.Id, upAmt); err != nil {
			tx.Rollback()
			log.Errorf("update fund error, fund: %s, %v", useFund, err)
			return nil, err
		}
		transCode = errcode.SuccessCode
	}

	// step5: insert tables: trade
	tradeObj := model.Trade{
		TransDate:   transDate,
		TransTime:   transTime,
		Uuid:        tradeUUID,
		FromUuid:    req.UUID,
		ToUuid:      tradeUUID,
		FromBid:     fromBranch[0].Id,
		FromCid:     c.Id,
		ToBid:       useBranch.Id,
		ToCid:       useBranch.Cid,
		IsDelay:     req.IsRealtime,
		IsLarge:     false,
		IsReject:    false,
		Amt:         req.MoneyAmt,
		Fee:         feeBalance,
		Remark:      req.Comment,
		MoneyType:   req.MoneyType,
		Errcode:     transCode,
		ConfirmDate: 0,
		ConfirmTime: 0,
		ConfirmAmt:  0,
		ConfirmOpid: 0,
	}

	if err := model.AddTradeTX(tx, tradeObj); err != nil {
		tx.Rollback()
		log.Errorf("add trade error, req: %v, %v", req, err)
		return nil, err
	}

	// setp6: update sequence
	sequence := *tradeSequence
	sequence.TransDate = tradeObj.TransDate
	sequence.TransTime = tradeObj.TransTime
	sequence.ToBid = tradeObj.ToBid
	sequence.Errcode = tradeObj.Errcode
	sequence.Uuid = tradeObj.Uuid
	sequence.ToUuid = tradeObj.ToUuid
	if err := model.UpdateTradeSeqTX(tx, sequence); err != nil {
		tx.Rollback()
		log.Errorf("add trade-sequence record error,  %v, %v", sequence, err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("commit error, req: %v, %v", req, err)
		return nil, err
	}

	return &tradeObj, nil
}

func HandleConfirmFund(req *types.RequestReceipt, c *model.Customer) (*model.Trade, error) {
	tradeSequence, err := model.GetTradeSeqByUUID(req.UUID)
	if err != nil {
		log.Errorf("get trade-sequence error, uuid: %s, %v", req.UUID, err)
		return nil, err
	}

	if tradeSequence.Errcode != errcode.InitSequenceCode && tradeSequence.Errcode != errcode.RiskTradeCode {
		log.Errorf("trade-sequence has been modified, uuid: %s, %v", req.UUID, err)
		return nil, errors.New("trade-sequence has been modified")
	}

	trade, err := model.GetTradeByUUId(req.UUID)
	if err != nil {
		log.Errorf("get trade information error, uuid: %s, %v", req.UUID, err)
		return nil, err
	}

	if trade.MoneyType != req.MoneyType || trade.Amt != req.MoneyAmt {
		log.Errorf("check money information error, req: %v, %v", req, err)
		return nil, errors.New("invalid money information")
	}

	risk, err := model.GetTradeRiskByUUID(req.UUID)
	if err != nil {
		log.Errorf("get trade information error, uuid: %s, %v", req.UUID, err)
		return nil, err
	}

	if req.RiskLevel <= 0 || risk.Level <= req.RiskLevel {
		log.Errorf("invalid risk level error, req: %v, %v", req, err)
		return nil, errors.New("invalid risk level")
	}
	//todo: check op_name information.(Assume that each risk transaction confirmation requires an auditor to confirm it)
	// dummy op_id is 5555

	fromFund, err := model.GetFundByBranchId(trade.FromBid)
	if err != nil {
		log.Errorf("get from-fund error, bid: %d, %v", trade.FromBid, err)
		return nil, err
	}

	useFund, err := model.GetFundByBranchId(trade.ToBid)
	if err != nil {
		log.Errorf("get to-fund error, bid: %d, %v", trade.ToBid, err)
		return nil, err
	}

	tx := dbclient.NewDBTransaction(dbclient.DBDefault())
	if err := tx.Begin(); err != nil {
		log.Errorf("new transaction error, account: uid: %d, %v", c.Id, err)
		return nil, err
	}

	transDate, transTime := GetBizDatetime()
	// step 1: sub balance
	if err := model.UpdateFundTXDownFreezeBalance(tx, fromFund.Id, req.MoneyAmt); err != nil {
		tx.Rollback()
		log.Errorf("update fund error, fund: %#v, %v", fromFund, err)
		return nil, err
	}

	// step3: update risk
	riskObj := *risk
	risk.Level = req.RiskLevel
	risk.Remark = req.Comment

	if err := model.UpdateTradeRiskTX(tx, riskObj); err != nil {
		tx.Rollback()
		log.Errorf("get branch error, obj: %v, %v", riskObj, err)
		return nil, err
	}

	// step4: add balance
	upAmt := req.MoneyAmt - trade.Fee
	if err := model.UpdateFundTXUpBalance(tx, useFund.Id, upAmt); err != nil {
		tx.Rollback()
		log.Errorf("update up-fund error, req: %v, %v", req, err)
		return nil, err
	}

	// step5: insert tables: trade
	tradeObj := *trade
	tradeObj.ConfirmDate = transDate
	tradeObj.ConfirmTime = transTime
	tradeObj.ConfirmAmt = req.MoneyAmt
	tradeObj.ConfirmOpid = 5555 //dummy op id
	tradeObj.ConfirmMoneyType = req.MoneyType
	tradeObj.Errcode = errcode.SuccessCode

	if err := model.UpdateTradeTX(tx, tradeObj); err != nil {
		tx.Rollback()
		log.Errorf("update trade error, tradeObj: %v, %v", tradeObj, err)
		return nil, err
	}

	// setp6: update sequence
	sequence := *tradeSequence
	if sequence.ToBid == 0 {
		sequence.ToBid = tradeObj.ToBid
	}
	if sequence.Uuid == "" {
		sequence.Uuid = tradeObj.Uuid
	}
	if sequence.ToUuid == "" {
		sequence.ToUuid = tradeObj.ToUuid
	}
	sequence.Errcode = tradeObj.Errcode

	if err := model.UpdateTradeSeqTX(tx, sequence); err != nil {
		tx.Rollback()
		log.Errorf("update trade-sequence record error,  %v, %v", sequence, err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("commit error, req: %v, %v", req, err)
		return nil, err
	}

	return &tradeObj, nil
}
