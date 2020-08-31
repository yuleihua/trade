package model

import "github.com/yuleihua/trade/pkg/dbclient"

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

//get real primary key name
func (customer *Customer) GetKey() string {
	return "id"
}

//get primary key in model
func (customer *Customer) GetKeyProperty() int64 {
	return customer.Id
}

//set primary key
func (customer *Customer) SetKeyProperty(id int64) {
	customer.Id = id
}

//get real table name
func (customer *Customer) TableName() string {
	return "customer"
}

func GetCustomerFirst() (*Customer, error) {
	var c Customer
	err := dbclient.DB(customerName).Model(&Customer{}).First(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func GetCustomerLast() (*Customer, error) {
	var c Customer
	err := dbclient.DB(customerName).Model(&Customer{}).Last(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func GetCustomerOne() (*Customer, error) {
	var c Customer
	err := dbclient.DB(customerName).Model(&Customer{}).Take(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func GetCustomerById(id int64) (*Customer, error) {
	var c Customer
	err := dbclient.DB(customerName).Model(&Customer{}).Where("id = ?", id).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func GetCustomerByPhone(nation, phone string) (*Customer, error) {
	var c Customer
	err := dbclient.DB(customerName).Model(&Customer{}).Where("nation = ? and phone = ?", nation, phone).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func GetCustomerByEmail(nation, email string) (*Customer, error) {
	var c Customer
	err := dbclient.DB(customerName).Model(&Customer{}).Where("nation = ? and email = ?", nation, email).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func GetCustomerByName(nation, name string) (*Customer, error) {
	var c Customer
	err := dbclient.DB(customerName).Model(&Customer{}).Where("nation = ? and name = ?", nation, name).Find(&c).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func GetCustomerAll() ([]*Customer, error) {
	var cs []*Customer
	err := dbclient.DB(customerName).Model(&Customer{}).Order("id desc").Find(&cs).Error
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func GetCustomer(where string, args ...interface{}) ([]*Customer, error) {
	var cs []*Customer
	err := dbclient.DB(customerName).Model(&Customer{}).Find(&cs, where, args).Error
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func GetCustomerList(page, limit int64, where string, args ...interface{}) ([]*Customer, error) {
	var cs []*Customer
	err := dbclient.DB(customerName).Model(&Customer{}).Limit(limit).Offset((page-1)*limit).Find(&cs, where, args).Error
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func (customer *Customer) Create() []error {
	return dbclient.DB(customerName).Model(&Customer{}).Create(customer).GetErrors()
}

func (customer *Customer) Update(c Customer) []error {
	return dbclient.DB(customerName).Model(&Customer{}).UpdateColumns(c).GetErrors()
}

func (customer *Customer) UpdateById(id int64) (int64, error) {
	ravDatabase := dbclient.DB(customerName).Model(&Customer{}).Where("id=?", id).Update(customer)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (customer *Customer) Delete() {
	dbclient.DB(customerName).Model(&Customer{}).Delete(customer)
}

func AddCustomerTX(t *dbclient.DBTransaction, c Customer) error {
	return t.GetTx().Create(c).Model(&Customer{}).Create(&c).Error
}

func UpdateCustomerTX(t *dbclient.DBTransaction, c Customer) error {
	return t.GetTx().Model(&Customer{}).Where("id=?", c.Id).Update(c).Error
}

func DeleteCustomerTX(t *dbclient.DBTransaction, c Customer) error {
	return t.GetTx().Model(&Customer{}).Where("id=?", c.Id).Delete(nil).Error
}
