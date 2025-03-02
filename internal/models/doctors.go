package models

import (
	"github.com/sev-2/raiden/pkg/db"	
)

type Doctors struct {
	db.ModelBase
	Id         int32  `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('doctors_id_seq')"`
	Nama       string `json:"nama,omitempty" column:"name:nama;type:varchar;nullable:false"`
	Speciality string `json:"speciality,omitempty" column:"name:speciality;type:varchar;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctors" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	ScheduleDoctors []*Schedule `json:"schedule_doctors,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
}
