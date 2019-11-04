package service

import (
	"context"
	"fmt"

	"valerian/library/jpush"
)

func (p *Service) pushSingleUser(c context.Context, aid int64, msg *jpush.Message) (msgID string, err error) {
	payload := &jpush.Payload{
		Platform: jpush.NewPlatform().All(),
		Audience: jpush.NewAudience().SetAlias(fmt.Sprintf("%d", aid)),
		Notification: &jpush.Notification{
			Alert: msg.Content,
		},
		Message: msg,
		Options: &jpush.Options{
			TimeLive:       60,
			ApnsProduction: false,
		},
	}
	return p.jp.Push(payload)
}
