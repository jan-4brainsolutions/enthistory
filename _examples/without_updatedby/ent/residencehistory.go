// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"_examples/without_updatedby/ent/residencehistory"

	"github.com/flume/enthistory"
)

// ResidenceHistory is the model entity for the ResidenceHistory schema.
type ResidenceHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation enthistory.OpType `json:"operation,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref uuid.UUID `json:"ref,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ResidenceHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case residencehistory.FieldID:
			values[i] = new(sql.NullInt64)
		case residencehistory.FieldOperation, residencehistory.FieldName:
			values[i] = new(sql.NullString)
		case residencehistory.FieldHistoryTime, residencehistory.FieldCreatedAt, residencehistory.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case residencehistory.FieldRef:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ResidenceHistory fields.
func (rh *ResidenceHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case residencehistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rh.ID = int(value.Int64)
		case residencehistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				rh.HistoryTime = value.Time
			}
		case residencehistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				rh.Operation = enthistory.OpType(value.String)
			}
		case residencehistory.FieldRef:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value != nil {
				rh.Ref = *value
			}
		case residencehistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rh.CreatedAt = value.Time
			}
		case residencehistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rh.UpdatedAt = value.Time
			}
		case residencehistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rh.Name = value.String
			}
		default:
			rh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ResidenceHistory.
// This includes values selected through modifiers, order, etc.
func (rh *ResidenceHistory) Value(name string) (ent.Value, error) {
	return rh.selectValues.Get(name)
}

// Update returns a builder for updating this ResidenceHistory.
// Note that you need to call ResidenceHistory.Unwrap() before calling this method if this ResidenceHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (rh *ResidenceHistory) Update() *ResidenceHistoryUpdateOne {
	return NewResidenceHistoryClient(rh.config).UpdateOne(rh)
}

// Unwrap unwraps the ResidenceHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rh *ResidenceHistory) Unwrap() *ResidenceHistory {
	_tx, ok := rh.config.driver.(*txDriver)
	if !ok {
		panic("ent: ResidenceHistory is not a transactional entity")
	}
	rh.config.driver = _tx.drv
	return rh
}

// String implements the fmt.Stringer.
func (rh *ResidenceHistory) String() string {
	var builder strings.Builder
	builder.WriteString("ResidenceHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(rh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", rh.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fmt.Sprintf("%v", rh.Ref))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(rh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(rh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(rh.Name)
	builder.WriteByte(')')
	return builder.String()
}

// ResidenceHistories is a parsable slice of ResidenceHistory.
type ResidenceHistories []*ResidenceHistory
