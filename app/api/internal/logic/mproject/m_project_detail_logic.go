package mproject

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MProjectDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMProjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MProjectDetailLogic {
	return &MProjectDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MProjectDetailLogic) MProjectDetail(req *types.IdReq) (resp *types.DetailMProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
