package userrepo

import (
	"basics/internal/config"
	userModels "basics/internal/models/user"
	"fmt"
)

func AssignUserRole(input *userModels.AssignRoleInput) (*userModels.AssignRoleResponse, error) {

	query := `
	UPDATE users
	SET scope = $2,
	company_id =$3,
	plant_id =$4,
	machine_id =$5
	WHERE user_id = $1
	RETURNING user_id,name,email,scope,company_id,plant_id,machine_id,created_at
	`
	var response userModels.AssignRoleResponse
	err := config.DB.QueryRow(
		query, input.UserId, input.Scope,
		input.CompanyId, input.PlantId,
		input.MachineId).Scan(
		&response.UserId, &response.UserName, &response.Email,
		&response.Scope, &response.CompanyId, &response.PlantId,
		&response.MachineId, &response.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to assign user role: %w", err)
	}

	return &response, nil

}
