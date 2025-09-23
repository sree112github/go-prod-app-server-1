package companyservice

import (
	companyModels "basics/internal/models/company"
	userModels "basics/internal/models/user"

	companyrepo "basics/internal/repository/company"
)



func GetAllCompanyBasedOnScope(company *userModels.AssignRoleInput, limit, page int) ([]companyModels.CompanyResponseModel, error) {

	return companyrepo.GetAllCompanyBasedOnScope(company, limit, page)
}
