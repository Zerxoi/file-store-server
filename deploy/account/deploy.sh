echo "编译生成二进制文件"
GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o account $GOPATH/src/file-store-server/service/account

echo "创建镜像 account"
docker build -t account .

echo "启动容器 account"
docker run -d --name=account account