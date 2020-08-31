package model

import "github.com/yuleihua/trade/pkg/dbclient"

var tradeSeqName = "trade_seq"

type TradeSeq struct {
	Id        int64  `gorm:"column:id;type:bigint(20)" json:"id"`
	TransDate int    `gorm:"column:trans_date;type:int(11);default:'0'" json:"trans_date"`
	TransTime int    `gorm:"column:trans_time;type:int(11);default:'0'" json:"trans_time"`
	FromUuid  string `gorm:"column:from_uuid;type:varchar(64);default:''" json:"from_uuid"` // 流水号
	ToUuid    string `gorm:"column:to_uuid;type:varchar(64);default:''" json:"to_uuid"`     // 流水号
	Uuid      string `gorm:"column:uuid;type:varchar(64);default:''" json:"uuid"`           // 流水号
	FromBid   int64  `gorm:"column:from_bid;type:bigint(20);default:'0'" json:"from_bid"`
	FromCid   int64  `gorm:"column:from_cid;type:bigint(20);default:'0'" json:"from_cid"`
	ToBid     int64  `gorm:"column:to_bid;type:bigint(20);default:'0'" json:"to_bid"`
	Errcode   int    `gorm:"column:errcode;type:int(10);default:'0'" json:"errcode"`
}

//get real primary key name
func (tradeSeq *TradeSeq) GetKey() string {
	return "id"
}

//get primary key in model
func (tradeSeq *TradeSeq) GetKeyProperty() int64 {
	return tradeSeq.Id
}

//set primary key
func (tradeSeq *TradeSeq) SetKeyProperty(id int64) {
	tradeSeq.Id = id
}

//get real table name
func (tradeSeq *TradeSeq) TableName() string {
	return "trade_seq"
}

func GetTradeSeqFirst() (*TradeSeq, error) {
	var t TradeSeq
	err := dbclient.DB(tradeSeqName).Model(&TradeSeq{}).First(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeSeqLast() (*TradeSeq, error) {
	var t TradeSeq
	err := dbclient.DB(tradeSeqName).Model(&TradeSeq{}).Last(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeSeqOne() (*TradeSeq, error) {
	var t TradeSeq
	err := dbclient.DB(tradeSeqName).Model(&TradeSeq{}).Take(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeSeqById(id int64) (*TradeSeq, error) {
	var t TradeSeq
	err := dbclient.DB(tradeSeqName).Model(&TradeSeq{}).Where("id = ?", id).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeSeqByFromUUID(uuid string) (*TradeSeq, error) {
	var t TradeSeq
	err := dbclient.DB(tradeSeqName).Model(&TradeSeq{}).Where("from_uuid = ?", uuid).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeSeqByUUID(uuid string) (*TradeSeq, error) {
	var t TradeSeq
	err := dbclient.DB(tradeSeqName).Model(&TradeSeq{}).Where("uuid = ?", uuid).Find(&t).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetTradeSeqAll() ([]*TradeSeq, error) {
	var ts []*TradeSeq
	err := dbclient.DB(tradeSeqName).Model(&TradeSeq{}).Order("id desc").Find(&ts).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func GetTradeSeq(where string, args ...interface{}) ([]*TradeSeq, error) {
	var ts []*TradeSeq
	err := dbclient.DB(tradeSeqName).Model(&TradeSeq{}).Find(&ts, where, args).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func GetTradeSeqList(page, limit int64, where string, args ...interface{}) ([]*TradeSeq, error) {
	var ts []*TradeSeq
	err := dbclient.DB(tradeSeqName).Model(&TradeSeq{}).Limit(limit).Offset((page-1)*limit).Find(&ts, where, args).Error
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (tradeSeq *TradeSeq) Create() error {
	return dbclient.DB(tradeSeqName).Model(&TradeSeq{}).Create(tradeSeq).Error
}

func (tradeSeq *TradeSeq) Update(t TradeSeq) []error {
	return dbclient.DB(tradeSeqName).Model(&TradeSeq{}).UpdateColumns(t).GetErrors()
}

func (tradeSeq *TradeSeq) UpdateById(id int64) (int64, error) {
	ravDatabase := dbclient.DB(tradeSeqName).Model(&TradeSeq{}).Where("id=?", id).Update(tradeSeq)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (tradeSeq *TradeSeq) Delete() {
	dbclient.DB(tradeSeqName).Model(&TradeSeq{}).Delete(tradeSeq)
}

func AddTradeSeqTX(tx *dbclient.DBTransaction, t TradeSeq) error {
	return tx.GetTx().Model(&TradeSeq{}).Create(&t).Error
}

func UpdateTradeSeqTX(tx *dbclient.DBTransaction, t TradeSeq) error {
	return tx.GetTx().Model(&TradeSeq{}).Where("id=?", t.Id).Update(t).Error
}

func DeleteTradeSeqTX(tx *dbclient.DBTransaction, t TradeSeq) error {
	return tx.GetTx().Model(&TradeSeq{}).Where("id=?", t.Id).Delete(nil).Error
}
