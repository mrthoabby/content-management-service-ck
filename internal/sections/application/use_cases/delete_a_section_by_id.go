package usecases

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/ports"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
)

func NewDeleteASectionByID(sectionProvider ports.SectionProvider) *DeleteASectionByID {
	return &DeleteASectionByID{
		SectionProvider: sectionProvider,
	}
}

type DeleteASectionByID struct {
	ports.SectionProvider
}

func (d *DeleteASectionByID) Execute(context context.Context, params dto.SectionIDDTO) {
	errorDeleting := d.SectionProvider.DeleteSectionByIDAsync(context, models.SectionID(params.SectionID))
	errorhandler.Handle(errorDeleting, d, "DeleteSectionByID")
}
