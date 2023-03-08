package mproject

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMProjectLogic {
	return &UpdateMProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMProjectLogic) UpdateMProject(req *types.UpdateMProjectReq) (resp *types.UpdateMProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
