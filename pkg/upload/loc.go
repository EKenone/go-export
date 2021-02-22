package upload

import (
	"go-export/internal/conf"
)

type Loc struct {
	Conf
	Folder string
}

func (loc *Loc) Upload() string {

	path := loc.Folder + "/" + loc.Filename + ".csv"

	return LocFileUrl(path)
}

func LocFileUrl(path string) string {
	return conf.Conf.Loc.Host + "/" + path
}
