package suppliers

import (
	"context"

	"github.com/go-rel/rel"
)

type delete struct {
	repository rel.Repository
}

func (d delete) Delete(ctx context.Context, supplier *Supplier) {
	d.repository.MustDelete(ctx, supplier)
}
