// Code generated by goctl. DO NOT EDIT.
package types

type AddReq struct {
	Book  string `form:"book"`
	Price int64  `form:"price"`
}

type AddResp struct {
	Ok bool `json:"ok"`
}

type CheckReq struct {
	Book string `form:"book"`
}

type CheckResp struct {
	Found bool  `json:"found"`
	Price int64 `json:"price"`
}

type JwtTokenRequest struct {
}

type JwtTokenResponse struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"` // 建议客户端刷新token的绝对时间
}

type GetUserRequest struct {
	UserId string `json:"userId"`
}

type GetUserResponse struct {
	Name string `json:"name"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
}

type RegisterResp struct {
	Msg string `json:"msg"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Username     string `json:"Username"`
	Age          int    `json:"age"`
	Gender       string `json:"gender"`
	Token        string `json:"token"`
	ExpireTime   int64  `json:"expire_time"`
	RefreshAfter int64  `json:"refreshAfter"`
}