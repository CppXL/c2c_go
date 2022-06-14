package config

type Controller struct {
	BackendListenAddr string `yaml:"BackendListenAddr"`
	BackendListenPort uint16 `yaml:"BackendListenPort"`
	User              string `yaml:"User"`
	Password          string `yaml:"Password"`
}
