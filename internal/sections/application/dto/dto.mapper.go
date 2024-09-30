package dto

import (
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
)

func MapSectionToSectionDTO(section models.Section) ResponseSectionDTO {
	return ResponseSectionDTO{
		ID:    string(section.ID),
		Name:  string(section.Name),
		Pages: MapPagesToPagesDTO(section.Pages),
	}

}

func MapPageToPageDTO(page models.Page) PageDTO {
	return PageDTO{
		ID:      string(page.ID),
		Name:    string(page.Name),
		Content: models.PageContent(page.Content).Data,
	}
}

func MapPagesToPagesDTO(pages []models.Page) []PageDTO {
	var pagesDTO []PageDTO

	for _, page := range pages {
		pagesDTO = append(pagesDTO, MapPageToPageDTO(page))
	}

	return pagesDTO
}

func MapSectionsToSectionDTO(sections []models.Section) []ResponseSectionDTO {
	var sectionsDTO []ResponseSectionDTO

	for _, section := range sections {
		sectionsDTO = append(sectionsDTO, MapSectionToSectionDTO(section))
	}

	return sectionsDTO
}

func MapPageContentToPageContentDTO(pageContent models.SectionPageIDContent) PageContentDTO {
	return PageContentDTO{
		SectionID: string(pageContent.SectionID),
		PageID:    string(pageContent.PageID),
		Content:   string(pageContent.Content.Data),
	}
}
