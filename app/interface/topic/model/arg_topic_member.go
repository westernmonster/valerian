package model

import validation "github.com/go-ozzo/ozzo-validation"

type ArgTopicMember struct {
	// 成员ID
	AccountID int64 `json:"account_id,string" swaggertype:"string"`
	// 角色
	// user 普通用户
	// admin 管理员
	// owner 主理人
	Role string `json:"role"`

	// 操作
	// C U D
	Opt string `json:"opt"`
}

func (p *ArgTopicMember) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.AccountID, validation.Required.Error(`请传入成员ID`)),
		validation.Field(&p.Role, validation.Required.Error(`请传入成员角色`),
			validation.In(MemberRoleUser, MemberRoleAdmin).Error("角色值不正确")),
		validation.Field(&p.Opt, validation.Required.Error(`请传入操作`),
			validation.In("C", "U", "D").Error("操作值不正确")),
	)
}

type ArgBatchSavedTopicMember struct {
	TopicID int64 `json:"topic_id,string" swaggertype:"string"`
	// 成员列表
	Members []*ArgTopicMember `json:"members"`
}

func (p *ArgBatchSavedTopicMember) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required.Error(`请传入话题ID`)),
		validation.Field(&p.Members, validation.Required.Error(`请传入成员`)),
	)
}

type ArgChangeOwner struct {
	TopicID     int64 `json:"topic_id,string" swaggertype:"string"`
	ToAccountID int64 `json:"to_account_id,string" swaggertype:"string"`
}

func (p *ArgChangeOwner) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TopicID, validation.Required.Error(`请传入话题ID`)),
		validation.Field(&p.ToAccountID, validation.Required.Error(`请传入成员ID`)),
	)
}
