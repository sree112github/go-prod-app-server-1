package companyservice

import (
	companyModels "basics/internal/models/company"
	companyrepo "basics/internal/repository/company"
)

func CreateCompany(company *companyModels.CompanyInputModel) (*companyModels.CompanyResponseModel, error) {

	return companyrepo.CreateCompany(company)
}