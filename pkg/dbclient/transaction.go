package dbclient

import (
	"errors"

	"github.com/jinzhu/gorm"
)

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
