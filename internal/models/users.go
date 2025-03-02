package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type Users struct {
	db.ModelBase
	Id       int32  `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('users_id_seq')"`
	Nama     string `json:"nama,omitempty" column:"name:nama;type:varchar;nullable:false"`
	Email    string `json:"email,omitempty" column:"name:email;type:varchar;nullable:false;unique"`
	Phone    string `json:"phone,omitempty" column:"name:phone;type:varchar;nullable:false;unique"`
	Password string `json:"password,omitempty" column:"name:password;type:varchar;nullable:false"`
	RoleId   int32  `json:"role_id,omitempty" column:"name:role_id;type:integer;nullable:false;default:3"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"users" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	PaymentsThroughReservationsUser  []*Payments     `json:"payments_through_reservations_user,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	ReservationUsers                 []*Reservations `json:"reservation_users,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	Role                             *Roles          `json:"role,omitempty" onUpdate:"no action" onDelete:"restrict" join:"joinType:hasOne;primaryKey:id;foreignKey:role_id"`
	SchedulesThroughReservationsUser []*Schedule     `json:"schedules_through_reservations_user,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
}
