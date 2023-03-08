package mproject

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMProjectLogic {
	return &CreateMProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMProjectLogic) CreateMProject(req *types.CreateMProjectReq) (resp *types.CreateMProjectResp, err error) {
	// todo: add your logic here and delete this line

	return
}
