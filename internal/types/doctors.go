package types

import (
	"github.com/sev-2/raiden"
)

type Doctors struct {
	raiden.TypeBase
}

func (t *Doctors) Name() string {
	return "_doctors"
}

func (r *Doctors) Format() string {
	return "doctors[]"
}

func (r *Doctors) Enums() []string {
	return []string{}
}

func (r *Doctors) Attributes() []string {
	return []string{}
}

func (r *Doctors) Comment() *string {
	return nil
}

