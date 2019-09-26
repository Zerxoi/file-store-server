package main

import (
	"file-store-server/service/upload/config"
	upProto "file-store-server/service/upload/proto"
	"file-store-server/service/upload/route"
	"file-store-server/service/upload/rpc"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

func startRPCService() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"172.17.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.service.upload"),
	)
	service.Init()
	upProto.RegisterUploadServiceHandler(service.Server(), new(rpc.Upload))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

// 提供外部用户上传文件功能
func startAPIService() {
	router := route.Router()
	router.Run(config.UploadServiceHost)
}

func main() {
	// rpc 服务
	go startRPCService()

	// api 服务
	startAPIService()
}
