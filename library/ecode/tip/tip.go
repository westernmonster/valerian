package tip

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"sync/atomic"
	"time"

	cmcd "valerian/library/ecode"
	"valerian/library/log"
	xhttp "valerian/library/net/http/mars"
	xtime "valerian/library/time"
)

const (
	_codeOk          = 200
	_codeNotModified = 304
	_checkURL        = "http://%s/x/internal/msm/codes/langs"
)

var (
	defualtEcodes = &ecodes{}
	defaultConfig = &Config{
		Domain: "internal.stonote.lan",
		All:    xtime.Duration(time.Hour),
		Diff:   xtime.Duration(time.Minute * 5),
		ClientConfig: &xhttp.ClientConfig{
			App: &xhttp.App{
				Key:    "3c4e41f926e51656",
				Secret: "26a2095b60c24154521d24ae62b885bb",
			},
			Dial:    xtime.Duration(5 * time.Second),
			Timeout: xtime.Duration(5 * time.Second),
		},
	}
)

// Config config.
type Config struct {
	// Domain server domain
	Domain string
	// All get all time slice
	All xtime.Duration
	// Diff get diff time slice
	Diff xtime.Duration
	//HTTPClient httpclient config
	ClientConfig *xhttp.ClientConfig
}

type ecodes struct {
	codes  atomic.Value
	client *xhttp.Client
	conf   *Config
}

type res struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  *data  `json:"result"`
}

type data struct {
	Ver  int64
	MD5  string
	Code map[int]map[string]string
}

// Init init ecode.
func Init(conf *Config) {
	if conf == nil {
		conf = defaultConfig
	} else {
		panic(`请删除配置文件内无用配置！！！perf、log、trace、report、ecode、httpServer
			http://info.stonote.cn/pages/viewpage.action?pageId=3671762
			`)
	}
	defualtEcodes.conf = conf
	defualtEcodes.client = xhttp.NewClient(conf.ClientConfig)
	defualtEcodes.codes.Store(make(map[int]map[string]string))
	ver, _ := defualtEcodes.update(0)
	go defualtEcodes.updateproc(ver)
}

func (e *ecodes) updateproc(lastVer int64) {
	var (
		ver  int64
		err  error
		last = time.Now()
		all  = time.Duration(e.conf.All)
		diff = time.Duration(e.conf.Diff)
	)
	if e.conf.All == 0 {
		all = time.Hour
	}
	if e.conf.Diff == 0 {
		diff = 5 * time.Minute
	}
	time.Sleep(diff)
	for {
		cur := time.Now()
		if cur.Sub(last) > all {
			if ver, err = e.update(0); err != nil {
				log.Error(fmt.Sprintf("e.update() error(%v)", err))
				time.Sleep(10 * time.Second)
				continue
			}
			last = cur
		} else {
			if ver, err = e.update(lastVer); err != nil {
				log.Error(fmt.Sprintf("e.update(%d) error(%v)", lastVer, err))
				time.Sleep(10 * time.Second)
				continue
			}
		}
		lastVer = ver
		time.Sleep(diff)
	}
}

func (e *ecodes) update(ver int64) (lver int64, err error) {
	var (
		res    = &res{}
		bytes  []byte
		params = url.Values{}
	)
	params.Set("ver", strconv.FormatInt(ver, 10))
	if err = e.client.Get(context.TODO(), fmt.Sprintf(_checkURL, e.conf.Domain), "", params, &res); err != nil {
		err = fmt.Errorf("e.client.Get(%v) error(%v)", fmt.Sprintf(_checkURL, e.conf.Domain), err)
		return
	}

	fmt.Println("update()")
	fmt.Println(res.Code)
	switch res.Code {
	case _codeOk:
		if res.Result == nil {
			err = fmt.Errorf("code get() response error result: %v", res)
			return
		}
	case _codeNotModified:
		return ver, nil
	default:
		err = cmcd.Int(res.Code)
		return
	}
	if bytes, err = json.Marshal(res.Result.Code); err != nil {
		return
	}
	mb := md5.Sum(bytes)
	if res.Result.MD5 != hex.EncodeToString(mb[:]) {
		err = fmt.Errorf("get codes fail,error md5")
		return
	}
	oCodes, ok := e.codes.Load().(map[int]map[string]string)
	if !ok {
		return
	}
	nCodes := copy(oCodes)
	for k, v := range res.Result.Code {
		nCodes[k] = v
	}
	cmcd.Register(nCodes)
	e.codes.Store(nCodes)
	return res.Result.Ver, nil
}

func copy(src map[int]map[string]string) (dst map[int]map[string]string) {
	dst = make(map[int]map[string]string)
	for k, v := range src {
		dst[k] = make(map[string]string)
		for k1, v1 := range v {
			dst[k][k1] = v1
		}
	}
	return
}
