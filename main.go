package main

import (
	"github.com/kensay98/vindecoder/app/handlers"
	"github.com/kensay98/vindecoder/config"
)

var conf *config.Config

func init() {
	conf = config.GetConfig()
}

func main() {
	addr := conf.Host + ":" + conf.Port
	handlers.GetApp().Run(addr)
}
