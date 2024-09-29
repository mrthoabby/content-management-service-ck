package application

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/ports/in"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
)

func NewSectionService(getItemByID in.UseCaseWithParamAndReturn[types.GetSectionByIDParams, models.Section]) *Service {
	return &Service{
		getItemByID: getItemByID,
	}
}

type Service struct {
	getItemByID in.UseCaseWithParamAndReturn[types.GetSectionByIDParams, models.Section]
}

func (s *Service) GetSectionByID(context context.Context, params types.GetSectionByIDParams) dto.SectionDTO {
	sections := s.getItemByID.Execute(context, params)

	sectionDTO := MapSectionToSectionDTO(sections)

	return sectionDTO
}
