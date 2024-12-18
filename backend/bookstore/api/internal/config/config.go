package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Add     zrpc.RpcClientConf // 手动代码
	Check   zrpc.RpcClientConf // 手动代码
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
}
