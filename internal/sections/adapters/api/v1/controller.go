package v1

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	ports "github.com/mrthoabby/content-management-service-ck/internal/sections/application/ports/out"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
	"github.com/mrthoabby/content-management-service-ck/pkg/util/validator"
)

const (
	fetchPagesQuery = "fetchPages"

	sectionIDParam = "section_id"
)

func NewSectionController(sectionService ports.SectionService) *SectionController {
	return &SectionController{
		sectionService: sectionService,
	}
}

type SectionController struct {
	sectionService ports.SectionService
}

func (s *SectionController) GetSectionByID(responseWriter http.ResponseWriter, request *http.Request) {

	sectionId := validator.IsNotEmptyString(mux.Vars(request)[sectionIDParam], "Section ID is required")
	getPages := strings.ToLower(request.URL.Query().Get(fetchPagesQuery)) == "true"

	sectionDTO := s.sectionService.GetSectionByID(request.Context(), types.GetSectionByIDParams{
		SectionID: sectionId,
		LoadPages: getPages,
	})

	sectionJSON, errorParsingToJSON := json.Marshal(sectionDTO)

	errorhandler.Handle(errorParsingToJSON)

	responseWriter.Write(sectionJSON)

}
