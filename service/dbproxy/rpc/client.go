package rpc

import (
	"file-store-server/service/dbproxy/proto"
	"log"

	"github.com/micro/go-micro"
)

var dbCli proto.DBProxyService

func init() {
	service := micro.NewService()
	service.Init()
	dbCli = proto.NewDBProxyService("go.micro.service.dbproxy", service.Client())
}

// Client 获取DBProxy RPC调用的客户端
func Client() proto.DBProxyService {
	if dbCli != nil {
		return dbCli
	}
	log.Println("DBProxy 未初始化")
	return nil
}
