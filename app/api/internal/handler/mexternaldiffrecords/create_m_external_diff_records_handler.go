package mexternaldiffrecords

import (
	"net/http"

	"github.com/maolinc/gencode/app/api/internal/logic/mexternaldiffrecords"
	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateMExternalDiffRecordsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMExternalDiffRecordsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := mexternaldiffrecords.NewCreateMExternalDiffRecordsLogic(r.Context(), svcCtx)
		resp, err := l.CreateMExternalDiffRecords(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
