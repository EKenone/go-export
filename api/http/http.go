package main

import (
	"flag"
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
	eg.Run(":" + conf.Conf.Ept.Port)
}
