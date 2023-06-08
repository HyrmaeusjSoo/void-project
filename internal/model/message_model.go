package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromId   uint      `gorm:"type:bigint(20);not null" json:"fromId,omitempty"`
	TargetId uint      `gorm:"type:bigint(20);not null" json:"targetId,omitempty"`
	Type     int8      `gorm:"type:int(2);default:1;not null" json:"type,omitempty"`
	Content  string    `gorm:"type:text" json:"content,omitempty"`
	SendTime time.Time `gorm:"type:datetime" json:"sendTime,omitempty"`
}
