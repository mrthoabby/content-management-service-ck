package usecases

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/ports"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
)

func NewGetAllSections(section ports.SectionProvider) *GetAllSections {
	return &GetAllSections{
		sections: section,
	}
}

type GetAllSections struct {
	sections ports.SectionProvider
}

func (g *GetAllSections) Execute(context context.Context, params types.GetAllSectionsParams) coredomain.PaginatedResult[[]models.Section] {
	if params.LoadPages {
		paginatedData, errorGettingSections := g.sections.FetchAllPartialSectionsAsync(context, params.Pagination)
		errorhandler.Handle(errorGettingSections)

		sections := make([]models.Section, 0, len(paginatedData.Data))
		for _, section := range paginatedData.Data {
			sections = append(sections, models.Section{
				ID:   section.ID,
				Name: section.Name,
			})
		}

		return coredomain.PaginatedResult[[]models.Section]{
			Data:        sections,
			CountTotal:  paginatedData.CountTotal,
			CurrentPage: paginatedData.CurrentPage,
			GroupedBy:   paginatedData.GroupedBy,
			TotalPages:  paginatedData.TotalPages,
		}
	}

	paginatedData, errorGettingSections := g.sections.FetchAllSectionsAsync(context, params.Pagination)
	errorhandler.Handle(errorGettingSections)

	return paginatedData

}
