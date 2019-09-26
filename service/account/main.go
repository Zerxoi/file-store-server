package main

import (
	"file-store-server/service/account/handler"
	"file-store-server/service/account/proto"
	"log"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"

	"github.com/micro/go-micro"
)

func main() {
	// 创建一个服务
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"172.17.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.service.user"),
	)
	service.Init()

	err := proto.RegisterUserServiceHandler(service.Server(), new(handler.User))
	if err != nil {
		log.Println(err)
	}

	if err = service.Run(); err != nil {
		log.Println(err)
	}
}
