package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey;comment:主键ID" json:"id" form:"id"`    // 主键ID
	CreatedAt time.Time      `gorm:"AUTO_CREATED_AT;comment :创建时间" json:"createdAt"` // 创建时间
	UpdatedAt time.Time      `gorm:"AUTO_UPDATED_AT;comment :更新时间" json:"updatedAt"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index;comment :删除时间" json:"-"`                   // 删除时间
}

//ID 是用来给当前的Go程序去使用和调用的，大写的原因是私有变量
// 请求入参 会映射 model 的属性
// form 参数是用来约束入参(查询)的格式的 json 是用来约束出参的格式的
