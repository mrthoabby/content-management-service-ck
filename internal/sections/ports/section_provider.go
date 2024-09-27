package ports

import (
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain"
	valueobjects "github.com/mrthoabby/content-management-service-ck/internal/sections/value_objects"
	commons "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
)

type SectionProvider interface {
	FetchSectionByID(string) (domain.Section, error)
	FetchNetSectionByID(string) (domain.NetSection, error)
	FetchPaginatedNetSections(commons.Pagination) ([]domain.NetSection, error)
	FetchSectionPageContentBySectionPageID(domain.SectionPageID) (valueobjects.PageContent, error)
	FetchNetSectionsByQueryPaginated(string) ([]domain.NetSection, error)

	CreateSection(domain.Section) error
	CreateSectionPage(valueobjects.PageIDName) error

	UpdateSectionPageContent(domain.SectionPageIDContent) error
	UpdateSectionPageName(domain.SectionPageIDPageName) error
	UpdateSectionName(domain.SectionPageIDName) error

	DeleteSectionPageByID(valueobjects.PageID) error
	DeleteSectionByID(domain.SectionID) error
}
