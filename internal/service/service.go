package service

import (
	"fmt"
	"go-export/internal/export"
)

func doEpt(f export.Form) {
	if f.Total == 0 {
		f.Total = len(f.Data)
	}

	ec := export.InitExportConf(f)
	for _, v := range f.GetExportList() {
		ec.WriteRow(v)
	}
}

// 获取进度条统一服务
func getProgress(mark string) (string, string, int) {
	data := export.CurrentProgress(mark)

	progress := fmt.Sprintf("%.2f", float64(data.Current)/float64(data.Total))

	if progress == "1.00" && data.Status == export.StatusWait {
		progress = "99%"
	} else {
		progress = fmt.Sprintf("%v", float64(data.Current*100)/float64(data.Total)) + "%"
	}

	return progress, data.Url, data.Status
}
