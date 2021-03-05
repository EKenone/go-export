package upload

import (
	"io"
	"log"
	"os"
)

type Loc struct {
	Conf
	Host     string
	HostPath string
}

func (loc *Loc) Upload() string {
	path := loc.Dir + loc.Filename + ".csv"
	hostFullPath := loc.HostPath + path
	dirPath := loc.HostPath + loc.Dir
	log.Println(dirPath)
	log.Println(hostFullPath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	src, _ := os.Open(loc.FilePath)
	defer src.Close()

	fi, _ := src.Stat()
	perm := fi.Mode()

	dst, _ := os.OpenFile(hostFullPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	defer dst.Close()

	io.Copy(dst, src)

	return loc.FileUrl(path)
}

func (loc *Loc) FileUrl(path string) string {
	return loc.Host + "/" + path
}
