package system

import api "github.com/DROWNING2003/gin-layout/api/v1"

type RouterGroup struct {
	JwtRouter
	BaseRouter
}

var (
	jwtApi              = api.ApiGroupApp.SystemApiGroup.JwtApi
	baseApi             = api.ApiGroupApp.SystemApiGroup.BaseApi
)
