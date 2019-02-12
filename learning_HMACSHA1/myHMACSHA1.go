package myHMACSHA1

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func GetStringFromHMAC(mac hash.hash) string {
	return hex.EncodeToString(mac.Sum(nil))
}

func GetHMACSHA1(secret, message []byte) hash.hash {

	key := []byte("secret")
	message := []byte("message")
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	return mac
}
