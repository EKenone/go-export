package upload

import (
	"go-export/internal/conf"
)

type Upload interface {
	Upload() string
}

type Conf struct {
	Filename string
	FilePath string
}

const (
	aliy  = "oss"
	qiniu = "qiniu"
	txy   = "txy"
)

func (cf *Conf) Uploader(up Upload) string {
	return up.Upload()
}

func ToYun(filename string, filePath string) string {
	s := Conf{
		Filename: filename,
		FilePath: filePath,
	}

	switch conf.Conf.Ept.UploadServer {
	case aliy:
		return s.Uploader(&Oss{
			Conf: s,
		})
	case qiniu:
		return ""
	case txy:
		return s.Uploader(&Txy{
			Conf: s,
		})
	}

	return s.Uploader(&Loc{
		Conf: s,
	})
}
