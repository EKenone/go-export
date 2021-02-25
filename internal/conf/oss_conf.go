package conf

import "sync"

type OssConf struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
	Dir             string `yaml:"dir"`
}

var (
	ossConf *OssConf
	ossOnce sync.Once
)

func GetOssConf() *OssConf {
	ossOnce.Do(func() {
		_ = GetViper().UnmarshalKey("oss", &ossConf)
	})
	return ossConf
}
