package svc

import (
	"bookstore/api/internal/config"
	"bookstore/api/model"
	"bookstore/rpc/add/adder"
	"bookstore/rpc/check/checker"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Adder     adder.Adder     // 手动代码
	Checker   checker.Checker // 手动代码
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:    c,
		Adder:     adder.NewAdder(zrpc.MustNewClient(c.Add)),       // 手动代码
		Checker:   checker.NewChecker(zrpc.MustNewClient(c.Check)), // 手动代码
		UserModel: model.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
