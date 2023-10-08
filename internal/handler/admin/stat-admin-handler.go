package admin

import (
	"blog_server/common/respx"
	"net/http"

	"blog_server/internal/logic/admin"
	"blog_server/internal/svc"
)

func StatAdminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewStatAdminLogic(r.Context(), svcCtx)
		resp, err, msg := l.StatAdmin()

		respx.Response(w, resp, err, msg)
	}
}
