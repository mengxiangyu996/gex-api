package jwt

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"gex-api/config"
	"strings"
	"time"
)

type Payload struct {
	Id     int
	Expire time.Time
}

// 加密
func Generate(payload *Payload) string {

	payloadByte, _ := json.Marshal(payload)

	return base64.StdEncoding.EncodeToString([]byte(string(payloadByte) + signature()))
}

// 解析
func Parse(token string) *Payload {

	var payload *Payload

	tokenByte, _ := base64.StdEncoding.DecodeString(token)

	json.Unmarshal([]byte(strings.ReplaceAll(string(tokenByte), signature(), "")), &payload)

	return payload
}

// 签名
func signature() string {

	appKey := md5.Sum([]byte(config.App.Key))

	return hex.EncodeToString(appKey[:])
}