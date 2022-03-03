package config

import (
	"testing"
)

func TestInitSrvConfig(t *testing.T) {

	InitSrvConfig("f:\\code\\go\\c2c\\server.yaml")
}

func TestLoadConfigFromConf(t *testing.T) {
	loadConfigFromConf("")
}

func TestCheckConfig(t *testing.T) {
	SConfig.Server.ListenPort = 5899
	SConfig.Server.ControlPort = 5899
	err := checkConfig()
	if err != nil {
		panic(err)
	}
}
