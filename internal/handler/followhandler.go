package handler

import (
	"blog_server/common/response"
	"net/http"

	"blog_server/internal/logic"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func followHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddAndUnFollowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFollowLogic(r.Context(), svcCtx)
		err, msg := l.Follow(&req)
		response.Response(w, nil, err, msg)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.Ok(w)
		//}
	}
}
