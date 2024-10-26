package logic

import (
	"context"
	"errors"
	"strings"
	"time"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	"bookstore/api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("用户名和密码不为空")
	}

	// 插入数据
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Name:       req.Username,
		Password:   req.Password,
		Age:        int64(req.Age),
		Gender:     req.Gender,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		Msg: "用户注册成功",
	}, nil

}
