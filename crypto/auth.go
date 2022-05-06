package crypto

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secret   = os.Getenv("SECRET_KEY") // secret key to encrypt header.payload
	lifetime = 30 * time.Minute        // token's lifetime is set to 30 minutes
)

// Auth は署名前の認証トークン情報を表す。
type Auth struct {
	UserID string
	Iat    int64
}

// JWTの生成
func GenerateToken(userID string, now time.Time) (string, error) {
	// 署名なしトークンの生成：header.payload encorded by Base64
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    userID,
		"issued_at":  now.Unix(),
		"expired_at": now.Add(lifetime).Unix(),
	})

	// 署名付きトークンの生成：header.payload.signature encorded by Base64
	// header.payload -> signature encrypted by HS256 with a secret key
	return token.SignedString([]byte(secret))
}

// JWTの検証と返却
func VerifyToken(signedToken string) (*jwt.Token, error) {
	// 秘密鍵を用いて暗号化したHeader+Payload部とSignature部を比較することで，トークンの整合性を確認
	token, err := jwt.Parse(signedToken, func(parsedToken *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
