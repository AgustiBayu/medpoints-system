package types

import (
	"github.com/sev-2/raiden"
)

type Reservations struct {
	raiden.TypeBase
}

func (t *Reservations) Name() string {
	return "_reservations"
}

func (r *Reservations) Format() string {
	return "reservations[]"
}

func (r *Reservations) Enums() []string {
	return []string{}
}

func (r *Reservations) Attributes() []string {
	return []string{}
}

func (r *Reservations) Comment() *string {
	return nil
}

