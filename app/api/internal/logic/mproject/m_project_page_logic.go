package mproject

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MProjectPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMProjectPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MProjectPageLogic {
	return &MProjectPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MProjectPageLogic) MProjectPage(req *types.SearchMProjectReq) (resp *types.SearchMProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
