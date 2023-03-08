package mtypemap

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MTypeMapDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMTypeMapDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MTypeMapDetailLogic {
	return &MTypeMapDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MTypeMapDetailLogic) MTypeMapDetail(req *types.IdReq) (resp *types.DetailMTypeMapResp, err error) {
	// todo: add your logic here and delete this line

	return
}
