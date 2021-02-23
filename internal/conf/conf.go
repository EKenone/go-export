package conf

import (
	"flag"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Ept   *EptConf
	Oss   *OssConf
	Txy   *TxyConf
	Loc   *LocConf
	Redis *RedisConf
}

type EptConf struct {
	Path         string `yaml:"path"`
	Port         string `yaml:"port"`
	UploadServer string `yaml:"uploadServer"`
}

type OssConf struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
	Dir             string `yaml:"dir"`
}

type TxyConf struct {
	Host      string `yaml:"host"`
	SecretId  string `yaml:"secretId"`
	SecretKey string `yaml:"secretKey"`
	Dir       string `yaml:"dir"`
}

type LocConf struct {
	Host string `yaml:"host"`
	Dir  string `yaml:"dir"`
}

type RedisConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pwd  string `yaml:"user"`
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
