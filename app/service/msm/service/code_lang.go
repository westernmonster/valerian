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

// CodesLangs get codes.
func (s *Service) CodesLangs(c context.Context, rver int64) (*model.CodesLangs, error) {
	var (
		e  *list.Element
		ok bool
	)
	codes := s.langsCodes.Load().(*model.CodesLangs)
	if rver == 0 {
		return codes, nil
	}
	if rver == codes.Ver {
		return nil, ecode.NotModified
	}
	code := make(map[int]map[string]string)
	s.langsLock.RLock()
	if e, ok = s.langsVersion.Map[rver]; !ok {
		s.langsLock.RUnlock()
		return codes, nil
	}
	for ; e != nil; e = e.Next() {
		ver := e.Value.(*model.CodeLangs)
		if ver.Ver <= rver {
			continue
		}
		code[ver.Code] = ver.Msg
	}
	s.langsLock.RUnlock()
	return s.newCodesLang(code, codes.Ver)
}

func (s *Service) allLang() (err error) {
	var (
		lcode  *model.CodeLangs
		ncodes *model.CodesLangs
	)

	code := make(map[int]map[string]string)
	var codes []*model.CodeLang
	if codes, err = s.d.GetCodeLangs(context.Background(), s.d.ApmDB()); err != nil {
		return
	}

	for _, v := range codes {
		code[int(v.Code)][v.Locale] = v.Message
	}

	if ncodes, err = s.newCodesLang(code, lcode.Ver); err != nil {
		return
	}
	s.langsCodes.Store(ncodes)
	// save last ver into list.
	s.langsLock.Lock()
	l := s.langsVersion.List
	m := s.langsVersion.Map
	if _, ok := m[lcode.Ver]; ok {
		s.langsLock.Unlock()
		return
	}
	l.PushBack(lcode)
	m[lcode.Ver] = l.Back()
	s.langsLock.Unlock()
	return
}

func (s *Service) newCodesLang(code map[int]map[string]string, ver int64) (*model.CodesLangs, error) {
	bytes, err := json.Marshal(code)
	if err != nil {
		log.Errorf("json.Marshal(%v) error(%v)", code, err)
		return nil, err
	}
	mb := md5.Sum(bytes)
	return &model.CodesLangs{Ver: ver, Code: code, MD5: hex.EncodeToString(mb[:])}, nil
}

func (s *Service) updateLangproc() {
	for {
		time.Sleep(300 * time.Second)
		s.allLang()
	}
}
