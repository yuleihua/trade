package model

import client "github.com/yuleihua/aaa/dbclient"

var tradeName = "trade"

type Trade struct {
	Id               int64  `gorm:"column:id;type:bigint(20)" json:"id"`
	TransDate        int    `gorm:"column:trans_date;type:int(11);default:'0'" json:"trans_date"`
	TransTime        int    `gorm:"column:trans_time;type:int(11);default:'0'" json:"trans_time"`
	Uuid             string `gorm:"column:uuid;type:varchar(64);default:''" json:"uuid"`           // 流水号
	FromUuid         string `gorm:"column:from_uuid;type:varchar(64);default:''" json:"from_uuid"` // 流水号
	ToUuid           string `gorm:"column:to_uuid;type:varchar(64);default:''" json:"to_uuid"`     // 流水号
	FromBid          int64  `gorm:"column:from_bid;type:bigint(20);default:'0'" json:"from_bid"`
	FromCid          int64  `gorm:"column:from_cid;type:bigint(20);default:'0'" json:"from_cid"`
	ToBid            int64  `gorm:"column:to_bid;type:bigint(20);default:'0'" json:"to_bid"`
	ToCid            int64  `gorm:"column:to_cid;type:bigint(20);default:'0'" json:"to_cid"`
	IsDelay          int8   `gorm:"column:is_delay;type:tinyint(1);default:'0'" json:"is_delay"`
	IsLarge          int8   `gorm:"column:is_large;type:tinyint(1);default:'0'" json:"is_large"`
	IsReject         int8   `gorm:"column:is_reject;type:tinyint(1);default:'0'" json:"is_reject"`
	Amt              int64  `gorm:"column:amt;type:bigint(20);default:'0'" json:"amt"`
	Fee              int64  `gorm:"column:fee;type:bigint(20);default:'0'" json:"fee"`
	Remark           string `gorm:"column:remark;type:varchar(100);default:''" json:"remark"`       // 备注
	MoneyType        string `gorm:"column:money_type;type:varchar(8);default:''" json:"money_type"` // 币种
	Errcode          int    `gorm:"column:errcode;type:int(10);default:'0'" json:"errcode"`
	ConfirmDate      int    `gorm:"column:confirm_date;type:int(11);default:'0'" json:"confirm_date"`
	ConfirmTime      int    `gorm:"column:confirm_time;type:int(11);default:'0'" json:"confirm_time"`
	ConfirmMoneyType string `gorm:"column:confirm_money_type;type:varchar(8);default:''" json:"confirm_money_type"`
	ConfirmAmt       int64  `gorm:"column:confirm_amt;type:bigint(20);default:'0'" json:"confirm_amt"`
	ConfirmOpid      int64  `gorm:"column:confirm_opid;type:bigint(20);default:'0'" json:"confirm_opid"`
}

// get primary key name
func (t *Trade) GetKey() string {
	return "id"
}

// get primary key in model
func (t *Trade) GetKeyProperty() int64 {
	return t.Id
}

// set primary key
func (t *Trade) SetKeyProperty(id int64) {
	t.Id = id
}

// get table name
func (t *Trade) TableName() string {
	return tradeName
}

func GetTradeFirst() (*Trade, error) {
	var obj Trade
	err := client.DB(tradeName).Model(&Trade{}).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeLast() (*Trade, error) {
	var obj Trade
	err := client.DB(tradeName).Model(&Trade{}).Last(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeOne() (*Trade, error) {
	var obj Trade
	err := client.DB(tradeName).Model(&Trade{}).Take(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeById(id int64) (*Trade, error) {
	var obj Trade
	err := client.DB(tradeName).Model(&Trade{}).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeByCustomerId(cid int64) ([]*Trade, error) {
	var objs []*Trade
	err := client.DB(tradeName).Model(&Trade{}).Where("cid = ?", cid).Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetTradeByAccount(account string) (*Trade, error) {
	var obj Trade
	err := client.DB(tradeName).Model(&Trade{}).Where("account = ? limit 1", account).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeAll() ([]*Trade, error) {
	var objs []*Trade
	err := client.DB(tradeName).Model(&Trade{}).Order("id desc").Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetTrade(where string, args ...interface{}) ([]*Trade, error) {
	var objs []*Trade
	err := client.DB(tradeName).Model(&Trade{}).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetTradeList(page, limit int64, where string, args ...interface{}) ([]*Trade, error) {
	var objs []*Trade
	err := client.DB(tradeName).Model(&Trade{}).Limit(limit).Offset((page-1)*limit).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (t *Trade) Create() []error {
	return client.DB(tradeName).Model(&Trade{}).Create(t).GetErrors()
}

func (t *Trade) Update(obj Trade) []error {
	return client.DB(tradeName).Model(&Trade{}).UpdateColumns(obj).GetErrors()
}

func (t *Trade) UpdateById(id int64) (int64, error) {
	ravDatabase := client.DB(tradeName).Model(&Trade{}).Where("id=?", id).Update(t)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (t *Trade) Delete() {
	client.DB(tradeName).Model(&Trade{}).Delete(t)
}

func AddTradeTX(t *client.DBTransaction, obj Trade) error {
	return t.GetTx().Model(&Trade{}).Create(&obj).Error
}

func UpdateTradeTX(t *client.DBTransaction, obj Trade) error {
	return t.GetTx().Model(&Trade{}).Where("id=?", obj.Id).Update(obj).Error
}

func DeleteTradeTX(t *client.DBTransaction, obj Trade) error {
	return t.GetTx().Model(&Trade{}).Where("id=?", obj.Id).Delete(nil).Error
}
