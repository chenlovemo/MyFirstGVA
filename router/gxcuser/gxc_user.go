package gxcuser

import (
	"github.com/gin-gonic/gin"
)

type GxcUserRouter struct{}

func (s *GxcUserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("gxcuser")
	//.Use(middleware.OperationRecord()

	userRouterWithoutRecord := Router.Group("gxcuser")
	{
		userRouter.POST("gxc_admin_register", gxcBaseApi.Register)       // 管理员注册账号
		userRouter.POST("gxc_changePassword", gxcBaseApi.ChangePassword) // 用户修改密码
		//userRouter.POST("gxc_setUserAuthority", gxcBaseApi.SetUserAuthority)     // 设置用户权限
		userRouter.DELETE("gxc_deleteUser", gxcBaseApi.DeleteUser) // 删除用户
		userRouter.PUT("gxc_setUserInfo", gxcBaseApi.SetUserInfo)  // 设置用户信息
		userRouter.PUT("gxc_setSelfInfo", gxcBaseApi.SetSelfInfo)  // 设置自身信息
		//userRouter.POST("gxc_setUserAuthorities", gxcBaseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.POST("gxc_resetPassword", gxcBaseApi.ResetPassword)  // 设置用户权限组
		userRouter.PUT("gxc_setSelfSetting", gxcBaseApi.SetSelfSetting) // 用户界面配置
	}
	{
		userRouterWithoutRecord.POST("gxc_getUserList", gxcBaseApi.GetUserList) // 分页获取用户列表
		userRouterWithoutRecord.GET("gxc_getUserInfo", gxcBaseApi.GetUserInfo)  // 获取自身信息

	}
}
