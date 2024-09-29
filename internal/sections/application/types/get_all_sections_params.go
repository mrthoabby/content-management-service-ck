package types

import coredomain "github.com/mrthoabby/content-management-service-ck/pkg/commons/domain"

type GetAllSectionsParams struct {
	LoadPages  bool
	Pagination coredomain.Pagination
}
