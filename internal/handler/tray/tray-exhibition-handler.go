package tray

import (
	"net/http"

	"blog_server/internal/logic/tray"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TrayExhibitionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TrayReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tray.NewTrayExhibitionLogic(r.Context(), svcCtx)
		resp, err := l.TrayExhibition(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
