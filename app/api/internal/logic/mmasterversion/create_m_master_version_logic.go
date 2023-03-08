package mmasterversion

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMMasterVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMMasterVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMMasterVersionLogic {
	return &CreateMMasterVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMMasterVersionLogic) CreateMMasterVersion(req *types.CreateMMasterVersionReq) (resp *types.CreateMMasterVersionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
