package service

import (
	"context"
	"testing"
	"valerian/app/interface/topic/model"
	"valerian/app/interface/topic/service/mocks"
	"valerian/library/net/metadata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDao struct {
	mock.Mock
}

var (
	acc = &model.Account{
		ID:       1,
		Mobile:   "17761292244",
		Email:    "a@a.com",
		UserName: "aaa",
		Role:     "user",
		Avatar:   "avatar",
		IDCert:   true,
		WorkCert: true,
		IsOrg:    false,
		IsVIP:    false,
	}
)

func TestGetTopicMembersPaged(t *testing.T) {
	mockDao := &mocks.IDao{}
	md := metadata.MD{}
	md["aid"] = int64(1)
	c := metadata.NewContext(context.Background(), md)

	srv := &Service{
		d:      mockDao,
		missch: make(chan func(), 1024),
	}

	go srv.cacheproc()

	data := []*model.TopicMember{
		&model.TopicMember{
			ID:        1,
			TopicID:   1,
			AccountID: 1,
			Role:      "owner",
		},
		&model.TopicMember{
			ID:        2,
			TopicID:   1,
			AccountID: 2,
			Role:      "admin",
		},
	}

	mockDao.On("DB").Return(nil, nil)
	mockDao.On("AccountCache", c, int64(1)).Return(acc, nil)
	mockDao.On("AccountCache", c, int64(2)).Return(acc, nil)
	mockDao.On("TopicMembersCache", c, int64(1), 1, 10).Return(100, data, nil)
	mockDao.On("GetTopicMembersPaged", c, mockDao.DB(), int64(1), 1, 10).Return(100, data, nil)
	mockDao.On("SetTopicMembersCache", c, int64(1), 100, 1, 10, data).Return(nil)

	resp, err := srv.GetTopicMembersPaged(c, int64(1), 1, 10)

	assert.Equal(t, 100, resp.Count)
	assert.Equal(t, 10, resp.PageSize)
	assert.Equal(t, nil, err)
	// assert.Equal(t, 2, len(resp.Data))

}
