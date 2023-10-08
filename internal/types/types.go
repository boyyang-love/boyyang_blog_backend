// Code generated by goctl. DO NOT EDIT.
package types

type CreateBlogReq struct {
	Title    string `form:"title"`
	SubTitle string `form:"sub_title"`
	Content  string `form:"content"`
	Cover    string `form:"cover"`
	UserId   uint   `form:"user_id,optional"`
	Tags     string `form:"tags,optional"`
}

type CreateBlogRes struct {
	Uid uint32 `json:"uid"`
}

type UpdateBlogReq struct {
	Uid      uint32 `form:"uid"`
	Title    string `form:"title"`
	SubTitle string `form:"sub_title"`
	Content  string `form:"content"`
	Cover    string `form:"cover,optional"`
}

type UpdateBlogRes struct {
	Msg string `json:"msg"`
}

type DeleteBlogReq struct {
	Uid uint32 `form:"uid"`
}

type BlogInfoReq struct {
	Uids  string `form:"uids,optional"`
	Page  string `form:"page,optional"`
	Limit string `form:"limit,optional"`
}

type BlogInfoRes struct {
	Count    int64      `json:"count"`
	BlogInfo []BlogInfo `json:"blog_info"`
}

type BlogInfo struct {
	Uid        uint32 `json:"uid"`
	Created    int    `json:"created"`
	Updated    int    `json:"updated"`
	Title      string `json:"title"`
	SubTitle   string `json:"sub_title"`
	Content    string `json:"content"`
	Cover      string `json:"cover"`
	ThumbsUp   *int   `json:"thumbs_up"`
	Collection *int   `json:"collection"`
	UserId     uint   `json:"user_id"`
	UserInfo   User   `json:"user_info,omitempty" gorm:"foreignKey:UserId;references:Uid"`
	Tag        string `json:"tags"`
}

type ThumbsUpBlogReq struct {
	Uid uint32 `form:"uid"`
}

type ThumbsUpBlogRes struct {
	Msg string `json:"msg"`
}

type User struct {
	Id              uint   `json:"id"`
	Uid             uint   `json:"uid"`
	Username        string `json:"username"`
	Gender          int    `json:"gender"`
	AvatarUrl       string `json:"avatar_url"`
	Age             int    `json:"age"`
	Email           string `json:"email"`
	Address         string `json:"address"`
	Tel             int    `json:"tel"`
	Qq              int    `json:"qq"`
	Wechat          string `json:"wechat"`
	GitHub          string `json:"git_hub"`
	BackgroundImage string `json:"background_image"`
	Motto           string `json:"motto"`
	Role            string `json:"role"`
}

type CreateBlogCommentReq struct {
	Content string `form:"content" gorm:"size:2000"`
	BlogId  uint32 `form:"blog_id"`
}

type CreateBlogCommentRes struct {
	Msg string `json:"msg"`
}

type DeleateBlogCommentReq struct {
	Uid uint32 `form:"uid"`
}

type DeleateBlogCommentRes struct {
	Msg string `json:"msg"`
}

type ThumbsUpBlogCommentReq struct {
	Uid uint32 `form:"uid"`
}

type ThumbsUpBlogCommentRes struct {
	Msg string `json:"msg"`
}

type DashboardRes struct {
	UserInfo    DashboardUserInfo     `json:"user_info"` // 用户信息
	Dashboard   []Dashboard           `json:"dashboard"`
	Exhibitions []DashboardExhibition `json:"exhibitions"`
}

type Dashboard struct {
	Name                    string `json:"name"`
	BlogPublishValue        string `json:"blog_publish_value"`
	ExhibitionsPublishValue string `json:"exhibitions_publish_value"`
}

type DashboardUserInfo struct {
	ThumbsUp      *int `json:"thumbs_up" gorm:"default:0"` // 获赞数
	Like          *int `json:"like" gorm:"default:0"`      // 收藏数
	Publish       *int `json:"publish" gorm:"default:0"`   // 上传数
	Following     *int `json:"following"`
	DashboardUser User `json:"dashboard_user"`
}

type DashboardExhibition struct {
	Uid    uint32 `json:"uid"`
	Title  string `json:"title"`
	Des    string `json:"des"`
	Cover  string `json:"cover"`
	UserId uint   `json:"user_id"`
}

type AddLikesReq struct {
	Uid       uint32 `form:"uid"`        // 图片ID 或者是博客ID
	LikesType int    `form:"likes_type"` // 1 true 0 false
	Type      int    `form:"type"`       // 1  图片 2 博客
}

