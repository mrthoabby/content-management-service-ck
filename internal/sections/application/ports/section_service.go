package ports

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
)

// TODO: Asegurase que aquí solo sean DTOs
type SectionService interface {
	GetSectionByID(context.Context, types.GetSectionByIDParams) dto.ResponseSectionDTO
	GetAllSections(context.Context, types.GetAllSectionsParams) coredomain.PaginatedResult[[]dto.ResponseSectionDTO]

	GetPageContentByPageID(context.Context, types.GetPageContentParams) dto.PageContentDTO
	GetSectionsByQuery(context.Context, types.GetSectionsByQuery) []dto.ResponseSectionDTO

	CreateSection(context.Context, dto.CreateSectionRequestDTO)
	CreateSectionPage(context.Context, dto.CreateSectionPageRequestDTO)

	UpdateSection(context.Context, dto.SectionToUpdateDTO)
	UpdateSectionPage(context.Context, dto.SectionPageToUpdateDTO)

	DeleteSectionPageByID(context.Context, dto.SectionIDPageIDDTO)
	DeleteSectionByID(context.Context, dto.SectionIDDTO)
}
