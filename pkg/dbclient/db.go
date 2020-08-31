package dbclient

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/yuleihua/trade/conf"
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

func NewDBClient(c *conf.Database) (*gorm.DB, error) {
	sqlConn := c.Dsn
	db, err := gorm.Open("mysql", sqlConn)
	if err != nil {
		log.Errorf("open mysql error, %s,%v", sqlConn, err)
		return nil, err
	}
	if err := db.DB().Ping(); err != nil {
		log.Errorf("ping mysql error, %s,%v", sqlConn, err)
		return nil, err
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.SetLogger(log.StandardLogger())
	db.LogMode(true)
	defaultDB = db

	return db, nil
}
