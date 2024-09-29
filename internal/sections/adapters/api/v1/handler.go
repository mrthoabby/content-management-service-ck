package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	ports "github.com/mrthoabby/content-management-service-ck/internal/sections/application/ports/out"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
	numbersutil "github.com/mrthoabby/content-management-service-ck/pkg/util/numbers_util"
	stringutil "github.com/mrthoabby/content-management-service-ck/pkg/util/string_util"
	"github.com/mrthoabby/content-management-service-ck/pkg/util/validator"
)

const (
	loadPagesQuery = "load_pages"

	sectionIDParam = "section_id"

	paginationCurrentPageQuery = "current_page"
	paginationGroupedByQuery   = "grouped_by"
)

func NewSectionHandler(sectionService ports.SectionService) *SectionController {
	return &SectionController{
		sectionService: sectionService,
	}
}

type SectionController struct {
	sectionService ports.SectionService
}

func (s *SectionController) GetSectionByID(responseWriter http.ResponseWriter, request *http.Request) {
	getPages := false
	getPagesQueryValue := request.URL.Query().Get(loadPagesQuery)

	if !stringutil.IsEmptyString(getPagesQueryValue) {
		getPages = validator.IsAValidBoolean(getPagesQueryValue, fmt.Sprintf("%s ", loadPagesQuery))
	}

	sectionId := validator.IsNotEmptyString(mux.Vars(request)[sectionIDParam], fmt.Sprintf("%s param is required", sectionIDParam))

	sectionDTO := s.sectionService.GetSectionByID(request.Context(), types.GetSectionByIDParams{
		SectionID: sectionId,
		LoadPages: getPages,
	})

	sectionJSON, errorParsingToJSON := json.Marshal(sectionDTO)

	errorhandler.Handle(errorParsingToJSON)

	responseWriter.Write(sectionJSON)

}

func (s *SectionController) GetAllSections(responseWriter http.ResponseWriter, request *http.Request) {
	getPages := false
	getPagesQueryValue := request.URL.Query().Get(loadPagesQuery)

	if !stringutil.IsEmptyString(getPagesQueryValue) {
		getPages = validator.IsAValidBoolean(getPagesQueryValue, fmt.Sprintf("%s ", loadPagesQuery))
	}

	currentPaginationPage := numbersutil.ForcePositiveValue(validator.IsAValidNumber(request.URL.Query().Get(paginationCurrentPageQuery), fmt.Sprintf("%s query param is required", paginationCurrentPageQuery)))
	paginationGroupedBy := numbersutil.ForcePositiveValue(validator.IsAValidNumber(request.URL.Query().Get(paginationGroupedByQuery), fmt.Sprintf("%s query param is required", paginationGroupedByQuery)))

	sectionsDTO := s.sectionService.GetAllSections(request.Context(), types.GetAllSectionsParams{
		LoadPages: getPages,
		Pagination: coredomain.Pagination{
			CurrentPage: currentPaginationPage,
			GroupBy:     paginationGroupedBy,
		},
	})

	sectionsJSON, errorParsingToJSON := json.Marshal(sectionsDTO)

	errorhandler.Handle(errorParsingToJSON)

	responseWriter.Write(sectionsJSON)

}
