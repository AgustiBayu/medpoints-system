package controllers

import (
	"medpointssystem/internal/models"
	"github.com/sev-2/raiden"
)

type DoctorController struct {
	raiden.ControllerBase
	Http  string `path:"/doctor" type:"rest"`
	Model models.Doctors
}