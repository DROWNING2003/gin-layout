package system

import "github.com/DROWNING2003/gin-layout/service"

type ApiGroup struct {
	JwtApi
	BaseApi
}

var (
	jwtService              = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
)
