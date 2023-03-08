package menv

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMEnvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMEnvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMEnvLogic {
	return &UpdateMEnvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMEnvLogic) UpdateMEnv(req *types.UpdateMEnvReq) (resp *types.UpdateMEnvResp, err error) {
	// todo: add your logic here and delete this line

	return
}
