package config

// tcp agent config struct
type TcpAgent struct {
	ListenAddr string `yaml:"ListenAddr" default:"0.0.0.0"`
	ListenPort uint16 `yaml:"ListenPort" default:"56689"`
	User       string `yaml:"User" default:"admin"`
	Password   string `yaml:"Password" default:"P@ssW0r3"`
}

// logging 结构体
type TcpLogging struct {
	Level string `yaml:"Level"`
	Path  string `yaml:"Path"`
}

// tcp agent 顶层结构体
type Config struct {
	Tcp struct {
		Agent   TcpAgent   `yaml:"Agent"`
		Logging TcpLogging `yaml:"Logging"`
	} `yaml:"Tcp"`
	Controller Controller `yaml:"Controller"`
}

// 默认配置文件路径
const (
	MaxPort = 65535
	MinPort = 1024
)

var (
	// TODO:初始化为程序运行路径拼接 .server.yaml
	defaultConfPath         = "/home/ub/code/c2c/c2c/config.yml"
	SConfig         *Config = newConfigPoint()
)
