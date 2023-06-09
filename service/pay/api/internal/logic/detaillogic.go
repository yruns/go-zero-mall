package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"mall/service/pay/rpc/types/pay"

	"mall/service/pay/api/internal/svc"
	"mall/service/pay/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	detailResponse, err := l.svcCtx.PayRpc.Detail(l.ctx, &pay.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	var result types.DetailResponse

	err = copier.Copy(&result, detailResponse)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &result, nil
}
