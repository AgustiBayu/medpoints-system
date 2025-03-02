package types

import (
	"github.com/sev-2/raiden"
)

type Users struct {
	raiden.TypeBase
}

func (t *Users) Name() string {
	return "_users"
}

func (r *Users) Format() string {
	return "users[]"
}

func (r *Users) Enums() []string {
	return []string{}
}

func (r *Users) Attributes() []string {
	return []string{}
}

func (r *Users) Comment() *string {
	return nil
}

