package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"mall/service/pay/rpc/types/pay"

	"mall/service/pay/api/internal/svc"
	"mall/service/pay/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CallbackLogic) Callback(req *types.CallbackRequest) (resp *types.CallbackResponse, err error) {
	var callbackRequest pay.CallbackRequest
	err = copier.CopyWithOption(&callbackRequest, req, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.PayRpc.Callback(l.ctx, &callbackRequest)
	if err != nil {
		return nil, err
	}

	return &types.CallbackResponse{}, nil
}
