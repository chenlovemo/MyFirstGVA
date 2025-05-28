package bbs

import (
	//"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

//type ArticleRouter struct{}

func (e *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	articleRouter := Router.Group("article") /*.Use(middleware.OperationRecord())*/
	articleRouterWithoutRecord := Router.Group("article")
	{
		articleRouter.POST("article_create", bbsArticleApi.CreateBbsArticle)   // 创建客户
		articleRouter.PUT("article_update", bbsArticleApi.UpdateBbsArticle)    // 更新客户
		articleRouter.DELETE("article_delete", bbsArticleApi.DeleteBbsArticle) // 删除客户
	}
	{
		articleRouterWithoutRecord.GET("article_getsingle", bbsArticleApi.GetBbsArticle)   // 获取单一客户信息
		articleRouterWithoutRecord.GET("article_getlist", bbsArticleApi.GetBbsArticleList) // 获取客户列表
	}
}
