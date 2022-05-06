package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// 受け取ったパスワードをハッシュ化
func HashPwd(pwd string) string {
	hashedPwd := ""
	b := []byte(pwd)
	hash := sha256.Sum256(b)
	hashedPwd = hex.EncodeToString(hash[:])
	return hashedPwd
}

// 受け取った平文をハッシュ化してデータベースのハッシュ値と照合
func CompHashPwd(hash, pwd string) bool {
	hashedPwd := HashPwd(pwd)
	if hashedPwd != hash {
		return false
	}
	return true
}
