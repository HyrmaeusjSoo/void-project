package model

import (
	"time"
)

type User struct {
	Model
	Account       string     `gorm:"unique;type:varchar(50);not null;comment:唯一账号id" json:"account,omitempty"`
	Password      string     `gorm:"type:varchar(100);default:null" json:"password,omitempty"`
	Name          string     `gorm:"type:varchar(100);default:null" json:"name,omitempty"`
	Avatar        string     `gorm:"type:varchar(500);default:null" json:"avatar,omitempty"`
	Gender        string     `gorm:"default:male;type:varchar(6);comment:male表示男，famale表示女" json:"gender,omitempty"`
	Phone         string     `gorm:"type:varchar(11);default:null" valid:"matches(^1[3-9]{1}\\d{9}$)" json:"phone,omitempty"`
	Email         string     `gorm:"type:varchar(255);default:null" valid:"email" json:"email,omitempty"`
	Identity      string     `gorm:"type:varchar(255);default:null" json:"identity,omitempty"`
	ClientIp      string     `gorm:"type:varchar(100);default:null" valid:"ipv4" json:"client_ip,omitempty"`
	ClientPort    string     `gorm:"type:varchar(100);default:null" json:"client_port,omitempty"`
	Salt          string     `gorm:"type:varchar(255);default:null" json:"salt,omitempty"`
	LoginTime     *time.Time `json:"login_time,omitempty"`
	HeartBeatTime *time.Time `json:"heart_beat_time,omitempty"`
	LoginOutTime  *time.Time `json:"login_out_time,omitempty"`
	IsLoginOut    bool       `gorm:"type:tinyint(1)" json:"is_login_out,omitempty"`
	DeviceInfo    string     `gorm:"type:varchar(255);default:null" json:"device_info,omitempty"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) SecureClear() {
	u.Password = ""
	u.Identity = ""
	u.Salt = ""
	u.ClientIp = ""
	u.ClientPort = ""
	u.HeartBeatTime = nil
	u.LoginTime = nil
	u.LoginOutTime = nil
	u.DeviceInfo = ""
}
