package crypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func Sha256(data string) string {
	t := sha256.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func MD5(byteMessage []byte) string {
	h := md5.New()
	h.Write(byteMessage)
	return hex.EncodeToString(h.Sum(nil))
}

func HamSha256(data string, key []byte) string {
	hmac := hmac.New(sha256.New, key)
	hmac.Write([]byte(data))

	return base64.StdEncoding.EncodeToString(hmac.Sum(nil))
}
