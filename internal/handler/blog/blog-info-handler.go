package blog

import (
	"blog_server/common/respx"
	"net/http"

	"blog_server/internal/logic/blog"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BlogInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BlogInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := blog.NewBlogInfoLogic(r.Context(), svcCtx)
		resp, err, msg := l.BlogInfo(&req)
		respx.Response(w, resp, err, msg)
	}
}
