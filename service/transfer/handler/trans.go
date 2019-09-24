package handler

import (
	"context"
	"encoding/json"
	"file-store-server/common"
	"file-store-server/config"
	"file-store-server/mq"
	"file-store-server/service/transfer/proto"
	"fmt"
	"log"
)

// Trans TransferServiceHandler 实现接口
type Trans struct{}

// Transfer 处理用户文件传输请求
func (t *Trans) Transfer(c context.Context, req *proto.ReqTrans, resp *proto.RespTrans) error {
	data := mq.TransferData{
		FileSha1:      req.FileSha1,
		CurLocation:   req.CurLocation,
		DestLocation:  req.DestLocation,
		DestStoreType: common.StoreType(req.DestStoreType),
	}
	pubData, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		resp.Code = -1
		resp.Message = "JOSN Marshal 失败"
		return err
	}

	ok := mq.Publish(config.TransExchangeName, config.TransOSSRoutingKey, pubData)
	log.Println("Producer publish", string(pubData), "to", config.TransExchangeName)
	if !ok {
		// TODO: 加入重试发送逻辑
		msg := "文件写入OSS失败"
		log.Println(msg)
		resp.Code = -2
		resp.Message = msg
		return fmt.Errorf(msg)
	}
	resp.Code = 0
	resp.Message = "文件传输请求已发布"
	return nil
}
