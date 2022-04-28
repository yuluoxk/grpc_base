package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_base/protos"
	"os"
)

func main() {
	// http2 安全
	cli, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	c := protos.NewEchoServiceClient(cli)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		req := protos.EchoRequest{
			Req: string(line),
		}
		res, err := c.GetUnaryEcho(context.Background(), &req)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
}
