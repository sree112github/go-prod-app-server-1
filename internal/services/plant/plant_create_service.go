package plantservice

import (
	plantModels "basics/internal/models/plant"

	companyrepo "basics/internal/repository/company"
	plantrepo "basics/internal/repository/plant"
	"fmt"
)

func CreatePlant(plant *plantModels.PlantInputModel) (*plantModels.PlantResponseModel, error) {

	companyUUID := plant.CompanyId.String()

	ok, err := companyrepo.IsCompanyExist(&companyUUID)

	if !ok {
		fmt.Println("error occured in the Service Layer During Plant Creation Due to company id")
		return nil, err
	}

	response, err := plantrepo.CreatePlant(plant)
	if err != nil {
		fmt.Println("error occured in the Service Layer During Plant Creation")
		return nil, err
	}

	return response, nil
}
