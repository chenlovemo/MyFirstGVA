package file

import (
	//"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

//type FileManageRouter struct{}

func (e *FileManageRouter) InitFileManageRouter(Router *gin.RouterGroup) {
	fileManageRouter := Router.Group("fileManage") /*.Use(middleware.OperationRecord())*/
	fileManageRouterWithoutRecord := Router.Group("fileManage")
	{
		//fileManageRouter.POST("create",    fileManageApi.CreateFileManage)   // 创建客户
		fileManageRouter.POST("upload", fileManageApi.FileManageUpload)   // 上传文件
		fileManageRouter.PUT("update", fileManageApi.UpdateFileManage)    // 更新客户
		fileManageRouter.DELETE("delete", fileManageApi.DeleteFileManage) // 删除客户
	}
	{
		fileManageRouterWithoutRecord.GET("getsingle", fileManageApi.GetFileManage)    // 获取单一客户信息
		fileManageRouterWithoutRecord.POST("getlist", fileManageApi.GetFileManageList) // 获取客户列表
	}
}
