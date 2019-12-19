package service

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"valerian/app/service/account/model"
)

func hashPassword(password string, salt string) (passwordHash string, err error) {
	mac := hmac.New(sha1.New, []byte(model.PasswordPepper))
	_, err = mac.Write([]byte(password))
	if err != nil {
		return
	}
	sha := hex.EncodeToString(mac.Sum(nil))

	h := sha1.New()
	_, err = h.Write([]byte(sha + salt))
	if err != nil {
		return
	}

	passwordHash = base64.URLEncoding.EncodeToString(h.Sum(nil))
	return
}
