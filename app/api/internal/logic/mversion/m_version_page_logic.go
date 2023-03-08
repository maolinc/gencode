package mversion

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MVersionPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMVersionPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MVersionPageLogic {
	return &MVersionPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MVersionPageLogic) MVersionPage(req *types.SearchMVersionReq) (resp *types.SearchMVersionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
