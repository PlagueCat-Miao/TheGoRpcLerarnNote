package client


import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"
	"testing"
	"github.com/PlagueCat-Miao/TheGoRpcLerarnNote/thrift_echo/gen-go/echo"

)

func TestClient(t *testing.T) {
	t.Log("Hellcat: thift-client 请求测试")
	//需要与服务端一致
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9898"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

    // plan B 0.11.0
	//iProtocol := protocolFactory.GetProtocol(transport)
	//oProtocol := protocolFactory.GetProtocol(transport)
	//tClient := thrift.NewTStandardClient(iProtocol, oProtocol)
	//client := echo.NewEchoClient(tClient)
	//plan A 0.9.2
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	useTransport, err := transportFactory.GetTransport(transport)
	client := echo.NewEchoClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9898", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	req := &echo.EchoReq{Msg:"You are welcome."}
	res, err := client.Echo(context.Background(),req)
	if err != nil {
		log.Println("Echo failed:", err)
		return
	}

	log.Println("response:", res.Msg)
	fmt.Println("well done")
	t.Log("Hellcat: thift-client 请求通过")
}