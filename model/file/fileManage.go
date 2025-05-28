package file

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type FileManageStruct struct {
	global.GVA_MODEL
	Name        string `gorm:"size:100;not null;default:'';comment:文件名" json:"name"`
	Description string `gorm:"size:400;not null;default:'';comment:文件描述" json:"description"`
	IsDelete    int8   `gorm:"size:1;not null;default:0;comment:0 未删除 1 删除" json:"isDelete" form:"isDelete" `
}

func (FileManageStruct) TableName() string {
	return "file_manage"
}
