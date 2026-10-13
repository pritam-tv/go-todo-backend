package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Fs02/go-todo-backend/suppliers"
	"github.com/go-chi/chi"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"go.uber.org/zap"
)

// Suppliers for suppliers endpoints.
type Suppliers struct {
	*chi.Mux
	repository rel.Repository
	suppliers  suppliers.Service
}

// Index handle GET /.
func (s Suppliers) Index(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		query  = r.URL.Query()
		result []suppliers.Supplier
		filter = suppliers.Filter{
			Keyword: query.Get("keyword"),
		}
	)

	s.suppliers.Search(ctx, &result, filter)
	render(w, result, 200)
}

// Create handle POST /
func (s Suppliers) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = r.Context()
		supplier suppliers.Supplier
	)

	if err := json.NewDecoder(r.Body).Decode(&supplier); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(w, ErrBadRequest, 400)
		return
	}

	if err := s.suppliers.Create(ctx, &supplier); err != nil {
		render(w, err, 422)
		return
	}

	w.Header().Set("Location", fmt.Sprint(r.RequestURI, "/", supplier.ID))
	render(w, supplier, 201)
}

// Show handle GET /{ID}
func (s Suppliers) Show(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = r.Context()
		supplier = ctx.Value(loadKey).(suppliers.Supplier)
	)

	render(w, supplier, 200)
}

// Update handle PATCH /{ID}
func (s Suppliers) Update(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = r.Context()
		supplier = ctx.Value(loadKey).(suppliers.Supplier)
		changes  = rel.NewChangeset(&supplier)
	)

	if err := json.NewDecoder(r.Body).Decode(&supplier); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(w, ErrBadRequest, 400)
		return
	}

	if err := s.suppliers.Update(ctx, &supplier, changes); err != nil {
		render(w, err, 422)
		return
	}

	render(w, supplier, 200)
}

// Destroy handle DELETE /{ID}
func (s Suppliers) Destroy(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = r.Context()
		supplier = ctx.Value(loadKey).(suppliers.Supplier)
	)

	s.suppliers.Delete(ctx, &supplier)
	render(w, nil, 204)
}

// Clear handle DELETE /
func (s Suppliers) Clear(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	s.suppliers.Clear(ctx)
	render(w, nil, 204)
}

// Load is middleware that loads suppliers to context.
func (s Suppliers) Load(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx      = r.Context()
			id, _    = strconv.Atoi(chi.URLParam(r, "ID"))
			supplier suppliers.Supplier
		)

		if err := s.repository.Find(ctx, &supplier, where.Eq("id", id)); err != nil {
			if errors.Is(err, rel.ErrNotFound) {
				render(w, err, 404)
				return
			}
			panic(err)
		}

		ctx = context.WithValue(ctx, loadKey, supplier)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewSuppliers handler.
func NewSuppliers(repository rel.Repository, suppliers suppliers.Service) Suppliers {
	h := Suppliers{
		Mux:        chi.NewMux(),
		repository: repository,
		suppliers:  suppliers,
	}

	h.Get("/", h.Index)
	h.Post("/", h.Create)
	h.With(h.Load).Get("/{ID}", h.Show)
	h.With(h.Load).Patch("/{ID}", h.Update)
	h.With(h.Load).Delete("/{ID}", h.Destroy)
	h.Delete("/", h.Clear)

	return h
}
