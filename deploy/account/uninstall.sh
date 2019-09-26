echo "删除二进制文件"
rm account

echo "结束容器 account"
docker stop account

echo "删除容器 account"
docker rm account

echo "删除镜像 fileserver/account"
docker image rm fileserver/account