package ports

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
)

type SectionProvider interface {
	FetchSectionByIDAsync(context.Context, models.SectionID) (*models.Section, error)
	FetchPartialSectionByIDAsync(context.Context, models.SectionID) (*models.PartialSection, error)
	FetchPaginatedPartialSectionsAsync(context.Context, coredomain.Pagination) ([]models.PartialSection, error)
	FetchSectionPageContentBySectionPageIDAsync(context.Context, models.SectionPageID) (*models.PageContent, error)
	FetchPartialSectionsByQueryPaginatedAsync(context.Context, models.SectionID) ([]models.PartialSection, error)

	CreateSectionAsync(context.Context, models.Section) error
	CreateSectionPageAsync(context.Context, models.PageIDName) error

	UpdateSectionPageContentAsync(context.Context, models.SectionPageIDContent) error
	UpdateSectionPageNameAsync(context.Context, models.SectionPageIDPageName) error
	UpdateSectionNameAsync(context.Context, models.SectionPageIDName) error

	DeleteSectionPageByIDAsync(context.Context, models.PageID) error
	DeleteSectionByIDAsync(context.Context, models.SectionID) error
}
