package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"mall/service/order/model"
	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *order.ListRequest) (*order.ListResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	// 查询订单
	//list, err := l.svcCtx.OrderModel.FindAllByUid(l.ctx, in.Uid)
	//if err != nil {
	//	if err == model.ErrNotFound {
	//		return nil, status.Error(100, "订单不存在")
	//	}
	//	return nil, status.Error(500, err.Error())
	//}
	var list []*model.Order
	err = l.svcCtx.Gorm.Table("order").Where("uid = ?", in.Uid).Find(&list).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(100, "订单不存在")
		}
	}

	orderList := make([]*order.DetailResponse, 0)
	for _, item := range list {
		var detailResp order.DetailResponse
		err := copier.CopyWithOption(&detailResp, item, copier.Option{
			IgnoreEmpty: true,
		})
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		orderList = append(orderList, &detailResp)
	}

	return &order.ListResponse{
		Data: orderList,
	}, nil
}
