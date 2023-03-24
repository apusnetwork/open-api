package config

import (
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var globalConfig Config

type Config struct {
	Env   Env
	Mysql Mysql `yaml:"Mysql"`
}

type Env struct {
	RunMode string // dev、prod
	Port    string
}

type Mysql struct {
	ReadDSN  string `yaml:"ReadDSN"`
	WriteDSN string `yaml:"WriteDSN"`
}

func init() {
	os.Setenv("APUS_RUN_MODE", "prod")

	viper.SetDefault("RunMode", "dev")
	viper.SetEnvPrefix("APUS")
	viper.BindEnv("RUN_MODE")

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		hlog.Fatalf("Fatal error config file: %s \n", err)
	}

	// 配置文件读取异常处理
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			hlog.Fatalf("no such config file: %s \n", err)
		} else {
			hlog.Fatalf("read config error: %s \n", err)
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		hlog.Infof("Config file changed: %s", e.Name)
	})

	globalConfig.Mysql = Mysql{
		WriteDSN: viper.GetString("Mysql.WriteDSN"),
		ReadDSN:  viper.GetString("Mysql.ReadDSN"),
	}
	globalConfig.Env = Env{
		RunMode: viper.GetString("RUN_MODE"),
		Port:    viper.GetString("Env.Port"),
	}

	hlog.Infof("Config file has loaded: %#v", globalConfig)
}

// Get 返回副本，以确保Config不被外部逻辑更改
func Get() Config {
	return globalConfig
}
