package biz

import "github.com/yuleihua/trade/transfer/types"

func VerifyTransferRequest(req *types.RequestTransfer) (bool, error) {
	return true, nil
}

func VerifyReceiptRequest(req *types.RequestReceipt) (bool, error) {
	return true, nil
}
