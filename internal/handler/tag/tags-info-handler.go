package tag

import (
	"blog_server/common/respx"
	"blog_server/internal/logic/tag"
	"blog_server/internal/types"
	"net/http"

	"blog_server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TagsInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TagsInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := tag.NewTagsInfoLogic(r.Context(), svcCtx)
		resp, err, msg := l.TagsInfo(&req)

		respx.Response(w, resp, err, msg)
	}
}
