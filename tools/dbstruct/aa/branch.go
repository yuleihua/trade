package model

import client "github.com/yuleihua/aaa/dbclient"

var branchName = "branch"

type Branch struct {
	Id        int64  `gorm:"column:id;type:bigint(20)" json:"id"`
	Cid       int64  `gorm:"column:cid;type:bigint(20);default:'0'" json:"cid"`
	ShortName string `gorm:"column:short_name;type:varchar(100);default:''" json:"short_name"` // 简称
	Name      string `gorm:"column:name;type:varchar(255);default:''" json:"name"`             // 名字
	Bankid    int64  `gorm:"column:bankid;type:bigint(20);default:'0'" json:"bankid"`          // 银行编号
	MoneyType string `gorm:"column:money_type;type:varchar(8);default:''" json:"money_type"`
	Account   string `gorm:"column:account;type:varchar(100);default:''" json:"account"` // 账户
	Status    int    `gorm:"column:status;type:int(10);default:'0'" json:"status"`
	IsMaster  int8   `gorm:"column:is_master;type:tinyint(1);default:'0'" json:"is_master"`
	Created   string `gorm:"column:created;type:timestamp;default:CURRENT_TIMESTAMP" json:"created"` // 创建时间
}

// get primary key name
func (b *Branch) GetKey() string {
	return "id"
}

// get primary key in model
func (b *Branch) GetKeyProperty() int64 {
	return b.Id
}

// set primary key
func (b *Branch) SetKeyProperty(id int64) {
	b.Id = id
}

// get table name
func (b *Branch) TableName() string {
	return branchName
}

func GetBranchFirst() (*Branch, error) {
	var obj Branch
	err := client.DB(branchName).Model(&Branch{}).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetBranchLast() (*Branch, error) {
	var obj Branch
	err := client.DB(branchName).Model(&Branch{}).Last(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetBranchOne() (*Branch, error) {
	var obj Branch
	err := client.DB(branchName).Model(&Branch{}).Take(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetBranchById(id int64) (*Branch, error) {
	var obj Branch
	err := client.DB(branchName).Model(&Branch{}).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetBranchByCustomerId(cid int64) ([]*Branch, error) {
	var objs []*Branch
	err := client.DB(branchName).Model(&Branch{}).Where("cid = ?", cid).Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetBranchByAccount(account string) (*Branch, error) {
	var obj Branch
	err := client.DB(branchName).Model(&Branch{}).Where("account = ? limit 1", account).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetBranchAll() ([]*Branch, error) {
	var objs []*Branch
	err := client.DB(branchName).Model(&Branch{}).Order("id desc").Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetBranch(where string, args ...interface{}) ([]*Branch, error) {
	var objs []*Branch
	err := client.DB(branchName).Model(&Branch{}).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetBranchList(page, limit int64, where string, args ...interface{}) ([]*Branch, error) {
	var objs []*Branch
	err := client.DB(branchName).Model(&Branch{}).Limit(limit).Offset((page-1)*limit).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (b *Branch) Create() []error {
	return client.DB(branchName).Model(&Branch{}).Create(b).GetErrors()
}

func (b *Branch) Update(obj Branch) []error {
	return client.DB(branchName).Model(&Branch{}).UpdateColumns(obj).GetErrors()
}

func (b *Branch) UpdateById(id int64) (int64, error) {
	ravDatabase := client.DB(branchName).Model(&Branch{}).Where("id=?", id).Update(b)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (b *Branch) Delete() {
	client.DB(branchName).Model(&Branch{}).Delete(b)
}

func AddBranchTX(t *client.DBTransaction, obj Branch) error {
	return t.GetTx().Model(&Branch{}).Create(&obj).Error
}

func UpdateBranchTX(t *client.DBTransaction, obj Branch) error {
	return t.GetTx().Model(&Branch{}).Where("id=?", obj.Id).Update(obj).Error
}

func DeleteBranchTX(t *client.DBTransaction, obj Branch) error {
	return t.GetTx().Model(&Branch{}).Where("id=?", obj.Id).Delete(nil).Error
}
