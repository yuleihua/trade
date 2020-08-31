package model

import "github.com/yuleihua/trade/pkg/dbclient"

var tradeRiskName = "trade_risk"

type TradeRisk struct {
	Id        int64  `gorm:"column:id;type:bigint(20)" json:"id"`
	TransDate int    `gorm:"column:trans_date;type:int(11);default:'0'" json:"trans_date"`
	TransTime int    `gorm:"column:trans_time;type:int(11);default:'0'" json:"trans_time"`
	Level     int    `gorm:"column:level;type:int(11);default:'0'" json:"level"`
	Uuid      string `gorm:"column:uuid;type:varchar(64);default:''" json:"uuid"` // 流水号
	Amt       int64  `gorm:"column:amt;type:bigint(20);default:'0'" json:"amt"`
	Fee       int64  `gorm:"column:fee;type:bigint(20);default:'0'" json:"fee"`
	Remark    string `gorm:"column:remark;type:varchar(100);default:''" json:"remark"` // 备注
}

//get real primary key name
func (tradeRisk *TradeRisk) GetKey() string {
	return "id"
}

//get primary key in model
func (tradeRisk *TradeRisk) GetKeyProperty() int64 {
	return tradeRisk.Id
}

//set primary key
func (tradeRisk *TradeRisk) SetKeyProperty(id int64) {
	tradeRisk.Id = id
}

//get real table name
func (tradeRisk *TradeRisk) TableName() string {
	return "trade_risk"
}

func GetTradeRiskFirst() (*TradeRisk, error) {
	var t TradeRisk
	err := dbclient.DB(tradeRiskName).Model(&TradeRisk{}).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeRiskLast() (*TradeRisk, error) {
	var t TradeRisk
	err := dbclient.DB(tradeRiskName).Model(&TradeRisk{}).Last(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeRiskOne() (*TradeRisk, error) {
	var t TradeRisk
	err := dbclient.DB(tradeRiskName).Model(&TradeRisk{}).Take(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeRiskById(id int64) (*TradeRisk, error) {
	var t TradeRisk
	err := dbclient.DB(tradeRiskName).Model(&TradeRisk{}).Where("id = ?", id).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeRiskByUUID(uuid string) (*TradeRisk, error) {
	var t TradeRisk
	err := dbclient.DB(tradeRiskName).Model(&TradeRisk{}).Where("uuid = ?", uuid).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeRiskAll() ([]*TradeRisk, error) {
	var ts []*TradeRisk
	err := dbclient.DB(tradeRiskName).Model(&TradeRisk{}).Order("id desc").Find(&ts).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func GetTradeRisk(where string, args ...interface{}) ([]*TradeRisk, error) {
	var ts []*TradeRisk
	err := dbclient.DB(tradeRiskName).Model(&TradeRisk{}).Find(&ts, where, args).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func GetTradeRiskList(page, limit int64, where string, args ...interface{}) ([]*TradeRisk, error) {
	var ts []*TradeRisk
	err := dbclient.DB(tradeRiskName).Model(&TradeRisk{}).Limit(limit).Offset((page-1)*limit).Find(&ts, where, args).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (tradeRisk *TradeRisk) Create() []error {
	return dbclient.DB(tradeRiskName).Model(&TradeRisk{}).Create(tradeRisk).GetErrors()
}

func (tradeRisk *TradeRisk) Update(t TradeRisk) []error {
	return dbclient.DB(tradeRiskName).Model(&TradeRisk{}).UpdateColumns(t).GetErrors()
}

func (tradeRisk *TradeRisk) UpdateById(id int64) (int64, error) {
	ravDatabase := dbclient.DB(tradeRiskName).Model(&TradeRisk{}).Where("id=?", id).Update(tradeRisk)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (tradeRisk *TradeRisk) Delete() {
	dbclient.DB(tradeRiskName).Model(&TradeRisk{}).Delete(tradeRisk)
}

func AddTradeRiskTX(tx *dbclient.DBTransaction, t TradeRisk) error {
	return tx.GetTx().Model(&TradeRisk{}).Create(&t).Error
}

func UpdateTradeRiskTX(tx *dbclient.DBTransaction, t TradeRisk) error {
	return tx.GetTx().Model(&TradeRisk{}).Where("id=?", t.Id).Update(t).Error
}

func DeleteTradeRiskTX(tx *dbclient.DBTransaction, t TradeRisk) error {
	return tx.GetTx().Model(&TradeRisk{}).Where("id=?", t.Id).Delete(nil).Error
}
