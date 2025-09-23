package userservice

import (
	userModels "basics/internal/models/user"
	companyrepo "basics/internal/repository/company"
	machinerepo "basics/internal/repository/machine"
	plantrepo "basics/internal/repository/plant"
	userrepo "basics/internal/repository/user"
)

func AssignUserRole(input *userModels.AssignRoleInput) (*userModels.AssignRoleResponse, error) {

	

	switch input.Scope {

	case "company_admin":
		isCompanyExist, err := companyrepo.IsCompanyExist(input.CompanyId)

		if !isCompanyExist {
			return nil, err
		}

	case "plant_admin":
		isPlantExists, err := plantrepo.IsPlantExists(*input.CompanyId,*input.PlantId)
		if !isPlantExists {
			return nil, err
		}

	case "machine_user":
		isMachineExists, err := machinerepo.IsMachineExists(*input.CompanyId,*input.PlantId,*input.MachineId)
		if !isMachineExists {
			return nil, err
		}
	}

	user, err := userrepo.AssignUserRole(input)
	if err != nil {
		return nil, err
	}

	return user, nil

}
