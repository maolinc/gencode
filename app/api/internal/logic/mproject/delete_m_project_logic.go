package mproject

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMProjectLogic {
	return &DeleteMProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMProjectLogic) DeleteMProject(req *types.IdsReq) (resp *types.DeleteMProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
