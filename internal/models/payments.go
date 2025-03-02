package models

import (
	"github.com/sev-2/raiden/pkg/db"	
	"time"
)

type Payments struct {
	db.ModelBase
	Id          int32     `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('payments_id_seq')"`
	Amount      int32     `json:"amount,omitempty" column:"name:amount;type:integer;nullable:false"`
	Status      string    `json:"status,omitempty" column:"name:status;type:varchar;nullable:false"`
	PaymentDate time.Time `json:"payment_date,omitempty" column:"name:payment_date;type:date;nullable:false;default:CURRENT_DATE"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"payments" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	ReservationPayments                 []*Reservations `json:"reservation_payments,omitempty" onUpdate:"no action" onDelete:"set null" join:"joinType:hasMany;primaryKey:id;foreignKey:payment_id"`
	SchedulesThroughReservationsPayment []*Schedule     `json:"schedules_through_reservations_payment,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:payment_id;targetPrimaryKey:id;targetForeign:payment_id"`
	UsersThroughReservationsPayment     []*Users        `json:"users_through_reservations_payment,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:payment_id;targetPrimaryKey:id;targetForeign:payment_id"`
}
