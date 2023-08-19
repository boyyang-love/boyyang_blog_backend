package tag

import (
	"net/http"

	"blog_server/internal/logic/tag"
	"blog_server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TagsInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tag.NewTagsInfoLogic(r.Context(), svcCtx)
		resp, err := l.TagsInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
