package upload

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"net/http"
	"net/url"
	"sync"
)

type TxyBucket struct {
	once sync.Once
	cli  *cos.Client
}

var txyBkt = &TxyBucket{
	once: sync.Once{},
	cli:  nil,
}

type Txy struct {
	Conf
	Host      string
	SecretId  string
	SecretKey string
}

func (txy *Txy) GetTxyBucket() *cos.Client {
	txyBkt.once.Do(func() {
		u, _ := url.Parse(txy.Host)
		b := &cos.BaseURL{BucketURL: u}
		txyBkt.cli = cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  txy.SecretId,
				SecretKey: txy.SecretKey,
			},
		})
	})
	return txyBkt.cli
}

func (txy *Txy) Upload() string {
	txyName := txy.Dir + txy.Filename + ".csv"

	_, err := txy.GetTxyBucket().Object.PutFromFile(context.Background(), txyName, txy.FilePath, nil)

	if err != nil {
		log.Println(err)
	}

	return txy.FileUrl(txyName)
}

func (txy *Txy) FileUrl(txyName string) string {
	return txy.Host + "/" + txyName
}
