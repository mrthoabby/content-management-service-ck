package service

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	usecases "github.com/mrthoabby/content-management-service-ck/internal/sections/application/use_cases"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
)

func NewSectionService(commands usecases.UseCasesCommands) *Service {
	return &Service{
		UseCasesCommands: commands,
	}
}

type Service struct {
	usecases.UseCasesCommands
}

func (s Service) GetSectionByID(context context.Context, params types.GetSectionByIDParams) dto.ResponseSectionDTO {
	sections := s.GetISectionByID.Execute(context, params)

	sectionDTO := dto.MapSectionToSectionDTO(sections)

	return sectionDTO
}

func (s Service) GetAllSections(context context.Context, params types.GetAllSectionsParams) coredomain.PaginatedResult[[]dto.ResponseSectionDTO] {
	pagination := s.GetSections.Execute(context, params)

	sectionsDTO := dto.MapSectionsToSectionDTO(pagination.Data)

	return coredomain.PaginatedResult[[]dto.ResponseSectionDTO]{
		Data:        sectionsDTO,
		CountTotal:  pagination.CountTotal,
		CurrentPage: pagination.CurrentPage,
		GroupedBy:   pagination.GroupedBy,
		TotalPages:  pagination.TotalPages,
	}
}

func (s Service) GetPageContentByPageID(context context.Context, params types.GetPageContentParams) dto.PageContentDTO {
	pageContent := s.GetPageContentByID.Execute(context, params)

	pageContentDTO := dto.MapPageContentToPageContentDTO(pageContent)

	return pageContentDTO
}

func (s Service) GetSectionsByQuery(context context.Context, params types.GetSectionsByQuery) []dto.ResponseSectionDTO {
	sections := s.GetSectionsWithQuery.Execute(context, params)

	sectionsDTO := dto.MapSectionsToSectionDTO(sections)

	return sectionsDTO
}

func (s Service) CreateSection(context context.Context, dto dto.CreateSectionRequestDTO) {
	s.CreateNewSection.Execute(context, dto)
}
