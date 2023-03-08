package mtypemap

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMTypeMapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMTypeMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMTypeMapLogic {
	return &UpdateMTypeMapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMTypeMapLogic) UpdateMTypeMap(req *types.UpdateMTypeMapReq) (resp *types.UpdateMTypeMapResp, err error) {
	// todo: add your logic here and delete this line

	return
}
