package core

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/adapters/api"
	v1 "github.com/mrthoabby/content-management-service-ck/internal/sections/adapters/api/v1"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/adapters/repository"
	"github.com/mrthoabby/content-management-service-ck/internal/sections/adapters/service"
	usecases "github.com/mrthoabby/content-management-service-ck/internal/sections/application/use_cases"
)

var cleaners []func(context.Context)

func initializeSections(globalRouter *mux.Router) {

	cleaners = append(cleaners, repository.CleanUp)
	repository := repository.NewSectionProvider()

	commands := usecases.UseCasesCommands{
		GetISectionByID:      usecases.NewGetSectionById(repository),
		GetSections:          usecases.NewGetAllSections(repository),
		GetPageContentByID:   usecases.NewGetPageContent(repository),
		GetSectionsWithQuery: usecases.NewGetSectionsByQuery(repository),
		CreateNewSection:     usecases.NewCreateSection(repository),
	}

	service := service.NewSectionService(commands)

	handler := v1.NewSectionHandler(service)

	router := api.NewSectionRouter(handler, globalRouter)

	router.InitialiceSectionRouter()
}

func RunIoc(globalRouter *mux.Router) {

	initializeSections(globalRouter)
}

func IOCCleanUp(context context.Context) {
	for _, cleaner := range cleaners {
		cleaner(context)
	}
}
