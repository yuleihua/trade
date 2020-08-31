package service

import (
	log "github.com/sirupsen/logrus"

	"github.com/yuleihua/trade/transfer/model"
	"github.com/yuleihua/trade/transfer/types"
)

func (t *Transfer) Login(req *types.RequestLogin) (int64, error) {
	customer, err := model.GetCustomerByName(types.DefaultNation, req.Name)
	if err != nil {
		log.Errorf("get customer error, name:%s, %v", req.Name, err)
		return 0, err
	}

	// todo: check password
	return customer.Id, nil
}
