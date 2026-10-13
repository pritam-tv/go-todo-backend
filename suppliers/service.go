package suppliers

import (
	"context"

	"github.com/go-rel/rel"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "suppliers")))
)

// Service instance for supplier's domain.
// Any operation done to any of object within this domain should use this service.
type Service interface {
	Search(ctx context.Context, suppliers *[]Supplier, filter Filter) error
	Create(ctx context.Context, supplier *Supplier) error
	Update(ctx context.Context, supplier *Supplier, changes rel.Changeset) error
	Delete(ctx context.Context, supplier *Supplier)
	Clear(ctx context.Context)
}

type service struct {
	search
	create
	update
	delete
	clear
}

var _ Service = (*service)(nil)

// New Suppliers service.
func New(repository rel.Repository) Service {
	return service{
		search: search{repository: repository},
		create: create{repository: repository},
		update: update{repository: repository},
		delete: delete{repository: repository},
		clear:  clear{repository: repository},
	}
}
