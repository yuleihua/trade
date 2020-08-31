package model

import "github.com/yuleihua/trade/pkg/dbclient"

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

//get real primary key name
func (fund *Fund) GetKey() string {
	return "id"
}

//get primary key in model
func (fund *Fund) GetKeyProperty() int64 {
	return fund.Id
}

//set primary key
func (fund *Fund) SetKeyProperty(id int64) {
	fund.Id = id
}

//get real table name
func (fund *Fund) TableName() string {
	return "fund"
}

func GetFundFirst() (*Fund, error) {
	var f Fund
	err := dbclient.DB(fundName).Model(&Fund{}).First(&f).Error
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func GetFundLast() (*Fund, error) {
	var f Fund
	err := dbclient.DB(fundName).Model(&Fund{}).Last(&f).Error
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func GetFundOne() (*Fund, error) {
	var f Fund
	err := dbclient.DB(fundName).Model(&Fund{}).Take(&f).Error
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func GetFundById(id int64) (*Fund, error) {
	var f Fund
	err := dbclient.DB(fundName).Model(&Fund{}).Where("id = ?", id).Find(&f).Error
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func GetFundByBranchId(bid int64) (*Fund, error) {
	var f Fund
	err := dbclient.DB(fundName).Model(&Fund{}).Where("bid = ?", bid).Find(&f).Error
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func GetFundAll() ([]*Fund, error) {
	var fs []*Fund
	err := dbclient.DB(fundName).Model(&Fund{}).Order("id desc").Find(&fs).Error
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func GetFund(where string, args ...interface{}) ([]*Fund, error) {
	var fs []*Fund
	err := dbclient.DB(fundName).Model(&Fund{}).Find(&fs, where, args).Error
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func GetFundList(page, limit int64, where string, args ...interface{}) ([]*Fund, error) {
	var fs []*Fund
	err := dbclient.DB(fundName).Model(&Fund{}).Limit(limit).Offset((page-1)*limit).Find(&fs, where, args).Error
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func (fund *Fund) Create() []error {
	return dbclient.DB(fundName).Model(&Fund{}).Create(fund).GetErrors()
}

func (fund *Fund) Update(f Fund) []error {
	return dbclient.DB(fundName).Model(&Fund{}).UpdateColumns(f).GetErrors()
}

func (fund *Fund) UpdateById(id int64) (int64, error) {
	ravDatabase := dbclient.DB(fundName).Model(&Fund{}).Where("id=?", id).Update(fund)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (fund *Fund) Delete() {
	dbclient.DB(fundName).Model(&Fund{}).Delete(fund)
}

func AddFundTX(t *dbclient.DBTransaction, f Fund) error {
	return t.GetTx().Model(&Fund{}).Create(&f).Error
}

func UpdateFundTX(t *dbclient.DBTransaction, f Fund) error {
	return t.GetTx().Model(&Fund{}).Where("id=?", f.Id).Update(f).Error
}

func UpdateFundTXDownBalance(t *dbclient.DBTransaction, fid, amt int64) error {
	return t.GetTx().Raw("UPDATE fund SET balance = (balance-?) WHERE id = ?", amt, fid).Error
}

func UpdateFundTXUpBalance(t *dbclient.DBTransaction, fid, amt int64) error {
	return t.GetTx().Raw("UPDATE fund SET balance = (balance+?) WHERE id = ?", amt, fid).Error
}

func UpdateFundTXDownFreezeBalance(t *dbclient.DBTransaction, fid, amt int64) error {
	return t.GetTx().Raw("UPDATE fund SET freeze_balance = (freeze_balance-?) WHERE id = ?", amt, fid).Error
}

func UpdateFundTXDownFreezeBalanceAll(t *dbclient.DBTransaction, fid, amt int64) error {
	return t.GetTx().Raw("UPDATE fund SET balance = (balance-?), freeze_balance = (freeze_balance+?) WHERE id = ?", amt, amt, fid).Error
}

func DeleteFundTX(t *dbclient.DBTransaction, f Fund) error {
	return t.GetTx().Model(&Fund{}).Where("id=?", f.Id).Delete(nil).Error
}
