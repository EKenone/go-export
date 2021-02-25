package conf

import "sync"

type QiniuConf struct {
	AccessKey  string `yaml:"accessKey"`
	SecretKey  string `yaml:"secretKey"`
	BucketName string `yaml:"bucketName"`
	Host       string `yaml:"host"`
	Dir        string `yaml:"dir"`
}

var (
	qnyConf *QiniuConf
	qnyOnce sync.Once
)

func GetQnyConf() *QiniuConf {
	qnyOnce.Do(func() {
		_ = GetViper().UnmarshalKey("qiniu", &qnyConf)
	})
	return qnyConf
}
