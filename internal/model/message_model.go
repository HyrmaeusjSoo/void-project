package model

import "void-project/internal/model/base"

type Message struct {
	base.Model
	FromId   uint       `gorm:"type:bigint(20);not null" json:"from_id,omitempty"`
	TargetId uint       `gorm:"type:bigint(20);not null" json:"target_id,omitempty"`
	Type     int8       `gorm:"type:int(2);default:1;not null" json:"type,omitempty"`
	Content  string     `gorm:"type:text" json:"content,omitempty"`
	SendTime base.Model `gorm:"type:datetime" json:"send_time,omitempty"`
}

func (u *Message) TableName() string {
	return "message"
}
