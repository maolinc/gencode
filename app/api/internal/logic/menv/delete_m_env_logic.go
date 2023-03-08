package menv

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMEnvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMEnvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMEnvLogic {
	return &DeleteMEnvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMEnvLogic) DeleteMEnv(req *types.IdsReq) (resp *types.DeleteMEnvResp, err error) {
	// todo: add your logic here and delete this line

	return
}
