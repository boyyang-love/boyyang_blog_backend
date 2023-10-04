package article

import (
	"blog_server/common/respx"
	"net/http"

	"blog_server/internal/logic/article"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewDeleteArticleLogic(r.Context(), svcCtx)
		err, msg := l.DeleteArticle(&req)

		respx.Response(w, nil, err, msg)
	}
}
