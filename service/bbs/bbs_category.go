package bbs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BbsCategoryService struct{}

var BbsCategoryServiceApp = new(BbsCategoryService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExaCustomer
//@description: 创建文章
//@param: e model.ExaCustomer
//@return: err error

func (exa *BbsCategoryService) CreateBbsCategory(e bbs.BbsCategory) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除文章
//@param: e model.ExaCustomer
//@return: err error

func (exa *BbsCategoryService) DeleteBbsCategory(e bbs.BbsCategory) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExaCustomer
//@description: 更新文章
//@param: e *model.ExaCustomer
//@return: err error

func (exa *BbsCategoryService) UpdateBbsCategory(e *bbs.BbsCategory) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExaCustomer
//@description: 获取文章信息
//@param: id uint
//@return: customer model.ExaCustomer, err error

func (exa *BbsCategoryService) GetBbsCategory(id uint) (category bbs.BbsCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&category).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取文章列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *BbsCategoryService) GetBbsCategoryInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//db := global.GVA_DB.Model(&example.ExaCustomer{})
	db_detail := global.GVA_DB.Model(&bbs.BbsCategory{})
	db_count := global.GVA_DB.Model(&bbs.BbsCategory{})
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
	var CategoryList []bbs.BbsCategory

	if info.Keyword != "" {
		err = db_detail.Limit(limit).Offset(offset).Where("title like ?", "%"+info.Keyword+"%").Find(&CategoryList).Error
		if err != nil {
			return CategoryList, total, err
		}

		//返回总数
		//comments add by gxc  2025年5月20日15:07:28
		//此处是查询所有记录，但是上面db_detail的查询拼接了offset，导致查询不到count，因此下面新建一个查询，重新查
		err = db_count.Where("title like ?", "%"+info.Keyword+"%").Count(&total).Error

	} else {
		//返回列表数据，写到CategoryList切片
		err = db_detail.Limit(limit).Offset(offset).Where("1=1").Find(&CategoryList).Error
		//返回总数

		//comments add by gxc  2025年5月20日15:07:28
		//此处是查询所有记录，但是上面db_detail的查询拼接了offset，导致查询不到count，因此下面新建一个查询，重新查
		err = db_count.Count(&total).Error
		//total = db.Where("1=1").RowsAffected

		//err = db.Where("1=1").Count(&total).Error
	}

	//err = db.Where("1=1").Count(&total).Error
	//if err != nil {
	//	return CategoryList, total, err
	//} else {
	//	err = db.Limit(limit).Offset(offset).Where("1=1").Find(&CategoryList).Error
	//}
	return CategoryList, total, err
}
