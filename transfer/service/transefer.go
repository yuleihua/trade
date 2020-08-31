package service

import (
	"context"
	"sync"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/yuleihua/trade/conf"
	"github.com/yuleihua/trade/pkg/dbclient"
	"github.com/yuleihua/trade/pkg/redis"
	"github.com/yuleihua/trade/transfer/types"
)

type Transfer struct {
	ctx       context.Context
	cancel    context.CancelFunc
	db        *gorm.DB
	cache     *redis.Client
	isRunning bool
	lock      sync.Mutex

	chanNotification chan []string
}

func NewTransfer() *Transfer {
	newCtx, newCancel := context.WithCancel(context.Background())
	return &Transfer{ctx: newCtx, cancel: newCancel}
}

func (t *Transfer) Init(c *conf.Config) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	dbClient, err := dbclient.NewDBClient(c.DB)
	if err != nil {
		log.Errorf("open database error, %v, %v", c.DB, err)
		return err
	}
	t.db = dbClient

	redisClient := redis.NewClient(c.Cache.Address, c.Cache.Password, c.Cache.DialTimeout, c.Cache.ReadTimeout, c.Cache.WriteTimeout, c.Cache.PoolSize)
	if err := redisClient.Open(); err != nil {
		log.Errorf("open redis error, %v, %v", c.Cache, err)
		return err
	}
	t.cache = redisClient

	return nil
}

func (t *Transfer) Start() error {
	t.lock.Lock()
	defer t.lock.Unlock()

	if !t.isRunning {
		go t.loop()
		t.isRunning = true
	}

	return nil
}

func (t *Transfer) Stop() error {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.isRunning {
		t.cancel()
		t.isRunning = false
	}

	return nil
}

func (t *Transfer) Name() string {
	return types.ServiceName
}

func (t *Transfer) loop() {
myLoop:
	for {
		select {
		case <-t.ctx.Done():
			break myLoop

		case item := <-t.chanNotification:
			// dummy sending email or sms
			log.Info(item)
		}
	}
	return
}
