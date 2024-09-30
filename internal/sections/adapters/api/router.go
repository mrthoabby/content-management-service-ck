package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/mrthoabby/content-management-service-ck/internal/sections/adapters/api/v1"
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

	mainPath.HandleFunc(fmt.Sprintf("/{%s}", v1.SectionIDParam), r.GetSectionByID).Methods(http.MethodGet)
	mainPath.HandleFunc("", r.GetAllSections).Methods(http.MethodGet)
	mainPath.HandleFunc(fmt.Sprintf("/{%s}/pages/{%s}", v1.SectionIDParam, v1.PageIDParam), r.GetPageContentByPageID).Methods(http.MethodGet)
	mainPath.HandleFunc("/search", r.GetSectionsByQuery).Methods(http.MethodGet)

}

func NewSectionRouter(sectionHandler ports.SectionHandler, router *mux.Router) *SectionRouter {
	return &SectionRouter{
		SectionHandler: sectionHandler,
		Router:         router,
	}
}
