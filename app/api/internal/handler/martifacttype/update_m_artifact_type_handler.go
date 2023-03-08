package martifacttype

import (
	"net/http"

	"github.com/maolinc/gencode/app/api/internal/logic/martifacttype"
	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateMArtifactTypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateMArtifactTypeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := martifacttype.NewUpdateMArtifactTypeLogic(r.Context(), svcCtx)
		resp, err := l.UpdateMArtifactType(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
