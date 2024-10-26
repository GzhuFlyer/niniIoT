package logic

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GetUserMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMsgLogic {
	return &GetUserMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMsgLogic) GetUserMsg(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("用户名和密码不为空")
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByName(l.ctx, req.Username)
	fmt.Println(userInfo)
	switch err {
	case nil:
	case sqlx.ErrNotFound:
		return nil, errors.New("用户不存在")
	default:
		return nil, err
	}
	if req.Password != userInfo.Password {
		return nil, errors.New("密码错误")
	}

	// jwt
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	token, err := l.getJwtToken(accessSecret,
		now,
		accessExpire,
		userInfo.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Username:     userInfo.Name,
		Age:          int(userInfo.Age),
		Gender:       userInfo.Gender,
		Token:        token,
		ExpireTime:   now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

// 获取token
func (l *GetUserMsgLogic) getJwtToken(key string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(key))
}
