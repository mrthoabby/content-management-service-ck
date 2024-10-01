package dto

import (
	"encoding/json"
	"io"
	"net/http"

	errortypes "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler/error_types"
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
