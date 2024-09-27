package valueobjects

type PageContent string
type PageID string
type PageName string

type PageIDName struct {
	ID   PageID
	Name PageName
}

type Page struct {
	ID      PageID
	Name    PageName
	Content PageContent
}
