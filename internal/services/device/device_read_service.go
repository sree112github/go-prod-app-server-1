package deviceservice

import (
	deviceModels "basics/internal/models/device"
	userModels "basics/internal/models/user"
	devicerepo "basics/internal/repository/device"
)

func GetAllDevicesBasedOnSCope(user *userModels.AssignRoleInput, limit, page int, search string) ([]deviceModels.GetAllDeviceResponseStructure, error) {

	return devicerepo.GetAllDevicesBasedOnSCope(user, limit, page, search)
}
