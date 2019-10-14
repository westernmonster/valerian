package service

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"
	"valerian/app/interface/comment/model"
	account "valerian/app/service/account/api"
	"valerian/library/database/sqalx"
	"valerian/library/ecode"
	"valerian/library/gid"
	"valerian/library/log"
	"valerian/library/net/metadata"
)

func (p *Service) GetCommentsPaged(c context.Context, resourceID int64, targetType string, limit, offset int) (resp *model.CommentListResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var data []*model.Comment
	if data, err = p.d.GetCommentsPaged(c, p.d.DB(), resourceID, targetType, offset, limit); err != nil {
		return
	}

	resp = &model.CommentListResp{
		Paging: &model.Paging{},
		Items:  make([]*model.CommentItem, len(data)),
	}

	for i, v := range data {
		item := &model.CommentItem{
			ID:        v.ID,
			Content:   v.Content,
			Featured:  bool(v.Featured),
			IsDelete:  bool(v.Deleted),
			CreatedAt: v.CreatedAt,
		}

		var stat *model.CommentStat
		if stat, err = p.d.GetCommentStatByID(c, p.d.DB(), v.ID); err != nil {
			return
		}

		item.LikeCount = stat.LikeCount
		item.DislikeCount = stat.DislikeCount
		item.ChildCommentsCount = stat.ChildrenCount

		if item.ChildCommentsCount > 0 {
			children := make([]*model.Comment, 0)
			if children, err = p.d.GetAllChildrenComments(c, p.d.DB(), v.ID); err != nil {
				return
			}

			item.ChildComments = make([]*model.ChildCommentItem, len(children))

			for j, x := range children {
				child := &model.ChildCommentItem{
					ID:        x.ID,
					Content:   x.Content,
					Featured:  bool(x.Featured),
					IsDelete:  bool(x.Deleted),
					CreatedAt: x.CreatedAt,
				}

				if child.Like, err = p.d.IsLike(c, aid, x.ID, model.TargetTypeComment); err != nil {
					return
				}

				if child.Dislike, err = p.d.IsDislike(c, aid, x.ID, model.TargetTypeComment); err != nil {
					return
				}

				var account *account.BaseInfoReply
				if account, err = p.d.GetAccountBaseInfo(c, x.CreatedBy); err != nil {
					return
				}
				child.Creator = &model.CommentCreator{
					ID:       account.ID,
					UserName: account.UserName,
					Avatar:   account.Avatar,
				}
				intro := account.GetIntroductionValue()
				child.Creator.Introduction = &intro
				item.ChildComments[j] = child
			}
		}

		if item.Like, err = p.d.IsLike(c, aid, v.ID, model.TargetTypeComment); err != nil {
			return
		}

		if item.Dislike, err = p.d.IsDislike(c, aid, v.ID, model.TargetTypeComment); err != nil {
			return
		}

		var account *account.BaseInfoReply
		if account, err = p.d.GetAccountBaseInfo(c, v.CreatedBy); err != nil {
			return
		}
		item.Creator = &model.CommentCreator{
			ID:       account.ID,
			UserName: account.UserName,
			Avatar:   account.Avatar,
		}
		intro := account.GetIntroductionValue()
		item.Creator.Introduction = &intro

		resp.Items[i] = item
	}

	if resp.Paging.Prev, err = genURL("/api/v1/comment/list/comments", url.Values{
		"resource_id": []string{strconv.FormatInt(resourceID, 10)},
		"target_type": []string{targetType},
		"limit":       []string{strconv.Itoa(limit)},
		"offset":      []string{strconv.Itoa(offset - limit)},
	}); err != nil {
		return
	}

	if resp.Paging.Next, err = genURL("/api/v1/comment/list/comments", url.Values{
		"resource_id": []string{strconv.FormatInt(resourceID, 10)},
		"target_type": []string{targetType},
		"limit":       []string{strconv.Itoa(limit)},
		"offset":      []string{strconv.Itoa(offset + limit)},
	}); err != nil {
		return
	}

	if len(resp.Items) < limit {
		resp.Paging.IsEnd = true
		resp.Paging.Next = ""
	}

	if offset == 0 {
		resp.Paging.Prev = ""
	}

	return
}

