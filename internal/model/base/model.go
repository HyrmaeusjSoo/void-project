package base

import (
	"gorm.io/gorm"
)

type Model struct {
	ID        uint            `gorm:"type:int(11);primaryKey;autoIncrement" json:"id,omitempty"`
	CreatedAt Time            `gorm:"type:datetime;not null;autoCreateTime" json:"create_at,omitempty"`
	UpdatedAt Time            `gorm:"type:datetime;not null;autoUpdateTime" json:"update_at,omitempty"`
	DeletedAt *gorm.DeletedAt `gorm:"type:datetime;index" json:"delete_at,omitempty"`
}
