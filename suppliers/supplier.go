package suppliers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

var (
	// SupplierURLPrefix to be returned when encoding supplier.
	SupplierURLPrefix = os.Getenv("URL") + "suppliers/"
	// ErrSupplierNameBlank validation error.
	ErrSupplierNameBlank = errors.New("Name can't be blank")
)

// Supplier represent a record stored in suppliers table.
type Supplier struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate supplier.
func (s Supplier) Validate() error {
	var err error
	switch {
	case len(s.Name) == 0:
		err = ErrSupplierNameBlank
	}

	return err
}

// MarshalJSON implement custom marshaller to marshal url.
func (s Supplier) MarshalJSON() ([]byte, error) {
	type Alias Supplier

	return json.Marshal(struct {
		Alias
		URL string `json:"url"`
	}{
		Alias: Alias(s),
		URL:   fmt.Sprint(SupplierURLPrefix, s.ID),
	})
}
