package config

import (
	"testing"
)

func TestInitSrvConfig(t *testing.T) {

	err := InitConfig("")
	if err != nil {
		t.Error(err)
	}
}

func TestLoadConfigFromConf(t *testing.T) {
	err := loadConfigFromConf("../config.yml")
	if err != nil {
		t.Error(err)
	}
}

func TestCheckConfig(t *testing.T) {
	err := checkConfig()
	if err != nil {
		t.Error(err)
	}
}
