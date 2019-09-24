package rpc

import (
	"context"
	"file-store-server/service/upload/config"
	upProto "file-store-server/service/upload/proto"
)

// Upload 用于实现 UploadServiceHandler 接口
type Upload struct{}

// UploadEntry 用于获取上传的入口地址
func (up *Upload) UploadEntry(c context.Context, req *upProto.ReqEntry, resp *upProto.RespEntry) error {
	resp.Entry = config.UploadEntry
	return nil
}
