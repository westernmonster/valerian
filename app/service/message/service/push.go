package service

import (
	"context"
	"fmt"
	"strconv"

	"valerian/library/jpush"
)

func (p *Service) pushSingleUser(c context.Context, aid int64, msgID int64, title, content, link string) (pushID string, err error) {
	payload := &jpush.Payload{
		Platform: jpush.NewPlatform().All(),
		Audience: jpush.NewAudience().SetAlias(fmt.Sprintf("%d", aid)),
		Notification: &jpush.Notification{
			Alert: title,
			Android: &jpush.AndroidNotification{
				Alert: title,
				Extras: map[string]interface{}{
					"id":   strconv.FormatInt(msgID, 10),
					"type": "link",
					"url":  link,
				},
				AlertType: 1,
			},
			Ios: &jpush.IosNotification{
				Alert: title,
				Extras: map[string]interface{}{
					"id":   strconv.FormatInt(msgID, 10),
					"type": "link",
					"url":  link,
				},
			},
		},
		Options: &jpush.Options{
			TimeLive:       60,
			ApnsProduction: true,
		},
	}

	return p.jp.Push(payload)
}
