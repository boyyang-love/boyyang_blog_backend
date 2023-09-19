package admin

import (
	"blog_server/common/respx"
	"net/http"

	"blog_server/internal/logic/admin"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ExhibitionAdminInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminExhibitionsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewExhibitionAdminInfoLogic(r.Context(), svcCtx)
		resp, err, msg := l.ExhibitionAdminInfo(&req)
		respx.Response(w, resp, err, msg)
	}
}
