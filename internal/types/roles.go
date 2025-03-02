package types

import (
	"github.com/sev-2/raiden"
)

type Roles struct {
	raiden.TypeBase
}

func (t *Roles) Name() string {
	return "_roles"
}

func (r *Roles) Format() string {
	return "roles[]"
}

func (r *Roles) Enums() []string {
	return []string{}
}

func (r *Roles) Attributes() []string {
	return []string{}
}

func (r *Roles) Comment() *string {
	return nil
}

