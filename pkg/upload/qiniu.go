package upload

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"sync"
)

type Qiniu struct {
	Conf
	AccessKey  string
	SecretKey  string
	BucketName string
	Host       string
}

type qnyBucket struct {
	once sync.Once
	cli  *storage.FormUploader
}

var qny = &qnyBucket{
	once: sync.Once{},
	cli:  nil,
}

func (q *Qiniu) GetQnyBucket() *storage.FormUploader {
	qny.once.Do(func() {
		cfg := storage.Config{
			Zone:          &storage.ZoneHuanan,
			UseCdnDomains: false,
		}
		qny.cli = storage.NewFormUploader(&cfg)
	})
	return qny.cli
}

func (q *Qiniu) Upload() string {
	ret := storage.PutRet{}
	qiniuName := q.Dir + q.Filename + ".csv"

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	err := q.GetQnyBucket().PutFile(context.Background(), &ret, q.GetToken(), qiniuName, q.FilePath, &putExtra)

	if err != nil {
		log.Println(err)
	}

	log.Println(ret)

	return q.FileUrl(qiniuName)
}

func (q *Qiniu) FileUrl(qiniuName string) string {
	return q.Host + "/" + qiniuName
}

func (q *Qiniu) GetToken() string {
	putPolicy := storage.PutPolicy{
		Scope: q.BucketName,
	}
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	return putPolicy.UploadToken(mac)
}
