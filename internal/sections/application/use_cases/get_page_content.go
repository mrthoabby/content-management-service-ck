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
		sections: sections,
	}
}

type GetPageContent struct {
	sections ports.SectionProvider
}

func (g *GetPageContent) Execute(context context.Context, params types.GetPageContentParams) models.SectionPageIDContent {
	content, errorGettingContent := g.sections.FetchPageContentByPageIDAsync(context, models.SectionPageID{
		SectionID: models.SectionID(params.SectionID),
		PageID:    models.PageID(params.PageID),
	})
	errorhandler.Handle(errorGettingContent)

	return *content
}
