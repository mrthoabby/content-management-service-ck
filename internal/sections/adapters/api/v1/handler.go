package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/ports"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
	numbersutil "github.com/mrthoabby/content-management-service-ck/pkg/util/numbers_util"
	stringutil "github.com/mrthoabby/content-management-service-ck/pkg/util/string_util"
	"github.com/mrthoabby/content-management-service-ck/pkg/util/validate"
)

const (
	APIVersion  = "v1"
	APIMainPath = "/api/" + APIVersion + "/sections"

	LoadPagesQuery = "load_pages"

	PageIDParam = "page_id"

	SectionIDParam = "section_id"

	QueryParam = "query"

	PaginationCurrentPageQuery = "current_page"
	PaginationGroupedByQuery   = "grouped_by"

	paramsRequiredMessage       = "%s param is required"
	queryParamIsRequiredMessage = "%s query param is required"
)

func needLoadPages(request *http.Request) bool {
	getPages := false

	getPagesQueryValue := request.URL.Query().Get(LoadPagesQuery)

	if !stringutil.IsEmptyString(getPagesQueryValue) {
		getPages = validate.IsAValidBoolean(getPagesQueryValue, fmt.Sprintf("%s ", LoadPagesQuery))
	}

	return getPages
}

func NewSectionHandler(sectionService ports.SectionService) *SectionController {
	return &SectionController{
		sectionService: sectionService,
	}
}

type SectionController struct {
	sectionService ports.SectionService
}

func (s SectionController) GetSectionByID(responseWriter http.ResponseWriter, request *http.Request) {
	sectionId := validate.IsNotEmptyString(mux.Vars(request)[SectionIDParam], fmt.Sprintf(paramsRequiredMessage, SectionIDParam))

	sectionDTO := s.sectionService.GetSectionByID(request.Context(), types.GetSectionByIDParams{
		SectionID: sectionId,
		LoadPages: needLoadPages(request),
	})

	sectionJSON, errorParsingToJSON := json.Marshal(sectionDTO)

	errorhandler.Handle(errorParsingToJSON, s)

	responseWriter.Write(sectionJSON)

}

func (s SectionController) GetAllSections(responseWriter http.ResponseWriter, request *http.Request) {
	currentPaginationPage := numbersutil.ForcePositiveValue(validate.IsAValidNumber(request.URL.Query().Get(PaginationCurrentPageQuery), fmt.Sprintf(queryParamIsRequiredMessage, PaginationCurrentPageQuery)))
	paginationGroupedBy := numbersutil.ForcePositiveValue(validate.IsAValidNumber(request.URL.Query().Get(PaginationGroupedByQuery), fmt.Sprintf(queryParamIsRequiredMessage, PaginationGroupedByQuery)))

	sectionsDTO := s.sectionService.GetAllSections(request.Context(), types.GetAllSectionsParams{
		LoadPages: needLoadPages(request),
		Pagination: coredomain.Pagination{
			CurrentPage: currentPaginationPage,
			GroupBy:     paginationGroupedBy,
		},
	})

	sectionsJSON, errorParsingToJSON := json.Marshal(sectionsDTO)

	errorhandler.Handle(errorParsingToJSON, s)

	responseWriter.Write(sectionsJSON)

}

func (s SectionController) GetPageContentByPageID(responseWriter http.ResponseWriter, request *http.Request) {
	pageID := validate.IsNotEmptyString(mux.Vars(request)[PageIDParam], fmt.Sprintf(paramsRequiredMessage, PageIDParam))
	sectionID := validate.IsNotEmptyString(mux.Vars(request)[SectionIDParam], fmt.Sprintf(paramsRequiredMessage, SectionIDParam))

	pageContentDTO := s.sectionService.GetPageContentByPageID(request.Context(), types.GetPageContentParams{
		PageID:    pageID,
		SectionID: sectionID,
	})

	pageContentJSON, errorParsingToJSON := json.Marshal(pageContentDTO)

	errorhandler.Handle(errorParsingToJSON, s)

	responseWriter.Write(pageContentJSON)

}

func (s SectionController) GetSectionsByQuery(responseWriter http.ResponseWriter, request *http.Request) {
	query := validate.IsNotEmptyString(request.URL.Query().Get(QueryParam), fmt.Sprintf(queryParamIsRequiredMessage, QueryParam))

	sectionsDTO := s.sectionService.GetSectionsByQuery(request.Context(), types.GetSectionsByQuery{
		Query:     query,
		LoadPages: needLoadPages(request),
	})

	sectionsJSON, errorParsingToJSON := json.Marshal(sectionsDTO)

	errorhandler.Handle(errorParsingToJSON, s)

	responseWriter.Write(sectionsJSON)

}

func (s SectionController) CreateSection(responseWriter http.ResponseWriter, request *http.Request) {
	sectionDTO, errorBuilding := dto.BuildCreateSectionRequestDTO(request.Body)
	errorhandler.Handle(errorBuilding, s, "error building section dto", "handler: CreateSection")

	validate.IsAValidStructure(sectionDTO, "section")

	s.sectionService.CreateSection(request.Context(), *sectionDTO)

	responseWriter.Header().Set("Location", fmt.Sprintf("%s/%s", APIMainPath, sectionDTO.ID))
	responseWriter.WriteHeader(http.StatusCreated)
}

func (s SectionController) CreateSectionPage(responseWriter http.ResponseWriter, request *http.Request) {
	sectionID := validate.IsNotEmptyString(mux.Vars(request)[SectionIDParam], fmt.Sprintf(paramsRequiredMessage, SectionIDParam))
	pageDTO, errorBuilding := dto.BuildCreateSectionPageRequestDTO(request.Body, sectionID)
	errorhandler.Handle(errorBuilding, s, "error building page dto", "handler: CreateSectionPage")

	validate.IsAValidStructure(pageDTO, "page")

	s.sectionService.CreateSectionPage(request.Context(), *pageDTO)

	responseWriter.Header().Set("Location", fmt.Sprintf("%s/%s/pages/%s", APIMainPath, sectionID, pageDTO.PageID))
	responseWriter.WriteHeader(http.StatusCreated)
}
