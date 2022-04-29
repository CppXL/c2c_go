package config

import (
	"c2c/common/logger"
	"c2c/common/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// 服务端配置结构体
type agent struct {
	ListenAddr  string `yaml:"ListenAddr" default:"0.0.0.0"`
	ControlAddr string `yaml:"ControlAddr" default:"0.0.0.0"`
	ListenPort  int64  `yaml:"ListenPort" default:"56689"`
	ControlPort int64  `yaml:"ControlPort" default:"57888"`
	User        string `yaml:"User" default:"admin"`
	Password    string `yaml:"Password" default:"P@ssW0r3"`
}

// logging 结构体
type logging struct {
	Level []string `yaml:"Level"`
	Path  string   `yaml:"Path"`
}

// 顶层结构体
type appConfig struct {
	Tcp struct {
		Agent   agent   `yaml:"Agent"`
		Logging logging `yaml:"Logging"`
	} `yaml:"Tcp"`
}

// 默认配置文件路径
const (
	MaxPort = 65535
	MinPort = 1024
)

var (
	// TODO:初始化为程序运行路径拼接 .server.yaml
	defaultConfPath = "../server.yml"
	SConfig         *appConfig
)

// 返回sconfig指针
func newServerConfigPoint() *appConfig {
	return &appConfig{}
}

// 初始化服务端配置 传入配置文件路径
func InitSrvConfig(confPath string) error {
	return loadConfigFromConf(confPath)
}

func init() {
	SConfig = newServerConfigPoint()
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
			panic(err)
		}
	}
	logger.Infof("try to load config from %s", confPath)

	// 判断文件是否存在 返回false可能是目录
	if !utils.IsFileExists(confPath) {
		// 路径无效或者是文件夹则尝试加载默认配置
		loadDefaultConf()
		return fmt.Errorf("path %s does not exists or is directory", confPath)
	}

	// 如果有效 尝试载入配置文件
	cont := []byte{}
	cont1, err := ioutil.ReadFile(confPath)
	if err != nil {
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
	scont, err := json.MarshalIndent(SConfig, "", "  ")
	logger.FatalIfError(err)
	logger.Infof("config is %s", scont)
	logger.Infof("scont is %s\n", scont)
	// fmt.Printf("config is %+v\n", *SConfig)
	//
	err = checkConfig()
	logger.FatalIfError(err)
	if err != nil {
		panic(err)
	}

	return nil
}

// 从struct中载入默认值
func loadDefaultConf() {

}
func checkConfig() error {
	// 检查配置文件配置是否正确
	// Listen端口是否在范围内
	if value, err := utils.IsNumInRange(SConfig.Tcp.Agent.ListenPort, MinPort, MaxPort, [2]byte{'(', ')'}); !value && err == nil {
		return errors.New("listen port out of range")
	}
	// 控制端口是否在范围内
	if value, err := utils.IsNumInRange(SConfig.Tcp.Agent.ControlPort, MinPort, MaxPort, [2]byte{'(', ')'}); !value && err == nil {
		return errors.New("control port out of range")
	}

	// 是否相等 端口
	if SConfig.Tcp.Agent.ListenPort == SConfig.Tcp.Agent.ControlPort {
		return errors.New("listen port and control port is equal")
	}

	// todo 检查IP地址是否合法
	//

	return nil
}
