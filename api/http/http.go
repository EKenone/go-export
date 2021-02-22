package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-export/internal/conf"
	"go-export/internal/export"
	"log"
)

func Ept(ctx *gin.Context) {
	var f export.Form

	if err := ctx.ShouldBindJSON(&f); err != nil {
		log.Println(err)
		ctx.JSON(200, gin.H{"code": 401, "msg": "错误数据格式"})
		return
	}

	if f.Total == 0 {
		f.Total = len(f.Data)
	}

	ec := export.InitExportConf(f)
	for _, v := range f.GetExportList() {
		ec.WriteRow(v)
	}
	ctx.JSON(200, gin.H{"code": 200, "msg": "ok"})
}

type ProgressResp struct {
	Progress string
	Url      string
	Status   int
}

func EptProgress(ctx *gin.Context) {
	mark, _ := ctx.GetQuery("mark")

	data := export.CurrentProgress(mark)

	progress := fmt.Sprintf("%.2f", float64(data.Current)/float64(data.Total))

	if progress == "1.00" && data.Status == export.StatusWait {
		progress = "99%"
	} else {
		progress = fmt.Sprintf("%v", float64(data.Current*100)/float64(data.Total)) + "%"
	}

	ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": ProgressResp{
		Progress: progress,
		Url:      data.Url,
		Status:   data.Status,
	}})
}

var env string

func init() {
	flag.StringVar(&env, "env", "debug", "default gin mode")
}

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Panic(err)
	}

	gin.SetMode(env)

	eg := gin.Default()
	eg.POST("/ept", Ept)
	eg.GET("/ept-progress", EptProgress)
	eg.Run(":" + conf.Conf.Ept.Port)
}
