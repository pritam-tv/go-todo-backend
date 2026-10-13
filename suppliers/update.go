package suppliers

import (
	"context"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

type update struct {
	repository rel.Repository
}

func (u update) Update(ctx context.Context, supplier *Supplier, changes rel.Changeset) error {
	if err := supplier.Validate(); err != nil {
		logger.Warn("validation error", zap.Error(err))
		return err
	}

	u.repository.MustUpdate(ctx, supplier, changes)
	return nil
}
