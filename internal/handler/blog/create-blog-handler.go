package blog

import (
	"blog_server/common/respx"
	"blog_server/internal/logic/blog"
	"net/http"

	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateBlogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateBlogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := blog.NewCreateBlogLogic(r.Context(), svcCtx)
		resp, err, msg := l.CreateBlog(&req)
		respx.Response(w, resp, err, msg)
	}
}
