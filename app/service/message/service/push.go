package service

import (
	"context"
	"fmt"

	"valerian/library/jpush"
)

func (p *Service) pushSingleUser(c context.Context, aid int64, message string) (msgID string, err error) {
	payload := &jpush.Payload{
		Platform: jpush.NewPlatform().All(),
		Audience: jpush.NewAudience().SetRegistrationId(fmt.Sprintf("%d", aid)),
		Notification: &jpush.Notification{
			Alert: message,
		},
		Options: &jpush.Options{
			TimeLive:       60,
			ApnsProduction: false,
		},
	}
	return p.jp.Push(payload)
}
