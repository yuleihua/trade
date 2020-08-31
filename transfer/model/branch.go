package model

import (
	"github.com/yuleihua/trade/pkg/dbclient"
)

var branchName = "branch"

type Branch struct {
	Id        int64  `gorm:"column:id;type:bigint(20)" json:"id"`
	Cid       int64  `gorm:"column:cid;type:bigint(20);default:'0'" json:"cid"`
	ShortName string `gorm:"column:short_name;type:varchar(100);default:''" json:"short_name"` // 简称
	Name      string `gorm:"column:name;type:varchar(255);default:''" json:"name"`             // 名字
	Bankid    int64  `gorm:"column:bankid;type:bigint(20);default:'0'" json:"bankid"`          // 银行编号
	Account   string `gorm:"column:account;type:varchar(100);default:''" json:"account"`       // 账户
	IsMaster  int8   `gorm:"column:is_master;type:tinyint(1);default:'0'" json:"is_master"`
	Status    int    `gorm:"column:status;type:int(10);default:'0'" json:"status"`
	Created   string `gorm:"column:created;type:timestamp;default:CURRENT_TIMESTAMP" json:"created"` // 创建时间
}

//get real primary key name
func (branch *Branch) GetKey() string {
	return "id"
}

//get primary key in model
func (branch *Branch) GetKeyProperty() int64 {
	return branch.Id
}

//set primary key
func (branch *Branch) SetKeyProperty(id int64) {
	branch.Id = id
}

//get real table name
func (branch *Branch) TableName() string {
	return "branch"
}

func GetBranchFirst() (*Branch, error) {
	var b Branch
	err := dbclient.DB(branchName).Model(&Branch{}).First(&b).Error
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func GetBranchLast() (*Branch, error) {
	var b Branch
	err := dbclient.DB(branchName).Model(&Branch{}).Last(&b).Error
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func GetBranchOne() (*Branch, error) {
	var b Branch
	err := dbclient.DB(branchName).Model(&Branch{}).Take(&b).Error
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func GetBranchById(id int64) (*Branch, error) {
	var b Branch
	err := dbclient.DB(branchName).Model(&Branch{}).Where("id = ?", id).Find(&b).Error
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func GetBranchByCustomerId(cid int64) ([]*Branch, error) {
	var bs []*Branch
	err := dbclient.DB(branchName).Model(&Branch{}).Where("cid = ?", cid).Find(&bs).Error
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func GetBranchByAccount(account string) (*Branch, error) {
	var b Branch
	err := dbclient.DB(branchName).Model(&Branch{}).Where("account = ? limit 1", account).Find(&b).Error
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func GetBranchAll() ([]*Branch, error) {
	var bs []*Branch
	err := dbclient.DB(branchName).Model(&Branch{}).Order("id desc").Find(&bs).Error
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func GetBranch(where string, args ...interface{}) ([]*Branch, error) {
	var bs []*Branch
	err := dbclient.DB(branchName).Model(&Branch{}).Find(&bs, where, args).Error
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func GetBranchList(page, limit int64, where string, args ...interface{}) ([]*Branch, error) {
	var bs []*Branch
	err := dbclient.DB(branchName).Model(&Branch{}).Limit(limit).Offset((page-1)*limit).Find(&bs, where, args).Error
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (branch *Branch) Create() []error {
	return dbclient.DB(branchName).Model(&Branch{}).Create(branch).GetErrors()
}

func (branch *Branch) Update(b Branch) []error {
	return dbclient.DB(branchName).Model(&Branch{}).UpdateColumns(b).GetErrors()
}

func (branch *Branch) UpdateById(id int64) (int64, error) {
	ravDatabase := dbclient.DB(branchName).Model(&Branch{}).Where("id=?", id).Update(branch)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (branch *Branch) Delete() {
	dbclient.DB(branchName).Model(&Branch{}).Delete(branch)
}

func AddBranchTX(t *dbclient.DBTransaction, b Branch) error {
	return t.GetTx().Model(&Branch{}).Create(&b).Error
}

func UpdateBranchTX(t *dbclient.DBTransaction, b Branch) error {
	return t.GetTx().Model(&Branch{}).Where("id=?", b.Id).Update(b).Error
}

func DeleteBranchTX(t *dbclient.DBTransaction, b Branch) error {
	return t.GetTx().Model(&Branch{}).Where("id=?", b.Id).Delete(nil).Error
}
