package mversion

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMVersionLogic {
	return &CreateMVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMVersionLogic) CreateMVersion(req *types.CreateMVersionReq) (resp *types.CreateMVersionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
