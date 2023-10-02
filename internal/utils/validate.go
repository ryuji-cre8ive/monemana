package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"
)

func IsValidSignature(signature string, body []byte) bool {
	channelSecret := os.Getenv("LINE_BOT_CHANNEL_SECRET")
	// 署名をBase64デコード
	decodedSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}

	// リクエストボディのダイジェストを計算
	hash := hmac.New(sha256.New, []byte(channelSecret))
	hash.Write(body)
	calculatedSignature := hash.Sum(nil)

	// 署名を比較
	return hmac.Equal(decodedSignature, calculatedSignature)
}
