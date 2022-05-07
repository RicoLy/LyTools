package handler

import (
	"net/http"

	"Testing/go-zero-demo/mall/order/api/internal/logic"
	"Testing/go-zero-demo/mall/order/api/internal/svc"
	"Testing/go-zero-demo/mall/order/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func addOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddOrderLogic(r.Context(), svcCtx)
		resp, err := l.AddOrder(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
