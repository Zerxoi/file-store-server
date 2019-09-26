package main

import (
	"file-store-server/service/dbproxy/proto"
	"file-store-server/service/dbproxy/rpc"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"172.17.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.service.dbproxy"),
	)
	service.Init()

	err := proto.RegisterDBProxyServiceHandler(service.Server(), new(rpc.DBProxy))
	if err != nil {
		log.Println(err)
	}
	if err = service.Run(); err != nil {
		log.Println(err)
	}
}
