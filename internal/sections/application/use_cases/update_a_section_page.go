package usecases

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/ports"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
)

func NewUpdateASectionPage(sectionProvider ports.SectionProvider) *UpdateASectionPage {
	return &UpdateASectionPage{
		SectionProvider: sectionProvider,
	}
}

type UpdateASectionPage struct {
	ports.SectionProvider
}

func (u *UpdateASectionPage) Execute(context context.Context, params dto.SectionPageToUpdateDTO) {
	errorUpdating := u.SectionProvider.UpdateSectionPageAsync(context, models.SectionIDPageIDContent{
		SectionID: models.SectionID(params.SectionID),
		PageID:    models.PageID(params.PageID),
		Content: models.PageContent{
			Data: params.PageContent,
		},
		PageName: models.PageName(params.PageName),
	})
	errorhandler.Handle(errorUpdating, u, "UpdateSectionPage")
}
