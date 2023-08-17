// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	blog "blog_server/internal/handler/blog"
	comment "blog_server/internal/handler/comment"
	dashboard "blog_server/internal/handler/dashboard"
	exhibition "blog_server/internal/handler/exhibition"
	like "blog_server/internal/handler/like"
	login "blog_server/internal/handler/login"
	star "blog_server/internal/handler/star"
	user "blog_server/internal/handler/user"
	"blog_server/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/blog/create",
				Handler: blog.CreateBlogHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/blog/update",
				Handler: blog.UpdateBlogHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/blog/delete",
				Handler: blog.DeleteBlogHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/blog/info",
				Handler: blog.BloginfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/blog/thumbsup",
				Handler: blog.ThumbsupBlogHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/blog/comment/create",
				Handler: comment.CreateCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/blog/comment/delete",
				Handler: comment.DeleteCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/blog/comment/thumbsup",
				Handler: comment.ThumbsupCommentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/dashboard",
				Handler: dashboard.DashboardHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/like",
				Handler: like.LikeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/likes/info",
				Handler: like.LikesInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/likes/info/ids",
				Handler: like.LikesInfoIdsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/follow",
				Handler: like.FollowHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/follow/info",
				Handler: like.FollowInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: login.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: login.RegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: exhibition.UploadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/cos/upload",
				Handler: exhibition.CosUploadHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/exhibition/create",
				Handler: exhibition.CreateExhibitionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/exhibition/info",
				Handler: exhibition.ExhibitionInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/exhibition/update",
				Handler: exhibition.UpdateExhibitionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/exhibition/approval",
				Handler: exhibition.ApprovalExhibitionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/exhibition/del",
				Handler: exhibition.DelExhibitionHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/users/info",
				Handler: user.UserinfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/users/update",
				Handler: user.UpdateUserinfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/users/update/password",
				Handler: user.UpdateUserPasswordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/star",
				Handler: star.StarHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
