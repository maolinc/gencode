package mtypemap

import (
	"net/http"

	"github.com/maolinc/gencode/app/api/internal/logic/mtypemap"
	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteMTypeMapHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := mtypemap.NewDeleteMTypeMapLogic(r.Context(), svcCtx)
		resp, err := l.DeleteMTypeMap(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
