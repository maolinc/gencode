package mtypemap

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMTypeMapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMTypeMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMTypeMapLogic {
	return &CreateMTypeMapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMTypeMapLogic) CreateMTypeMap(req *types.CreateMTypeMapReq) (resp *types.CreateMTypeMapResp, err error) {
	// todo: add your logic here and delete this line

	return
}
