package export

import (
	"go-export/internal/conf"
	"go-export/pkg/upload"
)

const (
	aliy  = "oss"
	qiniu = "qiniu"
	txy   = "txy"
)

func toYun(filename string, filePath string) string {
	s := upload.Conf{
		Filename: filename,
		FilePath: filePath,
	}

	switch conf.Conf.Ept.UploadServer {
	case aliy:
		s.Dir = conf.Conf.Oss.Dir
		return s.Uploader(&upload.Oss{
			Conf:            s,
			Endpoint:        conf.Conf.Oss.Endpoint,
			AccessKeyId:     conf.Conf.Oss.AccessKeyId,
			AccessKeySecret: conf.Conf.Oss.AccessKeySecret,
			BucketName:      conf.Conf.Oss.BucketName,
		})
	case qiniu:
		s.Dir = conf.Conf.Qiniu.Dir
		return s.Uploader(&upload.Qiniu{
			Conf:       s,
			AccessKey:  conf.Conf.Qiniu.AccessKey,
			SecretKey:  conf.Conf.Qiniu.SecretKey,
			BucketName: conf.Conf.Qiniu.BucketName,
			Host:       conf.Conf.Qiniu.Host,
		})
	case txy:
		s.Dir = conf.Conf.Txy.Dir
		return s.Uploader(&upload.Txy{
			Conf:      s,
			Host:      conf.Conf.Txy.Host,
			SecretId:  conf.Conf.Txy.SecretId,
			SecretKey: conf.Conf.Txy.SecretKey,
		})
	}

	s.Dir = conf.Conf.Loc.Dir
	return s.Uploader(&upload.Loc{
		Conf: s,
		Host: conf.Conf.Loc.Host,
	})
}
