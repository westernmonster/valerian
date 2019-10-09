package service

import (
	"valerian/app/service/feed/def"
	"valerian/library/log"

	"github.com/nats-io/stan.go"
)

func (p *Service) onReviseAdded(m *stan.Msg) {
	var err error
	info := new(def.MsgReviseAdded)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onReviseAdded Unmarshal failed %#v", err)
		return
	}

	// var tx sqalx.Node
	// if tx, err = p.d.DB().Beginx(context.Background()); err != nil {
	// 	log.Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
	// 	return
	// }

	// defer func() {
	// 	if err != nil {
	// 		if err1 := tx.Rollback(); err1 != nil {
	// 			log.Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
	// 		}
	// 		return
	// 	}
	// }()

	// if err = tx.Commit(); err != nil {
	// 	log.Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	// 	return
	// }

	m.Ack()
}

func (p *Service) onReviseDeleted(m *stan.Msg) {
	var err error
	info := new(def.MsgReviseDeleted)
	if err = info.Unmarshal(m.Data); err != nil {
		log.Errorf("onReviseAdded Unmarshal failed %#v", err)
		return
	}

	// 	var tx sqalx.Node
	// 	if tx, err = p.d.DB().Beginx(context.Background()); err != nil {
	// 		log.Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
	// 		return
	// 	}

	// 	defer func() {
	// 		if err != nil {
	// 			if err1 := tx.Rollback(); err1 != nil {
	// 				log.Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
	// 			}
	// 			return
	// 		}
	// 	}()

	// 	if err = tx.Commit(); err != nil {
	// 		log.Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	// 		return
	// 	}

	m.Ack()
}
