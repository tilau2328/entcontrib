// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/invalidfieldmessage"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/schema"
)

// InvalidFieldMessage is the model entity for the InvalidFieldMessage schema.
type InvalidFieldMessage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// JSON holds the value of the "json" field.
	JSON         *schema.SomeJSON `json:"json,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InvalidFieldMessage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case invalidfieldmessage.FieldJSON:
			values[i] = new([]byte)
		case invalidfieldmessage.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InvalidFieldMessage fields.
func (ifm *InvalidFieldMessage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case invalidfieldmessage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ifm.ID = int(value.Int64)
		case invalidfieldmessage.FieldJSON:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field json", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ifm.JSON); err != nil {
					return fmt.Errorf("unmarshal field json: %w", err)
				}
			}
		default:
			ifm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the InvalidFieldMessage.
// This includes values selected through modifiers, order, etc.
func (ifm *InvalidFieldMessage) Value(name string) (ent.Value, error) {
	return ifm.selectValues.Get(name)
}

// Update returns a builder for updating this InvalidFieldMessage.
// Note that you need to call InvalidFieldMessage.Unwrap() before calling this method if this InvalidFieldMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (ifm *InvalidFieldMessage) Update() *InvalidFieldMessageUpdateOne {
	return NewInvalidFieldMessageClient(ifm.config).UpdateOne(ifm)
}

// Unwrap unwraps the InvalidFieldMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ifm *InvalidFieldMessage) Unwrap() *InvalidFieldMessage {
	_tx, ok := ifm.config.driver.(*txDriver)
	if !ok {
		panic("ent: InvalidFieldMessage is not a transactional entity")
	}
	ifm.config.driver = _tx.drv
	return ifm
}

// String implements the fmt.Stringer.
func (ifm *InvalidFieldMessage) String() string {
	var builder strings.Builder
	builder.WriteString("InvalidFieldMessage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ifm.ID))
	builder.WriteString("json=")
	builder.WriteString(fmt.Sprintf("%v", ifm.JSON))
	builder.WriteByte(')')
	return builder.String()
}

// InvalidFieldMessages is a parsable slice of InvalidFieldMessage.
type InvalidFieldMessages []*InvalidFieldMessage
