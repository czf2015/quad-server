package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"goserver/libs/conf"
)

var appKey = []byte(conf.GetSectionKey("app", "APP_KEY").String())

func EncryptPassword(password string) string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, appKey)
	// Write Data to it
	h.Write([]byte(password))
	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}
