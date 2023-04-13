// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/implicitskippedmessage"
)

// ImplicitSkippedMessage is the model entity for the ImplicitSkippedMessage schema.
type ImplicitSkippedMessage struct {
	config
	// ID of the ent.
	ID                         int `json:"id,omitempty"`
	depends_on_skipped_skipped *int
	selectValues               sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ImplicitSkippedMessage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case implicitskippedmessage.FieldID:
			values[i] = new(sql.NullInt64)
		case implicitskippedmessage.ForeignKeys[0]: // depends_on_skipped_skipped
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ImplicitSkippedMessage fields.
func (ism *ImplicitSkippedMessage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case implicitskippedmessage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ism.ID = int(value.Int64)
		case implicitskippedmessage.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field depends_on_skipped_skipped", value)
			} else if value.Valid {
				ism.depends_on_skipped_skipped = new(int)
				*ism.depends_on_skipped_skipped = int(value.Int64)
			}
		default:
			ism.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ImplicitSkippedMessage.
// This includes values selected through modifiers, order, etc.
func (ism *ImplicitSkippedMessage) Value(name string) (ent.Value, error) {
	return ism.selectValues.Get(name)
}

// Update returns a builder for updating this ImplicitSkippedMessage.
// Note that you need to call ImplicitSkippedMessage.Unwrap() before calling this method if this ImplicitSkippedMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (ism *ImplicitSkippedMessage) Update() *ImplicitSkippedMessageUpdateOne {
	return NewImplicitSkippedMessageClient(ism.config).UpdateOne(ism)
}

// Unwrap unwraps the ImplicitSkippedMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ism *ImplicitSkippedMessage) Unwrap() *ImplicitSkippedMessage {
	_tx, ok := ism.config.driver.(*txDriver)
	if !ok {
		panic("ent: ImplicitSkippedMessage is not a transactional entity")
	}
	ism.config.driver = _tx.drv
	return ism
}

// String implements the fmt.Stringer.
func (ism *ImplicitSkippedMessage) String() string {
	var builder strings.Builder
	builder.WriteString("ImplicitSkippedMessage(")
	builder.WriteString(fmt.Sprintf("id=%v", ism.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ImplicitSkippedMessages is a parsable slice of ImplicitSkippedMessage.
type ImplicitSkippedMessages []*ImplicitSkippedMessage
