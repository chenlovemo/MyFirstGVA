package file

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	//ArticleApi
	FileManageApi
}

var (
	//customerService = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	//articleService  = service.ServiceGroupApp.BbsServiceGroup.BbsArticleService
	//categoryService = service.ServiceGroupApp.BbsServiceGroup.BbsCategoryService
	fileManageService = service.ServiceGroupApp.FileManageGroup.FileManageService
)
