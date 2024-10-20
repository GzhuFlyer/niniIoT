#go-zero知识点回顾

1，运行

```
1, 执行 etcd命令 (etcd推荐版本使用  etcd Version: 3.4.9)
2, 进入niniIoT/data/mysql,执行  docker-compose up 命令
3, 进入niniIoT/data/redis,执行  docker-compose up 命令 
4, 进入niniIoT/backend/shorturl/rpc/transform, 执行 go run transform.go -f etc/transform.yaml
5, 进入niniIoT/backend/shorturl/api, 执行 go run shorturl.go -f etc/shorturl-api.yaml 
6, 运行 curl -i "http://localhost:8888/shorten?url=https://go-zero.dev", 查看效果，curl -i "http://localhost:8888/expand?shorten=b0434f"

```

上面是使用 go-zero做的一个 短链服务，用于实习回顾 go-zero的一些知识点。