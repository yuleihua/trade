package model

import client "github.com/yuleihua/aaa/dbclient"

var tradeRiskName = "tradeRisk"

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

// get primary key name
func (t *TradeRisk) GetKey() string {
	return "id"
}

// get primary key in model
func (t *TradeRisk) GetKeyProperty() int64 {
	return t.Id
}

// set primary key
func (t *TradeRisk) SetKeyProperty(id int64) {
	t.Id = id
}

// get table name
func (t *TradeRisk) TableName() string {
	return tradeRiskName
}

func GetTradeRiskFirst() (*TradeRisk, error) {
	var obj TradeRisk
	err := client.DB(tradeRiskName).Model(&TradeRisk{}).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeRiskLast() (*TradeRisk, error) {
	var obj TradeRisk
	err := client.DB(tradeRiskName).Model(&TradeRisk{}).Last(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeRiskOne() (*TradeRisk, error) {
	var obj TradeRisk
	err := client.DB(tradeRiskName).Model(&TradeRisk{}).Take(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeRiskById(id int64) (*TradeRisk, error) {
	var obj TradeRisk
	err := client.DB(tradeRiskName).Model(&TradeRisk{}).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeRiskByCustomerId(cid int64) ([]*TradeRisk, error) {
	var objs []*TradeRisk
	err := client.DB(tradeRiskName).Model(&TradeRisk{}).Where("cid = ?", cid).Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetTradeRiskByAccount(account string) (*TradeRisk, error) {
	var obj TradeRisk
	err := client.DB(tradeRiskName).Model(&TradeRisk{}).Where("account = ? limit 1", account).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeRiskAll() ([]*TradeRisk, error) {
	var objs []*TradeRisk
	err := client.DB(tradeRiskName).Model(&TradeRisk{}).Order("id desc").Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetTradeRisk(where string, args ...interface{}) ([]*TradeRisk, error) {
	var objs []*TradeRisk
	err := client.DB(tradeRiskName).Model(&TradeRisk{}).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetTradeRiskList(page, limit int64, where string, args ...interface{}) ([]*TradeRisk, error) {
	var objs []*TradeRisk
	err := client.DB(tradeRiskName).Model(&TradeRisk{}).Limit(limit).Offset((page-1)*limit).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (t *TradeRisk) Create() []error {
	return client.DB(tradeRiskName).Model(&TradeRisk{}).Create(t).GetErrors()
}

func (t *TradeRisk) Update(obj TradeRisk) []error {
	return client.DB(tradeRiskName).Model(&TradeRisk{}).UpdateColumns(obj).GetErrors()
}

func (t *TradeRisk) UpdateById(id int64) (int64, error) {
	ravDatabase := client.DB(tradeRiskName).Model(&TradeRisk{}).Where("id=?", id).Update(t)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (t *TradeRisk) Delete() {
	client.DB(tradeRiskName).Model(&TradeRisk{}).Delete(t)
}

func AddTradeRiskTX(t *client.DBTransaction, obj TradeRisk) error {
	return t.GetTx().Model(&TradeRisk{}).Create(&obj).Error
}

func UpdateTradeRiskTX(t *client.DBTransaction, obj TradeRisk) error {
	return t.GetTx().Model(&TradeRisk{}).Where("id=?", obj.Id).Update(obj).Error
}

func DeleteTradeRiskTX(t *client.DBTransaction, obj TradeRisk) error {
	return t.GetTx().Model(&TradeRisk{}).Where("id=?", obj.Id).Delete(nil).Error
}
