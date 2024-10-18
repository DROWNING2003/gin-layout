package system

import (
	"github.com/DROWNING2003/gin-layout/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
