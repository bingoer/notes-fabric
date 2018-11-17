fabric环境



1 源码下载
--Go get
go get -u github.com/hyperledger/fabric
cd fabric
git checkout release-1.2

--Git clone
git clone -b release-1.2 https://github.com/hyperledger/fabric.git
cd fabric
git checkout release-1.2
cp -R fabric/ /home/gopath/src/github.com/hyperledger/




2 下载依赖
--依赖
go get -u github.com/golang/protobuf/protoc-gen-go
mkdir -p $GOPATH/src/github.com/hyperledger/fabric/build/docker/gotools/bin
cp  $GOPATH/bin/protoc-gen-go $GOPATH/src/github.com/hyperledger/fabric/build/docker/gotools/bin

--tools
mkdir -p /home/gopath/src/github.com/hyperledger/fabric/release/linux-amd64/bin/
cd fabric-sample/bin
cp * /home/gopath/src/github.com/hyperledger/fabric/release/linux-amd64/bin/ 
chmod u+x /home/gopath/src/github.com/hyperledger/fabric/release/linux-amd64/bin/*




3 docker 下载与备份
--镜像下载
/home/gopath/src/github.com/hyperledger/fabric/examples/e2e_cli
chmod -R u+x /home/gopath/src/github.com/hyperledger/fabric/release/linux-amd64/bin/*
./download-dockerimages.sh

--单独下载
docker pull hyperledger/fabric-orderer:1.2.0
docker pull hyperledger/fabric-peer:1.2.0
docker pull hyperledger/fabric-ca:1.2.0
docker pull hyperledger/fabric-tools:1.2.0
docker pull hyperledger/fabric-ccenv:1.2.0
docker pull hyperledger/fabric-zookeeper:0.4.10
docker pull hyperledger/fabric-kafka:0.4.10
docker pull hyperledger/fabric-couchdb:0.4.10
docker pull hyperledger/fabric-baseimage:amd64-0.4.10
docker pull hyperledger/fabric-baseos:amd64-0.4.10

--导出镜像
docker save  45b65950cbb5 > /data/fabric1.2/fabric-orderer.tar
docker save  0d75e744dd41 > /data/fabric1.2/fabric-peer.tar 
docker save  4baf7789a8ec > /data/fabric1.2/fabric-orderer-2.tar
docker save  82c262e65984 > /data/fabric1.2/fabric-peer-2.tar 
docker save  66cc132bd09c > /data/fabric1.2/fabric-ca.tar 
docker save  379602873003 > /data/fabric1.2/fabric-tools.tar
docker save  6acf31e2d9a4 > /data/fabric1.2/fabric-ccenv.tar
docker save  2b51158f3898 > /data/fabric1.2/fabric-zookeeper.tar
docker save  936aef6db0e6 > /data/fabric1.2/fabric-kafka.tar
docker save  3092eca241fc > /data/fabric1.2/fabric-couchdb.tar
docker save  62513965e238 > /data/fabric1.2/fabric-baseimage.tar
docker save  52190e831002 > /data/fabric1.2/fabric-baseos.tar

--导入镜像
docker load < /data/fabric1.2/fabric-orderer-2.tar
docker load < /data/fabric1.2/fabric-peer-2.tar 
docker load < /data/fabric1.2/fabric-orderer.tar
docker load < /data/fabric1.2/fabric-peer.tar 
docker load < /data/fabric1.2/fabric-ca.tar 
docker load < /data/fabric1.2/fabric-tools.tar
docker load < /data/fabric1.2/fabric-ccenv.tar
docker load < /data/fabric1.2/fabric-zookeeper.tar
docker load < /data/fabric1.2/fabric-kafka.tar
docker load < /data/fabric1.2/fabric-couchdb.tar
docker load < /data/fabric1.2/fabric-baseimage.tar
docker load < /data/fabric1.2/fabric-baseos.tar

--重命名
docker tag  45b65950cbb5 hyperledger/fabric-orderer:amd64-1.2.1-snapshot-78a3a8d
docker tag  0d75e744dd41 hyperledger/fabric-peer:amd64-1.2.1-snapshot-78a3a8d
docker tag  4baf7789a8ec hyperledger/fabric-orderer:1.2.0
docker tag  82c262e65984 hyperledger/fabric-peer:1.2.0 
docker tag  66cc132bd09c hyperledger/fabric-ca:1.2.0
docker tag  379602873003 hyperledger/fabric-tools:1.2.0
docker tag  6acf31e2d9a4 hyperledger/fabric-ccenv:1.2.0
docker tag  2b51158f3898 hyperledger/fabric-zookeeper:0.4.10
docker tag  936aef6db0e6 hyperledger/fabric-kafka:0.4.10
docker tag  3092eca241fc hyperledger/fabric-couchdb:0.4.10 
docker tag  62513965e238 hyperledger/fabric-baseimage:amd64-0.4.10 
docker tag  52190e831002 hyperledger/fabric-baseos:amd64-0.4.10 


docker tag  45b65950cbb5 hyperledger/fabric-orderer:amd64-latest
docker tag  0d75e744dd41 hyperledger/fabric-peer:amd64-latest
docker tag  4baf7789a8ec hyperledger/fabric-orderer:latest
docker tag  82c262e65984 hyperledger/fabric-peer:latest
docker tag  66cc132bd09c hyperledger/fabric-ca:latest
docker tag  379602873003 hyperledger/fabric-tools:latest
docker tag  6acf31e2d9a4 hyperledger/fabric-ccenv:latest
docker tag  2b51158f3898 hyperledger/fabric-zookeeper:latest
docker tag  936aef6db0e6 hyperledger/fabric-kafka:latest
docker tag  3092eca241fc hyperledger/fabric-couchdb:latest
docker tag  62513965e238 hyperledger/fabric-baseimage:latest
docker tag  52190e831002 hyperledger/fabric-baseos:latest




4 运行实例
cd /home/gopath/src/github.com/hyperledger/fabric/examples/e2e_cli
bash network_setup.sh down
bash network_setup.sh up






--问题
generateArtifacts.sh:行125: /home/gopath/src/github.com/hyperledger/fabric/release/linux-amd64/bin/configtxgen: 没有那个文件或目录

--解决
从fabric-sample/bin 拷贝到 /home/gopath/src/github.com/hyperledger/fabric/release/linux-amd64/bin/ 
> chmod u+x *