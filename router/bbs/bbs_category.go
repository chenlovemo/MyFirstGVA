package bbs

import (
	//"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

//type CategoryRouter struct{}

func (e *CategoryRouter) InitCategoryRouter(Router *gin.RouterGroup) {
	categoryRouter := Router.Group("category") /*.Use(middleware.OperationRecord())*/
	categoryRouterWithoutRecord := Router.Group("category")
	{
		categoryRouter.POST("create", bbsCategoryApi.CreateBbsCategory)   // 创建客户
		categoryRouter.PUT("update", bbsCategoryApi.UpdateBbsCategory)    // 更新客户
		categoryRouter.DELETE("delete", bbsCategoryApi.DeleteBbsCategory) // 删除客户
	}
	{
		categoryRouterWithoutRecord.GET("getsingle", bbsCategoryApi.GetBbsCategory)    // 获取单一客户信息
		categoryRouterWithoutRecord.POST("getlist", bbsCategoryApi.GetBbsCategoryList) // 获取客户列表
	}
}
