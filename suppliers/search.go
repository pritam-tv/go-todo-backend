package suppliers

import (
	"context"

	"github.com/go-rel/rel"
)

// Filter for search.
type Filter struct {
	Keyword string
}

type search struct {
	repository rel.Repository
}

func (s search) Search(ctx context.Context, suppliers *[]Supplier, filter Filter) error {
	var (
		query = rel.Select().SortAsc("name")
	)

	if filter.Keyword != "" {
		query = query.Where(rel.Like("name", "%"+filter.Keyword+"%"))
	}

	s.repository.MustFindAll(ctx, suppliers, query)
	return nil
}
