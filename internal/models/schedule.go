package models

import (
	"fmt"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

// Custom type for Date to format it as a string in JSON
type CustomDate time.Time

func (d *CustomDate) MarshalJSON() ([]byte, error) {
	// Format to YYYY-MM-DD
	return []byte(fmt.Sprintf("\"%s\"", time.Time(*d).Format("2006-01-02"))), nil
}

func (d *CustomDate) UnmarshalJSON(b []byte) error {
	// Remove quotes around the string and parse the date format
	parsedTime, err := time.Parse("\"2006-01-02\"", string(b))
	if err != nil {
		return err
	}
	*d = CustomDate(parsedTime)
	return nil
}

// Custom type for Time to format it as a string in JSON
type CustomTime time.Time

func (t *CustomTime) MarshalJSON() ([]byte, error) {
	// Format to HH:MM:SS
	return []byte(fmt.Sprintf("\"%s\"", time.Time(*t).Format("15:04:05"))), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) error {
	// Remove quotes around the string and parse the time format
	parsedTime, err := time.Parse("\"15:04:05\"", string(b))
	if err != nil {
		return err
	}
	*t = CustomTime(parsedTime)
	return nil
}

type Schedule struct {
	db.ModelBase
	Id        int32        `json:"id,omitempty" column:"name:id;type:integer;primaryKey;nullable:false;default:nextval('schedule_id_seq')"`
	DoctorId  int32        `json:"doctor_id,omitempty" column:"name:doctor_id;type:integer;nullable:false"`
	Date      CustomDate   `json:"date,omitempty" column:"name:date;type:date;nullable:false"`
	StartTime CustomTime   `json:"start_time,omitempty" column:"name:start_time;type:time;nullable:false"`
	EndTime   CustomTime   `json:"end_time,omitempty" column:"name:end_time;type:time;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"schedule" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor                              *Doctors        `json:"doctor,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	PaymentsThroughReservationsSchedule []*Payments     `json:"payments_through_reservations_schedule,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:schedule_id;targetPrimaryKey:id;targetForeign:schedule_id"`
	ReservationSchedules                []*Reservations `json:"reservation_schedules,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:schedule_id"`
	UsersThroughReservationsSchedule    []*Users        `json:"users_through_reservations_schedule,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:schedule_id;targetPrimaryKey:id;targetForeign:schedule_id"`
}
