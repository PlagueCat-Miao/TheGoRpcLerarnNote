package main

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"

	"github.com/PlagueCat-Miao/TheGoRpcLerarnNote/thrift_echo/gen-go/echo"

)


type EchoServer struct {
}

func (e *EchoServer) Echo(ctx context.Context,req *echo.EchoReq) (*echo.EchoRes, error) {
	fmt.Printf("message from client: %v\n", req.GetMsg())

	resp := &echo.EchoRes{
		Msg: "success",
	}

	return resp, nil
}

func main() {
	transport, err := thrift.NewTServerSocket(":9898")
	if err != nil {
		panic(err)
	}

	handler := &EchoServer{}
	processor := echo.NewEchoProcessor(handler)
	//需要与客户端一致
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()//需要与客户端一致
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	if err := server.Serve(); err != nil {
		panic(err)
	}
}