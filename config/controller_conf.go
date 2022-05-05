package config

type Controller struct {
	BackendListenAddr string `yaml:"BackendListenAddr"`
	BackendListenPort int64  `yaml:"BackendListenPort"`
	User              string `yaml:"User"`
	Password          string `yaml:"Password"`
}
