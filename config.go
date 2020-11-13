package wua

import (
	"fmt"
	"github.com/wuapp/rj"
)

var configFile = "config.rj"

type config struct {
	*rj.Node
}

var Config = loadConfig()

func loadConfig() *config {
	cfg := new(config)

	var err error
	cfg.Node, err = rj.Load(configFile)

	if err != nil {
		fmt.Println("Open config file failed, filename:", configFile, ", error:", err.Error())
	}

	return cfg
}

func GetConfig(c interface{}) (err error) {
	//Config.GetStruct()
	return Config.Node.ToStruct(c)
}
