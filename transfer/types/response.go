package types

type ResponseTradeFee struct {
	TransDate int    `json:"trans_date"`
	TransTime int    `json:"trans_time"`
	Uuid      string `json:"uuid"`
	MoneyType string `json:"money_type"`
	Amt       int64  `json:"amt"`
	Fee       int64  `json:"fee"`
	Remark    string `json:"remark"`
}

type ResponseTradeDetail struct {
	TransDate   int                 `json:"trans_date"`
	TransTime   int                 `json:"trans_time"`
	Uuid        string              `json:"uuid"`
	FromUuid    string              `json:"from_uuid"`
	ToUuid      string              `json:"to_uuid"`
	FromBid     int64               `json:"from_bid"`
	FromCid     int64               `json:"from_cid"`
	ToBid       int64               `json:"to_bid"`
	ToCid       int64               `json:"to_cid"`
	IsDelay     bool                `json:"is_delay"`
	IsLarge     bool                `json:"is_large"`
	IsReject    bool                `json:"is_reject"`
	Amt         int64               `json:"amt"`
	Fee         int64               `json:"fee"`
	Remark      string              `json:"remark"`
	MoneyType   string              `json:"money_type"`
	Errcode     int                 `json:"errcode"`
	ConfirmDate int                 `json:"confirm_date"`
	ConfirmTime int                 `json:"confirm_time"`
	ConfirmAmt  int64               `json:"confirm_amt"`
	ConfirmOpid int64               `json:"confirm_opid"`
	FeeDetails  []*ResponseTradeFee ` json:"fees"`
}

type ResponseTradeDetails []ResponseTradeDetail

type ResponseTradeLite struct {
	TransDate   int    `json:"trans_date"`
	TransTime   int    `json:"trans_time"`
	Uuid        string `json:"uuid"`
	FromUuid    string `json:"from_uuid"`
	ToUuid      string `json:"to_uuid"`
	FromBid     int64  `json:"from_bid"`
	FromCid     int64  `json:"from_cid"`
	ToBid       int64  `json:"to_bid"`
	ToCid       int64  `json:"to_cid"`
	IsDelay     bool   `json:"is_delay"`
	IsLarge     bool   `json:"is_large"`
	IsReject    bool   `json:"is_reject"`
	Amt         int64  `json:"amt"`
	Fee         int64  `json:"fee"`
	Remark      string `json:"remark"`
	MoneyType   string `json:"money_type"`
	Errcode     int    `json:"errcode"`
	ConfirmDate int    `json:"confirm_date"`
	ConfirmTime int    `json:"confirm_time"`
	ConfirmAmt  int64  `json:"confirm_amt"`
	ConfirmOpid int64  `json:"confirm_opid"`
}

type ResponseTradeReceipt struct {
	TransDate   int    `json:"trans_date"`
	TransTime   int    `json:"trans_time"`
	Uuid        string `json:"uuid"`
	FromUuid    string `json:"from_uuid"`
	ToUuid      string `json:"to_uuid"`
	FromBid     int64  `json:"from_bid"`
	FromCid     int64  `json:"from_cid"`
	ToBid       int64  `json:"to_bid"`
	ToCid       int64  `json:"to_cid"`
	IsDelay     bool   `json:"is_delay"`
	IsLarge     bool   `json:"is_large"`
	IsReject    bool   `json:"is_reject"`
	Amt         int64  `json:"amt"`
	Fee         int64  `json:"fee"`
	Remark      string `json:"remark"`
	MoneyType   string `json:"money_type"`
	Errcode     int    `json:"errcode"`
	ConfirmDate int    `json:"confirm_date"`
	ConfirmTime int    `json:"confirm_time"`
	ConfirmAmt  int64  `json:"confirm_amt"`
	ConfirmOpid int64  `json:"confirm_opid"`
}
