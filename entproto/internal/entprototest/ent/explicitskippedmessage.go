// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/explicitskippedmessage"
)

// ExplicitSkippedMessage is the model entity for the ExplicitSkippedMessage schema.
type ExplicitSkippedMessage struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExplicitSkippedMessage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case explicitskippedmessage.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExplicitSkippedMessage fields.
func (esm *ExplicitSkippedMessage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case explicitskippedmessage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			esm.ID = int(value.Int64)
		default:
			esm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ExplicitSkippedMessage.
// This includes values selected through modifiers, order, etc.
func (esm *ExplicitSkippedMessage) Value(name string) (ent.Value, error) {
	return esm.selectValues.Get(name)
}

// Update returns a builder for updating this ExplicitSkippedMessage.
// Note that you need to call ExplicitSkippedMessage.Unwrap() before calling this method if this ExplicitSkippedMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (esm *ExplicitSkippedMessage) Update() *ExplicitSkippedMessageUpdateOne {
	return NewExplicitSkippedMessageClient(esm.config).UpdateOne(esm)
}

// Unwrap unwraps the ExplicitSkippedMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (esm *ExplicitSkippedMessage) Unwrap() *ExplicitSkippedMessage {
	_tx, ok := esm.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExplicitSkippedMessage is not a transactional entity")
	}
	esm.config.driver = _tx.drv
	return esm
}

// String implements the fmt.Stringer.
func (esm *ExplicitSkippedMessage) String() string {
	var builder strings.Builder
	builder.WriteString("ExplicitSkippedMessage(")
	builder.WriteString(fmt.Sprintf("id=%v", esm.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ExplicitSkippedMessages is a parsable slice of ExplicitSkippedMessage.
type ExplicitSkippedMessages []*ExplicitSkippedMessage
