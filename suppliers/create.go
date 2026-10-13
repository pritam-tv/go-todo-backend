package suppliers

import (
	"context"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type create struct {
	repository rel.Repository
}

func (c create) Create(ctx context.Context, supplier *Supplier) error {
	if err := supplier.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	c.repository.MustInsert(ctx, supplier)
	return nil
}
