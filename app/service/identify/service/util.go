package service

import (
	"fmt"
	"math/big"
	"net"
	"strings"

	"github.com/ztrue/tracerr"
)

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

// asteriskEmailName 处理Email用户名
func asteriskEmailName(email string) string {
	components := strings.Split(email, "@")
	newUserName := components[0]
	return newUserName
}

// asteriskMobile 处理手机用户名
func asteriskMobile(mobile string) string {
	runes := []rune(mobile)

	newMobile := "*" + string(runes[len(runes)-4:])

	return newMobile
}

// comparePassword 比较密码
func comparePassword(password, passwordHash, salt string) (identical bool, err error) {
	hash, err := hashPassword(password, salt)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	if hash == passwordHash {
		identical = true
	}

	return
}
