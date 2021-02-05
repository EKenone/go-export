package conf

import (
	"flag"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Ept *EptConf
}

type EptConf struct {
	Path string `yaml:"path"`
	Port string `yaml:"port"`
}

var (
	confPath string
	Conf     *Config
	viperIns *viper.Viper
	newOnce  = sync.Once{}
)

func init() {
	flag.StringVar(&confPath, "conf", "conf.yaml", "default config path.")
}

func Init() (err error) {
	Conf = new(Config)
	err = GetViper().Unmarshal(&Conf)
	return
}

func GetViper() *viper.Viper {
	newOnce.Do(func() {
		viperIns = newViper()
	})

	return viperIns
}

func newViper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile(confPath)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("无法找到配置文件")
		} else {
			panic(err)
		}
	}
	return v
}
