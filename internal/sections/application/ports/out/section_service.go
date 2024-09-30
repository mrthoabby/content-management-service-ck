package ports

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
)

type SectionService interface {
	GetSectionByID(context.Context, types.GetSectionByIDParams) dto.SectionDTO
	GetAllSections(context.Context, types.GetAllSectionsParams) coredomain.PaginatedResult[[]dto.SectionDTO]

	GetPageContentByPageID(context.Context, types.GetPageContentParams) dto.PageContentDTO
	GetSectionsByQuery(context.Context, types.GetSectionsByQuery) []dto.SectionDTO
	// GetPartialSectionsByQueryPaginated(string) []models.PartialSection

	// CreateSection(models.Section)
	// CreateSectionPage(models.PageIDName)

	// UpdateSectionPageContent(models.SectionPageIDContent)
	// UpdateSectionPageName(models.SectionPageIDPageName)
	// UpdateSectionName(models.SectionPageIDName)

	// DeleteSectionPageByID(models.PageID)
	// DeleteSectionByID(models.SectionID)
}
