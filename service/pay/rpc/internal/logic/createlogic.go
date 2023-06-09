package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"mall/service/order/rpc/types/order"
	"mall/service/pay/model"
	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"
	"mall/service/user/rpc/types/user"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pay.CreateRequest) (*pay.CreateResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	// 查询订单是否存在
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, err
	}

	// 查询订单已经创建支付
	var newPay model.Pay
	err = l.svcCtx.Gorm.Table("pay").Where("oid = ?", in.Oid).Find(&newPay).Error
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	if newPay.Id != 0 {
		return nil, status.Error(100, "订单已支付")
	}

	err = copier.CopyWithOption(&newPay, in, copier.Option{
		IgnoreEmpty: true,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	newPay.Status = 0
	newPay.Source = 0

	// 记录订单
	err = l.svcCtx.Gorm.Table("pay").Create(&newPay).Error
	if err != nil {
		return nil, status.Error(500, "支付失败")
	}

	return &pay.CreateResponse{
		Id: newPay.Id,
	}, nil
}
