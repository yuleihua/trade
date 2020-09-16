package model

import client "github.com/yuleihua/aaa/dbclient"

var tradeSeqName = "tradeSeq"

type TradeSeq struct {
	Id        int64  `gorm:"column:id;type:bigint(20)" json:"id"`
	TransDate int    `gorm:"column:trans_date;type:int(11);default:'0'" json:"trans_date"`
	TransTime int    `gorm:"column:trans_time;type:int(11);default:'0'" json:"trans_time"`
	FromUuid  string `gorm:"column:from_uuid;type:varchar(64);default:''" json:"from_uuid"` // 流水号
	FromBid   int64  `gorm:"column:from_bid;type:bigint(20);default:'0'" json:"from_bid"`
	FromCid   int64  `gorm:"column:from_cid;type:bigint(20);default:'0'" json:"from_cid"`
	ToUuid    string `gorm:"column:to_uuid;type:varchar(64);default:''" json:"to_uuid"`
	ToBid     int64  `gorm:"column:to_bid;type:bigint(20);default:'0'" json:"to_bid"`
	Uuid      string `gorm:"column:uuid;type:varchar(64);default:''" json:"uuid"`
	Errcode   int    `gorm:"column:errcode;type:int(4);default:'0'" json:"errcode"`
}

// get primary key name
func (t *TradeSeq) GetKey() string {
	return "id"
}

// get primary key in model
func (t *TradeSeq) GetKeyProperty() int64 {
	return t.Id
}

// set primary key
func (t *TradeSeq) SetKeyProperty(id int64) {
	t.Id = id
}

// get table name
func (t *TradeSeq) TableName() string {
	return tradeSeqName
}

func GetTradeSeqFirst() (*TradeSeq, error) {
	var obj TradeSeq
	err := client.DB(tradeSeqName).Model(&TradeSeq{}).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeSeqLast() (*TradeSeq, error) {
	var obj TradeSeq
	err := client.DB(tradeSeqName).Model(&TradeSeq{}).Last(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeSeqOne() (*TradeSeq, error) {
	var obj TradeSeq
	err := client.DB(tradeSeqName).Model(&TradeSeq{}).Take(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeSeqById(id int64) (*TradeSeq, error) {
	var obj TradeSeq
	err := client.DB(tradeSeqName).Model(&TradeSeq{}).Where("id = ?", id).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeSeqByCustomerId(cid int64) ([]*TradeSeq, error) {
	var objs []*TradeSeq
	err := client.DB(tradeSeqName).Model(&TradeSeq{}).Where("cid = ?", cid).Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetTradeSeqByAccount(account string) (*TradeSeq, error) {
	var obj TradeSeq
	err := client.DB(tradeSeqName).Model(&TradeSeq{}).Where("account = ? limit 1", account).Find(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func GetTradeSeqAll() ([]*TradeSeq, error) {
	var objs []*TradeSeq
	err := client.DB(tradeSeqName).Model(&TradeSeq{}).Order("id desc").Find(&objs).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetTradeSeq(where string, args ...interface{}) ([]*TradeSeq, error) {
	var objs []*TradeSeq
	err := client.DB(tradeSeqName).Model(&TradeSeq{}).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func GetTradeSeqList(page, limit int64, where string, args ...interface{}) ([]*TradeSeq, error) {
	var objs []*TradeSeq
	err := client.DB(tradeSeqName).Model(&TradeSeq{}).Limit(limit).Offset((page-1)*limit).Find(&objs, where, args).Error
	if err != nil {
		return nil, err
	}
	return objs, nil
}

func (t *TradeSeq) Create() []error {
	return client.DB(tradeSeqName).Model(&TradeSeq{}).Create(t).GetErrors()
}

func (t *TradeSeq) Update(obj TradeSeq) []error {
	return client.DB(tradeSeqName).Model(&TradeSeq{}).UpdateColumns(obj).GetErrors()
}

func (t *TradeSeq) UpdateById(id int64) (int64, error) {
	ravDatabase := client.DB(tradeSeqName).Model(&TradeSeq{}).Where("id=?", id).Update(t)
	return ravDatabase.RowsAffected, ravDatabase.Error
}

func (t *TradeSeq) Delete() {
	client.DB(tradeSeqName).Model(&TradeSeq{}).Delete(t)
}

func AddTradeSeqTX(t *client.DBTransaction, obj TradeSeq) error {
	return t.GetTx().Model(&TradeSeq{}).Create(&obj).Error
}

func UpdateTradeSeqTX(t *client.DBTransaction, obj TradeSeq) error {
	return t.GetTx().Model(&TradeSeq{}).Where("id=?", obj.Id).Update(obj).Error
}

func DeleteTradeSeqTX(t *client.DBTransaction, obj TradeSeq) error {
	return t.GetTx().Model(&TradeSeq{}).Where("id=?", obj.Id).Delete(nil).Error
}
