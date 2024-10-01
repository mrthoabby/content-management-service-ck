package usecases

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/ports"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
)

func NewDeleteASectionPageByID(sectionProvider ports.SectionProvider) *DeleteASectionPageByID {
	return &DeleteASectionPageByID{
		SectionProvider: sectionProvider,
	}
}

type DeleteASectionPageByID struct {
	ports.SectionProvider
}

func (d *DeleteASectionPageByID) Execute(context context.Context, params dto.SectionIDPageIDDTO) {
	errorDeleting := d.SectionProvider.DeleteSectionPageByIDAsync(context, models.SectionPageID{
		SectionID: models.SectionID(params.SectionID),
		PageID:    models.PageID(params.PageID),
	})
	errorhandler.Handle(errorDeleting, d, "DeleteSectionPageByID")
}
