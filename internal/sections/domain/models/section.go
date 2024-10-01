package models

import (
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
	errortypes "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler/error_types"
	stringutil "github.com/mrthoabby/content-management-service-ck/pkg/util/string_util"
)

type SectionID string

func (s SectionID) Validate() errorhandler.Commons {
	if stringutil.IsEmptyString(string(s)) {
		return errortypes.NewDomainError("SectionID is required")
	}

	return nil
}

type SectionName string

func (s SectionName) Validate() errorhandler.Commons {
	if stringutil.IsEmptyString(string(s)) {
		return errortypes.NewDomainError("SectionName is required")
	}

	return nil
}

type PartialSection struct {
	ID    SectionID
	Name  SectionName
	Pages []PageIDName
}

type Section struct {
	ID    SectionID
	Name  SectionName
	Pages []Page
}

type SectionPageID struct {
	SectionID SectionID
	PageID    PageID
}

type SectionPageIDContent struct {
	SectionID SectionID
	PageID    PageID
	Content   PageContent
}

type SectionPageIDPageName struct {
	SectionID SectionID
	NetPageID PageIDName
}

type SectionIDPageIDContent struct {
	SectionID SectionID
	PageID    PageID
	PageName  PageName
	Content   PageContent
}

type SectionIDName struct {
	SectionID   SectionID
	SectionName SectionName
}
