package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/ports"
)

type SectionRouter struct {
	ports.SectionHandler
	*mux.Router
}

func (r *SectionRouter) InitialiceSectionRouter(middlewares ...*mux.MiddlewareFunc) {
	if len(middlewares) > 0 {
		for _, middleware := range middlewares {
			r.Use(*middleware)
		}
	}

	mainPath := r.PathPrefix("/api/v1/sections").Subrouter()

	mainPath.HandleFunc("/{section_id}", r.GetSectionByID).Methods(http.MethodGet)
}

func NewSectionRouter(sectionHandler ports.SectionHandler, router *mux.Router) *SectionRouter {
	return &SectionRouter{
		SectionHandler: sectionHandler,
		Router:         router,
	}
}
