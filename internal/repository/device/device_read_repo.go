package devicerepo

import (
	"basics/internal/config"
	deviceModels "basics/internal/models/device"
	userModels "basics/internal/models/user"
	"database/sql"
	"fmt"
)

func GetAllDevicesBasedOnSCope(user *userModels.AssignRoleInput, limit, page int, search string) ([]deviceModels.GetAllDeviceResponseStructure, error) {

	var rows *sql.Rows
	var err error

	offset := (page - 1) * limit

	switch user.Scope {

	case "super_admin":

		if search == "" {
			rows, err = config.DB.Query(
				`Select d.device_id,d.name,m.machine_id,m.machine_name,p.plant_id,p.plant_name,c.company_id,c.company_name from devices d
			join machines m on d.machine_id = m.machine_id
			join plants p on m.plant_id = p.plant_id
			join companies c on p.company_id = c.company_id 
			order by d.created_at desc limit $1 offset $2`, limit, offset,
			)
		} else {
			rows, err = config.DB.Query(
				`Select d.device_id,d.name,m.machine_id,m.machine_name,p.plant_id,p.plant_name,c.company_id,c.company_name from devices d
			join machines m on d.machine_id = m.machine_id
			join plants p on m.plant_id = p.plant_id
			join companies c on p.company_id = c.company_id
			where d.name ILIKE '%' || $1 || '%' or m.machine_name ILIKE '%' || $1 || '%'
			order by d.created_at desc limit $2 offset $3`, search, limit, offset,
			)
		}
	case "company_admin":
		rows, err = config.DB.Query(
			`Select d.device_id,d.name,m.machine_id,m.machine_name,p.plant_id,p.plant_name,c.company_id,c.company_name from devices d
			join machines m on d.machine_id = m.machine_id
			join plants p on m.plant_id = p.plant_id
			join companies c on p.company_id = c.company_id 
			where c.company_id = $1
			order by d.created_at desc limit $2 offset $3`, user.CompanyId, limit, offset,
		)
	case "plant_admin":
		rows, err = config.DB.Query(
			`Select d.device_id,d.name,m.machine_id,m.machine_name,p.plant_id,p.plant_name,c.company_id,c.company_name from devices d
			join machines m on d.machine_id = m.machine_id
			join plants p on m.plant_id = p.plant_id
			join companies c on p.company_id = c.company_id
			where p.plant_id =$1
			order by d.created_at desc limit $2 offset $3`, user.PlantId, limit, offset,
		)
	case "machine_admin":
		rows, err = config.DB.Query(
			`Select d.device_id,d.name,m.machine_id,m.machine_name,p.plant_id,p.plant_name,c.company_id,c.company_name from devices d
			join machines m on d.machine_id = m.machine_id
			join plants p on m.plant_id = p.plant_id
			join companies c on p.company_id = c.company_id 
			where m.machine_id = $1
			order by d.created_at desc limit $2 offset $3`, user.MachineId, limit, offset,
		)

	default:
		return nil, fmt.Errorf("error due to undinined %v", err)
	}

	if err != nil {
		return nil, fmt.Errorf("the error Occured in the scope %v Due to %v", user.Scope, err)
	}
	var devices []deviceModels.GetAllDeviceResponseStructure

	for rows.Next() {
		var d deviceModels.GetAllDeviceResponseStructure

		if err := rows.Scan(&d.DeviceId, &d.DeviceName, &d.MachineId, &d.MachineName, &d.PlantId, &d.PlantName, &d.CompanyId, &d.CompanyName); err != nil {
			return nil, err
		}

		devices = append(devices, d)
	}
	return devices, nil
}
