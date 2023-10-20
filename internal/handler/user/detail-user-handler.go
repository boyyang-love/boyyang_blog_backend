package user

import (
	"blog_server/common/respx"
	"net/http"

	"blog_server/internal/logic/user"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DetailUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDetailUserLogic(r.Context(), svcCtx)
		resp, err, msg := l.DetailUser(&req)
		respx.Response(w, resp, err, msg)
	}
}
