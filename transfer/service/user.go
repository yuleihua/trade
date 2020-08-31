package service

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yuleihua/trade/transfer/biz"
	"github.com/yuleihua/trade/transfer/model"
	"github.com/yuleihua/trade/transfer/types"
)

func (t *Transfer) Login(req *types.RequestLogin) (string, error) {
	customer, err := model.GetCustomerByName(types.DefaultNation, req.Name)
	if err != nil {
		log.Errorf("get customer error, name:%s, %v", req.Name, err)
		return "", err
	}

	// todo: check password

	uid := strconv.FormatInt(customer.Id, 10)
	token, err := biz.MakeJWTToken(uid, req.Name)
	if err != nil {
		log.Errorf("make token error, uid:%s, name:%s, %v", uid, req.Name, err)
		return "", err
	}

	if err := t.cache.Set(req.Name, token, types.JWTTokenTTL); err != nil {
		log.Errorf("cache token error, uid:%s, name:%s, %v", uid, req.Name, err)
		return "", err
	}

	return token, nil
}
