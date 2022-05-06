package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPwd(pwd string) string {
	hashedPwd := ""
	b := []byte(pwd)
	hash := sha256.Sum256(b)
	hashedPwd = hex.EncodeToString(hash[:])
	return hashedPwd
}

func CompHashPwd(hash, pwd string) bool {
	hashedPwd := HashPwd(pwd)
	if hashedPwd != hash {
		return false
	}
	return true
}
