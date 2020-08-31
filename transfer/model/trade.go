package model

import "github.com/yuleihua/trade/pkg/dbclient"

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
	IsDelay          bool   `gorm:"column:is_delay;type:tinyint(1);default:'0'" json:"is_delay"`
	IsLarge          bool   `gorm:"column:is_large;type:tinyint(1);default:'0'" json:"is_large"`
	IsReject         bool   `gorm:"column:is_reject;type:tinyint(1);default:'0'" json:"is_reject"`
	Amt              int64  `gorm:"column:amt;type:bigint(20);default:'0'" json:"amt"`
	Fee              int64  `gorm:"column:fee;type:bigint(20);default:'0'" json:"fee"`
	Remark           string `gorm:"column:remark;type:varchar(100);default:''" json:"remark"`       // 备注
	MoneyType        string `gorm:"column:money_type;type:varchar(8);default:''" json:"money_type"` // 币种
	Errcode          int    `gorm:"column:errcode;type:int(10);default:'0'" json:"errcode"`
	ConfirmDate      int    `gorm:"column:confirm_date;type:int(11);default:'0'" json:"confirm_date"`
	ConfirmTime      int    `gorm:"column:confirm_time;type:int(11);default:'0'" json:"confirm_time"`
	ConfirmAmt       int64  `gorm:"column:confirm_amt;type:bigint(20);default:'0'" json:"confirm_amt"`
	ConfirmOpid      int64  `gorm:"column:confirm_opid;type:bigint(20);default:'0'" json:"confirm_opid"`
	ConfirmMoneyType string `gorm:"column:confirm_money_type;type:varchar(8);default:''" json:"confirm_money_type"` // 币种
}

//get real primary key name
func (trade *Trade) GetKey() string {
	return "id"
}

//get primary key in model
func (trade *Trade) GetKeyProperty() int64 {
	return trade.Id
}

//set primary key
func (trade *Trade) SetKeyProperty(id int64) {
	trade.Id = id
}

//get real table name
func (trade *Trade) TableName() string {
	return "trade"
}

func GetTradeFirst() (*Trade, error) {
	var t Trade
	err := dbclient.DB(tradeName).Model(&Trade{}).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeLast() (*Trade, error) {
	var t Trade
	err := dbclient.DB(tradeName).Model(&Trade{}).Last(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeOne() (*Trade, error) {
	var t Trade
	err := dbclient.DB(tradeName).Model(&Trade{}).Take(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeById(id int64) (*Trade, error) {
	var t Trade
	err := dbclient.DB(tradeName).Model(&Trade{}).Where("id = ?", id).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeByUUId(uuid string) (*Trade, error) {
	var t Trade
	err := dbclient.DB(tradeName).Model(&Trade{}).Where("uuid = ?", uuid).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeAll() ([]*Trade, error) {
	var ts []*Trade
	err := dbclient.DB(tradeName).Model(&Trade{}).Order("id desc").Find(&ts).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func GetTrade(where string, args ...interface{}) ([]*Trade, error) {
	var ts []*Trade
	err := dbclient.DB(tradeName).Model(&Trade{}).Find(&ts, where, args).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func GetTradeList(page, limit int64, where string, args ...interface{}) ([]*Trade, error) {
	var ts []*Trade
	err := dbclient.DB(tradeName).Model(&Trade{}).Limit(limit).Offset((page-1)*limit).Find(&ts, where, args).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (trade *Trade) Create() []error {
	return dbclient.DB(tradeName).Model(&Trade{}).Create(trade).GetErrors()
}

func (trade *Trade) Update(t Trade) []error {
	return dbclient.DB(tradeName).Model(&Trade{}).UpdateColumns(t).GetErrors()
}

func (trade *Trade) UpdateById(id int64) (int64, error) {
	ravDatabase := dbclient.DB(tradeName).Model(&Trade{}).Where("id=?", id).Update(trade)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (trade *Trade) Delete() {
	dbclient.DB(tradeName).Model(&Trade{}).Delete(trade)
}

func AddTradeTX(tx *dbclient.DBTransaction, t Trade) error {
	return tx.GetTx().Model(&Trade{}).Create(&t).Error
}

func UpdateTradeTX(tx *dbclient.DBTransaction, t Trade) error {
	return tx.GetTx().Model(&Trade{}).Where("id=?", t.Id).Update(t).Error
}

func DeleteTradeTX(tx *dbclient.DBTransaction, t Trade) error {
	return tx.GetTx().Model(&Trade{}).Where("id=?", t.Id).Delete(nil).Error
}
