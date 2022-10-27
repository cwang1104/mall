// Code generated by goctl. DO NOT EDIT.
package types

type SendEmailRequest struct {
	Email string `form:"email"`
}

type SendEmailResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type RegisterRequest struct {
	Email      string `form:"email"`
	Captche    string `form:"catpche"`
	Password   string `form:"password"`
	Repassword string `form:"repassword"`
}

type CommonResp struct {
	Msg  string `json:"msg"`
	Code int32  `json:"code"`
}

type LoginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type LoginResponse struct {
	Code         int32  `json:"code"`
	Msg          string `json:"msg"`
	Token        string `json:"token"`
	UserName     string `json:"username"`
	AccessExpire int64  `json:"accessExpire"`
}

type AdminLoginReq struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

type AdminLoginResponse struct {
	Code         int32  `json:"code"`
	Msg          string `json:"msg"`
	AdminToken   string `json:"admin_token"`
	UserName     string `json:"user_name"`
	AccessExpire int64  `json:"accessExpire"`
}

type GetUserListReq struct {
	CurrentPage int32 `form:"currentPage,default=1"`
	PageSize    int32 `form:"pageSize,default=10"`
}

type GetUserListResponse struct {
	Code        int32  `json:"code"`
	Msg         string `json:"msg"`
	Total       int32  `json:"total"`
	CurrentPage int32  `json:"current_page"`
	PageSize    int32  `json:"page_size"`
}
