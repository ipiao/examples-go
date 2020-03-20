package main

import (
	"context"
	"fmt"
	"log"

	greeter "github.com/ipiao/examples-go/example_gomicro/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	logger.Init(logger.WithLevel(logger.TraceLevel))
	srv := micro.NewService(
		micro.Name("example.gomicro.greeter"),
		micro.Metadata(map[string]string{"NodeId": "1"}),
		micro.Registry(consul.NewRegistry()))

	srv.Init()

	greeter.RegisterSayHandler(srv.Server(), new(Greeter))
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

type Greeter struct {
}

func (*Greeter) Hello(ctx context.Context, req *greeter.Request, resp *greeter.Response) error {
	resp.Msg = fmt.Sprint("hello, %s!", req.Name)
	return nil
}
