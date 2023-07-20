package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

// GenerateLower 加密后返回小写值
func GenerateLower(code string) string {
	m := md5.New()
	io.WriteString(m, code)
	return hex.EncodeToString(m.Sum(nil))
}

// GenerateUpper 加密后返回大写值
func GenerateUpper(code string) string {
	return strings.ToUpper(GenerateLower(code))
}

// SaltPassword 密码加盐
func SaltPassword(pwd string, salt string) string {
	return GenerateLower(fmt.Sprintf("%s$%s", pwd, salt))
}

// CheckPassword 校验密码
func CheckPassword(rpwd, salt, pwd string) bool {
	return pwd == SaltPassword(rpwd, salt)
}
