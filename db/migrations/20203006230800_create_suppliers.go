package migrations

import (
	"github.com/go-rel/rel"
)

// MigrateCreateSuppliers definition
func MigrateCreateSuppliers(schema *rel.Schema) {
	schema.CreateTable("suppliers", func(t *rel.Table) {
		t.ID("id")
		t.DateTime("created_at")
		t.DateTime("updated_at")
		t.String("name")
		t.String("email")
		t.String("phone")
		t.String("address")
	})

	schema.CreateIndex("suppliers", "name", []string{"name"})
}

// RollbackCreateSuppliers definition
func RollbackCreateSuppliers(schema *rel.Schema) {
	schema.DropTable("suppliers")
}
