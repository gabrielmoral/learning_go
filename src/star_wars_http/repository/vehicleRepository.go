package repository

import (
	"star_wars_http/types"
)

func GetVehicles(url string) types.VehiclesResult {
	vehiclesRes := types.VehiclesResult{}

	get(url, &vehiclesRes)

	return vehiclesRes
}
