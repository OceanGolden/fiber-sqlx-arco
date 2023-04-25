package auth

type LoginRequest struct {
	Username string `zh:"用户名称" json:"username" validate:"required"`
	Password string `zh:"用户密码" json:"password" validate:"required"`
}

type LoginResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Access   string `json:"access"`
	Refresh  string `json:"refresh"`
	//ExpiredAt int64  `json:"expired_at"`
}
