package crypto

import (
	"go/token"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	secret   = "secret"         // 署名時の秘密鍵
	lifetime = 30 * time.Minute // トークンの有効期限は30分
)

func GenerateToken(userID string, now time.Time) (string, error) {
	// 署名なしトークンの生成：header.payload
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user_id":    userID,
		"issued_at":  now.Unix(),
		"expired_at": now.Add(lifetime).Unix(),
	})

	// 署名付きトークンの生成：header.payload.secret
	return token.SignedString([]byte(secret))
}

func VerifyToken(signedToken string) bool {
	token, err := jwt.Parse(signedToken, func)
	return true
}
