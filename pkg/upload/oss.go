package upload

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go-export/internal/conf"
	"log"
	"os"
	"sync"
)

const PathPrefix = "export/"

type OssBucket struct {
	once *sync.Once
	cli  *oss.Bucket
}

var bkt = &OssBucket{
	once: &sync.Once{},
	cli:  nil,
}

func getOssBucket() *oss.Bucket {
	bkt.once.Do(func() {
		cli, err := oss.New(conf.Conf.Oss.Endpoint, conf.Conf.Oss.AccessKeyId, conf.Conf.Oss.AccessKeySecret)
		if err != nil {
			log.Panicln(err)
		}
		bucket, _ := cli.Bucket(conf.Conf.Oss.BucketName)
		bkt.cli = bucket
	})

	return bkt.cli
}

type Oss struct {
	Conf
}

func (oss *Oss) Upload() string {
	file, _ := os.Open(oss.FilePath)

	ossName := PathPrefix + oss.Filename + ".csv"

	err := getOssBucket().PutObject(ossName, file)

	if err != nil {
		log.Println(err)
	}

	return ossFileUrl(ossName)
}

func ossFileUrl(ossName string) string {
	return "https://" + conf.Conf.Oss.BucketName + "." + conf.Conf.Oss.Endpoint + "/" + ossName
}
