package global

import (
	"gorm.io/gorm"
	"time"
)

// GpModel 统一表结构字段
type GpModel struct {
	ID        string         `json:"id" gorm:"size:32;primary_key;comment:'主键'"`
	CreatedAt time.Time      `json:"created_at" gorm:"comment:'创建时间'"`
	CreatedBy string         `json:"created_by" gorm:"size:32;comment:'创建实体'"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"comment:'更新时间'"`
	UpdatedBy string         `json:"updated_by" gorm:"size:32;comment:'更新实体'"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index:idx_deleted;comment:'删除时间'"`
	DeletedBy string         `json:"deleted_by" gorm:"size:32;comment:'删除实体'"`
}
