package model

import client "github.com/yuleihua/aaa/dbclient"

var fundName = "fund"

type Fund struct {
	Id            int64  `gorm:"column:id;type:bigint(20)" json:"id"`
	Bid           int64  `gorm:"column:bid;type:bigint(20);default:'0'" json:"bid"`
	MoneyType     string `gorm:"column:money_type;type:varchar(8);default:''" json:"money_type"` // 源IP地址
	Balance       int64  `gorm:"column:balance;type:bigint(20);default:'0'" json:"balance"`
	FreezeBalance int64  `gorm:"column:freeze_balance;type:bigint(20);default:'0'" json:"freeze_balance"`
	LastBalance   int64  `gorm:"column:last_balance;type:bigint(20);default:'0'" json:"last_balance"`
	Created       string `gorm:"column:created;type:timestamp;default:CURRENT_TIMESTAMP" json:"created"` // 创建时间
}

// get primary key name
func (f *Fund) GetKey() string {
	return "id"
}

// get primary key in model
func (f *Fund) GetKeyProperty() int64 {
	return f.Id
}

// set primary key
func (f *Fund) SetKeyProperty(id int64) {
	f.Id = id
}

// get table name
func (f *Fund) TableName() string {
	return fundName
}

func GetFundFirst() (*Fund, error) {
	var obj Fund
	err := client.DB(fundName).Model(&Fund{}).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetFundLast() (*Fund, error) {
	var obj Fund
	err := client.DB(fundName).Model(&Fund{}).Last(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetFundOne() (*Fund, error) {
	var obj Fund
	err := client.DB(fundName).Model(&Fund{}).Take(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetFundById(id int64) (*Fund, error) {
	var obj Fund
	err := client.DB(fundName).Model(&Fund{}).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetFundByCustomerId(cid int64) ([]*Fund, error) {
	var objs []*Fund
	err := client.DB(fundName).Model(&Fund{}).Where("cid = ?", cid).Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetFundByAccount(account string) (*Fund, error) {
	var obj Fund
	err := client.DB(fundName).Model(&Fund{}).Where("account = ? limit 1", account).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetFundAll() ([]*Fund, error) {
	var objs []*Fund
	err := client.DB(fundName).Model(&Fund{}).Order("id desc").Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetFund(where string, args ...interface{}) ([]*Fund, error) {
	var objs []*Fund
	err := client.DB(fundName).Model(&Fund{}).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetFundList(page, limit int64, where string, args ...interface{}) ([]*Fund, error) {
	var objs []*Fund
	err := client.DB(fundName).Model(&Fund{}).Limit(limit).Offset((page-1)*limit).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (f *Fund) Create() []error {
	return client.DB(fundName).Model(&Fund{}).Create(f).GetErrors()
}

func (f *Fund) Update(obj Fund) []error {
	return client.DB(fundName).Model(&Fund{}).UpdateColumns(obj).GetErrors()
}

func (f *Fund) UpdateById(id int64) (int64, error) {
	ravDatabase := client.DB(fundName).Model(&Fund{}).Where("id=?", id).Update(f)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (f *Fund) Delete() {
	client.DB(fundName).Model(&Fund{}).Delete(f)
}

func AddFundTX(t *client.DBTransaction, obj Fund) error {
	return t.GetTx().Model(&Fund{}).Create(&obj).Error
}

func UpdateFundTX(t *client.DBTransaction, obj Fund) error {
	return t.GetTx().Model(&Fund{}).Where("id=?", obj.Id).Update(obj).Error
}

func DeleteFundTX(t *client.DBTransaction, obj Fund) error {
	return t.GetTx().Model(&Fund{}).Where("id=?", obj.Id).Delete(nil).Error
}
