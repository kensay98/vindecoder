package storage

import (
	"github.com/kensay98/vindecoder/logger"
	"github.com/kensay98/vindecoder/config"
)

type Storage interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

var storage Storage
var log = logger.GetLogger()

func init() {
	conf := config.GetConfig()
	host := conf.Redis.Host + ":" + conf.Redis.Port
	storage = NewRedisStorage(host, conf.Redis.Password, conf.Redis.DB)
}

func GetStorage() Storage {
	return storage
}
