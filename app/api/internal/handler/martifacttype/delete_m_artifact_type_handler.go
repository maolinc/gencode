package martifacttype

import (
	"net/http"

	"github.com/maolinc/gencode/app/api/internal/logic/martifacttype"
	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteMArtifactTypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := martifacttype.NewDeleteMArtifactTypeLogic(r.Context(), svcCtx)
		resp, err := l.DeleteMArtifactType(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
