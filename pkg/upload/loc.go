package upload

import (
	"go-export/internal/conf"
)

type Loc struct {
	Conf
}

func (loc *Loc) Upload() string {

	path := conf.Conf.Loc.Dir + loc.Filename + ".csv"

	return LocFileUrl(path)
}

func LocFileUrl(path string) string {
	return conf.Conf.Loc.Host + "/" + path
}
