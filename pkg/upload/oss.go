package upload

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"os"
	"sync"
)

type OssBucket struct {
	once sync.Once
	cli  *oss.Bucket
}

var bkt = &OssBucket{
	once: sync.Once{},
	cli:  nil,
}

type Oss struct {
	Conf
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

func (o *Oss) GetOssBucket() *oss.Bucket {
	bkt.once.Do(func() {
		cli, err := oss.New(o.Endpoint, o.AccessKeyId, o.AccessKeySecret)
		if err != nil {
			log.Println(err)
		}
		bucket, _ := cli.Bucket(o.BucketName)
		bkt.cli = bucket
	})

	return bkt.cli
}

func (o *Oss) Upload() string {
	file, _ := os.Open(o.FilePath)
	defer file.Close()

	ossName := o.Dir + o.Filename + ".csv"

	err := o.GetOssBucket().PutObject(ossName, file)

	if err != nil {
		log.Println(err)
	}

	return o.FileUrl(ossName)
}

func (o *Oss) FileUrl(ossName string) string {
	return "https://" + o.BucketName + "." + o.Endpoint + "/" + ossName
}
