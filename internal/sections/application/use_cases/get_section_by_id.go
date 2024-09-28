package usecases

import (
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/ports"
	valueobjects "github.com/mrthoabby/content-management-service-ck/internal/sections/domain/value_objects"
	"github.com/mrthoabby/content-management-service-ck/pkg/commons/errorhandler"
)

type GetSectionById struct {
	sectionProvide ports.SectionProvider
}

func (g *GetSectionById) Execute(params valueobjects.GetSectionByIDParams) models.Section {
	if params.LoadPages {
		secion, errorGettingSection := g.sectionProvide.FetchSectionByIDAsync(params.SectionID)
		errorhandler.Handle(errorGettingSection)

		return secion
	}

	section, errorGettingSection := g.sectionProvide.FetchPartialSectionByIDAsync(params.SectionID)
	errorhandler.Handle(errorGettingSection)

	return models.Section{
		ID:   section.ID,
		Name: section.Name,
	}
}
