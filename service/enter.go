package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/file"
	"github.com/flipped-aurora/gin-vue-admin/server/service/gxcuser"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	GxcuserServiceGroup gxcuser.ServiceGroup
	BbsServiceGroup     bbs.BbsServiceGroup
	FileManageGroup     file.FileManageServiceGroup
	//ArticleServiceGroup bbs.ServiceGroup
	//bbsServiceGroup bbs.ServiceGroup
}
