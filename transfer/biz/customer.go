package biz

import (
	"github.com/yuleihua/trade/transfer/model"
	"github.com/yuleihua/trade/transfer/types"
)

type CustomerStatus int

const (
	StatusInit CustomerStatus = 0
	StatusVerify
	StatusNormal
	StatusRisk
	StatusCancel
	StatusNoExisted
)

// CheckCustomer check customer's status by name, phone, email
func CheckCustomer(nation, name, phone, email string) (CustomerStatus, error) {
	status := StatusInit

	nationZone := nation
	if nation == "" {
		nationZone = types.DefaultNation
	}

	var err error
	var client *model.Customer

	if name != "" {
		client, err = model.GetCustomerByName(nationZone, name)
	} else if phone != "" {
		client, err = model.GetCustomerByPhone(nationZone, phone)
	} else if email != "" {
		client, err = model.GetCustomerByEmail(nationZone, email)
	}

	if err != nil {
		return status, err
	}

	if client != nil && client.Id != 0 {
		customers, err := model.GetBranchByCustomerId(client.Id)
		if err != nil {
			return status, err
		}

		if len(customers) == 0 {
			return StatusNoExisted, nil
		}

		status = CustomerStatus(customers[0].Status)
	}

	return status, nil
}
