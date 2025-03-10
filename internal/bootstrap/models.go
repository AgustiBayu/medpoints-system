// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden/pkg/resource"
	"medpointssystem/internal/models"
)

func RegisterModels() {
	resource.RegisterModels(
		&models.Doctors{},
		&models.Payments{},
		&models.Reservations{},
		&models.Roles{},
		&models.Schedule{},
		&models.Users{},
	)
}
