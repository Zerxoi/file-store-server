echo "编译生成二进制文件"
GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o transfer $GOPATH/src/file-store-server/service/transfer

echo "创建镜像 fileserver/transfer"
docker build -t fileserver/transfer .

echo "启动容器 transfer"
docker run -d --name=transfer -v /home/zerxoi/fileserver:/fileserver fileserver/transfer