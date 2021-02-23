package upload

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go-export/internal/conf"
	"log"
	"net/http"
	"net/url"
	"sync"
)

type TxyBucket struct {
	once *sync.Once
	cli  *cos.Client
}

var txyBkt = &TxyBucket{
	once: &sync.Once{},
	cli:  nil,
}

func getTxyBucket() *cos.Client {
	txyBkt.once.Do(func() {
		u, _ := url.Parse(conf.Conf.Txy.Host)
		b := &cos.BaseURL{BucketURL: u}
		txyBkt.cli = cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  conf.Conf.Txy.SecretId,
				SecretKey: conf.Conf.Txy.SecretKey,
			},
		})
	})
	return txyBkt.cli
}

type Txy struct {
	Conf
}

func (txy *Txy) Upload() string {
	txyName := conf.Conf.Txy.Dir + txy.Filename + ".csv"

	_, err := getTxyBucket().Object.PutFromFile(context.Background(), txyName, txy.FilePath, nil)

	if err != nil {
		log.Println(err)
	}

	return txyFileUrl(txyName)
}

func txyFileUrl(txyName string) string {
	return conf.Conf.Txy.Host + "/" + txyName
}
