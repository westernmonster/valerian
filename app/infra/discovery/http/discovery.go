package http

import (
	"net/http"
	"strconv"
	"time"

	"valerian/app/infra/discovery/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
	"valerian/library/net/http/mars/binding"
)

const (
	_pollWaitSecond = 30 * time.Second
)

func register(c *mars.Context) {
	arg := new(model.ArgRegister)
	if err := c.BindWith(arg, binding.Form); err != nil {
		c.JSON(nil, err)
		return
	}

	i := model.NewInstance(arg)
	if i.Status == 0 || i.Status > 2 {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	// register replication
	if arg.DirtyTimestamp > 0 {
		i.DirtyTimestamp = arg.DirtyTimestamp
	}
	svr.Register(c, i, arg.LatestTimestamp, arg.Replication)
	c.JSON(nil, nil)
}

func renew(c *mars.Context) {
	arg := new(model.ArgRenew)
	if err := c.BindWith(arg, binding.Form); err != nil {
		c.JSON(nil, err)
		return
	}
	// renew
	c.JSON(svr.Renew(c, arg))
}

func cancel(c *mars.Context) {
	arg := new(model.ArgCancel)
	if err := c.BindWith(arg, binding.Form); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(nil, svr.Cancel(c, arg))
}

func fetchAll(c *mars.Context) {
	c.JSON(svr.FetchAll(c), nil)
}

func fetch(c *mars.Context) {
	arg := new(model.ArgFetch)
	if err := c.BindWith(arg, binding.Query); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(svr.Fetch(c, arg))
}

func fetchs(c *mars.Context) {
	arg := new(model.ArgFetchs)
	if err := c.BindWith(arg, binding.Query); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(svr.Fetchs(c, arg))
}

func poll(c *mars.Context) {
	arg := new(model.ArgPolls)
	if err := c.BindWith(arg, binding.Query); err != nil {
		c.JSON(nil, err)
		return
	}
	ch, new, err := svr.Polls(c, arg)
	if err != nil && err != ecode.NotModified {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	// wait for instance change
	select {
	case e := <-ch:
		if len(arg.Treeid) != 0 {
			c.JSON(e[strconv.FormatInt(arg.Treeid[0], 10)], nil)
		} else {
			c.JSON(e[arg.Appid[0]], nil)
		}
		if !new {
			svr.DelConns(arg) // broadcast will delete all connections of appid
		}
	case <-time.After(_pollWaitSecond):
		c.JSON(nil, ecode.NotModified)
		svr.DelConns(arg)
	case <-c.Writer.(http.CloseNotifier).CloseNotify():
		c.JSON(nil, ecode.NotModified)
		svr.DelConns(arg)
	}
}

func polls(c *mars.Context) {
	arg := new(model.ArgPolls)
	if err := c.BindWith(arg, binding.Query); err != nil {
		c.JSON(nil, err)
		return
	}
	if len(arg.Treeid) != len(arg.LatestTimestamp) && len(arg.Appid) != len(arg.LatestTimestamp) {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	ch, new, err := svr.Polls(c, arg)
	if err != nil && err != ecode.NotModified {
		c.JSON(nil, err)
		return
	}
	// wait for instance change
	select {
	case e := <-ch:
		c.JSON(e, nil)
		if !new {
			svr.DelConns(arg) // broadcast will delete all connections of appid
		}
	case <-time.After(_pollWaitSecond):
		c.JSON(nil, ecode.NotModified)
		svr.DelConns(arg)
	case <-c.Writer.(http.CloseNotifier).CloseNotify():
		c.JSON(nil, ecode.NotModified)
		svr.DelConns(arg)
	}

}

func polling(c *mars.Context) {
	arg := new(model.ArgPolling)
	if err := c.BindWith(arg, binding.Query); err != nil {
		c.JSON(nil, err)
		return
	}
	c.JSON(svr.Polling(c, arg))
}

func set(c *mars.Context) {
	var (
		arg = new(model.ArgSet)
	)
	if err := c.BindWith(arg, binding.Form); err != nil {
		c.JSON(nil, err)
		return
	}
	// len of status,metadata must equal to len of hostname or be zero
	if (len(arg.Hostname) != len(arg.Status) && len(arg.Status) != 0) ||
		(len(arg.Hostname) != len(arg.Metadata) && len(arg.Metadata) != 0) {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, svr.Set(c, arg))
}

func nodes(c *mars.Context) {
	c.JSON(svr.Nodes(c), nil)
}
