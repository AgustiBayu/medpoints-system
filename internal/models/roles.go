package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type Roles struct {
	db.ModelBase
	Id   int32  `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('roles_id_seq')"`
	Name string `json:"name,omitempty" column:"name:name;type:varchar;nullable:false;unique"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"roles" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	UserRoles []*Users `json:"user_roles,omitempty" onUpdate:"no action" onDelete:"restrict" join:"joinType:hasMany;primaryKey:id;foreignKey:role_id"`
}
