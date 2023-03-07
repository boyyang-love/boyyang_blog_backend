package handler

import (
	"blog_server/common/response"
	"net/http"

	"blog_server/internal/logic"
	"blog_server/internal/svc"
)

func followInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFollowInfoLogic(r.Context(), svcCtx)
		resp, err, msg := l.FollowInfo()
		response.Response(w, resp, err, msg)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
