package usecases

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/ports"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
)

func NewCreateSection(sections ports.SectionProvider) *CreateSection {
	return &CreateSection{
		SectionProvider: sections,
	}
}

type CreateSection struct {
	ports.SectionProvider
}

func (c *CreateSection) Execute(context context.Context, params dto.CreateSectionRequestDTO) {
	errorCreating := c.CreateSectionAsync(context, models.SectionIDName{
		SectionID:   models.SectionID(params.ID),
		SectionName: models.SectionName(params.Name),
	})
	errorhandler.Handle(errorCreating, c, "error creating section", "usecase.create_section")
}
