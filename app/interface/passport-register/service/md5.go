package service

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"math/rand"
	"valerian/app/interface/passport-register/model"
)

func md5Hex(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// MD52IntStr converts MD5 checksum bytes to big int string in base 10.
func MD52IntStr(md5d []byte) (res string) {
	b := big.NewInt(0)
	b.SetBytes(md5d)
	res = b.Text(10)
	return
}

// IntStr2Md5 converts big int string in base 10 to MD5 checksum bytes.
func IntStr2Md5(intStr string) (res []byte) {
	b := big.NewInt(0)
	b.SetString(intStr, 10)
	return b.Bytes()
}

// hexEncode
func hexEncode(b []byte) string {
	return hex.EncodeToString(b)
}

// hexDecode
func hexDecode(s string) (res []byte, err error) {
	return hex.DecodeString(s)
}

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

func generateSalt(n uint32) (salt string, err error) {
	b := make([]byte, n)
	_, err = rand.Read(b)
	if err != nil {
		return
	}

	salt = base64.URLEncoding.EncodeToString(b)
	return
}
