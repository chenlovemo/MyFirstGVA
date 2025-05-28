package bbs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type BbsCategory struct {
	global.GVA_MODEL
	Title       string `gorm:"size:100;not null;default:'';comment:分类名称" json:"title"`
	Description string `gorm:"size:400;not null;default:'';comment:描述" json:"description"`
	Parent_id   uint   `gorm:"not null;default:0;comment:父ID" json:"parentId"`
	Sorted      int8   `gorm:"not null;default:1;comment:排序" json:"sorted"`
	Status      int8   `gorm:"size:1;not null;default:1;comment:0 未发布 1 发布" json:"status"`
	IsDelete    int8   `gorm:"size:1;not null;default:0;comment:0 未删除 1 删除" json:"isDelete" form:"isDelete" `
}

func (BbsCategory) TableName() string {
	return "bbs_category"
}
