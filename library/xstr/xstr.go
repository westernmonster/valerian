package xstr

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/blang/semver"
	"github.com/pkg/errors"
)

var (
	bfPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

func UrlVerPrefix(appVersion string) (urlPrefix string) {
	ver, err := semver.Make(appVersion)
	if err != nil {
		panic(errors.WithMessage(err, "wrong app version"))
	}

	return fmt.Sprintf("/api/v%d", ver.Major)
}

// JoinInts format int64 slice like:n1,n2,n3.
func JoinInts(is []int64) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return strconv.FormatInt(is[0], 10)
	}
	buf := bfPool.Get().(*bytes.Buffer)
	for _, i := range is {
		buf.WriteString(strconv.FormatInt(i, 10))
		buf.WriteByte(',')
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	s := buf.String()
	buf.Reset()
	bfPool.Put(buf)
	return s
}

// SplitInts split string into int64 slice.
func SplitInts(s string) ([]int64, error) {
	if s == "" {
		return nil, nil
	}
	sArr := strings.Split(s, ",")
	res := make([]int64, 0, len(sArr))
	for _, sc := range sArr {
		i, err := strconv.ParseInt(sc, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}

func Excerpt(s string) string {
	str := []rune(s)
	if len(str) > 100 {
		return string(str[:100])
	}

	return s
}

func Int64Array2StringArray(req []int64) (resp []string) {
	resp = make([]string, 0)

	for _, v := range req {
		resp = append(resp, strconv.FormatInt(v, 10))
	}

	return
}

func StringArray2Int64Array(req []string) (resp []int64, err error) {
	resp = make([]int64, 0)

	for _, v := range req {
		if x, e := strconv.ParseInt(v, 10, 64); e != nil {
			err = e
			return
		} else {
			resp = append(resp, x)
		}
	}

	return
}
