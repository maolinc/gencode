package mversion

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMVersionLogic {
	return &DeleteMVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMVersionLogic) DeleteMVersion(req *types.IdsReq) (resp *types.DeleteMVersionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
