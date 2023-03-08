package mmasterversion

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMMasterVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMMasterVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMMasterVersionLogic {
	return &UpdateMMasterVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMMasterVersionLogic) UpdateMMasterVersion(req *types.UpdateMMasterVersionReq) (resp *types.UpdateMMasterVersionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
