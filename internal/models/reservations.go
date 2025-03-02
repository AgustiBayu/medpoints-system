package models

import (
	"github.com/sev-2/raiden/pkg/db"	
)

type Reservations struct {
	db.ModelBase
	Id         int32  `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('reservations_id_seq')"`
	UserId     int32  `json:"user_id,omitempty" column:"name:user_id;type:integer;nullable:false"`
	ScheduleId int32  `json:"schedule_id,omitempty" column:"name:schedule_id;type:integer;nullable:false"`
	Status     string `json:"status,omitempty" column:"name:status;type:varchar;nullable:false"`
	PaymentId  *int32 `json:"payment_id,omitempty" column:"name:payment_id;type:integer;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"reservations" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Payment  *Payments `json:"payment,omitempty" onUpdate:"no action" onDelete:"set null" join:"joinType:hasOne;primaryKey:id;foreignKey:payment_id"`
	Schedule *Schedule `json:"schedule,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:schedule_id"`
	User     *Users    `json:"user,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
