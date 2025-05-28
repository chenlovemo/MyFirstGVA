package bbs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BbsArticleService struct{}

var BbsArticleServiceApp = new(BbsArticleService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaCustomer
//@description: 创建文章
//@param: e model.ExaCustomer
//@return: err error

func (exa *BbsArticleService) CreateBbsArticle(e bbs.BbsArticle) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除文章
//@param: e model.ExaCustomer
//@return: err error

func (exa *BbsArticleService) DeleteBbsArticle(e bbs.BbsArticle) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExaCustomer
//@description: 更新文章
//@param: e *model.ExaCustomer
//@return: err error

func (exa *BbsArticleService) UpdateBbsArticle(e *bbs.BbsArticle) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaCustomer
//@description: 获取文章信息
//@param: id uint
//@return: customer model.ExaCustomer, err error

func (exa *BbsArticleService) GetBbsArticle(id uint) (article bbs.BbsArticle, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&article).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取文章列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *BbsArticleService) GetBbsArticleInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//db := global.GVA_DB.Model(&example.ExaCustomer{})
	db := global.GVA_DB.Model(&bbs.BbsArticle{})
	//var a system.SysAuthority
	//a.AuthorityId = sysUserAuthorityID
	//auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	//if err != nil {
	//	return
	//}
	//var dataId []uint
	//for _, v := range auth.DataAuthorityId {
	//	dataId = append(dataId, v.AuthorityId)
	//}
	var ArticleList []bbs.BbsArticle
	err = db.Where("1=1").Count(&total).Error
	if err != nil {
		return ArticleList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Where("1=1").Find(&ArticleList).Error
	}
	return ArticleList, total, err
}
