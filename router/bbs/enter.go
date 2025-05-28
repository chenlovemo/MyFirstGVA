package bbs

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type CategoryRouter struct{}
type ArticleRouter struct{}

type RouterGroup struct {
	ArticleRouter
	CategoryRouter
}

var (
	//exaCustomerApi= api.ApiGroupApp.ExampleApiGroup.CustomerApi
	bbsArticleApi  = api.ApiGroupApp.BbsApiGroup.ArticleApi
	bbsCategoryApi = api.ApiGroupApp.BbsApiGroup.CategoryApi
)
