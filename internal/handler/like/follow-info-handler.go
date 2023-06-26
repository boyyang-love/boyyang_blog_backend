package like

import (
	"blog_server/common/respx"
	logic "blog_server/internal/logic/like"
	"blog_server/internal/svc"
	"net/http"
)

func FollowInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFollowInfoLogic(r.Context(), svcCtx)
		resp, err, msg := l.FollowInfo()
		respx.Response(w, resp, err, msg)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
