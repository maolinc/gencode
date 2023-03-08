package mexternaltask

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMExternalTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMExternalTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMExternalTaskLogic {
	return &CreateMExternalTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMExternalTaskLogic) CreateMExternalTask(req *types.CreateMExternalTaskReq) (resp *types.CreateMExternalTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}