type LikesInfoReq struct {
	ExhibitionId uint32 `form:"exhibition_id,optional"`
}

type LikesInfoRes struct {
	LikesInfo []LikesInfo `json:"likes_info"`
}

type LikesInfo struct {
	Uid      uint32 `json:"uid"`
	Created  int    `json:"created"`
	Updated  int    `json:"updated"`
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
	Des      string `json:"des"`
	Cover    string `json:"cover"`                      // 图片上传路径
	ThumbsUp *int   `json:"thumbs_up" gorm:"default:0"` // 点赞数
	UserId   uint32 `json:"user_id"`                    // 该图片上传者 id
}

type LikesInfoIds struct {
	Uids []uint32 `json:"uids"`
}

type AddAndUnFollowReq struct {
	FollowId   uint32 `form:"follow_id"`
	FollowType int    `form:"follow_type"` // 1 添加 0 取消
}

type FollowInfoRes struct {
	FollowingUser []User `json:"following_user"`
}

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginRes struct {
	Info  User   `json:"info"`
	Token string `json:"token"`
}

type RegisterReq struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	Tel       int    `form:"tel"`
	AvatarUrl string `form:"avatar_url"`
}

type RegisterRes struct {
	Uid uint32 `json:"uid"`
}

type UploadReq struct {
	Hash     string `json:"hash,optional"`
	FileName string `json:"file_name,optional"`
	Ext      string `json:"ext,optional"`
	Size     int64  `json:"size,optional"`
	FilePath string `json:"file_path,optional"`
}

type UploadRes struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}

type CosUploadReq struct {
	FileName string `form:"file_name"`
	Path     string `form:"path,optional"`
}

type CosUploadRes struct {
	Url           string `json:"url"`
	Token         string `json:"token"`
	Authorization string `json:"authorization"`
	FileId        string `json:"file_id"`
	CosFileId     string `json:"cos_file_id"`
	Key           string `json:"key"`
}

type CreateExhibitionReq struct {
	Title    string `form:"title"`
	SubTitle string `form:"sub_title,optional"`
	Des      string `form:"des"`
	Cover    string `form:"cover"`
	Tags     string `form:"tags,optional"`
	Type     string `form:"type,optional"`
	Size     int    `form:"size,optional"`
	Px       string `form:"px"`
	Rgb      string `form:"rgb,optional"`
	Palette  string `form:"palette,optional"`
}

type CreateExhibitionRes struct {
	Uid uint `json:"uid"`
}

type ExhibitionInfoReq struct {
	Uids     string `form:"uids,optional"`
	Page     int    `form:"page,optional"`
	Limit    int    `form:"limit,optional"`
	Type     int    `form:"type,optional"`
	Public   bool   `form:"public,optional"`
	IsLike   bool   `form:"is_like,optional"`
	Sort     string `form:"sort,optional"` // Created  ThumbsUp
	Keywords string `form:"keywords,optional"`
	Tags     string `form:"tags,optional"`
}

type ExhibitionInfoRes struct {
	Count          int              `json:"count"`
	Exhibitions    []ExhibitionInfo `json:"exhibitions"`
	InReview       int              `json:"in_review"`       // 审核中
	Approved       int              `json:"approved"`        // 审核通过
	ReviewRjection int              `json:"review_rjection"` //审核驳回
	LikesIds       []int            `json:"likes_ids"`       // 收藏ID集合
	StarIds        []int            `json:"star_ids"`        // star ID集合
}

type ExhibitionInfo struct {
	Uid       uint   `json:"uid"`
	Created   int    `json:"created"`
	Title     string `json:"title"`
	SubTitle  string `json:"sub_title"`
	Des       string `json:"des"`
	Cover     string `json:"cover"`
	Tags      string `json:"tags"`
	UserId    uint   `json:"user_id"`
	ThumbsUp  int    `json:"thumbs_up"`
	Download  int    `json:"download"`
	Status    int    `json:"status"`     // 1待审核 2审核通过 3未通过审核
	RejectRes string `json:"reject_res"` // 驳回原因
	UserInfo  User   `json:"user_info,omitempty" gorm:"foreignKey:UserId;references:uid"`
	Type      string `json:"type"`
	Size      int    `json:"size"`
	Px        string `json:"px"`
	Rgb       string `json:"rgb"`
	Palette   string `json:"palette"`
}

