package like

import (
	"blog_server/common/respx"
	logic "blog_server/internal/logic/like"
	"net/http"

	"blog_server/internal/svc"
)

func LikesInfoIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLikesInfoIdsLogic(r.Context(), svcCtx)
		resp, err, msg := l.LikesInfoIds()
		respx.Response(w, resp, err, msg)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
