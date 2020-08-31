package conf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Common struct {
	IsRelease bool   `yaml:"is_release"`
	LogLevel  int    `yaml:"log_level"`
	CPUs      int    `yaml:"cpus"`
	LogPath   string `yaml:"log_path"`
	LogFile   string `yaml:"log_file"`
	AppFile   string `yaml:"app_file"`
}

type WebServer struct {
	AppName   string   `yaml:"app"`
	Addr      string   `yaml:"address"`
	DomainApi string   `yaml:"domain_api"`
	Origins   []string `yaml:"origins"`
}

type Database struct {
	IsDebug  bool   `yaml:"is_debug"`
	Dsn      string `yaml:"dsn"`
	SqlFile  string `yaml:"sql_log"`
	IdleConn int    `yaml:"idle_conn"`
	MaxConn  int    `yaml:"max_conn"`
	LifeTime int    `yaml:"lifeTime"`
}

type Cache struct {
	Address      string `yaml:"address"`
	Password     string `yaml:"password"`
	DialTimeout  int    `yaml:"dial_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	PoolSize     int    `yaml:"pool_size"`
}

type Session struct {
	Type     string `yaml:"type"`
	Server   string `yaml:"server"`
	Password string `yaml:"password"`
}

type Opentracing struct {
	IsOnline    bool   `yaml:"is_online"`
	Type        string `yaml:"type"`
	ServiceName string `yaml:"service_name"`
	Address     string `yaml:"address"`
}

type Metrics struct {
	IsOnline bool          `yaml:"is_online"`
	FreqSec  time.Duration `yaml:"freq_sec"`
	Address  string        `yaml:"address"`
}

type Config struct {
	Cmn         *Common      `yaml:"common"`
	Server      *WebServer   `yaml:"server"`
	DB          *Database    `yaml:"database"`
	Cache       *Cache       `yaml:"cache"`
	OpenTracing *Opentracing `yaml:"tracing"`
	Metric      *Metrics     `yaml:"metric"`
}

func Setup(c string) (*Config, error) {
	if _, err := os.Stat(c); err != nil {
		return nil, fmt.Errorf("config %s error:%v", c, err)
	}

	log.Infof("load config from file:%s", c)

	configBytes, err := ioutil.ReadFile(c)
	if err != nil {
		return nil, errors.New("config load err:" + err.Error())
	}

	var conf Config
	if err := yaml.Unmarshal(configBytes, &conf); err != nil {
		log.Errorf("yaml unmarshal error:%v\n", err)
		return nil, err
	}
	log.Infof("config data:%#v", conf)

	return &conf, nil
}
