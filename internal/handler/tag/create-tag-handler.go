package tag

import (
	"blog_server/common/respx"
	"net/http"

	"blog_server/internal/logic/tag"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateTagReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tag.NewCreateTagLogic(r.Context(), svcCtx)
		err, msg := l.CreateTag(&req)

		respx.Response(w, nil, err, msg)
	}
}
