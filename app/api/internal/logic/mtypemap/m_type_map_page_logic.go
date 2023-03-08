package mtypemap

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MTypeMapPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMTypeMapPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MTypeMapPageLogic {
	return &MTypeMapPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MTypeMapPageLogic) MTypeMapPage(req *types.SearchMTypeMapReq) (resp *types.SearchMTypeMapResp, err error) {
	// todo: add your logic here and delete this line

	return
}
