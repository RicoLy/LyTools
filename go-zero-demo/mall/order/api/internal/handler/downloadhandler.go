package handler

import (
	"net/http"

	"Testing/go-zero-demo/mall/order/api/internal/logic"
	"Testing/go-zero-demo/mall/order/api/internal/svc"
	"Testing/go-zero-demo/mall/order/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDownloadLogic(r.Context(), svcCtx)
		err := l.Download(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
