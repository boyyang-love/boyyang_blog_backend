package like

import (
	"blog_server/common/respx"
	"blog_server/internal/logic/like"
	"net/http"

	"blog_server/internal/svc"
)

func LikesInfoIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := like.NewLikesInfoIdsLogic(r.Context(), svcCtx)
		resp, err, msg := l.LikesInfoIds()
		respx.Response(w, resp, err, msg)

	}
}
