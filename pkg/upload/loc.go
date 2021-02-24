package upload

type Loc struct {
	Conf
	Host string
}

func (loc *Loc) Upload() string {
	path := loc.Dir + loc.Filename + ".csv"

	return loc.FileUrl(path)
}

func (loc *Loc) FileUrl(path string) string {
	return loc.Host + "/" + path
}
