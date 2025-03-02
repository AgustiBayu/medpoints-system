package controllers

import (
	"medpointssystem/internal/models"
	"github.com/sev-2/raiden"
)

type ScheduleController struct {
	raiden.ControllerBase
	Http  string `path:"/schedule" type:"rest"`
	Model models.Schedule
}

