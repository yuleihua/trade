package model

var GormTpl = `
func Get{{object}}First() (*{{object}}, error) {
	var obj {{object}}
	err := client.DB({{entry}}Name).Model(&{{object}}{}).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func Get{{object}}Last() (*{{object}}, error) {
	var obj {{object}}
	err := client.DB({{entry}}Name).Model(&{{object}}{}).Last(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func Get{{object}}One() (*{{object}}, error) {
	var obj {{object}}
	err := client.DB({{entry}}Name).Model(&{{object}}{}).Take(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func Get{{object}}ById(id int64) (*{{object}}, error) {
	var obj {{object}}
	err := client.DB({{entry}}Name).Model(&{{object}}{}).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func Get{{object}}ByCustomerId(cid int64) ([]*{{object}}, error) {
	var objs []*{{object}}
	err := client.DB({{entry}}Name).Model(&{{object}}{}).Where("cid = ?", cid).Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func Get{{object}}ByAccount(account string) (*{{object}}, error) {
	var obj {{object}}
	err := client.DB({{entry}}Name).Model(&{{object}}{}).Where("account = ? limit 1", account).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func Get{{object}}All() ([]*{{object}}, error) {
	var objs []*{{object}}
	err := client.DB({{entry}}Name).Model(&{{object}}{}).Order("id desc").Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func Get{{object}}(where string, args ...interface{}) ([]*{{object}}, error) {
	var objs []*{{object}}
	err := client.DB({{entry}}Name).Model(&{{object}}{}).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func Get{{object}}List(page, limit int64, where string, args ...interface{}) ([]*{{object}}, error) {
	var objs []*{{object}}
	err := client.DB({{entry}}Name).Model(&{{object}}{}).Limit(limit).Offset((page-1)*limit).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func ({{shortName}} *{{object}}) Create() []error {
	return client.DB({{entry}}Name).Model(&{{object}}{}).Create({{shortName}}).GetErrors()
}

func ({{shortName}} *{{object}}) Update(obj {{object}}) []error {
	return client.DB({{entry}}Name).Model(&{{object}}{}).UpdateColumns(obj).GetErrors()
}

func ({{shortName}} *{{object}}) UpdateById(id int64) (int64, error) {
	ravDatabase := client.DB({{entry}}Name).Model(&{{object}}{}).Where("id=?", id).Update({{shortName}})
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func ({{shortName}} *{{object}}) Delete() {
	client.DB({{entry}}Name).Model(&{{object}}{}).Delete({{shortName}})
}

func Add{{object}}TX(t *client.DBTransaction, obj {{object}}) error {
	return t.GetTx().Model(&{{object}}{}).Create(&obj).Error
}

func Update{{object}}TX(t *client.DBTransaction, obj {{object}}) error {
	return t.GetTx().Model(&{{object}}{}).Where("id=?", obj.Id).Update(obj).Error
}

func Delete{{object}}TX(t *client.DBTransaction, obj {{object}}) error {
	return t.GetTx().Model(&{{object}}{}).Where("id=?", obj.Id).Delete(nil).Error
}
`

var GormInit = `
package {{package}}

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

const (
	ErrorDBRecordNotFound = "db record not found"
	ErrorDBDuplicated     = "db duplicate entry"
)

var defaultDB *gorm.DB

func DBDefault() *gorm.DB {
	if defaultDB == nil {
		log.Fatalf("db is nil")
	}
	return defaultDB
}

func DB(table string) *gorm.DB {
	if defaultDB == nil {
		log.Fatalf("db is nil")
	}
	return defaultDB
}

func NewDBClient(dsn string, idleConns, maxConns int) (*gorm.DB, error) {
	sqlConn := dsn
	db, err := gorm.Open("mysql", sqlConn)
	if err != nil {
		log.Errorf("open mysql error, %s,%v", sqlConn, err)
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		log.Errorf("ping mysql error, %s,%v", sqlConn, err)
		return nil, err
	}

	db.DB().SetMaxIdleConns(idleConns)
	db.DB().SetMaxOpenConns(maxConns)

	defaultDB = db

	return db, nil
}

type DBTransaction struct {
	uuid     string
	Database *gorm.DB
	Tx       *gorm.DB
}

func NewDBTransaction(db *gorm.DB) *DBTransaction {
	return &DBTransaction{
		Database: db,
	}
}

func (t *DBTransaction) Begin() error {
	if t.Database == nil {
		return errors.New("db is nil")
	}
	tx := t.Database.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	t.Tx = tx
	return nil
}

func (t *DBTransaction) Commit() error {
	if t.Database == nil {
		return errors.New("db is nil")
	}

	return t.Tx.Commit().Error
}

func (t *DBTransaction) Rollback() error {
	if t.Database == nil {
		return errors.New("db is nil")
	}

	return t.Tx.Rollback().Error
}

func (t *DBTransaction) GetDatabase() *gorm.DB {
	return t.Database
}

func (t *DBTransaction) GetTx() *gorm.DB {
	return t.Tx
}

`
