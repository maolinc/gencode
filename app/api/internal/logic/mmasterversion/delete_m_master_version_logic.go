package mmasterversion

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMMasterVersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMMasterVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMMasterVersionLogic {
	return &DeleteMMasterVersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMMasterVersionLogic) DeleteMMasterVersion(req *types.IdsReq) (resp *types.DeleteMMasterVersionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
