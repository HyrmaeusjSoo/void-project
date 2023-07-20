package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// 生成密码加密
func GeneratePassword(password string) (string, error) {
	pwdByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwdByte), err
}

// 对比密码
func ComparePassword(hashPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)) == nil
}
