package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

// Md5encoder 加密后返回小写值
func Md5encoder(code string) string {
	m := md5.New()
	io.WriteString(m, code)
	return hex.EncodeToString(m.Sum(nil))
}

// Md5StrToUpper 加密后返回大写值
func Md5StrToUpper(code string) string {
	return strings.ToUpper(Md5encoder(code))
}

// SaltPassword 密码加盐
func SaltPassword(pwd string, salt string) string {
	saltPwd := fmt.Sprintf("%s$%s", Md5encoder(pwd), salt)
	return saltPwd
}

// CheckPassword 校验密码
func CheckPassword(rpwd, salt, pwd string) bool {
	return pwd == SaltPassword(rpwd, salt)
}
