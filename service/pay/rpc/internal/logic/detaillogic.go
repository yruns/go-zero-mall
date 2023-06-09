package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"mall/service/pay/model"

	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *pay.DetailRequest) (*pay.DetailResponse, error) {
	var newPay model.Pay
	err := l.svcCtx.Gorm.Table("pay").Where("id = ?", in.Id).First(&newPay).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(100, "支付不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	var result pay.DetailResponse
	err = copier.Copy(&result, newPay)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
