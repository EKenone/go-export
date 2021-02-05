package main

import (
	"context"
	"encoding/json"
	"flag"
	"go-export/api/rpc/pb"
	"go-export/internal/conf"
	"go-export/internal/export"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
}

func (s *Server) Ept(ctx context.Context, req *pb.EptRequest) (*pb.EptReply, error) {

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

	if f.Total == 0 {
		f.Total = len(f.Data)
	}

	ec := export.InitExportConf(f)
	for _, v := range f.GetExportList() {
		ec.WriteRow(v)
	}

	return &pb.EptReply{
		Code: 200,
		Msg:  "ok",
	}, nil
}

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Panic(err)
	}

	l, err := net.Listen("tcp", ":"+conf.Conf.Ept.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs := grpc.NewServer()

	pb.RegisterExportServer(gs, &Server{})

	if err := gs.Serve(l); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
