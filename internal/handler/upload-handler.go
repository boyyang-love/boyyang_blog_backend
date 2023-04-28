package handler

import (
	"blog_server/common/helper"
	"blog_server/common/respx"
	"blog_server/models"
	"net/http"
	"path"

	"blog_server/internal/logic"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func uploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		token := r.Header.Get("Authorization")
		claims, _ := helper.ParseJwtToken(token, svcCtx.Config.Auth.AccessSecret)

		// 文件处理
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}

		hash, _ := helper.MakeFileHash(file, fileHeader)
		fileInfo := new(models.Upload)
		has := svcCtx.DB.
			Model(&models.Upload{}).
			Where("hash = ?", hash).
			First(&fileInfo)

		if has.RowsAffected != 0 {
			respx.Response(w, &types.UploadRes{
				FileName: fileInfo.FileName,
				FilePath: fileInfo.FilePath,
			}, nil, respx.SucMsg{Msg: "上传成功"})

			return
		}

		url, err := helper.CosFileUpload(svcCtx.Client, file, fileHeader, "images")
		if err != nil {
			httpx.Error(w, err)

			return
		}

		svcCtx.DB.Model(&models.Upload{}).Create(&models.Upload{
			Hash:     hash,
			FileName: fileHeader.Filename,
			FilePath: url,
			Ext:      path.Ext(fileHeader.Filename),
			Size:     fileHeader.Size,
			UserId:   uint(claims.Id),
		})

		req.FilePath = url
		req.FileName = fileHeader.Filename
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Ext = path.Ext(fileHeader.Filename)

		l := logic.NewUploadLogic(r.Context(), svcCtx)
		resp, err, msg := l.Upload(&req)
		respx.Response(w, resp, err, msg)

		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
