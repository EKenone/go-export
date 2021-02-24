package upload

type Upload interface {
	Upload() string
}

type Conf struct {
	Filename string
	FilePath string
	Dir      string
}

func (cf *Conf) Uploader(up Upload) string {
	return up.Upload()
}
