echo "删除二进制文件"
rm transfer

echo "结束容器 transfer"
docker stop transfer

echo "删除容器 transfer"
docker rm transfer

echo "删除镜像 transfer"
docker image rm transfer