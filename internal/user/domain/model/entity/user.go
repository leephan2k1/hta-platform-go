package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"column:id;primaryKey;type:text" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:LOCALTIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:LOCALTIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`

	FirstName string `gorm:"column:first_name;type:varchar(255)"`
	LastName  string `gorm:"column:last_name;type:varchar(255)"`
	Email     string `gorm:"column:email;type:varchar(255);unique"`
	Picture   string `gorm:"column:picture;type:text"`
}

// GORM override table name
func (User) TableName() string {
	return "hta.user"
}
