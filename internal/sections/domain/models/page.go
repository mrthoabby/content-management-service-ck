package models

import (
	errorhandler "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler"
	errortypes "github.com/mrthoabby/content-management-service-ck/pkg/commons/error_handler/error_types"
	stringutil "github.com/mrthoabby/content-management-service-ck/pkg/util/string_util"
)

type PageID string

func (p PageID) Validate() errorhandler.Commons {
	if stringutil.IsEmptyString(string(p)) {
		return errortypes.NewDomainError("PageID is required")
	}

	return nil
}

type PageContent struct {
	Content string
}

type PageName string

func (p PageName) Validate() errorhandler.Commons {
	if stringutil.IsEmptyString(string(p)) {
		return errortypes.NewDomainError("PageName is required")
	}

	return nil
}

type PageIDName struct {
	ID   PageID
	Name PageName
}

type Page struct {
	ID      PageID
	Name    PageName
	Content PageContent
}
