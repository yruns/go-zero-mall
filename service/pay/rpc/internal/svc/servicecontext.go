package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"mall/common/database"
	"mall/service/order/rpc/orderclient"
	"mall/service/pay/rpc/internal/config"
	"mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	UserRpc  userclient.User
	OrderRpc orderclient.Order

	Gorm *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := database.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		Gorm:     db,
	}
}
