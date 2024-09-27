package domain

import valueobjects "github.com/mrthoabby/content-management-service-ck/internal/sections/value_objects"

type SectionID string
type SectionName string

type NetSection struct {
	ID    SectionID
	Name  SectionName
	Pages []valueobjects.PageIDName
}

type Section struct {
	ID    SectionID
	Name  SectionName
	Pages []valueobjects.Page
}

type SectionPageID struct {
	SectionID SectionID
	PageID    valueobjects.PageID
}

type SectionPageIDContent struct {
	SectionID SectionID
	PageID    valueobjects.PageID
	Content   valueobjects.PageContent
}

type SectionPageIDPageName struct {
	SectionID SectionID
	NetPageID valueobjects.PageIDName
}

type SectionPageIDName struct {
	SectionID   SectionID
	SectionName SectionName
}
