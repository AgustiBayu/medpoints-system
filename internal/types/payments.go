package types

import (
	"github.com/sev-2/raiden"
)

type Payments struct {
	raiden.TypeBase
}

func (t *Payments) Name() string {
	return "_payments"
}

func (r *Payments) Format() string {
	return "payments[]"
}

func (r *Payments) Enums() []string {
	return []string{}
}

func (r *Payments) Attributes() []string {
	return []string{}
}

func (r *Payments) Comment() *string {
	return nil
}

