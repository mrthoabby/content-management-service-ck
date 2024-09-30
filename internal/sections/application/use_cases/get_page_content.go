package usecases

import (
	"context"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/application/types"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/ports"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
)

func NewGetPageContent(sections ports.SectionProvider) *GetPageContent {
	return &GetPageContent{
		SectionProvider: sections,
	}
}

type GetPageContent struct {
	ports.SectionProvider
}

func (g *GetPageContent) Execute(context context.Context, params types.GetPageContentParams) models.SectionPageIDContent {
	content, errorGettingContent := g.FetchPageContentByPageIDAsync(context, models.SectionPageID{
		SectionID: models.SectionID(params.SectionID),
		PageID:    models.PageID(params.PageID),
	})
	errorhandler.Handle(errorGettingContent, g, "error getting page content", "usecase.get_page_content")

	return *content
}
