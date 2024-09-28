package ports

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
)

type SectionProvider interface {
	FetchSectionByIDAsync(context.Context, string) (models.Section, errorhandler.Commons)
	FetchPartialSectionByIDAsync(context.Context, string) (models.PartialSection, errorhandler.Commons)
	FetchPaginatedPartialSectionsAsync(context.Context, coredomain.Pagination) ([]models.PartialSection, errorhandler.Commons)
	FetchSectionPageContentBySectionPageIDAsync(context.Context, models.SectionPageID) (models.PageContent, errorhandler.Commons)
	FetchPartialSectionsByQueryPaginatedAsync(context.Context, string) ([]models.PartialSection, errorhandler.Commons)

	CreateSectionAsync(context.Context, models.Section) errorhandler.Commons
	CreateSectionPageAsync(context.Context, models.PageIDName) errorhandler.Commons

	UpdateSectionPageContentAsync(context.Context, models.SectionPageIDContent) errorhandler.Commons
	UpdateSectionPageNameAsync(context.Context, models.SectionPageIDPageName) errorhandler.Commons
	UpdateSectionNameAsync(context.Context, models.SectionPageIDName) errorhandler.Commons

	DeleteSectionPageByIDAsync(context.Context, models.PageID) errorhandler.Commons
	DeleteSectionByIDAsync(context.Context, models.SectionID) errorhandler.Commons
}
