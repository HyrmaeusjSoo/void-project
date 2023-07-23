package model

import (
	"void-project/internal/model/base"
)

type Message struct {
	base.Model
	FromId     uint      `gorm:"type:bigint(20);not null" json:"from_id"`
	From       *User     `gorm:"foreignKey:from_id;references:id" json:"from,omitempty"`
	FromName   string    `gorm:"-:migration;<-:false" json:"from_name"`
	TargetId   uint      `gorm:"type:bigint(20);not null" json:"target_id"`
	Target     *User     `gorm:"foreignKey:target_id;references:id" json:"target,omitempty"`
	TargetName string    `gorm:"-:migration;<-:false" json:"target_name"`
	Type       int8      `gorm:"type:int(2);default:1;not null" json:"type,omitempty"`
	Content    string    `gorm:"type:text" json:"content,omitempty"`
	SendTime   base.Time `gorm:"type:datetime" json:"send_time,omitempty"`
}

func (u *Message) TableName() string {
	return "message"
}
