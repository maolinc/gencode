package mmasterversion

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MMasterVersionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMMasterVersionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MMasterVersionDetailLogic {
	return &MMasterVersionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MMasterVersionDetailLogic) MMasterVersionDetail(req *types.IdReq) (resp *types.DetailMMasterVersionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
