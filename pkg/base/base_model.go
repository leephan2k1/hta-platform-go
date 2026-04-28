package base

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:LOCALTIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:LOCALTIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
}

type AuditModel struct {
	CreatedBy string `gorm:"column:created_by" json:"created_by"`
	UpdatedBy string `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy string `gorm:"column:deleted_by" json:"deleted_by"`
}

type CommonModel struct {
	BaseModel
	AuditModel
}
