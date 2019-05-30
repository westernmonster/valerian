package http

import (
	"valerian/app/interface/account/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

func changePassword(c *mars.Context) {
	arg := new(model.ArgChangePassword)
	c.Bind(arg)

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	aid, _ := c.Get("aid")
	c.JSON(nil, srv.ChangePassword(c, aid.(int64), arg))
}

func getProfile(c *mars.Context) {
	aid, _ := c.Get("aid")
	c.JSON(srv.GetProfile(c, aid.(int64)))
}

func updateProfile(c *mars.Context) {
	arg := new(model.ArgUpdateProfile)
	c.Bind(arg)

	aid, _ := c.Get("aid")
	c.JSON(nil, srv.UpdateProfile(c, aid.(int64), arg))
}

func forgetPassword(c *mars.Context) {
	arg := new(model.ArgForgetPassword)
	c.Bind(arg)

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.ForgetPassword(c, arg))
}

func resetPassword(c *mars.Context) {
	arg := new(model.ArgResetPassword)
	c.Bind(arg)

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(nil, srv.ResetPassword(c, arg))
}
