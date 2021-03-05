package export

import (
	"go-export/internal/conf"
	"go-export/pkg/upload"
)

const (
	aliy  = "oss"
	qiniu = "qiniu"
	txy   = "txy"
	loc   = "loc"
)

func toYun(filename string, filePath string) string {
	s := upload.Conf{
		Filename: filename,
		FilePath: filePath,
	}

	switch conf.Conf.Ept.UploadServer {
	case aliy:
		oss := conf.GetOssConf()
		s.Dir = oss.Dir
		return s.Uploader(&upload.Oss{
			Conf:            s,
			Endpoint:        oss.Endpoint,
			AccessKeyId:     oss.AccessKeyId,
			AccessKeySecret: oss.AccessKeySecret,
			BucketName:      oss.BucketName,
		})
	case qiniu:
		qny := conf.GetQnyConf()
		s.Dir = qny.Dir
		return s.Uploader(&upload.Qiniu{
			Conf:       s,
			AccessKey:  qny.AccessKey,
			SecretKey:  qny.SecretKey,
			BucketName: qny.BucketName,
			Host:       qny.Host,
		})
	case txy:
		txy := conf.GetTxyConf()
		s.Dir = txy.Dir
		return s.Uploader(&upload.Txy{
			Conf:      s,
			Host:      txy.Host,
			SecretId:  txy.SecretId,
			SecretKey: txy.SecretKey,
		})
	default:
		loc := conf.GetLocConf()
		s.Dir = loc.Dir
		return s.Uploader(&upload.Loc{
			Conf:     s,
			Host:     loc.Host,
			HostPath: loc.HostPath,
		})
	}
}
