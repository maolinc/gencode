package mmasterversion

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MMasterVersionPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMMasterVersionPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MMasterVersionPageLogic {
	return &MMasterVersionPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MMasterVersionPageLogic) MMasterVersionPage(req *types.SearchMMasterVersionReq) (resp *types.SearchMMasterVersionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
