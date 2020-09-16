package model

import client "github.com/yuleihua/aaa/dbclient"

var tradeFeeName = "tradeFee"

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

// get primary key name
func (t *TradeFee) GetKey() string {
	return "id"
}

// get primary key in model
func (t *TradeFee) GetKeyProperty() int64 {
	return t.Id
}

// set primary key
func (t *TradeFee) SetKeyProperty(id int64) {
	t.Id = id
}

// get table name
func (t *TradeFee) TableName() string {
	return tradeFeeName
}

func GetTradeFeeFirst() (*TradeFee, error) {
	var obj TradeFee
	err := client.DB(tradeFeeName).Model(&TradeFee{}).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeFeeLast() (*TradeFee, error) {
	var obj TradeFee
	err := client.DB(tradeFeeName).Model(&TradeFee{}).Last(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeFeeOne() (*TradeFee, error) {
	var obj TradeFee
	err := client.DB(tradeFeeName).Model(&TradeFee{}).Take(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeFeeById(id int64) (*TradeFee, error) {
	var obj TradeFee
	err := client.DB(tradeFeeName).Model(&TradeFee{}).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeFeeByCustomerId(cid int64) ([]*TradeFee, error) {
	var objs []*TradeFee
	err := client.DB(tradeFeeName).Model(&TradeFee{}).Where("cid = ?", cid).Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetTradeFeeByAccount(account string) (*TradeFee, error) {
	var obj TradeFee
	err := client.DB(tradeFeeName).Model(&TradeFee{}).Where("account = ? limit 1", account).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeFeeAll() ([]*TradeFee, error) {
	var objs []*TradeFee
	err := client.DB(tradeFeeName).Model(&TradeFee{}).Order("id desc").Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetTradeFee(where string, args ...interface{}) ([]*TradeFee, error) {
	var objs []*TradeFee
	err := client.DB(tradeFeeName).Model(&TradeFee{}).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetTradeFeeList(page, limit int64, where string, args ...interface{}) ([]*TradeFee, error) {
	var objs []*TradeFee
	err := client.DB(tradeFeeName).Model(&TradeFee{}).Limit(limit).Offset((page-1)*limit).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (t *TradeFee) Create() []error {
	return client.DB(tradeFeeName).Model(&TradeFee{}).Create(t).GetErrors()
}

func (t *TradeFee) Update(obj TradeFee) []error {
	return client.DB(tradeFeeName).Model(&TradeFee{}).UpdateColumns(obj).GetErrors()
}

func (t *TradeFee) UpdateById(id int64) (int64, error) {
	ravDatabase := client.DB(tradeFeeName).Model(&TradeFee{}).Where("id=?", id).Update(t)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (t *TradeFee) Delete() {
	client.DB(tradeFeeName).Model(&TradeFee{}).Delete(t)
}

func AddTradeFeeTX(t *client.DBTransaction, obj TradeFee) error {
	return t.GetTx().Model(&TradeFee{}).Create(&obj).Error
}

func UpdateTradeFeeTX(t *client.DBTransaction, obj TradeFee) error {
	return t.GetTx().Model(&TradeFee{}).Where("id=?", obj.Id).Update(obj).Error
}

func DeleteTradeFeeTX(t *client.DBTransaction, obj TradeFee) error {
	return t.GetTx().Model(&TradeFee{}).Where("id=?", obj.Id).Delete(nil).Error
}
