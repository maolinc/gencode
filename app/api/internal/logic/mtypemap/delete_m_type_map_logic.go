package mtypemap

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMTypeMapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMTypeMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMTypeMapLogic {
	return &DeleteMTypeMapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMTypeMapLogic) DeleteMTypeMap(req *types.IdsReq) (resp *types.DeleteMTypeMapResp, err error) {
	// todo: add your logic here and delete this line

	return
}
