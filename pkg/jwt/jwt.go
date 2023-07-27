package jwt

import (
	"strconv"
	"time"
	"void-project/global"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenSignatureInvalid = jwt.ErrTokenSignatureInvalid
	ErrTokenExpired          = jwt.ErrTokenExpired
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uint) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "HyrmaeusjSoo",
			Subject:   "void-project",
			Audience:  jwt.ClaimStrings{strconv.Itoa(int(userId))},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * global.Config.System.AuthTokenExpire)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	return t.SignedString([]byte(global.Config.System.AuthJwtSecret))
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(global.Config.System.AuthJwtSecret), nil
	})

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, err
}
