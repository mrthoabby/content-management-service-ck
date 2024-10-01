package usecases

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/ports"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
)

func NewCreateNewSectionPage(sectionProvider ports.SectionProvider) *CreateNewSectionPage {
	return &CreateNewSectionPage{
		SectionProvider: sectionProvider,
	}
}

type CreateNewSectionPage struct {
	ports.SectionProvider
}

func (c *CreateNewSectionPage) Execute(context context.Context, params dto.CreateSectionPageRequestDTO) {
	errorCreating := c.CreateSectionPageAsync(context, models.SectionPageIDPageName{
		SectionID: models.SectionID(params.SectionID),
		NetPageID: models.PageIDName{
			ID:   models.PageID(params.PageID),
			Name: models.PageName(params.Name),
		},
	})
	errorhandler.Handle(errorCreating, c, "error creating section", "usecase.create_section")
}
