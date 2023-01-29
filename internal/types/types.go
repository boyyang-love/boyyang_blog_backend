// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginRes struct {
	Info  User   `json:"info"`
	Token string `json:"token"`
}

type RegisterReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Tel      int    `form:"tel"`
}

type RegisterRes struct {
	Id int `json:"id"`
}

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Gender    int    `json:"gender"`
	AvatarUrl string `json:"avatarUrl"`
	Tel       int    `json:"tel"`
}

type Comment struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	Content  string `json:"des" gorm:"size:2000"`
	BlogId   uint   `json:"blog_id"`
	UserId   uint   `json:"user_id"`
	Tag      string `json:"tag,omitempty"`
	ThumbsUp *int   `json:"thumbs_up" gorm:"default:0"`
	UserInfo User   `json:"userInfo"`
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

type UserInfoReq struct {
	Id int `form:"id"`
}

type UserInfoRes struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Age       int    `json:"age,omitempty"`
	Gender    int    `json:"gender"`
	AvatarUrl string `json:"avatarUrl"`
	Tel       int    `json:"tel"`
	Email     string `json:"email,omitempty"`
	Address   string `json:"address,omitempty"`
}

type UpdateUserInfoReq struct {
	Id        int    `form:"id"`
	Username  string `form:"username,optional"`
	Age       int    `form:"age,optional"`
	Gender    int    `form:"gender,optional"`
	AvatarUrl string `form:"avatarUrl,optional"`
	Tel       int    `form:"tel,optional"`
	Email     string `form:"email,optional"`
	Address   string `form:"address,optional"`
}

type UpdateUserInfoRes struct {
	Msg string `json:"msg"`
}

type CreateExhibitionReq struct {
	Title    string `form:"title"`
	SubTitle string `form:"sub_title"`
	Des      string `form:"des"`
	Cover    string `form:"cover"`
}

type CreateExhibitionRes struct {
	Id int `json:"id"`
}

type ExhibitionInfoReq struct {
	Ids   string `form:"ids,optional"`
	Page  string `form:"page,optional"`
	Limit string `form:"limit,optional"`
}

type ExhibitionInfoRes struct {
	Count       int              `json:"count"`
	Exhibitions []ExhibitionInfo `json:"exhibitions"`
}

type ExhibitionInfo struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	SubTitle  string `json:"sub_title"`
	Des       string `json:"des"`
	Cover     string `json:"cover"`
	UserId    uint   `json:"user_id"`
	Status    int    `json:"status"`     // 1待审核 2审核通过 3未通过审核
	RejectRes string `json:"reject_res"` // 驳回原因
	UserInfo  User   `json:"user_info" gorm:"foreignKey:UserId"`
}

type UpdateExhibitionReq struct {
	Id       uint   `form:"id"`
	Title    string `form:"title,optional"`
	SubTitle string `form:"sub_title,optional"`
	Des      string `form:"des,optional"`
}

type UpdateExhibitionRes struct {
	Id uint `json:"id"`
}

type ApprovalReq struct {
	Id     uint   `form:"id"`
	Status int    `form:"status"`
	Reason string `form:"reason"`
}

type ApprovalRes struct {
	Id uint `json:"id"`
}

type CreateBlogReq struct {
	Title    string `form:"title"`
	SubTitle string `form:"sub_title"`
	Content  string `form:"content"`
	Cover    string `form:"cover"`
	UserId   int    `form:"user_id,optional"`
}

type CreateBlogRes struct {
	Id int `json:"id"`
}

type UpdateBlogReq struct {
	Id       int    `form:"id"`
	Title    string `form:"title"`
	SubTitle string `form:"sub_title"`
	Content  string `form:"content"`
	Cover    string `form:"cover"`
}

type UpdateBlogRes struct {
	Msg string `json:"msg"`
}

type DeleteBlogReq struct {
	Id int `form:"id"`
}

type DeleteBlogRes struct {
	Msg string `json:"msg"`
}

type BlogInfoReq struct {
	Ids   string `form:"ids,optional"`
	Page  string `form:"page,optional"`
	Limit string `form:"limit,optional"`
}

type BlogInfoRes struct {
	Count    int        `json:"count"`
	BlogInfo []BlogInfo `json:"blog_info"`
}

type BlogInfo struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	SubTitle string    `json:"sub_title"`
	Content  string    `json:"content"`
	Cover    string    `json:"cover"`
	ThumbsUp *int      `json:"thumbs_up"`
	UserInfo User      `json:"user_info"`
	Comments []Comment `json:"comments,omitempty"`
}

type ThumbsUpBlogReq struct {
	Id int `form:"id"`
}

type ThumbsUpBlogRes struct {
	Msg string `json:"msg"`
}

type CreateBlogCommentReq struct {
	Content string `form:"content" gorm:"size:2000"`
	BlogId  uint   `form:"blog_id"`
}

type CreateBlogCommentRes struct {
	Msg string `json:"msg"`
}

type DeleateBlogCommentReq struct {
	Id uint `form:"id"`
}

type DeleateBlogCommentRes struct {
	Msg string `json:"msg"`
}

type ThumbsUpBlogCommentReq struct {
	Id uint `form:"id"`
}

type ThumbsUpBlogCommentRes struct {
	Msg string `json:"msg"`
}
