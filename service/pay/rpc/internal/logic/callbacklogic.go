package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"mall/service/order/rpc/types/order"
	"mall/service/pay/model"
	"mall/service/user/rpc/types/user"

	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
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

	// 查询支付是否存在
	var newPay model.Pay
	err = l.svcCtx.Gorm.Table("pay").Where("id = ?", in.Id).First(&newPay).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(100, "支付不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	// 支付金额和订单金额是否一致
	if in.Amount != newPay.Amount {
		return nil, status.Error(100, "支付金额与订单金额不符")
	}
	newPay.Source = in.Source
	newPay.Status = in.Status
	fmt.Println(newPay)
	err = l.svcCtx.Gorm.Table("pay").Save(&newPay).Error

	// 更新订单支付状态
	_, err = l.svcCtx.OrderRpc.Paid(l.ctx, &order.PaidRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, err
	}

	return &pay.CallbackResponse{}, nil
}
