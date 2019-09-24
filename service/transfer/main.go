package main

import (
	"encoding/json"
	"file-store-server/config"
	"file-store-server/db"
	"file-store-server/mq"
	"file-store-server/service/transfer/handler"
	"file-store-server/service/transfer/proto"
	"file-store-server/store/oss"
	"log"
	"os"

	"github.com/micro/go-micro"
)

// ProcesstTransfer 处理文件的真正逻辑
func ProcesstTransfer(msg []byte) bool {
	// 1. 解析msg
	pubData := mq.TransferData{}
	err := json.Unmarshal(msg, &pubData)
	if err != nil {
		log.Println("ProcesstTransfer:", err)
		return false
	}

	// 2. 根据临时存储文件路径,创建文件句柄
	f, err := os.Open(pubData.CurLocation)
	if err != nil {
		log.Println("ProcessTransfer:", err)
		return false
	}
	defer f.Close()

	// 3. 通过文件句柄将文件内容读出来并且上传到OSS
	oss.Bucket().PutObject(pubData.DestLocation, f)
	if err != nil {
		log.Println("ProcessTransfer:", err)
		return false
	}

	// 4. 更新文件存储路径到文件表
	ok := db.UpdateFileLocation(pubData.FileSha1, pubData.DestLocation)
	if !ok {
		log.Println("ProcesstTransfer: Failed to update file's location.")
		return false
	}
	err = os.Remove(pubData.CurLocation)
	if err != nil {
		log.Println("云端上传完成,删除本地文件失败!")
		return false
	}
	return true
}

func main() {
	go func() {
		log.Println("开始监听转移任务队列")
		mq.StartConsume(config.TransOSSQueueName, "Transfer_OSS", ProcesstTransfer)
	}()

	// 创建一个服务
	service := micro.NewService(
		micro.Name("go.micro.service.transfer"),
	)
	service.Init()

	proto.RegisterTransferServiceHandler(service.Server(), new(handler.Trans))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
