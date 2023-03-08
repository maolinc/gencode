package mtypemap

import (
	"net/http"

	"github.com/maolinc/gencode/app/api/internal/logic/mtypemap"
	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateMTypeMapHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMTypeMapReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := mtypemap.NewCreateMTypeMapLogic(r.Context(), svcCtx)
		resp, err := l.CreateMTypeMap(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
