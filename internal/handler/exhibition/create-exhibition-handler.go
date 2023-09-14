package exhibition

import (
	"blog_server/common/respx"
	"blog_server/internal/logic/exhibition"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func CreateExhibitionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateExhibitionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := exhibition.NewCreateExhibitionLogic(r.Context(), svcCtx)
		resp, err, msg := l.CreateExhibition(&req)
		respx.Response(w, resp, err, msg)
	}
}
