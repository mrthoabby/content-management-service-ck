package repository

import (
	"context"
	"os"
	"time"

	"github.com/mrthoabby/content-management-service-ck/internal/sections/domain/models"
	coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
	errortypes "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler/error_types"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var property struct {
	userName     string
	password     string
	aside        string
	databaseName string
}

var currentClient *mongo.Client

func loadProperties() {
	property.userName = os.Getenv("SECTION_PROVIDER_REPOSITORY_USERNAME")
	property.password = os.Getenv("SECTION_PROVIDER_REPOSITORY_PASSWORD")
	property.aside = os.Getenv("SECTION_PROVIDER_REPOSITORY_ASIDE")
	property.databaseName = os.Getenv("SECTION_PROVIDER_REPOSITORY_DATABASE_NAME")
}

func loadMongoDb() {
	context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + property.userName + ":" + property.password + "@" + property.aside).SetServerAPIOptions(serverAPI)

	client, errorConnecting := mongo.Connect(context, clientOptions)
	errorhandler.Handle(errorConnecting)

	currentClient = client

	pingDB(context)
}

func CleanUp(context context.Context) {
	if err := currentClient.Disconnect(context); err != nil {
		errorhandler.Handle(err)
	} else {
		logrus.Println("Disconnected from MongoDB!")
	}
}

func pingDB(ctx context.Context) {
	if err := currentClient.Database(property.databaseName).RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		errorhandler.Handle(err)
	} else {
		logrus.Println("Pinged. You successfully connected to MongoDB!")
	}
}

func NewSectionProvider() *SectionProvider {
	loadProperties()
	loadMongoDb()
	return &SectionProvider{
		Collection: currentClient.Database(property.databaseName).Collection("sections"),
	}
}

type SectionProvider struct {
	*mongo.Collection
}

const (
	sectionIdProperty   = "id"
	sectionNameProperty = "name"

	pageIdProperty      = "id"
	pageContentProperty = "pages.content"
)

func (s *SectionProvider) FetchSectionByIDAsync(context context.Context, sectionId models.SectionID) (*models.Section, error) {
	var section *models.Section

	filter := bson.D{{Key: sectionIdProperty, Value: sectionId}}
	errorFinding := s.Collection.FindOne(context, filter).Decode(section)
	if errorFinding != nil {
		if errorFinding == mongo.ErrNoDocuments {
			return nil, errortypes.NewNotFoundError("Section not found")
		}
		return nil, errorFinding
	}

	return section, nil
}

func (s *SectionProvider) FetchPartialSectionByIDAsync(context context.Context, sectionId models.SectionID) (*models.PartialSection, error) {
	var section *models.PartialSection

	filter := bson.D{{Key: sectionIdProperty, Value: sectionId}}

	projection := bson.D{
		{Key: sectionIdProperty, Value: 1},
		{Key: sectionNameProperty, Value: 1},
	}

	opts := options.FindOne().SetProjection(projection)

	errorFinding := s.Collection.FindOne(context, filter, opts).Decode(section)
	if errorFinding != nil {
		if errorFinding == mongo.ErrNoDocuments {
			return nil, errortypes.NewNotFoundError("Section not found")
		}
		return nil, errorFinding
	}

	return section, nil

}

func (s *SectionProvider) FetchAllSectionsAsync(context context.Context, pagination coredomain.Pagination) (coredomain.PaginatedResult[[]models.Section], error) {

	sections := make([]models.Section, 0)
	paginationResult := coredomain.PaginatedResult[[]models.Section]{
		CurrentPage: pagination.CurrentPage,
		Data:        sections,
		GroupedBy:   pagination.GroupBy,
		CountTotal:  0,
		TotalPages:  0,
	}

	skip := (pagination.CurrentPage - 1) * pagination.GroupBy

	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(pagination.GroupBy))

	cursor, errorFinding := s.Collection.Find(context, bson.D{}, opts)
	if errorFinding != nil {
		if errorFinding == mongo.ErrNoDocuments {
			return paginationResult, nil
		}
		return paginationResult, errorFinding
	}
	defer cursor.Close(context)

	if errorDecoding := cursor.All(context, &sections); errorDecoding != nil {
		return paginationResult, errorDecoding
	}

	totalDocs, errorCountingDocument := s.Collection.CountDocuments(context, bson.D{})
	if errorCountingDocument != nil {
		return paginationResult, errorCountingDocument
	}

	totalGroups := int(totalDocs) / pagination.GroupBy
	if totalDocs%int64(pagination.GroupBy) != 0 {
		totalGroups++
	}

	return coredomain.PaginatedResult[[]models.Section]{
		CurrentPage: pagination.CurrentPage,
		Data:        sections,
		GroupedBy:   pagination.GroupBy,
		CountTotal:  int(totalDocs),
		TotalPages:  totalGroups,
	}, nil
}

