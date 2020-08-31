package biz

import (
	log "github.com/sirupsen/logrus"
	"github.com/yuleihua/trade/pkg/dbclient"

	"github.com/yuleihua/trade/transfer/errcode"
	"github.com/yuleihua/trade/transfer/model"
	"github.com/yuleihua/trade/transfer/types"
)

func AddTradeSequence(req *types.RequestTransfer, c *model.Customer) error {
	branch, err := model.GetBranchByCustomerId(c.Id)
	if err != nil {
		log.Errorf("get branch error, %d, %v", c.Id, err)
		return err
	}

	nowDate, nowTime := GetBizDatetime()
	obj := model.TradeSeq{
		TransDate: nowDate,
		TransTime: nowTime,
		FromUuid:  req.UUID,
		FromBid:   branch[0].Id,
		FromCid:   c.Id,
		Errcode:   errcode.InitSequenceCode,
	}

	if err := obj.Create(); err != nil {
		log.Errorf("add trade-sequence record error,  %#v, %v", obj, err)
		return err
	}

	return nil
}

func UpdateTradeSequence(tx *dbclient.DBTransaction, t *model.Trade) error {
	sequence, err := model.GetTradeSeqByUUID(t.FromUuid)
	if err != nil {
		log.Errorf("get trade-sequence error, uuid: %s, %v", t.FromUuid, err)
		return err
	}

	sequence.TransDate = t.TransDate
	sequence.TransTime = t.TransTime
	sequence.ToBid = t.ToBid
	sequence.Errcode = t.Errcode
	sequence.Uuid = t.Uuid
	sequence.ToUuid = t.ToUuid

	if err := model.UpdateTradeSeqTX(tx, *sequence); err != nil {
		log.Errorf("add trade-sequence record error,  %v, %v", sequence, err)
		return err
	}

	return nil
}
