package plantservice

import (
	plantModels "basics/internal/models/plant"
	companyrepo "basics/internal/repository/company"
	machinerepo "basics/internal/repository/machine"
	plantrepo "basics/internal/repository/plant"
	"fmt"
)

func GetAllPlantsBasedOnScope(plant *plantModels.PlantInputModel, limit int, page int) ([]plantModels.PlantResponseModel, error) {

	companyId := plant.CompanyId.String()
	plantId := plant.PlantId.String()
	machineId := plant.MachineId.String()

	switch plant.Scope {

	case "company_admin":

		isCompanyExist, err := companyrepo.IsCompanyExist(&companyId)

		if !isCompanyExist {
			return nil, err
		}

	case "plant_admin":
		isPlantExists, err := plantrepo.IsPlantExists(companyId, plantId)
		fmt.Println("the Plant check success",isPlantExists)
		if !isPlantExists {
			return nil, err
		}

	case "machine_user":
		isMachineExists, err := machinerepo.IsMachineExists(companyId, plantId, machineId)
		if !isMachineExists {
			return nil, err
		}
	}

	plants, err := plantrepo.GetAllPlantsBasedOnScope(plant, limit, page)

	if err != nil {
		fmt.Println("Something went wrong in the SCope repo:", err)

		return nil, err
	}

	return plants, nil
}
