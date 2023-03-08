package mexternaltask

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMExternalTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMExternalTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMExternalTaskLogic {
	return &UpdateMExternalTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMExternalTaskLogic) UpdateMExternalTask(req *types.UpdateMExternalTaskReq) (resp *types.UpdateMExternalTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}
