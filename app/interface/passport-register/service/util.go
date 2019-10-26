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

func asteriskEmailName(email string) string {
	components := strings.Split(email, "@")
	username := components[0]

	newUserName := ""
	count := len(username)
	for i, ch := range username {
		if count-i <= 4 {
			newUserName += "*"
		} else {
			newUserName += string(ch)
		}
	}

	return newUserName
}

func asteriskMobile(mobile string) string {
	newUserName := ""
	count := len(mobile)
	for i, ch := range mobile {
		if count-i <= 4 {
			newUserName += string(ch)
		} else {
			newUserName += "*"
		}
	}

	return newUserName
}

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
