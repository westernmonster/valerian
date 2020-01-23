package service

import (
	"container/list"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"time"

	"valerian/app/service/msm/model"
	"valerian/library/ecode"
	"valerian/library/log"
)

// Codes get codes.
func (s *Service) Codes(c context.Context, rver int64) (*model.ErrCodes, error) {
	var (
		e  *list.Element
		ok bool
	)
	codes := s.codes.Load().(*model.ErrCodes)
	if rver == 0 {
		return codes, nil
	}
	if rver == codes.Ver {
		return nil, ecode.NotModified
	}
	code := make(map[int]string)
	s.lock.RLock()
	if e, ok = s.version.Map[rver]; !ok {
		s.lock.RUnlock()
		return codes, nil
	}
	for ; e != nil; e = e.Next() {
		ver := e.Value.(*model.ErrCode)
		if ver.Ver <= rver {
			continue
		}
		code[ver.Code] = ver.Msg
	}
	s.lock.RUnlock()
	return s.newCodes(code, codes.Ver)
}

// updateproc update proc.每2分钟拉一次增量，每小时拉一次全量
func (s *Service) updateproc() {
	var last = time.Now()
	for {
		cur := time.Now()
		if cur.Sub(last) > time.Hour {
			if err := s.all(); err != nil {
				time.Sleep(2 * time.Second)
				continue
			}
			last = cur
		} else {
			if err := s.diff(); err != nil {
				time.Sleep(2 * time.Second)
				continue
			}
		}
		time.Sleep(120 * time.Second)
	}
}

func (s *Service) newCodes(code map[int]string, ver int64) (*model.ErrCodes, error) {
	bytes, err := json.Marshal(code)
	if err != nil {
		log.Errorf("json.Marshal(%v) error(%v)", code, err)
		return nil, err
	}
	mb := md5.Sum(bytes)
	return &model.ErrCodes{Ver: ver, Code: code, MD5: hex.EncodeToString(mb[:])}, nil
}

// all get all codes.//全量获取
func (s *Service) all() (err error) {
	var (
		lcode  *model.ErrCode
		ncodes *model.ErrCodes
		code   map[int]string
	)
	if code, lcode, err = s.d.Codes(context.Background(), s.d.ApmDB()); err != nil {
		return
	}
	if ncodes, err = s.newCodes(code, lcode.Ver); err != nil {
		return
	}
	s.codes.Store(ncodes)
	// save last ver into list.
	s.lock.Lock()
	l := s.version.List
	m := s.version.Map
	if _, ok := m[lcode.Ver]; ok {
		s.lock.Unlock()
		return
	}
	l.PushBack(lcode)
	m[lcode.Ver] = l.Back()
	s.lock.Unlock()
	return
}

// diff get change code.增量获取
func (s *Service) diff() (err error) {
	var (
		vers   *list.List
		ncodes *model.ErrCodes
		ocodes = s.codes.Load().(*model.ErrCodes)
		code   = copy(ocodes.Code)
	)
	if vers, err = s.d.Diff(context.Background(), s.d.ApmDB(), ocodes.Ver); err != nil {
		return
	} else if vers.Len() == 0 {
		return
	}
	// merge diff ecode
	for e := vers.Front(); e != nil; e = e.Next() {
		ver := e.Value.(*model.ErrCode)
		if ver.Ver < ocodes.Ver {
			continue
		}
		code[ver.Code] = ver.Msg
	}
	// save global ecode
	if ncodes, err = s.newCodes(code, vers.Back().Value.(*model.ErrCode).Ver); err != nil {
		return
	}
	s.codes.Store(ncodes)
	// push diff to vers list and trim
	s.lock.Lock()
	m := s.version.Map
	l := s.version.List
	for e := vers.Front(); e != nil; e = e.Next() {
		l.PushBack(e.Value)
		ver := e.Value.(*model.ErrCode).Ver
		m[ver] = l.Back()
	}
	for i := 0; i < l.Len()-_maxVerNum; i++ {
		e := l.Front()
		l.Remove(e)
		delete(m, e.Value.(*model.ErrCode).Ver)
	}
	s.lock.Unlock()
	return
}

func copy(src map[int]string) (dst map[int]string) {
	dst = make(map[int]string)
	for k, v := range src {
		dst[k] = v
	}
	return
}
