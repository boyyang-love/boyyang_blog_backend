package star

import (
	"blog_server/common/respx"
	"net/http"

	"blog_server/internal/logic/star"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StarReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := star.NewStarLogic(r.Context(), svcCtx)
		err, msg := l.Star(&req)
		respx.Response(w, nil, err, msg)
	}
}
