package crypto

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	secret   = ""               // secret key to encrypt header.payload
	lifetime = 30 * time.Minute // token's lifetime is set to 30 minutes
)

// Auth は署名前の認証トークン情報を表す。
type Auth struct {
	UserID string
	Iat    int64
}

// JWTの生成
func GenerateToken(userID string, now time.Time) (string, error) {
	secret, err := readSecretKey(".secret.key")
	if err != nil {
		fmt.Println("Couldn't read secret key.: ", err)
		return "", err
	}

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
	secret, err := readSecretKey(".secret.key")
	if err != nil {
		fmt.Println("Couldn't read secret key.: ", err)
		return nil, err
	}

	// 秘密鍵を用いて暗号化したHeader+Payload部とSignature部を比較することで，トークンの整合性を確認
	token, err := jwt.Parse(signedToken, func(parsedToken *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// keyファイルから秘密鍵を取得
func readSecretKey(filePath string) (string, error) {
	keyStr := ""

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// ファイルを1行ごとに読み込み
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		// 開始行と終了行以外を対象に文字列を取得し連結
		if text == "-----BEGIN SECRET KEY-----" || text == "-----END SECRET KEY-----" {
			continue
		}
		keyStr = keyStr + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return keyStr, nil
}
