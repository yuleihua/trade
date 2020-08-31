package service

import (
	"context"
	"errors"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yuleihua/trade/transfer/biz"
	"github.com/yuleihua/trade/transfer/model"
	"github.com/yuleihua/trade/transfer/types"
)

// api handle
func (t *Transfer) AddTransfer(ctx context.Context, uid string, req *types.RequestTransfer) (*types.ResponseTradeReceipt, error) {
	isPass, err := biz.VerifyTransferRequest(req)
	if err != nil {
		log.Errorf("verify request error, %#v, %v", req, err)
		return nil, err
	}

	if !isPass {
		return nil, errors.New("failed to check parameters")
	}

	// Get customer information
	cid, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	customer, err := model.GetCustomerById(cid)
	if err != nil {
		log.Errorf("add trade-sequence error, %#v, %v", req, err)
		return nil, err
	}

	// Add trade sequence
	if err := biz.AddTradeSequence(req, customer); err != nil {
		log.Errorf("add trade-sequence error, %#v, %v", req, err)
		return nil, err
	}

	// Handle fund balance
	trade, err := biz.HandleTransferFund(req, customer)
	if err != nil {
		log.Errorf("handle trade-fund error, %#v, %v", req, err)
		return nil, err
	}

	receipt := biz.ConvertTradeReceipt(trade)
	return receipt, nil
}

// api handle
func (t *Transfer) ConfirmTransfer(ctx context.Context, uid string, req *types.RequestReceipt) (*types.ResponseTradeReceipt, error) {
	isPass, err := biz.VerifyReceiptRequest(req)
	if err != nil {
		log.Errorf("verify request error, %#v, %v", req, err)
		return nil, err
	}

	if !isPass {
		return nil, errors.New("failed to check parameters")
	}

	// Get customer information
	cid, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	customer, err := model.GetCustomerById(cid)
	if err != nil {
		log.Errorf("add trade-sequence error, %#v, %v", req, err)
		return nil, err
	}

	// Handle fund balance
	trade, err := biz.HandleConfirmFund(req, customer)
	if err != nil {
		log.Errorf("handle trade-fund error, %#v, %v", req, err)
		return nil, err
	}

	receipt := biz.ConvertTradeReceipt(trade)
	return receipt, nil
}

// api handle
func (t *Transfer) GetTransfer(ctx context.Context, uid int64, account, name, uuid string) ([]*types.ResponseTradeDetail, error) {
	customer, err := model.GetCustomerById(uid)
	if err != nil {
		return nil, err
	}

	var trades []*model.Trade
	if name != "" {
		toUser, err := model.GetCustomerByName(types.DefaultNation, name)
		if err != nil {
			return nil, err
		}

		trades, err = model.GetTrade("from_cid = ? and to_cid = ?", customer.Id, toUser.Id)
	} else if account != "" {
		toBranch, err := model.GetBranchByAccount(account)
		if err != nil {
			return nil, err
		}

		trades, err = model.GetTrade("from_cid = ? and to_bid = ?", customer.Id, toBranch.Id)
	} else if uuid != "" {
		trades, err = model.GetTrade("uuid = ?", uuid)
	} else {
		trades, err = model.GetTrade("from_cid = ?", customer.Id)
	}

	if err != nil {
		return nil, err
	}

	if len(trades) == 0 {
		return nil, nil
	}

	results := make([]*types.ResponseTradeDetail, 0, len(trades))
	for _, trade := range trades {
		fees, err := model.GetTradeFee("uuid = ?", trade.Uuid)
		if err != nil {
			return nil, err
		}

		obj := biz.ConvertTradeDetail(trade, fees)
		results = append(results, obj)
	}

	return results, nil
}
