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

type CreateMArtifactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMArtifactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMArtifactLogic {
	return &CreateMArtifactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMArtifactLogic) CreateMArtifact(req *types.CreateMArtifactReq) (resp *types.CreateMArtifactResp, err error) {
	// todo: add your logic here and delete this line

	data := &model.MArtifact{}
	_ = copier.Copy(data, req)
	err = l.svcCtx.MArtifactModel.Insert(l.ctx, nil, data)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %v", req)
	}

	return &types.CreateMArtifactResp{}, nil
}
