package mexternaltask

import (
	"net/http"

	"github.com/maolinc/gencode/app/api/internal/logic/mexternaltask"
	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteMExternalTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := mexternaltask.NewDeleteMExternalTaskLogic(r.Context(), svcCtx)
		resp, err := l.DeleteMExternalTask(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
