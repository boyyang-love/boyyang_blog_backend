package user

import (
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type Detail struct {
	exhibition int64
	thumbsUp   int64
	download   int64
	likes      int64
	follow     int64
	followIds  []int64
}

func NewDetailUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailUserLogic {
	return &DetailUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailUserLogic) DetailUser(req *types.InfoUserReq) (resp *types.InfoUserRes, err error, msg respx.SucMsg) {
	err, detail := l.detailInfo(req)
	if err != nil {
		return nil, err, msg
	}

	err, count, exs := l.exhibition(req)
	if err != nil {
		return nil, err, msg
	}

	err, info := l.userInfo(req)
	if err != nil {
		return nil, err, msg
	}
	return &types.InfoUserRes{
		Upload:     detail.exhibition,
		Star:       detail.thumbsUp,
		Like:       detail.likes,
		Download:   detail.download,
		Follow:     detail.follow,
		Count:      count,
		Exhibition: exs,
		UserInfo:   *info,
		FollowIds:  detail.followIds,
	}, nil, respx.SucMsg{Msg: "获取成功"}
}

func (l *DetailUserLogic) detailInfo(req *types.InfoUserReq) (err error, detail *Detail) {
	// 图片上传数
	var exhibition int64
	if err = l.svcCtx.DB.
		Model(&models.Exhibition{}).
		Where("user_id = ? and status = ?", req.Uid, 2).
		Count(&exhibition).
		Error; err != nil {
		return err, nil
	}
	// 图片点赞数 下载数
	var thumbsUpAndDownload struct {
		ThumbsUp int64 `json:"thumbsUp"`
		Download int64 `json:"download"`
	}
	if err = l.svcCtx.DB.
		Model(&models.Exhibition{}).
		Select("sum(thumbs_up) as thumbs_up, sum(download) as download").
		Where("user_id = ? and status = ?", req.Uid, 2).
		Scan(&thumbsUpAndDownload).
		Error; err != nil {
		return err, nil
	}
	// 图片收藏数
	var likes int64
	if err = l.svcCtx.DB.
		Model(&models.Likes{}).
		Where("user_id = ? and type = ? and likes_type = ?", req.Uid, 1, true).
		Count(&likes).
		Error; err != nil {
		return err, nil
	}
	// 粉丝数
	var follow int64
	var followIds []int64
	if err = l.svcCtx.DB.
		Model(&models.Follow{}).
		Select("user_id").
		Where("follow_user_id = ? and follow_type = ?", req.Uid, true).
		Find(&followIds).
		Count(&follow).
		Error; err != nil {
		return err, nil
	}

	return nil, &Detail{
		exhibition: exhibition,
		thumbsUp:   thumbsUpAndDownload.ThumbsUp,
		download:   thumbsUpAndDownload.Download,
		likes:      likes,
		follow:     follow,
		followIds:  followIds,
	}
}

func (l *DetailUserLogic) exhibition(req *types.InfoUserReq) (err error, count int64, ex []types.InfoExhibition) {
	DB := l.svcCtx.DB.
		Order("created desc").
		Model(&models.Exhibition{}).
		Offset((req.Page - 1) * req.Limit).
		Limit(req.Limit)
	var exs []types.InfoExhibition

	if req.Type == 1 {
		DB = DB.
			Where("user_id = ? and status = ?", req.Uid, 2).
			Scan(&exs)
	}

	if req.Type == 2 {
		var ids []int64
		if err = l.svcCtx.DB.
			Table("likes").
			Select("likes_id").
			Where("user_id = ? and type = ? and likes_type = ?", req.Uid, 1, true).
			Scan(&ids).
			Error; err != nil {
			return err, count, ex
		}
		DB = DB.
			Where("uid in ? and status = ?", ids, 2).
			Find(&exs)
	}

	if err = DB.
		Offset(-1).
		Count(&count).
		Error; err != nil {
		return err, count, nil
	}

	return nil, count, exs
}

func (l *DetailUserLogic) userInfo(req *types.InfoUserReq) (err error, info *types.User) {
	var userInfo *types.User
	if err = l.svcCtx.DB.
		Model(&models.User{}).
		Where("uid = ?", req.Uid).
		Select("uid", "username", "avatar_url", "background_image", "gender").
		First(&userInfo).
		Error; err != nil {
		return err, nil
	}

	return nil, userInfo
}
