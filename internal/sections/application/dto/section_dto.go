package dto

//TODO: Limpiar, organizar y mapear los DTOs
import (
	"encoding/json"
	"io"
	"net/http"

	errortypes "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler/error_types"
	stringutil "github.com/mrthoabby/content-management-service-ck/pkg/util/string_util"
)

type ResponseSectionDTO struct {
	ID    string    `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Pages []PageDTO `json:"pages,omitempty"`
}

type CreateSectionPageRequestDTO struct {
	SectionID string `json:"section_id" validate:"required"`
	PageID    string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

func BuildCreateSectionPageRequestDTO(body io.ReadCloser, sectionId string) (*CreateSectionPageRequestDTO, error) {
	if body == http.NoBody {
		return nil, errortypes.NewValidationError("body is required")
	}

	sectionDTO := &CreateSectionPageRequestDTO{}
	errorParsingBody := json.NewDecoder(body).Decode(sectionDTO)
	if errorParsingBody != nil {
		if _, ok := errorParsingBody.(*json.SyntaxError); ok {
			return nil, errortypes.NewInvalidFormatError("Invalid JSON format")
		}
		return nil, errorParsingBody
	}
	sectionDTO.SectionID = sectionId

	return sectionDTO, nil
}

type CreateSectionRequestDTO struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

func BuildCreateSectionRequestDTO(body io.ReadCloser) (*CreateSectionRequestDTO, error) {
	if body == http.NoBody {
		return nil, errortypes.NewValidationError("body is required")
	}

	sectionDTO := &CreateSectionRequestDTO{}
	errorParsingBody := json.NewDecoder(body).Decode(sectionDTO)
	if errorParsingBody != nil {
		if _, ok := errorParsingBody.(*json.SyntaxError); ok {
			return nil, errortypes.NewInvalidFormatError("Invalid JSON format")
		}
		return nil, errorParsingBody
	}

	return sectionDTO, nil
}

type UpdaterSectionDTO struct {
	Name string `json:"name,omitempty"`
}

func BuildUpdaterSectionDTO(body io.ReadCloser) (*UpdaterSectionDTO, error) {
	if body == http.NoBody {
		return nil, errortypes.NewValidationError("body is required")
	}

	sectionDTO := &UpdaterSectionDTO{}
	errorParsingBody := json.NewDecoder(body).Decode(sectionDTO)
	if errorParsingBody != nil {
		if _, ok := errorParsingBody.(*json.SyntaxError); ok {
			return nil, errortypes.NewInvalidFormatError("Invalid JSON format")
		}
		return nil, errorParsingBody
	}

	if stringutil.IsEmptyString(sectionDTO.Name) {
		return nil, errortypes.NewValidationError("name is required")
	}

	return sectionDTO, nil
}

type SectionToUpdateDTO struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name,omitempty"`
}

func BuildSectionToUpdateDTO(body io.ReadCloser, sectionId string) (*SectionToUpdateDTO, error) {
	if body == http.NoBody {
		return nil, errortypes.NewValidationError("body is required")
	}

	sectionDTO := &SectionToUpdateDTO{}
	errorParsingBody := json.NewDecoder(body).Decode(sectionDTO)
	if errorParsingBody != nil {
		if _, ok := errorParsingBody.(*json.SyntaxError); ok {
			return nil, errortypes.NewInvalidFormatError("Invalid JSON format")
		}
		return nil, errorParsingBody
	}

	if stringutil.IsEmptyString(sectionDTO.Name) {
		return nil, errortypes.NewValidationError("name is required")
	}

	sectionDTO.ID = sectionId

	return sectionDTO, nil
}

type UpdaterSectionPageDTO struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
}

func BuildUpdaterSectionPageDTO(body io.ReadCloser) (*UpdaterSectionPageDTO, error) {
	if body == http.NoBody {
		return nil, errortypes.NewValidationError("body is required")
	}

	sectionDTO := &UpdaterSectionPageDTO{}
	errorParsingBody := json.NewDecoder(body).Decode(sectionDTO)
	if errorParsingBody != nil {
		if _, ok := errorParsingBody.(*json.SyntaxError); ok {
			return nil, errortypes.NewInvalidFormatError("Invalid JSON format")
		}
		return nil, errorParsingBody
	}

	if stringutil.IsEmptyString(sectionDTO.Name) && stringutil.IsEmptyString(sectionDTO.Content) {
		return nil, errortypes.NewValidationError("name or content is required")
	}

	return sectionDTO, nil
}

type SectionPageToUpdateDTO struct {
	SectionID   string `json:"section_id" validate:"required"`
	PageID      string `json:"id" validate:"required"`
	PageName    string `json:"name,omitempty"`
	PageContent string `json:"content,omitempty"`
}

func BuildSectionPageToUpdateDTO(body io.ReadCloser, sectionId, pageId string) (*SectionPageToUpdateDTO, error) {
	if body == http.NoBody {
		return nil, errortypes.NewValidationError("body is required")
	}

	sectionDTO := &SectionPageToUpdateDTO{}
	errorParsingBody := json.NewDecoder(body).Decode(sectionDTO)
	if errorParsingBody != nil {
		if _, ok := errorParsingBody.(*json.SyntaxError); ok {
			return nil, errortypes.NewInvalidFormatError("Invalid JSON format")
		}
		return nil, errorParsingBody
	}

	if stringutil.IsEmptyString(sectionDTO.PageName) && stringutil.IsEmptyString(sectionDTO.PageContent) {
		return nil, errortypes.NewValidationError("name or content is required")
	}

	sectionDTO.SectionID = sectionId
	sectionDTO.PageID = pageId

	return sectionDTO, nil
}

type SectionIDPageIDDTO struct {
	SectionID string `json:"section_id" validate:"required"`
	PageID    string `json:"id" validate:"required"`
}

func NewSectionIDPageIDDTO(sectionId, pageId string) *SectionIDPageIDDTO {
	return &SectionIDPageIDDTO{
		SectionID: sectionId,
		PageID:    pageId,
	}
}

type SectionIDDTO struct {
	SectionID string `json:"section_id" validate:"required"`
}

func NewSectionIDDTO(sectionId string) *SectionIDDTO {
	return &SectionIDDTO{
		SectionID: sectionId,
	}
}

type SectionPageIDDTO struct {
	PageID string `json:"id" validate:"required"`
}
