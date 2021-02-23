package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"go-export/internal/conf"
	"go-export/internal/service"
	"log"
)

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

	s := service.HttpService{}

	eg := gin.Default()
	eg.POST("/ept", s.Ept)
	eg.GET("/ept-progress", s.EptProgress)
	eg.Run(":" + conf.Conf.Ept.Port)
}