func (s *SectionProvider) FetchAllPartialSectionsAsync(context context.Context, pagination coredomain.Pagination) (coredomain.PaginatedResult[[]models.PartialSection], error) {

	sections := make([]models.PartialSection, 0)
	paginationResult := coredomain.PaginatedResult[[]models.PartialSection]{
		CurrentPage: pagination.CurrentPage,
		Data:        sections,
		GroupedBy:   pagination.GroupBy,
		CountTotal:  0,
		TotalPages:  0,
	}

	skip := (pagination.CurrentPage - 1) * pagination.GroupBy

	projection := bson.D{
		{Key: pageContentProperty, Value: 0},
	}

	opts := options.Find().SetProjection(projection).SetSkip(int64(skip)).SetLimit(int64(pagination.GroupBy))

	cursor, errorFinding := s.Collection.Find(context, bson.D{}, opts)
	if errorFinding != nil {
		if errorFinding == mongo.ErrNoDocuments {
			return paginationResult, nil
		}
		return paginationResult, errorFinding
	}
	defer cursor.Close(context)

	if errorDecoding := cursor.All(context, &sections); errorDecoding != nil {
		return paginationResult, errorDecoding
	}

	totalDocs, errorCountingDocument := s.Collection.CountDocuments(context, bson.D{})
	if errorCountingDocument != nil {
		return paginationResult, errorCountingDocument
	}

	totalGroups := int(totalDocs) / pagination.GroupBy
	if totalDocs%int64(pagination.GroupBy) != 0 {
		totalGroups++
	}

	return coredomain.PaginatedResult[[]models.PartialSection]{
		CurrentPage: pagination.CurrentPage,
		Data:        sections,
		GroupedBy:   pagination.GroupBy,
		CountTotal:  int(totalDocs),
		TotalPages:  totalGroups,
	}, nil
}

func (s *SectionProvider) FetchSectionPageContentBySectionPageIDAsync(_ context.Context, _ models.SectionPageID) (*models.PageContent, error) {
	panic("not implemented") // TODO: Implement
}

func (s *SectionProvider) FetchPartialSectionsByQueryPaginatedAsync(_ context.Context, _ models.SectionID) ([]models.PartialSection, error) {
	panic("not implemented") // TODO: Implement
}

func (s *SectionProvider) CreateSectionAsync(_ context.Context, _ models.Section) error {
	panic("not implemented") // TODO: Implement
}

func (s *SectionProvider) CreateSectionPageAsync(_ context.Context, _ models.PageIDName) error {
	panic("not implemented") // TODO: Implement
}

func (s *SectionProvider) UpdateSectionPageContentAsync(_ context.Context, _ models.SectionPageIDContent) error {
	panic("not implemented") // TODO: Implement
}

func (s *SectionProvider) UpdateSectionPageNameAsync(_ context.Context, _ models.SectionPageIDPageName) error {
	panic("not implemented") // TODO: Implement
}

func (s *SectionProvider) UpdateSectionNameAsync(_ context.Context, _ models.SectionPageIDName) error {
	panic("not implemented") // TODO: Implement
}

func (s *SectionProvider) DeleteSectionPageByIDAsync(_ context.Context, _ models.PageID) error {
	panic("not implemented") // TODO: Implement
}

func (s *SectionProvider) DeleteSectionByIDAsync(_ context.Context, _ models.SectionID) error {
	panic("not implemented") // TODO: Implement
}
