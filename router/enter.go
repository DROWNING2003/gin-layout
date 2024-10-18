package router

import (
	"github.com/DROWNING2003/gin-layout/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
}
