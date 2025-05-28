package file

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/file"
	fileRes "github.com/flipped-aurora/gin-vue-admin/server/model/file/response"
	upload "github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"time"

	//"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	//exampleRes "github.com/flipped-aurora/gin-vue-admin/server/model/example/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileManageApi struct{}

func (e *FileManageApi) FileManageUpload(c *gin.Context) {

	//noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	//classId, _ := strconv.Atoi(c.DefaultPostForm("classId", "0"))
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}

	switch global.GVA_CONFIG.System.OssType {
	case "local":
		e.uploadLocal(c, header)

	case "minio":
		e.uploadMinio(c, header)

	default:
		e.uploadLocal(c, header)

	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", header.Filename))

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
func (e *FileManageApi) DeleteFileManage(c *gin.Context) {
	var fileManage file.FileManageStruct
	err := c.ShouldBindJSON(&fileManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(fileManage.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = fileManageService.DeleteFileManage(fileManage)
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
func (e *FileManageApi) UpdateFileManage(c *gin.Context) {
	var fileManage file.FileManageStruct
	err := c.ShouldBindJSON(&fileManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(fileManage.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(fileManage, utils.CustomerVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileManage.CreatedAt = time.Now()
	err = fileManageService.UpdateFileManage(&fileManage)
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
func (e *FileManageApi) GetFileManage(c *gin.Context) {
	var fileManage file.FileManageStruct
	err := c.ShouldBindQuery(&fileManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if fileManage.ID == 0 {
		fileManage.ID = 1
	}

	err = utils.Verify(fileManage.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := fileManageService.GetFileManage(fileManage.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	//response.OkWithDetailed(exampleRes.ExaCustomerResponse{FileManage: data}, "获取成功", c)
	response.OkWithDetailed(fileRes.FileManageResponse{FileManage: data}, "获取成功", c)

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
func (e *FileManageApi) GetFileManageList(c *gin.Context) {
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
	customerList, total, err := fileManageService.GetFileManageInfoList(utils.GetUserAuthorityId(c), pageInfo)
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

func (e *FileManageApi) uploadLocal(c *gin.Context, header *multipart.FileHeader) {
	//var file example.ExaFileUploadAndDownload
	var fileManage file.FileManageStruct
	// 单文件
	//comments add by gxc 2025年5月24日22:33:15
	//此处 file必须是file四个字母。这个是根据前台传递进来的参数决定的；
	//下面   c.Request.FormFile("file")  和 c.FormFile("file") 效果一样。file对象里面有文件信息（Filename记录文件名，Header里面记录文件文件类型，也包含文件名），
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	//_, header, _ := c.Request.FormFile("file")
	//fmt.Println(header)

	// 读取文件后缀
	ext := filepath.Ext(file.Filename)
	fmt.Println("ext=", ext)

	// 读取文件名并加密
	name_origin := strings.TrimSuffix(file.Filename, ext)
	name_MD5 := utils.MD5V([]byte(name_origin))
	// 拼接新文件名
	filename := name_MD5 + "_" + name_origin + "_" + time.Now().Format("20060102150405") + ext

	// 拼接路径和文件名   StorePath是临时目录，暂时不用   Local.Path是实际存储目录，这俩是yaml配置文件里面配置
	//p := global.GVA_CONFIG.Local.StorePath + "/" + filename
	fullFilepath := global.GVA_CONFIG.Local.Path + "/" + filename

	dst := fullFilepath
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	response.OkWithDetailed(fileRes.FileManageResponse{FileManage: fileManage}, "上传成功", c)
}
func (e *FileManageApi) uploadMinio(c *gin.Context, header *multipart.FileHeader) (string, string, error) {

	MinioClt, err := upload.GetMinio(global.GVA_CONFIG.Minio.Endpoint, global.GVA_CONFIG.Minio.AccessKeyId, global.GVA_CONFIG.Minio.AccessKeySecret, global.GVA_CONFIG.Minio.BucketName, global.GVA_CONFIG.Minio.UseSSL)
	if err != nil {
		global.GVA_LOG.Warn("你配置了使用minio，但是初始化失败，请检查minio可用性或安全配置: " + err.Error())
		panic("minio初始化失败") // 建议这样做，用户自己配置了minio，如果报错了还要把服务开起来，使用起来也很危险
	}

	f, openError := header.Open()
	// mutipart.File to os.File
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}

	filecontent := bytes.Buffer{}
	_, err = io.Copy(&filecontent, f)
	if err != nil {
		global.GVA_LOG.Error("读取文件失败", zap.Any("err", err.Error()))
		return "", "", errors.New("读取文件失败, err:" + err.Error())
	}
	f.Close() // 创建文件 defer 关闭

	// 对文件名进行加密存储
	ext := filepath.Ext(header.Filename)
	filename := utils.MD5V([]byte(strings.TrimSuffix(header.Filename, ext))) + ext
	var filePathres string
	if global.GVA_CONFIG.Minio.BasePath == "" {
		filePathres = "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + filename
	} else {
		filePathres = global.GVA_CONFIG.Minio.BasePath + "/" + time.Now().Format("2006-01-02") + "/" + filename
	}

	// 设置超时10分钟
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	// Upload the file with PutObject   大文件自动切换为分片上传
	info, err := MinioClt.Client.PutObject(ctx, global.GVA_CONFIG.Minio.BucketName, filePathres, &filecontent, header.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		global.GVA_LOG.Error("上传文件到minio失败", zap.Any("err", err.Error()))
		return "", "", errors.New("上传文件到minio失败, err:" + err.Error())
	}
	return global.GVA_CONFIG.Minio.BucketUrl + "/" + info.Key, filePathres, nil

}
