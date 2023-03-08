package mversion

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MVersionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMVersionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MVersionDetailLogic {
	return &MVersionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MVersionDetailLogic) MVersionDetail(req *types.IdReq) (resp *types.DetailMVersionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
