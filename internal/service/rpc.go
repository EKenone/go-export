package service

import (
	"context"
	"encoding/json"
	"go-export/api/rpc/pb"
	"go-export/internal/export"
)

type RpcService struct {
}

func (s *RpcService) Ept(ctx context.Context, req *pb.EptRequest) (*pb.EptReply, error) {
	var (
		header []export.HeaderField
		data   []map[string]interface{}
	)

	_ = json.Unmarshal([]byte(req.Header), &header)
	_ = json.Unmarshal([]byte(req.Data), &data)

	f := export.Form{
		HashMark: req.HashMark,
		Header:   header,
		Data:     data,
		Total:    int(req.Total),
	}

	// 这里协程调用就行，直接返回，查看进度条可以看情况
	//go doEpt(f)
	doEpt(f)

	return &pb.EptReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}

func (s *RpcService) EptProgress(ctx context.Context, req *pb.EptProgressRequest) (*pb.EptProgressReply, error) {
	progress, url, status := getProgress(req.Mark)

	return &pb.EptProgressReply{
		Progress: progress,
		Url:      url,
		Status:   int32(status),
	}, nil
}
