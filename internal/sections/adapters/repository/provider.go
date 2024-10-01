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

const (
	sectionIdProperty   = "id"
	sectionNameProperty = "name"

	pagesProperty               = "pages"
	pageIdProperty              = "pages.id"
	pageNameProperty            = "pages.name"
	pageListNameProperty        = "pages.$.name"
	pageContentProperty         = "pages.content"
	PageListContentDataProperty = "pages.$.content"
)

var env struct {
	userName       string
	password       string
	aside          string
	databaseName   string
	collectionName string
}

var currentClient *mongo.Client

func loadProperties() {
	env.userName = os.Getenv("SECTION_PROVIDER_REPOSITORY_USERNAME")
	env.password = os.Getenv("SECTION_PROVIDER_REPOSITORY_PASSWORD")
	env.aside = os.Getenv("SECTION_PROVIDER_REPOSITORY_ASIDE")
	env.databaseName = os.Getenv("SECTION_PROVIDER_REPOSITORY_DATABASE_NAME")
	env.collectionName = os.Getenv("SECTION_PROVIDER_REPOSITORY_COLLECTION_NAME")
}

func ensureIndexes(ctx context.Context) {
	db := currentClient.Database(env.databaseName)
	collection := db.Collection(env.collectionName)

	idIndexModel := mongo.IndexModel{
		Options: options.Index().SetUnique(true),
		Keys:    bson.M{sectionIdProperty: 1},
	}

	searcher := mongo.IndexModel{
		Options: options.Index().SetName("text_searcher").SetDefaultLanguage("spanish"),

		Keys: bson.D{
			{Key: sectionNameProperty, Value: 1},
			{Key: pageNameProperty, Value: 1},
			{Key: pageContentProperty, Value: 2},
		},
	}

	index, errorSettingIndex := collection.Indexes().CreateMany(ctx, []mongo.IndexModel{idIndexModel, searcher})
	errorhandler.Handle(errorSettingIndex, collection, "creating indexes", "provider")

	logrus.Println("Indexes  : ", index)
}

func loadMongoDb() {
	context, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + env.userName + ":" + env.password + "@" + env.aside).SetServerAPIOptions(serverAPI)

	client, errorConnecting := mongo.Connect(context, clientOptions)
	errorhandler.Handle(errorConnecting, clientOptions, "loading mongo db", "provider")

	currentClient = client

	ensureIndexes(context)

	pingDB(context)
}

func CleanUp(context context.Context) {
	if err := currentClient.Disconnect(context); err != nil {
		errorhandler.Handle(err, currentClient, "cleaning up", "provider")
	} else {
		logrus.Println("Disconnected from MongoDB!")
	}
}

func pingDB(ctx context.Context) {
	if err := currentClient.Database(env.databaseName).RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		errorhandler.Handle(err, currentClient, "pinging db", "provider")
	} else {
		logrus.Println("Pinged. You successfully connected to MongoDB!")
	}
}

func NewSectionProvider() *SectionProvider {
	loadProperties()
	loadMongoDb()
	return &SectionProvider{
		Collection: currentClient.Database(env.databaseName).Collection("sections"),
	}
}

type SectionProvider struct {
	*mongo.Collection
}

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

func (s *SectionProvider) FetchPageContentByPageIDAsync(context context.Context, params models.SectionPageID) (*models.SectionPageIDContent, error) {
	var pageContent *models.SectionPageIDContent

	filter := bson.D{{Key: sectionIdProperty, Value: params.SectionID}, {Key: pageIdProperty, Value: params.PageID}}

	projection := bson.D{
		{Key: sectionIdProperty, Value: 1},
		{Key: pageIdProperty, Value: 1},
		{Key: pageContentProperty, Value: 1},
	}

	opts := options.FindOne().SetProjection(projection)

	errorFinding := s.Collection.FindOne(context, filter, opts).Decode(pageContent)
	if errorFinding != nil {
		if errorFinding == mongo.ErrNoDocuments {
			return nil, errortypes.NewNotFoundError("Page not found")
		}
		return nil, errorFinding
	}

	return pageContent, nil
}

func (s *SectionProvider) FetchSectionsByQueryAsync(context context.Context, query string) ([]models.Section, error) {
	const regexTag = "$regex"

	filter := bson.M{
		"$or": bson.A{
			bson.M{sectionNameProperty: bson.M{regexTag: query}},
			bson.M{pageNameProperty: bson.M{regexTag: query}},
			bson.M{pageContentProperty: bson.M{regexTag: query}},
		},
	}

	options := options.Find()

	cursor, errorFinding := s.Collection.Find(context, filter, options)
	if errorFinding != nil {
		if errorFinding == mongo.ErrNoDocuments {
			return nil, errortypes.NewNotFoundError("Sections not found")
		}
		return nil, errorFinding
	}
	defer cursor.Close(context)

	var sections []models.Section
	if errorDecoding := cursor.All(context, &sections); errorDecoding != nil {
		return nil, errorDecoding
	}

	return sections, nil
}

