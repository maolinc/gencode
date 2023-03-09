package martifact

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/maolinc/gencode/app/model"
	"github.com/pkg/errors"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMArtifactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMArtifactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMArtifactLogic {
	return &UpdateMArtifactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMArtifactLogic) UpdateMArtifact(req *types.UpdateMArtifactReq) (resp *types.UpdateMArtifactResp, err error) {

	var data model.MArtifact
	_ = copier.Copy(&data, req)
	err = l.svcCtx.MArtifactModel.Update(l.ctx, nil, &data)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %v", req)
	}

	return &types.UpdateMArtifactResp{}, nil
}
