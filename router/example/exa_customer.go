package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("customer")
	{
		/*comments add by gxc  2025年5月24日11:46:41  这里是通过路由结构体（customerRouter）的POST方法来调用具体的结构体（exaCustomerApi）方法
		CreateExaCustomer 。这里面的CreateExaCustomer是结构体方法的名字，通过结构体方法的名字来调用。

		*/
		customerRouter.POST("customer", exaCustomerApi.CreateExaCustomer)   // 创建客户
		customerRouter.PUT("customer", exaCustomerApi.UpdateExaCustomer)    // 更新客户
		customerRouter.DELETE("customer", exaCustomerApi.DeleteExaCustomer) // 删除客户
	}
	{
		customerRouterWithoutRecord.GET("customer", exaCustomerApi.GetExaCustomer)          // 获取单一客户信息
		customerRouterWithoutRecord.GET("customerList?", exaCustomerApi.GetExaCustomerList) // 获取客户列表
	}
}
