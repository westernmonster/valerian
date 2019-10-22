package tag

import "valerian/library/net/http/mars"

// Tag create a tag into Keys field of context
type Tag struct {
	Name   string
	policy Policy
}

// Policy is a tag policy defined by custom
type Policy interface {
	Tag(ctx *mars.Context) string
}

// PolicyFunc is a policy function
type PolicyFunc func(ctx *mars.Context) string

// Tag calls p(ctx)
func (p PolicyFunc) Tag(ctx *mars.Context) string {
	return p(ctx)
}

// New create a new tag
func New(name string, p Policy) (tag *Tag) {
	if p == nil {
		panic("policy can not be nil")
	}

	tag = new(Tag)
	tag.Name = name
	tag.policy = p

	return
}

// ServeHTTP implements from Handler interface
func (t *Tag) ServeHTTP(ctx *mars.Context) {
	ctx.Keys[t.Name] = t.policy.Tag(ctx)
}

// Value will return tag value from context by input tag name
func Value(ctx *mars.Context, name string) (value string, ok bool) {
	value, ok = ctx.Keys[name].(string)
	return
}
