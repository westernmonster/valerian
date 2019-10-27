package api

import (
	"valerian/app/service/discuss/model"
)

func FromCategory(v *model.DiscussCategory) *CategoryInfo {
	reply := &CategoryInfo{
		ID:      v.ID,
		TopicID: v.TopicID,
		Name:    v.Name,
		Seq:     int32(v.Seq),
	}

	return reply
}
