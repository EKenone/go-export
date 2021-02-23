package main

import (
	"flag"
	"go-export/api/rpc/pb"
	"go-export/internal/conf"
	"go-export/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

const RpcMaxData = 1024 * 1024 * 1024 // 1GB

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Panic(err)
	}

	l, err := net.Listen("tcp", ":"+conf.Conf.Ept.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs := grpc.NewServer(grpc.MaxRecvMsgSize(RpcMaxData), grpc.MaxSendMsgSize(RpcMaxData))

	pb.RegisterExportServer(gs, &service.RpcService{})

	if err := gs.Serve(l); err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
