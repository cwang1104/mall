package seckill

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mall/apps/app/api/internal/logic/seckill"
	"mall/apps/app/api/internal/svc"
)

func GetSeckillResHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := seckill.NewGetSeckillResLogic(r.Context(), svcCtx)
		resp, err := l.GetSeckillRes()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
