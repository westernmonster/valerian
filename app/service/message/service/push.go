package service

import (
	"context"
	"fmt"
	"strconv"

	"valerian/app/service/message/model"
	"valerian/library/jpush"
)

func (p *Service) pushSingleUser(c context.Context, item *model.PushMessage) (pushID string, err error) {
	payload := &jpush.Payload{
		Platform: jpush.NewPlatform().All(),
		Audience: jpush.NewAudience().SetAlias(fmt.Sprintf("%d", item.Aid)),
		Notification: &jpush.Notification{
			Alert: item.Title,
			Android: &jpush.AndroidNotification{
				Alert: item.Title,
				Extras: map[string]interface{}{
					"id":   strconv.FormatInt(item.MsgID, 10),
					"type": "link",
					"url":  item.Link,
				},
				AlertType: 1,
			},
			Ios: &jpush.IosNotification{
				Alert: item.Title,
				Extras: map[string]interface{}{
					"id":   strconv.FormatInt(item.MsgID, 10),
					"type": "link",
					"url":  item.Link,
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
