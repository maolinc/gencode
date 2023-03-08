package martifact

import (
	"context"
	"github.com/maolinc/gencode/app/api/internal/tools/copier"

	"github.com/maolinc/gencode/app/model"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MArtifactPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMArtifactPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MArtifactPageLogic {
	return &MArtifactPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MArtifactPageLogic) MArtifactPage(req *types.SearchMArtifactReq) (resp *types.SearchMArtifactResp, err error) {

	var cond model.MArtifactQuery

	_ = copier.CopierWithOptions(&cond, req)

	list, err := l.svcCtx.MArtifactModel.FindListByCursor(l.ctx, &cond)
	if err != nil {
		return nil, err
	}

	var res types.SearchMArtifactResp
	_ = copier.CopierWithOptions(&res.List, list)

	return &res, nil
}
