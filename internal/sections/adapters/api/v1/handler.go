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
	LoadPagesQuery = "load_pages"

	PageIDParam = "page_id"

	SectionIDParam = "section_id"

	PaginationCurrentPageQuery = "current_page"
	PaginationGroupedByQuery   = "grouped_by"

	paramsRequiredMessage       = "%s param is required"
	queryParamIsRequiredMessage = "%s query param is required"
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
	getPagesQueryValue := request.URL.Query().Get(LoadPagesQuery)

	if !stringutil.IsEmptyString(getPagesQueryValue) {
		getPages = validator.IsAValidBoolean(getPagesQueryValue, fmt.Sprintf("%s ", LoadPagesQuery))
	}

	sectionId := validator.IsNotEmptyString(mux.Vars(request)[SectionIDParam], fmt.Sprintf(paramsRequiredMessage, SectionIDParam))

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
	getPagesQueryValue := request.URL.Query().Get(LoadPagesQuery)

	if !stringutil.IsEmptyString(getPagesQueryValue) {
		getPages = validator.IsAValidBoolean(getPagesQueryValue, fmt.Sprintf("%s ", LoadPagesQuery))
	}

	currentPaginationPage := numbersutil.ForcePositiveValue(validator.IsAValidNumber(request.URL.Query().Get(PaginationCurrentPageQuery), fmt.Sprintf(queryParamIsRequiredMessage, PaginationCurrentPageQuery)))
	paginationGroupedBy := numbersutil.ForcePositiveValue(validator.IsAValidNumber(request.URL.Query().Get(PaginationGroupedByQuery), fmt.Sprintf(queryParamIsRequiredMessage, PaginationGroupedByQuery)))

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

func (s *SectionController) GetPageContentByPageID(responseWriter http.ResponseWriter, request *http.Request) {
	pageID := validator.IsNotEmptyString(mux.Vars(request)[PageIDParam], fmt.Sprintf(paramsRequiredMessage, PageIDParam))
	sectionID := validator.IsNotEmptyString(mux.Vars(request)[SectionIDParam], fmt.Sprintf(paramsRequiredMessage, SectionIDParam))

	pageContentDTO := s.sectionService.GetPageContentByPageID(request.Context(), types.GetPageContentParams{
		PageID:    pageID,
		SectionID: sectionID,
	})

	pageContentJSON, errorParsingToJSON := json.Marshal(pageContentDTO)

	errorhandler.Handle(errorParsingToJSON)

	responseWriter.Write(pageContentJSON)

}
