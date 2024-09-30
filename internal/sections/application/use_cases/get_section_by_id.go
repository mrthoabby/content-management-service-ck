package usecases

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/ports"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
)

func NewGetSectionById(sectionProvide ports.SectionProvider) *GetSectionById {
	return &GetSectionById{
		SectionProvider: sectionProvide,
	}
}

type GetSectionById struct {
	ports.SectionProvider
}

func (g *GetSectionById) Execute(context context.Context, params types.GetSectionByIDParams) models.Section {
	if params.LoadPages {
		secion, errorGettingSection := g.FetchSectionByIDAsync(context, models.SectionID(params.SectionID))
		errorhandler.Handle(errorGettingSection)

		return *secion
	}

	section, errorGettingSection := g.FetchPartialSectionByIDAsync(context, models.SectionID(params.SectionID))
	errorhandler.Handle(errorGettingSection)

	return models.Section{
		ID:   section.ID,
		Name: section.Name,
	}
}
