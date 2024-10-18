package v1

import (
	"github.com/DROWNING2003/gin-layout/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
}
