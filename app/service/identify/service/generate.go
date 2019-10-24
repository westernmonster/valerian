package service

import (
	"fmt"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

func generateAccess(accountID int64, ct time.Time, dc int) string {
	return generateAK(accountID, int(ct.Month()), dc)
}

func generateRefresh(accountID int64, ct time.Time, dc int) string {
	return generateRK(accountID, int(ct.Month()), dc)
}

func generateCSRF(accountID int64) (res string) {
	return md5Hex(fmt.Sprintf("%d%d%d", rand.Int63n(100000000), time.Now().Nanosecond(), accountID))
}

func generateRK(accountID int64, month, dc int) string {
	return generateByAdditional(accountID, month, dc, "refresh")
}

func generateAK(accountID int64, month, dc int) string {
	return generateByAdditional(accountID, month, dc, "access")
}

func generateSD(accountID int64, month, dc int) string {
	return generateByAdditional(accountID, month, dc, "session")
}

func generateByAdditional(accountID int64, month, dc int, additional string) string {
	t := md5Hex(fmt.Sprintf("%s,%d,%s", uuid.NewV4().String(), accountID, additional))
	// [0, 29] + 1 + 1
	return t[:30] + formatHex(month) + formatHex(dc)
}

func formatHex(n int) string {
	return fmt.Sprintf("%x", n)
}
