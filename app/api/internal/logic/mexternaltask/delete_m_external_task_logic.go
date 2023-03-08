package mexternaltask

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMExternalTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMExternalTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMExternalTaskLogic {
	return &DeleteMExternalTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMExternalTaskLogic) DeleteMExternalTask(req *types.IdsReq) (resp *types.DeleteMExternalTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}
