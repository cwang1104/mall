package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mall/apps/app/api/internal/logic/user"
	"mall/apps/app/api/internal/svc"
	"mall/apps/app/api/internal/types"
)

func RegisterHandleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewRegisterHandleLogic(r.Context(), svcCtx)
		resp, err := l.RegisterHandle(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