type UpdateExhibitionReq struct {
	Uid      uint   `form:"uid"`
	Title    string `form:"title,optional"`
	SubTitle string `form:"sub_title,optional"`
	Des      string `form:"des,optional"`
}

type UpdateExhibitionRes struct {
	Uid uint `json:"uid"`
}

type ApprovalReq struct {
	Uid    uint   `form:"uid"`
	Status int    `form:"status"`
	Reason string `form:"reason,optional"`
}

type ApprovalRes struct {
	Uid uint `json:"uid"`
}

type DelExhibitionReq struct {
	Uid uint `form:"uid"`
}

type DelUploadReq struct {
	Key string `form:"key"`
}

type UpdateDownloadReq struct {
	Uid uint `form:"uid"`
}

type SimilarReq struct {
	Tag   string `form:"tag"`
	Color string `form:"color"`
}

type SimilarRes struct {
	Infos []ExhibitionInfo `json:"infos"`
}

type UserInfoReq struct {
	Uid uint32 `form:"uid,optional"`
}

type UserInfoRes struct {
	UserInfo      User          `json:"user_info"`   // 用户基本信息
	UserOtherInfo UserOtherInfo `json:"user_detail"` // 其它信息
}

type UserOtherInfo struct {
	Publish  int `json:"publish"`   // 发布数
	Likes    int `json:"likes"`     // 收藏数
	Follows  int `json:"follows"`   // 粉丝数
	ThumbsUp int `json:"thumbs_up"` // 点赞数
}

type UpdateUserInfoReq struct {
	Uid             uint32 `form:"uid"`
	Username        string `form:"username,optional"`
	Age             int    `form:"age,optional"`
	Gender          int    `form:"gender,optional"`
	AvatarUrl       string `form:"avatar_url,optional"`
	Tel             int    `form:"tel,optional"`
	Email           string `form:"email,optional"`
	Address         string `form:"address,optional"`
	BackgroundImage string `form:"background_image,optional"`
	Motto           string `form:"motto,optional"`
	Qq              int    `form:"qq,optional"`
	Wechat          string `form:"wechat,optional"`
	GitHub          string `form:"git_hub,optional"`
}

type UpdateUserPasswordReq struct {
	Password string `form:"password"`
}

type StarReq struct {
	Uid      uint32 `form:"uid"`
	StarType int    `form:"star_type"` // 0 取消star 1 star
	Type     int    `form:"type"`      // 1 图片 2 博客
}

type CreateTagReq struct {
	Name string `form:"name"`
	Type string `form:"type"`
}

type TagsInfoReq struct {
	Type string `form:"type"`
	Uids string `form:"uids,optional"`
}

type TagsInfoRes struct {
	TagsInfo []TagInfo `json:"tags_info"`
}

