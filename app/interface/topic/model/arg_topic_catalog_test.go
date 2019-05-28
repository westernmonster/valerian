package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var refId = int64(33)
var roottests = []struct {
	in  *TopicLevel1Catalog
	out bool
}{
	{
		in: &TopicLevel1Catalog{
			Name:     "",
			Seq:      1,
			Type:     "taxonomy",
			RefID:    nil,
			Children: []*TopicLevel2Catalog{},
		},
		out: true,
	},
	{
		in: &TopicLevel1Catalog{
			Name:     "dddd",
			Seq:      1,
			Type:     "asdasd",
			RefID:    nil,
			Children: []*TopicLevel2Catalog{},
		},
		out: true,
	},
	{
		in: &TopicLevel1Catalog{
			Name:     "dddd",
			Seq:      1,
			Type:     "article",
			RefID:    nil,
			Children: []*TopicLevel2Catalog{},
		},
		out: true,
	},
	{
		in: &TopicLevel1Catalog{
			Name:     "dddd",
			Seq:      1,
			Type:     "test_set",
			RefID:    nil,
			Children: []*TopicLevel2Catalog{},
		},
		out: true,
	},

	{
		in: &TopicLevel1Catalog{
			Name:     "dddd",
			Seq:      1,
			Type:     "test_set",
			RefID:    &refId,
			Children: []*TopicLevel2Catalog{},
		},
		out: false,
	},
}

func TestLevel1Validation(t *testing.T) {
	for idx, tt := range roottests {
		t.Run(fmt.Sprintf("validation test: %d", idx), func(t *testing.T) {
			err := tt.in.Validate()
			if tt.out {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

var level2tests = []struct {
	in  *TopicLevel2Catalog
	out bool
}{
	{
		in: &TopicLevel2Catalog{
			Name:     "",
			Seq:      1,
			Type:     "taxonomy",
			RefID:    nil,
			Children: []*TopicChildCatalog{},
		},
		out: true,
	},
	{
		in: &TopicLevel2Catalog{
			Name:     "dddd",
			Seq:      1,
			Type:     "asdasd",
			RefID:    nil,
			Children: []*TopicChildCatalog{},
		},
		out: true,
	},
	{
		in: &TopicLevel2Catalog{
			Name:     "dddd",
			Seq:      1,
			Type:     "article",
			RefID:    nil,
			Children: []*TopicChildCatalog{},
		},
		out: true,
	},
	{
		in: &TopicLevel2Catalog{
			Name:     "dddd",
			Seq:      1,
			Type:     "test_set",
			RefID:    nil,
			Children: []*TopicChildCatalog{},
		},
		out: true,
	},

	{
		in: &TopicLevel2Catalog{
			Name:     "dddd",
			Seq:      1,
			Type:     "test_set",
			RefID:    &refId,
			Children: []*TopicChildCatalog{},
		},
		out: false,
	},
}

func TestLevel2Validation(t *testing.T) {
	for idx, tt := range level2tests {
		t.Run(fmt.Sprintf("validation test: %d", idx), func(t *testing.T) {
			err := tt.in.Validate()
			if tt.out {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

var level1WithChildrenTests = []struct {
	in  *TopicLevel1Catalog
	out bool
}{
	{
		in: &TopicLevel1Catalog{
			Name:  "333",
			Seq:   1,
			Type:  "article",
			RefID: &refId,
			Children: []*TopicLevel2Catalog{
				&TopicLevel2Catalog{
					Name:     "dddd",
					Seq:      1,
					Type:     "article",
					RefID:    &refId,
					Children: []*TopicChildCatalog{},
				},
			},
		},
		out: true,
	},
	{
		in: &TopicLevel1Catalog{
			Name:  "333",
			Seq:   1,
			Type:  "test_set",
			RefID: &refId,
			Children: []*TopicLevel2Catalog{
				&TopicLevel2Catalog{
					Name:     "dddd",
					Seq:      1,
					Type:     "article",
					RefID:    &refId,
					Children: []*TopicChildCatalog{},
				},
			},
		},
		out: true,
	},
}

func TestLevel1WithChildrenValidation(t *testing.T) {
	for idx, tt := range level1WithChildrenTests {
		t.Run(fmt.Sprintf("validation test: %d", idx), func(t *testing.T) {
			err := tt.in.Validate()
			if tt.out {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

var level2WithChildrenTests = []struct {
	in  *TopicLevel1Catalog
	out bool
}{
	{
		in: &TopicLevel1Catalog{
			Name:  "333",
			Seq:   1,
			Type:  "taxonomy",
			RefID: nil,
			Children: []*TopicLevel2Catalog{
				&TopicLevel2Catalog{
					Name:  "dddd",
					Seq:   1,
					Type:  "article",
					RefID: &refId,
					Children: []*TopicChildCatalog{
						&TopicChildCatalog{
							Name:  "dddd",
							Seq:   1,
							Type:  "article",
							RefID: &refId,
						},
					},
				},
			},
		},
		out: true,
	},
	{
		in: &TopicLevel1Catalog{
			Name:  "333",
			Seq:   1,
			Type:  "taxonomy",
			RefID: nil,
			Children: []*TopicLevel2Catalog{
				&TopicLevel2Catalog{
					Name:  "dddd",
					Seq:   1,
					Type:  "test_set",
					RefID: &refId,
					Children: []*TopicChildCatalog{
						&TopicChildCatalog{
							Name:  "dddd",
							Seq:   1,
							Type:  "article",
							RefID: &refId,
						},
					},
				},
			},
		},
		out: true,
	},

	{
		in: &TopicLevel1Catalog{
			Name:  "333",
			Seq:   1,
			Type:  "taxonomy",
			RefID: nil,
			Children: []*TopicLevel2Catalog{
				&TopicLevel2Catalog{
					Name:  "dddd",
					Seq:   1,
					Type:  "taxonomy",
					RefID: nil,
					Children: []*TopicChildCatalog{
						&TopicChildCatalog{
							Name:  "dddd",
							Seq:   1,
							Type:  "article",
							RefID: &refId,
						},
					},
				},
			},
		},
		out: false,
	},
}

func TestLevel2WithChildrenValidation(t *testing.T) {
	for idx, tt := range level2WithChildrenTests {
		t.Run(fmt.Sprintf("validation test: %d", idx), func(t *testing.T) {
			err := tt.in.Validate()
			if tt.out {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
