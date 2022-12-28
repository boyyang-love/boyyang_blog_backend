package handler

import (
	"blog_server/common/response"
	"net/http"

	"blog_server/internal/logic"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createExhibitionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateExhibitionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateExhibitionLogic(r.Context(), svcCtx)
		resp, err := l.CreateExhibition(&req)
		response.Response(w, resp, err)

		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
