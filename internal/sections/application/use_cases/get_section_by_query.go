package usecases

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/ports"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
)

func NewGetSectionsByQuery(section ports.SectionProvider) *GetSectionByQuery {
	return &GetSectionByQuery{
		SectionProvider: section,
	}
}

type GetSectionByQuery struct {
	ports.SectionProvider
}

func (g *GetSectionByQuery) Execute(context context.Context, params types.GetSectionsByQuery) []models.Section {
	if params.LoadPages {
		sections, errorGettingSections := g.FetchSectionsByQueryAsync(context, params.Query)
		errorhandler.Handle(errorGettingSections, g, "error getting sections", "usecase.get_sections_by_query")

		return sections

	}

	partialSections, errorGettingSections := g.FetchPartialSectionsByQueryAsync(context, params.Query)
	errorhandler.Handle(errorGettingSections, g, "error getting partial sections", "usecase.get_sections_by_query")

	sections := make([]models.Section, 0, len(partialSections))
	for i, section := range partialSections {
		sections[i] = models.Section{
			ID:   section.ID,
			Name: section.Name,
		}
	}

	return sections
}
