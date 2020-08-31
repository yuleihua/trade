package redis

import (
	"time"

	"github.com/go-redis/redis"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	Address      string
	Password     string
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	PoolSize     int
	client       *redis.Client
}

func NewClient(address, password string, dialTime, readTime, writeTime, poolSize int) *Client {
	return &Client{
		Address:      address,
		Password:     password,
		DialTimeout:  dialTime,
		ReadTimeout:  readTime,
		WriteTimeout: writeTime,
		PoolSize:     poolSize,
	}
}

func (c *Client) Open() error {
	c.client = redis.NewClient(&redis.Options{
		Addr:               c.Address,
		DialTimeout:        time.Duration(c.DialTimeout) * time.Second,
		ReadTimeout:        time.Duration(c.ReadTimeout) * time.Second,
		WriteTimeout:       time.Duration(c.WriteTimeout) * time.Second,
		PoolSize:           c.PoolSize,
		PoolTimeout:        10 * time.Second,
		IdleTimeout:        time.Second,
		IdleCheckFrequency: time.Second,
		Password:           c.Password,
	})

	if _, err := c.client.Ping().Result(); err != nil {
		log.Errorf("cluster ping error, address:%s, error:%v", c.Address, err)
		return err
	}
	return nil
}

func (c *Client) Close() error {
	if err := c.client.Close(); err != nil {
		log.Warnf("cluster close error, address:%s, error:%v", c.Address, err)
		return err
	}
	return nil
}

func (c *Client) Delete(key string) error {
	if err := c.client.Del(key).Err(); err != nil {
		log.Errorf("cluster delete error, key:%s, error:%v", key, err)
		return err
	}
	return nil
}

func (c *Client) Get(key string) (string, error) {
	value, err := c.client.Get(key).Result()
	if err != nil && err != redis.Nil {
		log.Errorf("cluster close error, key:%s, error:%v", key, err)
		return "", err
	}
	return value, nil
}

func (c *Client) Set(key string, val interface{}, expiration time.Duration) error {
	if err := c.client.Set(key, val, expiration).Err(); err != nil {
		log.Errorf("cluster close error, key:%s, error:%v", key, err)
		return err
	}
	return nil
}

func (c *Client) SetNX(key string, val interface{}, expiration time.Duration) error {
	if err := c.client.SetNX(key, val, expiration).Err(); err != nil {
		log.Errorf("cluster close error, key:%s, error:%v", key, err)
		return err
	}
	return nil
}

func (c *Client) HGetAll(key string) (map[string]string, error) {
	info, err := c.client.HGetAll(key).Result()
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (c *Client) HGetAllWithStruct(key string, value interface{}) error {
	info, err := c.client.HGetAll(key).Result()
	if err != nil {
		return err
	}

	if err := ToStruct(info, value); err != nil {
		return err
	}

	return nil
}

func (c *Client) HMGet(key string, fields ...string) ([]interface{}, error) {
	return c.client.HMGet(key, fields...).Result()
}

func (c *Client) HMSet(key string, value interface{}) error {
	hmv, err := ToMap(value)
	if err != nil {
		return err
	}
	if err := c.client.HMSet(key, hmv).Err(); err != nil {
		return err
	}
	return nil
}

func (c *Client) HMSetWithExpire(key string, expire time.Duration, value interface{}) error {
	hmv, err := ToMap(value)
	if err != nil {
		return err
	}

	if err := c.client.HMSet(key, hmv).Err(); err != nil {
		return err
	}
	return c.client.Expire(key, expire).Err()
}

func (c *Client) HDel(key string, field string) error {
	return c.client.HDel(key, field).Err()
}

func (c *Client) SAdd(key string, members ...interface{}) error {
	err := c.client.SAdd(key, members...).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SAddWithExpire(key string, expire time.Duration, members ...interface{}) error {
	err := c.client.SAdd(key, members...).Err()
	if err != nil {
		return err
	}
	return c.client.Expire(key, expire).Err()
}

func (c *Client) SMembers(key string) ([]string, error) {
	return c.client.SMembers(key).Result()
}

func (c *Client) Expire(key string, expiration time.Duration) (bool, error) {
	return c.client.Expire(key, expiration).Result()
}
