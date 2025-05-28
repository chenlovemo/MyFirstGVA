package bbs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
	bbsRes "github.com/flipped-aurora/gin-vue-admin/server/model/bbs/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"time"

	//"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	//exampleRes "github.com/flipped-aurora/gin-vue-admin/server/model/example/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryApi struct{}

// CreateExaCustomer
// @Tags      ExaCustomer
// @Summary   创建文章
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "文章用户名, 文章手机号码"
// @Success   200   {object}  response.Response{msg=string}  "创建文章"
// @Router    /customer/customer [post]
func (e *CategoryApi) CreateBbsCategory(c *gin.Context) {
	var category bbs.BbsCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(category, utils.CustomerVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//category.UserId = utils.GetUserID(c)
	//category..SysUserAuthorityID = utils.GetUserAuthorityId(c)
	err = categoryService.CreateBbsCategory(category)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteExaCustomer
// @Tags      ExaCustomer
// @Summary   删除文章
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "文章ID"
// @Success   200   {object}  response.Response{msg=string}  "删除文章"
// @Router    /customer/customer [delete]
func (e *CategoryApi) DeleteBbsCategory(c *gin.Context) {
	var category bbs.BbsCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(category.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = categoryService.DeleteBbsCategory(category)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateExaCustomer
// @Tags      ExaCustomer
// @Summary   更新文章信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "文章ID, 文章信息"
// @Success   200   {object}  response.Response{msg=string}  "更新文章信息"
// @Router    /customer/customer [put]
func (e *CategoryApi) UpdateBbsCategory(c *gin.Context) {
	var category bbs.BbsCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(category.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(category, utils.CustomerVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	category.CreatedAt = time.Now()
	err = categoryService.UpdateBbsCategory(&category)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetExaCustomer
// @Tags      ExaCustomer
// @Summary   获取单一文章信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     example.ExaCustomer                                                true  "文章ID"
// @Success   200   {object}  response.Response{data=exampleRes.ExaCustomerResponse,msg=string}  "获取单一文章信息,返回包括文章详情"
// @Router    /customer/customer [get]
func (e *CategoryApi) GetBbsCategory(c *gin.Context) {
	var category bbs.BbsCategory
	err := c.ShouldBindQuery(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if category.ID == 0 {
		category.ID = 1
	}

	err = utils.Verify(category.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := categoryService.GetBbsCategory(category.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	//response.OkWithDetailed(exampleRes.ExaCustomerResponse{Category: data}, "获取成功", c)
	response.OkWithDetailed(bbsRes.BbsResponse{Category: data}, "获取成功", c)

}

// GetExaCustomerList
// @Tags      ExaCustomer
// @Summary   分页获取权限文章列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限文章列表,返回包括列表,总数,页码,每页数量"
// @Router    /customer/customerList [get]
func (e *CategoryApi) GetBbsCategoryList(c *gin.Context) {
	var pageInfo request.PageInfo

	/*
			comments add by gxc 2025年5月19日16:05:40
		对于一个像这样的请求体 {"name":"John", "age":30}，ShouldBindJSON 它将会被正确地解析并绑定到 Person 结构体
		如果请求为 GET /?name=John&age=30，ShouldBindQuery 则会成功解析并绑定到 Person 结构体
		BindJson是处理json，BindQqery是处理url参数的
	*/

	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	////如果前台不传Pagesize和Page信息那么给默认值 ,前台没有传进来吗？？？？
	//if pageInfo.PageSize == 0 {
	//	pageInfo.PageSize = 2
	//}
	//if pageInfo.Page == 0 {
	//	pageInfo.Page = 1
	//}

	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	customerList, total, err := categoryService.GetBbsCategoryInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     customerList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
