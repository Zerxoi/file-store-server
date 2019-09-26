echo "编译生成二进制文件"
GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o upload $GOPATH/src/file-store-server/service/upload

echo "创建镜像 fileserver/upload"
docker build -t fileserver/upload .

echo "启动容器 upload"
docker run -d --name=upload -p 28000:28000 -v /home/zerxoi/fileserver:/fileserver fileserver/upload