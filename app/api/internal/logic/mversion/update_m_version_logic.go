package mversion

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMVersionLogic {
	return &UpdateMVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMVersionLogic) UpdateMVersion(req *types.UpdateMVersionReq) (resp *types.UpdateMVersionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
