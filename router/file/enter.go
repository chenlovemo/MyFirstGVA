package file

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type FileManageRouter struct{}

type RouterGroup struct {
	FileManageRouter
}

var (
	//bbsCategoryApi = api.ApiGroupApp.BbsApiGroup.CategoryApi
	fileManageApi = api.ApiGroupApp.FileManageGroup.FileManageApi
)
