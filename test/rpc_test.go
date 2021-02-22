package test

import (
	"context"
	"encoding/json"
	"go-export/api/rpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func TestRpcEpt(t *testing.T) {
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		t.Log(err)
	}
	s, _ := status.FromError(err)
	c := pb.NewExportClient(conn)

	//a, _ := c.EptProgress(context.Background(), &pb.EptProgressRequest{
	//	Mark: "20210222175038XVlBz",
	//})
	//
	//t.Log(a)
	//return

	total := 30
	mark := time.Now().Format("20060102150405") + RandStringBytes(5)
	header := []map[string]string{
		{"field": "number", "title": "编号"},
		{"field": "name", "title": "姓名"},
		{"field": "other", "title": "特长"},
		{"field": "age", "title": "年龄"},
	}
	headerStr, _ := json.Marshal(header)
	list := make([]map[string]interface{}, 0)

	for i := 0; i < total; i++ {
		list = append(list, map[string]interface{}{
			"number": "test", "name": "没问题", "other": "无形装逼", "age": 28,
		})
	}
	dataStr, _ := json.Marshal(list)

	var res *pb.EptReply

	res, err = c.Ept(context.Background(), &pb.EptRequest{
		HashMark: mark,
		Total:    int32(total),
		Header:   string(headerStr),
		Data:     string(dataStr),
	})

	t.Log(s)
	if err != nil {
		t.Log(err)
	}

	t.Log(res)
}
