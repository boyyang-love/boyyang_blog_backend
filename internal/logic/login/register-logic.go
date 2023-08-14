package login

import (
	"blog_server/common/helper"
	"blog_server/common/respx"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"blog_server/models"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error, msg respx.SucMsg) {
	resp, err = l.register(req)
	if err != nil {
		return nil, err, msg
	}

	return resp, nil, respx.SucMsg{Msg: "账号注册成功"}
}

func (l *RegisterLogic) register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	DB := l.svcCtx.DB

	var user struct {
		Id       uint   `json:"id"`
		Username string `json:"username"`
		Tel      int    `json:"tel"`
	}

	isExistDB := DB.
		Model(&models.User{}).
		Where("username = ? or tel = ?", req.Username, req.Tel).
		First(&user)

	if err = isExistDB.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			createDB := DB.
				Model(&models.User{}).
				Create(
					&models.User{
						Username:  req.Username,
						Password:  helper.MakeHash(req.Password),
						Tel:       req.Tel,
						AvatarUrl: req.AvatarUrl,
					},
				).Scan(&user)

			if err = createDB.Error; err != nil {
				return nil, errors.New("注册失败，请稍后重试")
			}

			return &types.RegisterRes{Id: user.Id}, nil
		}

		return resp, err
	}

	if user.Username == req.Username {
		return nil, errors.New("用户名不可用")
	}

	if user.Tel == req.Tel {
		return nil, errors.New("该手机号已经被注册")
	}

	return resp, nil
}
