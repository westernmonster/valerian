package service

import "context"

func (p *Service) Init(c context.Context) (err error) {
	if err = p.d.CreateAccountIndices(c); err != nil {
		return
	}

	if err = p.d.CreateTopicIndices(c); err != nil {
		return
	}

	if err = p.d.CreateArticleIndices(c); err != nil {
		return
	}

	if err = p.d.CreateDiscussionIndices(c); err != nil {
		return
	}

	return
}
