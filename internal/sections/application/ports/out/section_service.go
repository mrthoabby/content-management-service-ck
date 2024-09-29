package ports

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
)

type SectionService interface {
	GetSectionByID(context context.Context, params types.GetSectionByIDParams) dto.SectionDTO
	GetAllSections(context context.Context, params types.GetAllSectionsParams) coredomain.PaginatedResult[[]dto.SectionDTO]

	// GetSectionPageContentBySectionIDAndPageID(models.SectionPageID) models.PageContent
	// GetPartialSectionsByQueryPaginated(string) []models.PartialSection

	// CreateSection(models.Section)
	// CreateSectionPage(models.PageIDName)

	// UpdateSectionPageContent(models.SectionPageIDContent)
	// UpdateSectionPageName(models.SectionPageIDPageName)
	// UpdateSectionName(models.SectionPageIDName)

	// DeleteSectionPageByID(models.PageID)
	// DeleteSectionByID(models.SectionID)
}
