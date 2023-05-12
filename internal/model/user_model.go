package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account       string `gorm:"unique;type:varchar(50);not null;comment:唯一账号id;default:null"`
	Password      string `gorm:"type:varchar(100);default:null"`
	Name          string `gorm:"type:varchar(100);default:null"`
	Avatar        string `gorm:"type:varchar(500);default:null"`
	Gender        string `gorm:"default:male;type:varchar(6);comment:male表示男，famale表示女"`
	Phone         string `gorm:"type:varchar(11);default:null" valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `gorm:"type:varchar(255);default:null" valid:"email"`
	Identity      string `gorm:"type:varchar(255);default:null"`
	ClientIp      string `gorm:"type:varchar(100);default:null" valid:"ipv4"`
	ClientPort    string `gorm:"type:varchar(100);default:null"`
	Salt          string `gorm:"type:varchar(255);default:null"`
	LoginTime     *time.Time
	HeartBeatTime *time.Time
	LoginOutTime  *time.Time
	IsLoginOut    bool   `gorm:"type:tinyint(1)"`
	DeviceInfo    string `gorm:"type:varchar(255);default:null"`
}

func (u *User) TableName() string {
	return "user"
}
