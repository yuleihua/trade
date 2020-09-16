package model

import client "github.com/yuleihua/aaa/dbclient"

var customerName = "customer"

type Customer struct {
	Id      int64  `gorm:"column:id;type:bigint(20)" json:"id"`
	Name    string `gorm:"column:name;type:varchar(255);default:''" json:"name"`       // 名字
	Nation  string `gorm:"column:nation;type:varchar(32);default:''" json:"nation"`    // 国家
	City    string `gorm:"column:city;type:varchar(32);default:''" json:"city"`        // 地址
	Address string `gorm:"column:address;type:varchar(255);default:''" json:"address"` // 地址
	Phone   string `gorm:"column:phone;type:varchar(100);default:''" json:"phone"`     // 电话
	Email   string `gorm:"column:email;type:varchar(100);default:''" json:"email"`     // 邮箱
	Remark  string `gorm:"column:remark;type:varchar(1024);default:null" json:"remark"`
	Created string `gorm:"column:created;type:timestamp;default:CURRENT_TIMESTAMP" json:"created"` // 创建时间
}

// get primary key name
func (c *Customer) GetKey() string {
	return "id"
}

// get primary key in model
func (c *Customer) GetKeyProperty() int64 {
	return c.Id
}

// set primary key
func (c *Customer) SetKeyProperty(id int64) {
	c.Id = id
}

// get table name
func (c *Customer) TableName() string {
	return customerName
}

func GetCustomerFirst() (*Customer, error) {
	var obj Customer
	err := client.DB(customerName).Model(&Customer{}).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCustomerLast() (*Customer, error) {
	var obj Customer
	err := client.DB(customerName).Model(&Customer{}).Last(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCustomerOne() (*Customer, error) {
	var obj Customer
	err := client.DB(customerName).Model(&Customer{}).Take(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCustomerById(id int64) (*Customer, error) {
	var obj Customer
	err := client.DB(customerName).Model(&Customer{}).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCustomerByCustomerId(cid int64) ([]*Customer, error) {
	var objs []*Customer
	err := client.DB(customerName).Model(&Customer{}).Where("cid = ?", cid).Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetCustomerByAccount(account string) (*Customer, error) {
	var obj Customer
	err := client.DB(customerName).Model(&Customer{}).Where("account = ? limit 1", account).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetCustomerAll() ([]*Customer, error) {
	var objs []*Customer
	err := client.DB(customerName).Model(&Customer{}).Order("id desc").Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetCustomer(where string, args ...interface{}) ([]*Customer, error) {
	var objs []*Customer
	err := client.DB(customerName).Model(&Customer{}).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetCustomerList(page, limit int64, where string, args ...interface{}) ([]*Customer, error) {
	var objs []*Customer
	err := client.DB(customerName).Model(&Customer{}).Limit(limit).Offset((page-1)*limit).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (c *Customer) Create() []error {
	return client.DB(customerName).Model(&Customer{}).Create(c).GetErrors()
}

func (c *Customer) Update(obj Customer) []error {
	return client.DB(customerName).Model(&Customer{}).UpdateColumns(obj).GetErrors()
}

func (c *Customer) UpdateById(id int64) (int64, error) {
	ravDatabase := client.DB(customerName).Model(&Customer{}).Where("id=?", id).Update(c)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (c *Customer) Delete() {
	client.DB(customerName).Model(&Customer{}).Delete(c)
}

func AddCustomerTX(t *client.DBTransaction, obj Customer) error {
	return t.GetTx().Model(&Customer{}).Create(&obj).Error
}

func UpdateCustomerTX(t *client.DBTransaction, obj Customer) error {
	return t.GetTx().Model(&Customer{}).Where("id=?", obj.Id).Update(obj).Error
}

func DeleteCustomerTX(t *client.DBTransaction, obj Customer) error {
	return t.GetTx().Model(&Customer{}).Where("id=?", obj.Id).Delete(nil).Error
}
