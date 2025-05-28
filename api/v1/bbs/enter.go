package bbs

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ArticleApi
	CategoryApi
}

var (
	//customerService = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	articleService  = service.ServiceGroupApp.BbsServiceGroup.BbsArticleService
	categoryService = service.ServiceGroupApp.BbsServiceGroup.BbsCategoryService
)
