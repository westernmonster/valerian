package service

import (
	"context"
	"fmt"
	"testing"
	"valerian/app/interface/topic/model"
	"valerian/app/interface/topic/service/mocks"
	"valerian/library/net/metadata"

	"github.com/stretchr/testify/mock"
)

type MockDao struct {
	mock.Mock
}

func TestGetTopicMembersPaged(t *testing.T) {
	mockDao := &mocks.IDao{}
	md := metadata.MD{}
	md["aid"] = int64(1)
	c := metadata.NewContext(context.Background(), md)

	srv := &Service{
		d: mockDao,
	}

	data := []*model.TopicMember{
		&model.TopicMember{
			ID:        1,
			TopicID:   1,
			AccountID: 1,
			Role:      "owner",
		},
	}
	// resp := &model.TopicMembersPagedResp{
	// 	Count:    100,
	// 	PageSize: 10,
	// 	Data:     make([]*model.TopicMemberResp, 0),
	// }
	// for i := 0; i < 10; i++ {
	// 	resp.Data = append(resp.Data, &model.TopicMemberResp{
	// 		AccountID: 1,
	// 		Role:      "owner",
	// 		Avatar:    "avatar",
	// 		UserName:  "username",
	// 	})
	// }

	mockDao.On("DB").Return(nil, nil)
	mockDao.On("TopicMembersCache", c, int64(1), 1, 10).Return(100, data, nil)
	mockDao.On("GetTopicMembersPaged", c, int64(1), 1, 10).Return(100, data, nil)
	mockDao.On("SetTopicMembersCache", c, int64(1), 100, 1, 10, data).Return(nil)
	mockDao.On("SetTopicMembersCache", c, int64(1), 100, 1, 10, data).Return(nil)

	resp, err := srv.GetTopicMembersPaged(c, int64(1), 1, 10)

	mockDao.AssertExpectations(t)

	fmt.Println(resp)
	fmt.Println(err)
}
