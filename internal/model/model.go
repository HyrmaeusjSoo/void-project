package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint            `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt *time.Time      `json:"create_at,omitempty"`
	UpdatedAt *time.Time      `json:"update_at,omitempty"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"delete_at,omitempty"`
}
