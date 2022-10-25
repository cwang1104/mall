package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mall/apps/app/api/internal/logic/user"
	"mall/apps/app/api/internal/svc"
	"mall/apps/app/api/internal/types"
)

func LoginHandleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLoginHandleLogic(r.Context(), svcCtx)
		resp, err := l.LoginHandle(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
