package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	//ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID
	//comments add by gxc 2025年5月14日22:07:10 此处 指定json的目的是为了当数据以json的形式返回到前端的时候，id是小写，前端可以通过小写的id去读取
	ID        uint           `gorm:"primarykey" json:"ID" form:"ID" ` // 主键ID ，
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间

}
