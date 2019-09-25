package main

import (
	"file-store-server/service/dbproxy/proto"
	"file-store-server/service/dbproxy/rpc"
	"log"

	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
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
