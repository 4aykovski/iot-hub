package v1

import "github.com/go-chi/chi/v5"

func New() *chi.Mux {

	return chi.NewMux()
}
