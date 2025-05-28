package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/file"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/gxcuser"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	GxcuserApiGroup gxcuser.ApiGroup

	BbsApiGroup     bbs.ApiGroup
	FileManageGroup file.ApiGroup
}
