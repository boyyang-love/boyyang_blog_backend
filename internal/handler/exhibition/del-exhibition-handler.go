package exhibition

import (
	"blog_server/common/respx"
	logic "blog_server/internal/logic/exhibition"
	"net/http"

	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelExhibitionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelExhibitionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDelExhibitionLogic(r.Context(), svcCtx)
		err, msg := l.DelExhibition(&req)

		respx.Response(w, nil, err, msg)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.Ok(w)
		//}
	}
}
