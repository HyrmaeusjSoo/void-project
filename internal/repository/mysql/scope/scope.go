package scope

import (
	"gorm.io/gorm"
)

// 排除User敏感信息
func OmitUserSensitivity(db *gorm.DB) *gorm.DB {
	return db.Omit("password", "identity", "salt", "client_ip", "client_port",
		"heart_beat_time", "login_time", "login_out_time", "device_info")
}
