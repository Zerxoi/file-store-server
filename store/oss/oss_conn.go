package oss

import (
	"file-store-server/config"
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var ossCli *oss.Client

// Client 创建client
func Client() *oss.Client {
	if ossCli != nil {
		return ossCli
	}
	ossCli, err := oss.New(config.OSSEndpoint, config.OSSAccessKeyID, config.OSSAccessKeySecret)
	if err != nil {
		log.Println("OSSClient:", err)
		return nil
	}
	return ossCli
}

// Bucket 获取bucket存储空间
func Bucket() *oss.Bucket {
	cli := Client()
	if cli != nil {
		bucket, err := cli.Bucket(config.OSSBucket)
		if err != nil {
			log.Println("Client:", err)
		}
		return bucket
	}
	return nil
}

// DownloadURL 临时授权下载URL
func DownloadURL(objName string) string {
	signedURL, err := Bucket().SignURL(objName, oss.HTTPGet, 3600)
	if err != nil {
		log.Println("DownloadURL:", err)
		return ""
	}
	return signedURL
}
