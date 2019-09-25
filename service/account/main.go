package main

import (
	"file-store-server/service/account/handler"
	"file-store-server/service/account/proto"
	"log"

	"github.com/micro/go-micro"
)

func main() {
	// 创建一个服务
	service := micro.NewService(
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