type TagInfo struct {
	Uid  uint32 `json:"uid"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type SearchReq struct {
	Keyword string `form:"keyword"`
	Type    int    `form:"type"` // 1 图片 2 博客
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
}

type SearchRes struct {
	Count          int64                   `json:"count"`
	ExhibitionInfo []SearchExhibitionInfos `json:"exhibitionInfo,omitempty"`
	BlogInfo       []SearchBlogInfos       `json:"blogInfo,omitempty"`
}

type SearchExhibitionInfos struct {
	Uid       uint32  `json:"uid" gorm:"primary_key"`
	Created   int     `json:"created" gorm:"autoCreateTime"`
	Updated   int     `json:"updated" gorm:"autoUpdateTime"`
	Title     string  `json:"title"`                      // 图片墙标题
	SubTitle  string  `json:"sub_title"`                  // 图片墙副标题
	Des       string  `json:"des"`                        // 图片描述
	Cover     string  `json:"cover"`                      // 图片上传路径
	Tags      *string `json:"tags"`                       // 图片标签
	ThumbsUp  *int    `json:"thumbs_up" gorm:"default:0"` // 点赞数
	UserId    uint32  `json:"user_id"`                    // 该图片上传者 id
	Status    int     `json:"status" gorm:"default:1"`    // 图片状态 1待审核 2审核通过 3未通过审核
	RejectRes string  `json:"reject_res"`                 // 状态为3时 驳回原因
}

type SearchBlogInfos struct {
	Uid          uint32 `json:"uid" gorm:"primary_key"`
	Created      int    `json:"created" gorm:"autoCreateTime"`
	Updated      int    `json:"updated" gorm:"autoUpdateTime"`
	Title        string `json:"title"`                       // 博客标题
	SubTitle     string `json:"sub_title"`                   // 博客副标题
	Content      string `json:"des" gorm:"size:10000"`       // 博客内容
	Cover        string `json:"cover,omitempty"`             // 背景图片
	UserId       uint32 `json:"user_id"`                     // 博客作者
	Tag          string `json:"tag,omitempty"`               // 博客标签
	ThumbsUp     int    `json:"thumbs_up" gorm:"default:0"`  // 点赞数
	ThumbsUpList string `json:"thumbs_up_list"`              // 点赞id集合
	Collection   *int   `json:"collection" gorm:"default:0"` // 收藏数
}

type TrayReq struct {
	Page  int `form:"page,optional"`
	Limit int `form:"limit,optional"`
}

type TrayRes struct {
	Count           int64                `json:"count"`
	TrayExhibitions []TrayExhibitionInfo `json:"exhibitions"`
}

type TrayExhibitionInfo struct {
	Uid   uint   `json:"uid"`
	Cover string `json:"cover"`
}

type AdminExhibitionsReq struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Type  int    `form:"type"`
	Sort  string `form:"sort,optional"`
}

type AdminExhibitionsRes struct {
	Count       int64                 `json:"count"`
	Exhibitions []AdminExhibitionInfo `json:"exhibitions"`
}

type AdminExhibitionInfo struct {
	Uid       uint   `json:"uid"`
	Created   int    `json:"created"`
	Title     string `json:"title"`
	SubTitle  string `json:"sub_title"`
	Des       string `json:"des"`
	Cover     string `json:"cover"`
	Tags      string `json:"tags"`
	UserId    uint   `json:"user_id"`
	ThumbsUp  int    `json:"thumbs_up"`
	Status    int    `json:"status"`     // 1待审核 2审核通过 3未通过审核
	RejectRes string `json:"reject_res"` // 驳回原因
	UserInfo  User   `json:"user_info,omitempty" gorm:"foreignKey:UserId;references:uid"`
	Type      string `json:"type"`
	Size      int    `json:"size"`
	Px        string `json:"px"`
}

type AdminStatRes struct {
	UserCount          int64 `json:"user_count"`
	ImageCount         int64 `json:"image_count"`
	ImageDownloadCount int64 `json:"image_download_count"`
	BlogCount          int64 `json:"blog_count"`
	ArticleCount       int64 `json:"article_count"`
}

type CreateArticleReq struct {
	Title    string `form:"title"`
	SubTitle string `form:"sub_title"`
	Content  string `form:"content"`
	Cover    string `form:"cover"`
	Tag      string `form:"tag"`
}

type UpdateArticleReq struct {
	Uid      uint32 `form:"uid"`
	Title    string `form:"title"`
	SubTitle string `form:"subtitle"`
	Content  string `form:"content"`
	Tag      string `form:"tag,optional"`
}

type DeleteArticleReq struct {
	Uid uint32 `form:"uid"`
}

type InfoArticleReq struct {
	Uid     uint32 `form:"uid,optional"`
	Page    int    `form:"page,optional"`
	Limit   int    `form:"limit,optional"`
	Keyword string `form:"keyword,optional"`
}

type InfoArticleRes struct {
	Count       int64         `json:"count,omitempty"`
	ArticleInfo []ArticleInfo `json:"article_info,omitempty"`
}

type ArticleInfo struct {
	Uid      uint32 `json:"uid"`
	Created  int    `json:"created"`
	Updated  int    `json:"updated"`
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
	Content  string `json:"content"`
	Cover    string `json:"cover"`
	UserId   uint   `json:"user_id"`
	UserInfo User   `json:"user_info,omitempty" gorm:"foreignKey:UserId;references:Uid"`
	Tag      string `json:"tags"`
}

type NoticeCreateReq struct {
	Content string `form:"content"`
}

type NoticeUpdateReq struct {
	Uid     uint32 `form:"uid"`
	Content string `form:"content"`
}

type NoticeDeleteReq struct {
	Uid uint32 `form:"uid"`
}

type NoticeInfoReq struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type NoticeInfoRes struct {
	Count int64        `json:"count"`
	Infos []NoticeInfo `json:"infos"`
}

type NoticeInfo struct {
	Uid     uint32 `json:"uid"`
	UserId  uint32 `json:"user_id"`
	Content string `json:"content"`
}
