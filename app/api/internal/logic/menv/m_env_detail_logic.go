package menv

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MEnvDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMEnvDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MEnvDetailLogic {
	return &MEnvDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MEnvDetailLogic) MEnvDetail(req *types.IdReq) (resp *types.DetailMEnvResp, err error) {
	// todo: add your logic here and delete this line

	return
}
