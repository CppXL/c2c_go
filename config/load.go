package config

import (
	"c2c/common/logger"
	"c2c/common/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// 返回sconfig指针
func newConfigPoint() *Config {
	return &Config{}
}

// 初始化服务端配置 传入配置文件路径
func InitConfig(confPath string) error {
	return loadConfigFromConf(confPath)
}

func loadConfigFromConf(confPath string) error {
	var err error
	// 如果路径为空则根据默认路径找配置文件
	if confPath == "" {
		confPath = defaultConfPath
	}

	// 判断是不是绝对路径
	// 如果不是绝对路径则转换为绝对路径
	if !filepath.IsAbs(confPath) {
		confPath, err = filepath.Abs(confPath)
		if err != nil {
			logger.Errorf(err.Error())
		}
	}
	logger.Infof("try to load config from %s", confPath)

	// 判断文件是否存在 返回false可能是目录
	if !utils.IsFileExists(confPath) {
		// 路径无效或者是文件夹则尝试加载默认配置
		loadDefaultConf()
		return fmt.Errorf("path %s does not exists or is directory", confPath)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile(confPath)
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Errorf("read config file error: %s", err.Error())
		} else {
			logger.Errorf("read config file error: %s", err.Error())
		}
	}

	// 如果有效 尝试载入配置文件
	cont := []byte{}
	// 读取文件
	cont1, err := ioutil.ReadFile(confPath)
	if err != nil {
		logger.Errorf(err.Error())
		return err
	}
	if cont1 != nil {
		cont = cont1
	}
	// 读取到了文件
	if len(cont) != 0 {
		// 反序列化配置文件内容
		err := yaml.UnmarshalStrict(cont, SConfig)
		// 如果报错则直接抛出
		logger.FatalIfError(err)
	}
	// 格式化
	scont, err := json.MarshalIndent(SConfig, "", "  ")
	logger.FatalIfError(err)
	logger.Infof("config is %s", scont)
	logger.Infof("scont is %s\n", scont)
	err = checkConfig()
	logger.FatalIfError(err)
	if err != nil {
		logger.Errorf(err.Error())
		return err
	}

	return nil
}

// 从tag中载入默认值
func loadDefaultConf() {

}
func checkConfig() error {
	// 检查配置文件配置是否正确
	// Listen端口是否在范围内
	if value, err := utils.IsNumInRange(int64(SConfig.Tcp.Agent.ListenPort), MinPort, MaxPort, [2]byte{'(', ')'}); !value && err == nil {
		return errors.New("listen port out of range")
	}
	// 控制端口是否在范围内
	// if value, err := utils.IsNumInRange(SConfig.Tcp.Agent.ControlPort, MinPort, MaxPort, [2]byte{'(', ')'}); !value && err == nil {
	// 	return errors.New("control port out of range")
	// }

	// // 是否相等 端口
	// if SConfig.Tcp.Agent.ListenPort == SConfig.Tcp.Agent.ControlPort {
	// 	return errors.New("listen port and control port is equal")
	// }

	// todo 检查IP地址是否合法
	//

	return nil
}
