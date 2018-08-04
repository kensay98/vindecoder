package storage

import (
	"bitbucket.org/babadjanov/vindecoder/logger"
	"bitbucket.org/babadjanov/vindecoder/config"
)

type Storage interface {
	Get(key string) string
	Set(key, value string)
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
