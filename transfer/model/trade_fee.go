package model

import "github.com/yuleihua/trade/pkg/dbclient"

var tradeFeeName = "trade_fee"

type TradeFee struct {
	Id        int64  `gorm:"column:id;type:bigint(20)" json:"id"`
	TransDate int    `gorm:"column:trans_date;type:int(11);default:'0'" json:"trans_date"`
	TransTime int    `gorm:"column:trans_time;type:int(11);default:'0'" json:"trans_time"`
	Uuid      string `gorm:"column:uuid;type:varchar(64);default:''" json:"uuid"`            // 流水号
	MoneyType string `gorm:"column:money_type;type:varchar(8);default:''" json:"money_type"` // 币种
	Amt       int64  `gorm:"column:amt;type:bigint(20);default:'0'" json:"amt"`
	Fee       int64  `gorm:"column:fee;type:bigint(20);default:'0'" json:"fee"`
	Remark    string `gorm:"column:remark;type:varchar(100);default:''" json:"remark"` // 备注
}

//get real primary key name
func (tradeFee *TradeFee) GetKey() string {
	return "id"
}

//get primary key in model
func (tradeFee *TradeFee) GetKeyProperty() int64 {
	return tradeFee.Id
}

//set primary key
func (tradeFee *TradeFee) SetKeyProperty(id int64) {
	tradeFee.Id = id
}

//get real table name
func (tradeFee *TradeFee) TableName() string {
	return "trade_fee"
}

func GetTradeFee(where string, args ...interface{}) ([]*TradeFee, error) {
	var ts []*TradeFee
	err := dbclient.DB(tradeFeeName).Model(&TradeFee{}).Find(&ts, where, args).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func GetTradeFeeList(page, limit int64, where string, args ...interface{}) ([]*TradeFee, error) {
	var ts []*TradeFee
	err := dbclient.DB(tradeFeeName).Model(&TradeFee{}).Limit(limit).Offset((page-1)*limit).Find(&ts, where, args).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (tradeFee *TradeFee) Create() error {
	return dbclient.DB(tradeFeeName).Model(&TradeFee{}).Create(tradeFee).Error
}

func (tradeFee *TradeFee) Update(t TradeFee) []error {
	return dbclient.DB(tradeFeeName).Model(&TradeFee{}).UpdateColumns(t).GetErrors()
}

func (tradeFee *TradeFee) UpdateById(id int64) (int64, error) {
	ravDatabase := dbclient.DB(tradeFeeName).Model(&TradeFee{}).Where("id=?", id).Update(tradeFee)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (tradeFee *TradeFee) Delete() {
	dbclient.DB(tradeFeeName).Model(&TradeFee{}).Delete(tradeFee)
}

func AddTradeFeeTX(tx *dbclient.DBTransaction, t TradeFee) error {
	return tx.GetTx().Model(&TradeFee{}).Create(&t).Error
}

func UpdateTradeFeeTX(tx *dbclient.DBTransaction, t TradeFee) error {
	return tx.GetTx().Model(&TradeFee{}).Where("id=?", t.Id).Update(t).Error
}

func DeleteTradeFeeTX(tx *dbclient.DBTransaction, t TradeFee) error {
	return tx.GetTx().Model(&TradeFee{}).Where("id=?", t.Id).Delete(nil).Error
}
