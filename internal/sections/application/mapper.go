package application

import (
	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/dto"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
)

func MapSectionToSectionDTO(section models.Section) dto.SectionDTO {
	return dto.SectionDTO{
		ID:    string(section.ID),
		Name:  string(section.Name),
		Pages: MapPagesToPagesDTO(section.Pages),
	}

}

func MapPageToPageDTO(page models.Page) dto.PageDTO {
	return dto.PageDTO{
		ID:      string(page.ID),
		Name:    string(page.Name),
		Content: models.PageContent(page.Content).Data,
	}
}

func MapPagesToPagesDTO(pages []models.Page) []dto.PageDTO {
	var pagesDTO []dto.PageDTO

	for _, page := range pages {
		pagesDTO = append(pagesDTO, MapPageToPageDTO(page))
	}

	return pagesDTO
}

func MapSectionsToSectionDTO(sections []models.Section) []dto.SectionDTO {
	var sectionsDTO []dto.SectionDTO

	for _, section := range sections {
		sectionsDTO = append(sectionsDTO, MapSectionToSectionDTO(section))
	}

	return sectionsDTO
}

func MapPageContentToPageContentDTO(pageContent models.SectionPageIDContent) dto.PageContentDTO {
	return dto.PageContentDTO{
		SectionID: string(pageContent.SectionID),
		PageID:    string(pageContent.PageID),
		Content:   string(pageContent.Content.Data),
	}
}
