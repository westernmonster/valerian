package helper

import (
	"crypto/rand"
	b64 "encoding/base64"
	"fmt"
	"io"
	"math/big"
	"net"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateValcode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func InetAtoN(ip string) int64 {
	if net.ParseIP(ip) == nil {
		return 0
	}

	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

// Base64Decode Base64 编码
func Base64Decode(from string) (to string, err error) {
	l, err := b64.URLEncoding.DecodeString(from)
	return string(l), err
}

func Base64Encode(from string) (to string) {
	to = b64.URLEncoding.EncodeToString([]byte(from))
	return
}
