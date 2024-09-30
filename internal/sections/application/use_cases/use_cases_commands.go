package usecases

import (
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/ports"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
)

type UseCasesCommands struct {
	GetISectionByID      ports.UseCaseWithParamAndReturn[types.GetSectionByIDParams, models.Section]
	GetSections          ports.UseCaseWithParamAndReturn[types.GetAllSectionsParams, coredomain.PaginatedResult[[]models.Section]]
	GetPageContentByID   ports.UseCaseWithParamAndReturn[types.GetPageContentParams, models.SectionPageIDContent]
	GetSectionsWithQuery ports.UseCaseWithParamAndReturn[types.GetSectionsByQuery, []models.Section]
}
