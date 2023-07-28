package model

import "void-project/internal/model/base"

// 用户
type User struct {
	base.Model
	Account       string     `gorm:"unique;type:varchar(50);not null;comment:唯一账号id" json:"account,omitempty"`
	Password      string     `gorm:"type:varchar(100)" json:"password,omitempty"`
	Name          string     `gorm:"type:varchar(100)" json:"name,omitempty"`
	Avatar        *string    `gorm:"type:varchar(500);default:null" json:"avatar,omitempty"`
	Gender        string     `gorm:"type:varchar(6);default:male;comment:male表示男，famale表示女" json:"gender,omitempty"`
	Phone         *string    `gorm:"type:varchar(20);default:null" valid:"matches(^1[3-9]{1}\\d{9}$)" json:"phone,omitempty"`
	Email         *string    `gorm:"type:varchar(255);default:null" valid:"email" json:"email,omitempty"`
	Identity      *string    `gorm:"type:varchar(255);default:null" json:"identity,omitempty"`
	ClientIp      *string    `gorm:"type:varchar(100);default:null" valid:"ipv4" json:"client_ip,omitempty"`
	ClientPort    *string    `gorm:"type:varchar(100);default:null" json:"client_port,omitempty"`
	Salt          *string    `gorm:"type:varchar(255);default:null" json:"salt,omitempty"`
	LoginTime     *base.Time `gorm:"type:datetime;default:null" json:"login_time,omitempty"`
	HeartBeatTime *base.Time `gorm:"type:datetime;default:null" json:"heart_beat_time,omitempty"`
	LoginOutTime  *base.Time `gorm:"type:datetime;default:null" json:"login_out_time,omitempty"`
	IsLoginOut    bool       `gorm:"type:tinyint(1)" json:"is_login_out,omitempty"`
	DeviceInfo    *string    `gorm:"type:varchar(255);default:null" json:"device_info,omitempty"`
}

func (u *User) TableName() string {
	return "user"
}

// 敏感信息字段清除
func (u *User) SecureClear() {
	u.Password = ""
	u.Identity = nil
	u.Salt = nil
	u.ClientIp = nil
	u.ClientPort = nil
	u.HeartBeatTime = nil
	u.LoginTime = nil
	u.LoginOutTime = nil
	u.DeviceInfo = nil
}
