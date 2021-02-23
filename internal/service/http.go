package service

import (
	"github.com/gin-gonic/gin"
	"go-export/internal/export"
)

type HttpService struct {
}

func (s *HttpService) Ept(ctx *gin.Context) {
	var f export.Form

	if err := ctx.ShouldBindJSON(&f); err != nil {
		ctx.JSON(200, gin.H{"code": 401, "msg": "错误数据格式"})
		return
	}

	doEpt(f)

	ctx.JSON(200, gin.H{"code": 200, "msg": "ok"})
}

type ProgressResp struct {
	Progress string `json:"progress"`
	Url      string `json:"url"`
	Status   int    `json:"status"`
}

func (s *HttpService) EptProgress(ctx *gin.Context) {
	mark, _ := ctx.GetQuery("mark")

	progress, url, status := getProgress(mark)

	ctx.JSON(200, gin.H{"code": 200, "msg": "ok", "data": ProgressResp{
		Progress: progress,
		Url:      url,
		Status:   status,
	}})
}
