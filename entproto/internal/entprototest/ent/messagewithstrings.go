// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/messagewithstrings"
)

// MessageWithStrings is the model entity for the MessageWithStrings schema.
type MessageWithStrings struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Strings holds the value of the "strings" field.
	Strings      []string `json:"strings,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MessageWithStrings) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case messagewithstrings.FieldStrings:
			values[i] = new([]byte)
		case messagewithstrings.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MessageWithStrings fields.
func (mws *MessageWithStrings) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case messagewithstrings.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mws.ID = int(value.Int64)
		case messagewithstrings.FieldStrings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field strings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &mws.Strings); err != nil {
					return fmt.Errorf("unmarshal field strings: %w", err)
				}
			}
		default:
			mws.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MessageWithStrings.
// This includes values selected through modifiers, order, etc.
func (mws *MessageWithStrings) Value(name string) (ent.Value, error) {
	return mws.selectValues.Get(name)
}

// Update returns a builder for updating this MessageWithStrings.
// Note that you need to call MessageWithStrings.Unwrap() before calling this method if this MessageWithStrings
// was returned from a transaction, and the transaction was committed or rolled back.
func (mws *MessageWithStrings) Update() *MessageWithStringsUpdateOne {
	return NewMessageWithStringsClient(mws.config).UpdateOne(mws)
}

// Unwrap unwraps the MessageWithStrings entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mws *MessageWithStrings) Unwrap() *MessageWithStrings {
	_tx, ok := mws.config.driver.(*txDriver)
	if !ok {
		panic("ent: MessageWithStrings is not a transactional entity")
	}
	mws.config.driver = _tx.drv
	return mws
}

// String implements the fmt.Stringer.
func (mws *MessageWithStrings) String() string {
	var builder strings.Builder
	builder.WriteString("MessageWithStrings(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mws.ID))
	builder.WriteString("strings=")
	builder.WriteString(fmt.Sprintf("%v", mws.Strings))
	builder.WriteByte(')')
	return builder.String()
}

// MessageWithStringsSlice is a parsable slice of MessageWithStrings.
type MessageWithStringsSlice []*MessageWithStrings
