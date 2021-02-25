package conf

import "sync"

type TxyConf struct {
	Host      string `yaml:"host"`
	SecretId  string `yaml:"secretId"`
	SecretKey string `yaml:"secretKey"`
	Dir       string `yaml:"dir"`
}

var (
	txyConf *TxyConf
	txyOnce sync.Once
)

func GetTxyConf() *TxyConf {
	txyOnce.Do(func() {
		_ = GetViper().UnmarshalKey("txy", &txyConf)
	})
	return txyConf
}
