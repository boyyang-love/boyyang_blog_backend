package handler

import (
	"blog_server/common/response"
	"net/http"

	"blog_server/internal/logic"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateBlogCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateCommentLogic(r.Context(), svcCtx)
		resp, err, msg := l.CreateComment(&req)
		response.Response(w, resp, err, msg)

		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
