package decoder

import (
	"github.com/kensay98/vindecoder/config"
)

type DecodedVin struct {
	Vin    string `json:"vin"`
	Make   string `json:"make"`
	Model  string `json:"model"`
	Year   string `json:"year"`
	Color  string `json:"color"`
	Weight string `json:"weight"`
	Type   string `json:"type"`
}

type Decoder interface {
	Decode(vin string) (DecodedVin, error)
}

var (
	UsableDecoder Decoder
	conf *config.Config
)

func init() {
	conf = config.GetConfig()
	UsableDecoder = NewDecodeThisDecoder(conf.DecodeThisApiKey)
}
