package usecases

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/ports"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
)

func NewUpdateASection(sectionProvider ports.SectionProvider) *UpdateASection {
	return &UpdateASection{
		SectionProvider: sectionProvider,
	}
}

type UpdateASection struct {
	ports.SectionProvider
}

func (u *UpdateASection) Execute(context context.Context, params dto.SectionToUpdateDTO) {
	errorUpdating := u.SectionProvider.UpdateSectionAsync(context, models.SectionIDName{
		SectionID:   models.SectionID(params.ID),
		SectionName: models.SectionName(params.Name),
	})
	errorhandler.Handle(errorUpdating, u, "UpdateSection")
}
