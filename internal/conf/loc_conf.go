package conf

import "sync"

type LocConf struct {
	Host     string `yaml:"host"`
	HostPath string `yaml:"hostPath"`
	Dir      string `yaml:"dir"`
}

var (
	locConf *LocConf
	locOnce sync.Once
)

func GetLocConf() *LocConf {
	locOnce.Do(func() {
		_ = GetViper().UnmarshalKey("loc", &locConf)
	})
	return locConf
}
