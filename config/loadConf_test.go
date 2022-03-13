package config

import (
	"testing"
)

func TestInitSrvConfig(t *testing.T) {

	InitSrvConfig("")
}

func TestLoadConfigFromConf(t *testing.T) {
	loadConfigFromConf("../server.yml")
}

func TestCheckConfig(t *testing.T) {
	err := checkConfig()
	if err != nil {
		panic(err)
	}
}
