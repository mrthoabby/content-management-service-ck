package ports

import (
	"net/http"
)

type SectionHandler interface {
	GetSectionByID(http.ResponseWriter, *http.Request)
	GetAllSections(http.ResponseWriter, *http.Request)
	GetPageContentByPageID(http.ResponseWriter, *http.Request)
	GetSectionsByQuery(http.ResponseWriter, *http.Request)

	CreateSection(http.ResponseWriter, *http.Request)
	CreateSectionPage(http.ResponseWriter, *http.Request)

	UpdateSection(http.ResponseWriter, *http.Request)
	UpdateSectionPage(http.ResponseWriter, *http.Request)

	DeleteSectionPageByID(http.ResponseWriter, *http.Request)
	DeleteSectionByID(http.ResponseWriter, *http.Request)
}
