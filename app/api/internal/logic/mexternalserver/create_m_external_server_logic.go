package mexternalserver

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMExternalServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMExternalServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMExternalServerLogic {
	return &CreateMExternalServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMExternalServerLogic) CreateMExternalServer(req *types.CreateMExternalServerReq) (resp *types.CreateMExternalServerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
