package model

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
