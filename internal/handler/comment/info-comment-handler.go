package comment

import (
	"blog_server/common/respx"
	"net/http"

	"blog_server/internal/logic/comment"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InfoCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewInfoCommentLogic(r.Context(), svcCtx)
		resp, err, msg := l.InfoComment(&req)
		respx.Response(w, resp, err, msg)
	}
}
