package mq

import (
	"file-store-server/common"
)

// TransferData 转移队列中消息载体的结构格式
type TransferData struct {
	FileSha1      string
	CurLocation   string
	DestLocation  string
	DestStoreType common.StoreType
}
