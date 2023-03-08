package mmasterversion

import (
	"net/http"

	"github.com/maolinc/gencode/app/api/internal/logic/mmasterversion"
	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteMMasterVersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := mmasterversion.NewDeleteMMasterVersionLogic(r.Context(), svcCtx)
		resp, err := l.DeleteMMasterVersion(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
