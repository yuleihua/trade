package types

type RequestTransfer struct {
	ReqID            string `form:"req_id" json:"req_id" `
	UUID             string `form:"uuid" json:"uuid" `
	ToAccount        string `form:"to_account" json:"to_account" `
	ToName           string `form:"to_name" json:"to_name" `
	FromAccount      string `form:"from_account" json:"from_account" `
	MoneyType        string `form:"money_type" json:"money_type" `
	MoneyAmt         int64  `form:"money_amt" json:"money_amt" `
	Password         string `form:"password" json:"password" `
	Comment          string `form:"comment" json:"comment" `
	IsRealtime       bool   `form:"is_realtime" json:"is_realtime" `
	OpDatetime       string `form:"op_datetime" json:"op_datetime" `
	OpTimezone       string `form:"op_timezone" json:"op_timezone" `
	NotificationType string `form:"notification_type" json:"notification_type" `
	Postscript       string `form:"postscript" json:"postscript" `
}

type RequestLogin struct {
	ReqID    string `form:"req_id" json:"req_id" `
	Name     string `form:"name" json:"name" validator:"required"`
	Password string `form:"password" json:"password" `
}

type RequestReceipt struct {
	ReqID     string `form:"req_id" json:"req_id" `
	UUID      string `form:"uuid" json:"uuid" `
	MoneyType string `form:"money_type" json:"money_type" `
	MoneyAmt  int64  `form:"money_amt" json:"money_amt" `
	OpName    string `form:"op_name" json:"op_name" `
	RiskLevel int    `form:"risk_level" json:"risk_level" `
	Comment   string `form:"comment" json:"comment" `
}
