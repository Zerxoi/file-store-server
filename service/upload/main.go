package main

import (
	"file-store-server/service/upload/config"
	upProto "file-store-server/service/upload/proto"
	"file-store-server/service/upload/route"
	"file-store-server/service/upload/rpc"
	"log"

	"github.com/micro/go-micro"
)

func startRPCService() {
	service := micro.NewService(
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