func (p *Service) AddComment(c context.Context, arg *model.ArgAddComment) (id int64, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}
	var tx sqalx.Node
	if tx, err = p.d.DB().Beginx(c); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.BeginTran() error(%+v)", err))
		return
	}

	defer func() {
		if err != nil {
			if err1 := tx.Rollback(); err1 != nil {
				log.For(c).Error(fmt.Sprintf("tx.Rollback() error(%+v)", err1))
			}
			return
		}
	}()

	item := &model.Comment{
		ID:         gid.NewID(),
		Content:    arg.Content,
		TargetType: arg.TargetType,
		ResourceID: arg.TargetID,
		CreatedBy:  aid,
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}

	if arg.TargetType == model.TargetTypeComment {
		var comment *model.Comment
		if comment, err = p.getComment(c, tx, arg.TargetID); err != nil {
			return
		}

		item.OwnerID = comment.OwnerID

		// 如果对象是子回复，则添加被回复人
		if comment.TargetType == model.TargetTypeComment {
			item.ReplyTo = &comment.CreatedBy
			item.ResourceID = comment.ResourceID
		} else {
			// 如果被回复对象是回复  则直接设置当前的资源ID为被回复的ID
			item.ResourceID = comment.ID
		}

	}

	if err = p.d.AddComment(c, tx, item); err != nil {
		return
	}

	if err = p.d.AddCommentStat(c, tx, &model.CommentStat{
		CommentID: item.ID,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}); err != nil {
		return
	}

	if err = p.incrStat(c, tx, item.ResourceID, arg); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		log.For(c).Error(fmt.Sprintf("tx.Commit() error(%+v)", err))
	}

	id = item.ID

	return
}

func (p *Service) sendNotify(c context.Context, node sqalx.Node, resourceID int64, arg *model.ArgAddComment) (err error) {
	switch arg.TargetType {
	case model.TargetTypeArticle:
		p.addCache(func() {
			// p.onArtcileCommented()
		})

		break
	case model.TargetTypeRevise:
		if err = p.d.IncrReviseStat(c, node, &model.ReviseStat{ReviseID: arg.TargetID, CommentCount: 1}); err != nil {
			return
		}
		p.addCache(func() {
			// p.onReviseCommented()
		})
		break
	case model.TargetTypeDiscussion:
		if err = p.d.IncrDiscussionStat(c, node, &model.DiscussionStat{DiscussionID: arg.TargetID, CommentCount: 1}); err != nil {
			return
		}
		p.addCache(func() {
			// p.onReviseCommented()
		})
		break
	case model.TargetTypeComment:
		p.addCache(func() {
			// p.onCommentReplied()
		})
		break
	}

	return
}

func (p *Service) incrStat(c context.Context, node sqalx.Node, resourceID int64, arg *model.ArgAddComment) (err error) {
	targetType := arg.TargetType
	if targetType == model.TargetTypeComment {
		var comment *model.Comment
		if comment, err = p.getComment(c, node, resourceID); err != nil {
			return
		}

		if err = p.d.IncrCommentStat(c, node, &model.CommentStat{CommentID: resourceID, ChildrenCount: 1}); err != nil {
			return
		}

		targetType = comment.TargetType
	}

	switch targetType {
	case model.TargetTypeArticle:
		if err = p.d.IncrArticleStat(c, node, &model.ArticleStat{ArticleID: arg.TargetID, CommentCount: 1}); err != nil {
			return
		}

		break
	case model.TargetTypeRevise:
		if err = p.d.IncrReviseStat(c, node, &model.ReviseStat{ReviseID: arg.TargetID, CommentCount: 1}); err != nil {
			return
		}
		break
	case model.TargetTypeDiscussion:
		if err = p.d.IncrDiscussionStat(c, node, &model.DiscussionStat{DiscussionID: arg.TargetID, CommentCount: 1}); err != nil {
			return
		}
		break
	}

	return

}

func (p *Service) DelComment(c context.Context, commentID int64) (err error) {
	// 因为被删除评论也会显示，所以不做其他处理
	return p.d.DelComment(c, p.d.DB(), commentID)
}

func (p *Service) getComment(c context.Context, node sqalx.Node, commentID int64) (item *model.Comment, err error) {
	// var addCache = true
	// if item, err = p.d.CommentCache(c, commentID); err != nil {
	// 	addCache = false
	// } else if item != nil {
	// 	return
	// }

	if item, err = p.d.GetCommentByID(c, p.d.DB(), commentID); err != nil {
		return
	} else if item == nil {
		err = ecode.CommentNotExist
		return
	}

	// if addCache {
	// 	p.addCache(func() {
	// 		p.d.SetCommentCache(context.TODO(), item)
	// 	})
	// }
	return
}