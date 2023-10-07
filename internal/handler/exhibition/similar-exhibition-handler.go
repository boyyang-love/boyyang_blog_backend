package exhibition

import (
	"blog_server/common/respx"
	"net/http"

	"blog_server/internal/logic/exhibition"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SimilarExhibitionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SimilarReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := exhibition.NewSimilarExhibitionLogic(r.Context(), svcCtx)
		resp, err, msg := l.SimilarExhibition(&req)

		respx.Response(w, resp, err, msg)
	}
}
