package http

import (
	"fmt"
	"strconv"

	"valerian/app/service/topic/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func createTopic(c *mars.Context) {
	arg := new(model.ArgCreateTopic)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	if topicID, err := srv.CreateTopic(c, arg); err != nil {
		c.JSON(nil, err)
	} else {
		c.JSON(strconv.FormatInt(topicID, 10), err)
	}

}

func editTopic(c *mars.Context) {
	arg := new(model.ArgUpdateTopic)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		fmt.Println(e)
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.UpdateTopic(c, arg))
}

func deleteTopic(c *mars.Context) {
	idStr := c.Request.Form.Get("id")
	if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	} else {
		c.JSON(nil, srv.DelTopic(c, id))
	}
}

func getTopic(c *mars.Context) {
	include := c.Request.Form.Get("include")
	idStr := c.Request.Form.Get("id")
	if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	} else {
		c.JSON(srv.GetTopic(c, id, include))
	}
}

func changeOwner(c *mars.Context) {
	arg := new(model.ArgChangeOwner)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.ChangeOwner(c, arg))
}

func favTopic(c *mars.Context) {
	idStr := c.Request.Form.Get("id")
	if id, err := strconv.ParseInt(idStr, 10, 64); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	} else if id == 0 {
		c.JSON(nil, ecode.RequestErr)
		return
	} else {
		c.JSON(srv.FavTopic(c, id))
	}
}

func topicsWithEditPermission(c *mars.Context) {
}