func (s *SectionProvider) FetchPartialSectionsByQueryAsync(context context.Context, query string) ([]models.PartialSection, error) {
	const regexTag = "$regex"

	filter := bson.M{
		"$or": bson.A{
			bson.M{sectionNameProperty: bson.M{regexTag: query}},
			bson.M{pageNameProperty: bson.M{regexTag: query}},
			bson.M{pageContentProperty: bson.M{regexTag: query}},
		},
	}

	options := options.Find().SetProjection(bson.D{
		{Key: pageContentProperty, Value: 0},
	})

	cursor, errorFinding := s.Collection.Find(context, filter, options)
	if errorFinding != nil {
		if errorFinding == mongo.ErrNoDocuments {
			return nil, errortypes.NewNotFoundError("Sections not found")
		}
		return nil, errorFinding
	}
	defer cursor.Close(context)

	var sections []models.PartialSection
	if errorDecoding := cursor.All(context, &sections); errorDecoding != nil {
		return nil, errorDecoding
	}

	return sections, nil
}

func (s *SectionProvider) CreateSectionAsync(context context.Context, model models.SectionIDName) error {
	insertionResult, errorInserting := s.Collection.InsertOne(context, bson.M{
		sectionIdProperty:   model.SectionID,
		sectionNameProperty: model.SectionName,
	})
	if errorInserting != nil {
		if mongo.IsDuplicateKeyError(errorInserting) {
			return errortypes.NewConflictError("Section already exists")
		}

		return errorInserting
	}

	logrus.Println("Section created: ", insertionResult.InsertedID)

	return nil
}

func (s *SectionProvider) CreateSectionPageAsync(context context.Context, model models.SectionPageIDPageName) error {
	filter := bson.M{
		sectionIdProperty: model.SectionID,
		pageIdProperty:    model.NetPageID.ID,
	}

	pageExistResult, errorFinding := s.Collection.CountDocuments(context, filter)
	if errorFinding != nil {
		return errorFinding
	}

	if pageExistResult > 0 {
		return errortypes.NewConflictError("Page already exists in section")
	}

	filter = bson.M{
		sectionIdProperty: model.SectionID,
	}

	update := bson.M{
		"$push": bson.M{
			pagesProperty: models.Page{
				ID:   model.NetPageID.ID,
				Name: model.NetPageID.Name,
			},
		},
	}

	updateResult, errorUpdating := s.Collection.UpdateOne(context, filter, update)
	if errorUpdating != nil {
		return errorUpdating
	}

	if updateResult.MatchedCount == 0 {
		return errortypes.NewNotFoundError("Section not found")
	}

	return nil
}

func (s *SectionProvider) UpdateSectionAsync(context context.Context, model models.SectionIDName) error {
	changes := 0

	filter := bson.M{
		sectionIdProperty: model.SectionID,
	}
	update := bson.M{}

	if model.SectionName != "" {
		changes++
		update[sectionNameProperty] = model.SectionName
	}

	if changes > 0 {
		updateResult, errorUpdating := s.Collection.UpdateOne(context, filter, bson.M{
			"$set": update,
		})
		if errorUpdating != nil {
			if errorUpdating == mongo.ErrNoDocuments {
				return errortypes.NewNotFoundError("Section not found")
			}
			return errorUpdating
		}

		if updateResult.MatchedCount == 0 {
			return errortypes.NewNotFoundError("Section not found")
		}
	}

	return nil

}

func (s *SectionProvider) UpdateSectionPageAsync(context context.Context, model models.SectionIDPageIDContent) error {
	changes := 0

	filter := bson.M{
		sectionIdProperty: model.SectionID,
		pageIdProperty:    model.PageID,
	}

	update := bson.M{}

	if model.Content.Data != "" {
		changes++
		update[PageListContentDataProperty] = bson.M{
			"data": model.Content.Data,
		}
	}

	if model.PageName != "" {
		changes++
		update[pageListNameProperty] = model.PageName
	}

	if changes > 0 {

		updateResult, errorUpdating := s.Collection.UpdateOne(context, filter, bson.M{
			"$set": update,
		})
		if errorUpdating != nil {
			if errorUpdating == mongo.ErrNoDocuments {
				return errortypes.NewNotFoundError("Section not found")
			}
			return errorUpdating
		}

		if updateResult.MatchedCount == 0 {
			return errortypes.NewNotFoundError("Page not found")
		}
	}

	return nil
}

func (s *SectionProvider) DeleteSectionPageByIDAsync(context context.Context, model models.SectionPageID) error {
	update := bson.M{
		"$pull": bson.M{
			pagesProperty: bson.M{
				"id": model.PageID,
			},
		},
	}

	result, errorUpdating := s.Collection.UpdateOne(context, bson.M{
		sectionIdProperty: model.SectionID,
	}, update)

	if errorUpdating != nil {
		return errorUpdating
	}

	if result.MatchedCount == 0 {
		return errortypes.NewNotFoundError("Page not found")
	}

	return nil
}

func (s *SectionProvider) DeleteSectionByIDAsync(context context.Context, sectionID models.SectionID) error {
	deleteResult, errorDeleting := s.Collection.DeleteOne(context, bson.M{
		sectionIdProperty: sectionID,
	})

	if errorDeleting != nil {
		return errorDeleting
	}

	if deleteResult.DeletedCount == 0 {
		return errortypes.NewNotFoundError("Section not found")
	}

	return nil
}
