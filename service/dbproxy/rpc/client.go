package rpc

import (
	"file-store-server/service/dbproxy/proto"
	"log"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"

	"github.com/micro/go-micro"
)

var dbCli proto.DBProxyService

func init() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"172.17.0.1:8500",
		}
	})

	service := micro.NewService(micro.Registry(reg))
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
