package jwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenSignatureInvalid = jwt.ErrTokenSignatureInvalid
	ErrTokenExpired          = jwt.ErrTokenExpired
	ErrMissingTokenExpire    = errors.New("missing token expire")
	ErrMissingSecret         = errors.New("missing secret")

	TokenExpire time.Duration // token过期时间
	Secret      string        // 密钥
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// 初始化Jwt
//
//	tokenExpire => token过期时间（小时）
//	secret => 密钥
func InitJwt(tokenExpire time.Duration, secret string) {
	TokenExpire = tokenExpire
	Secret = secret
}

// 生成标准Claims模式的Token
func GenerateToken(userId uint) (string, error) {
	if TokenExpire == 0 {
		return "", ErrMissingTokenExpire
	}
	if Secret == "" {
		return "", ErrMissingSecret
	}
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "HyrmaeusjSoo",
			Subject:   "void-project",
			Audience:  jwt.ClaimStrings{strconv.Itoa(int(userId))},
			ExpiresAt: jwt.NewNumericDate(now.Add(TokenExpire)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})
	return t.SignedString([]byte(Secret))
}

// 转换Token为标准Claims
func ParseToken(token string) (*Claims, error) {
	if Secret == "" {
		return nil, ErrMissingSecret
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(Secret), nil
	})

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, err
}
