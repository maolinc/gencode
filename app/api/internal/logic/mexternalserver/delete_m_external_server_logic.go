package mexternalserver

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMExternalServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMExternalServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMExternalServerLogic {
	return &DeleteMExternalServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMExternalServerLogic) DeleteMExternalServer(req *types.IdsReq) (resp *types.DeleteMExternalServerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
