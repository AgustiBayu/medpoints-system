package types

import (
	"github.com/sev-2/raiden"
)

type Schedule struct {
	raiden.TypeBase
}

func (t *Schedule) Name() string {
	return "_schedule"
}

func (r *Schedule) Format() string {
	return "schedule[]"
}

func (r *Schedule) Enums() []string {
	return []string{}
}

func (r *Schedule) Attributes() []string {
	return []string{}
}

func (r *Schedule) Comment() *string {
	return nil
}

