package service

import (
	"github.com/DROWNING2003/gin-layout/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
}
