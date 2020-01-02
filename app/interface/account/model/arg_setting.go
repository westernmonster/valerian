package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ArgActivitySetting struct {
	// 赞
	Like *bool `json:"like"`
	// 评论
	Comment *bool `json:"comment"`
	// 关注话题
	FollowTopic *bool `json:"follow_topic"`
	// 关注成员
	FollowMember *bool `json:"follow_member"`
}

type ArgNotifySetting struct {
	// 赞
	Like *bool `json:"like"`
	// 评论
	Comment *bool `json:"comment"`
	// 新粉丝
	NewFans *bool `json:"new_fans"`
	// 话题新成员
	NewMember *bool `json:"new_member"`
}

type ArgLanguageSetting struct {
	// 系统语言
	Language *string `json:"language"`
}

type ArgSetting struct {
	// 系统语言
	Language *string `json:"language"`

	// 动态设置
	Activity *ArgActivitySetting `json:"activity,omitempty"`

	// 通知设置
	Notify *ArgNotifySetting `json:"notify,omitempty"`
}

type ArgDeactiveAccount struct {
	// 验证码 6位数字
	Valcode  string `json:"valcode"`
	Identity string `json:"identity"`
	Prefix   string `json:"prefix"`
	// 标识类型, 1手机, 2邮件
	IdentityType int32 `json:"identity_type"`
}

func (p *ArgDeactiveAccount) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Valcode, validation.Required, validation.RuneLength(6, 6), is.Digit),
		validation.Field(&p.Identity, validation.Required, ValidateIdentity(p.IdentityType, p.Prefix)),
		validation.Field(&p.IdentityType, validation.Required, validation.In(IdentityEmail, IdentityMobile)),
	)
}
