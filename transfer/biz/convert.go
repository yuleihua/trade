package biz

import (
	"github.com/yuleihua/trade/transfer/model"
	"github.com/yuleihua/trade/transfer/types"
)

func ConvertTradeDetail(trade *model.Trade, fees []*model.TradeFee) *types.ResponseTradeDetail {
	if trade == nil {
		return nil
	}

	res := &types.ResponseTradeDetail{
		TransDate:   trade.TransDate,
		TransTime:   trade.TransTime,
		Uuid:        trade.Uuid,
		FromUuid:    trade.FromUuid,
		ToUuid:      trade.ToUuid,
		FromBid:     trade.FromBid,
		FromCid:     trade.FromCid,
		ToBid:       trade.ToBid,
		ToCid:       trade.ToCid,
		IsDelay:     trade.IsDelay,
		IsLarge:     trade.IsLarge,
		IsReject:    trade.IsReject,
		Amt:         trade.Amt,
		Fee:         trade.Fee,
		Remark:      trade.Remark,
		MoneyType:   trade.MoneyType,
		Errcode:     trade.Errcode,
		ConfirmDate: trade.ConfirmDate,
		ConfirmTime: trade.ConfirmTime,
		ConfirmAmt:  trade.ConfirmAmt,
		ConfirmOpid: trade.ConfirmOpid,
		FeeDetails:  make([]*types.ResponseTradeFee, 0, len(fees)),
	}

	for _, fee := range fees {
		obj := &types.ResponseTradeFee{
			Uuid:      fee.Uuid,
			MoneyType: fee.MoneyType,
			Amt:       fee.Amt,
			Fee:       fee.Fee,
			Remark:    fee.Remark,
		}
		res.FeeDetails = append(res.FeeDetails, obj)
	}

	return res
}

func ConvertTradeReceipt(trade *model.Trade) *types.ResponseTradeReceipt {
	if trade == nil {
		return nil
	}

	res := &types.ResponseTradeReceipt{
		TransDate:   trade.TransDate,
		TransTime:   trade.TransTime,
		Uuid:        trade.Uuid,
		FromUuid:    trade.FromUuid,
		ToUuid:      trade.ToUuid,
		FromBid:     trade.FromBid,
		FromCid:     trade.FromCid,
		ToBid:       trade.ToBid,
		ToCid:       trade.ToCid,
		IsDelay:     trade.IsDelay,
		IsLarge:     trade.IsLarge,
		IsReject:    trade.IsReject,
		Amt:         trade.Amt,
		Fee:         trade.Fee,
		Remark:      trade.Remark,
		MoneyType:   trade.MoneyType,
		Errcode:     trade.Errcode,
		ConfirmDate: trade.ConfirmDate,
		ConfirmTime: trade.ConfirmTime,
		ConfirmAmt:  trade.ConfirmAmt,
		ConfirmOpid: trade.ConfirmOpid,
	}

	return res
}
