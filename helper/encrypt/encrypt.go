package encrypt

import (
	"breeze-api/config"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

// 加密
func Generate(password string) string {

	passwordByte, _ := json.Marshal(password)

	return base64.StdEncoding.EncodeToString([]byte(string(passwordByte) + signature()))
}

// 比较
func Compare(hashPassword, password string) bool {

	passwordByte, _ := json.Marshal(password)

	return base64.StdEncoding.EncodeToString([]byte(string(passwordByte) + signature())) == hashPassword
}

// 签名
func signature() string {

	appKey := md5.Sum([]byte(config.App.Key))

	return hex.EncodeToString(appKey[:])
}
