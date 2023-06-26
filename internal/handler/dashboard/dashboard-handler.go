package dashboard

import (
	"blog_server/common/respx"
	logic "blog_server/internal/logic/dashboard"
	"blog_server/internal/svc"
	"net/http"
)

func DashboardHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDashboardLogic(r.Context(), svcCtx)
		resp, err, msg := l.Dashboard()
		respx.Response(w, resp, err, msg)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
