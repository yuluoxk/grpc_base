package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_base/protos"
	"net"
)

//type EchoServiceServer interface {
//	GetUnaryEcho(context.Context, *EchoRequest) (*EchoResponse, error)
//	mustEmbedUnimplementedEchoServiceServer()
//}

type echoService struct {
	protos.UnimplementedEchoServiceServer
}

func (es *echoService) GetUnaryEcho(ctx context.Context, req *protos.EchoRequest) (*protos.EchoResponse, error) {
	res := "received: " + req.GetReq()
	fmt.Println(res)
	return &protos.EchoResponse{Res: res}, nil
}

func main() {
	rpcs := grpc.NewServer()
	protos.RegisterEchoServiceServer(rpcs, &echoService{})
	// 实现网络功能
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)
}
