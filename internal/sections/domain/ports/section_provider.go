package ports

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
)

type SectionProvider interface {
	FetchSectionByIDAsync(context.Context, models.SectionID) (*models.Section, error)
	FetchPartialSectionByIDAsync(context.Context, models.SectionID) (*models.PartialSection, error)

	FetchAllSectionsAsync(context.Context, coredomain.Pagination) (coredomain.PaginatedResult[[]models.Section], error)
	FetchAllPartialSectionsAsync(context.Context, coredomain.Pagination) (coredomain.PaginatedResult[[]models.PartialSection], error)

	FetchSectionsByQueryAsync(context.Context, string) ([]models.Section, error)
	FetchPartialSectionsByQueryAsync(context.Context, string) ([]models.PartialSection, error)

	FetchPageContentByPageIDAsync(context.Context, models.SectionPageID) (*models.SectionPageIDContent, error)

	CreateSectionAsync(context.Context, models.SectionIDName) error
	CreateSectionPageAsync(context.Context, models.SectionPageIDPageName) error

	UpdateSectionAsync(context.Context, models.SectionIDName) error
	UpdateSectionPageAsync(context.Context, models.SectionIDPageIDContent) error

	DeleteSectionPageByIDAsync(context.Context, models.SectionPageID) error
	DeleteSectionByIDAsync(context.Context, models.SectionID) error
}
