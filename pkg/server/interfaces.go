package server

import "github.com/yuleihua/trade/conf"

type Service interface {
	Init(c *conf.Config) error
	Start() error
	Stop() error
	Name() string
}
