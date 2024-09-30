package application

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/ports/in"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
)

func NewSectionService(getItemByID in.UseCaseWithParamAndReturn[types.GetSectionByIDParams, models.Section], getAllSections in.UseCaseWithParamAndReturn[types.GetAllSectionsParams, coredomain.PaginatedResult[[]models.Section]], getPageContentByID in.UseCaseWithParamAndReturn[types.GetPageContentParams, models.SectionPageIDContent]) *Service {
	return &Service{
		getItemByID:        getItemByID,
		getAllSections:     getAllSections,
		getPageContentByID: getPageContentByID,
	}
}

type Service struct {
	getItemByID        in.UseCaseWithParamAndReturn[types.GetSectionByIDParams, models.Section]
	getAllSections     in.UseCaseWithParamAndReturn[types.GetAllSectionsParams, coredomain.PaginatedResult[[]models.Section]]
	getPageContentByID in.UseCaseWithParamAndReturn[types.GetPageContentParams, models.SectionPageIDContent]
	getSectionsByQuery in.UseCaseWithParamAndReturn[types.GetSectionsByQuery, []models.Section]
}

func (s *Service) GetSectionByID(context context.Context, params types.GetSectionByIDParams) dto.SectionDTO {
	sections := s.getItemByID.Execute(context, params)

	sectionDTO := MapSectionToSectionDTO(sections)

	return sectionDTO
}

func (s *Service) GetAllSections(context context.Context, params types.GetAllSectionsParams) coredomain.PaginatedResult[[]dto.SectionDTO] {
	pagination := s.getAllSections.Execute(context, params)

	sectionsDTO := MapSectionsToSectionDTO(pagination.Data)

	return coredomain.PaginatedResult[[]dto.SectionDTO]{
		Data:        sectionsDTO,
		CountTotal:  pagination.CountTotal,
		CurrentPage: pagination.CurrentPage,
		GroupedBy:   pagination.GroupedBy,
		TotalPages:  pagination.TotalPages,
	}
}

func (s *Service) GetPageContentByPageID(context context.Context, params types.GetPageContentParams) dto.PageContentDTO {
	pageContent := s.getPageContentByID.Execute(context, params)

	pageContentDTO := MapPageContentToPageContentDTO(pageContent)

	return pageContentDTO
}

func (s *Service) GetSectionsByQuery(context context.Context, params types.GetSectionsByQuery) []dto.SectionDTO {
	sections := s.getSectionsByQuery.Execute(context, params)

	sectionsDTO := MapSectionsToSectionDTO(sections)

	return sectionsDTO
}
