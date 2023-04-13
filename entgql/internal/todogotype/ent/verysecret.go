// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/tilau2328/entcontrib/entgql/internal/todogotype/ent/verysecret"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// VerySecret is the model entity for the VerySecret schema.
type VerySecret struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Password holds the value of the "password" field.
	Password     string `json:"password,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VerySecret) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case verysecret.FieldID, verysecret.FieldPassword:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VerySecret fields.
func (vs *VerySecret) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case verysecret.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				vs.ID = value.String
			}
		case verysecret.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				vs.Password = value.String
			}
		default:
			vs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VerySecret.
// This includes values selected through modifiers, order, etc.
func (vs *VerySecret) Value(name string) (ent.Value, error) {
	return vs.selectValues.Get(name)
}

// Update returns a builder for updating this VerySecret.
// Note that you need to call VerySecret.Unwrap() before calling this method if this VerySecret
// was returned from a transaction, and the transaction was committed or rolled back.
func (vs *VerySecret) Update() *VerySecretUpdateOne {
	return NewVerySecretClient(vs.config).UpdateOne(vs)
}

// Unwrap unwraps the VerySecret entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vs *VerySecret) Unwrap() *VerySecret {
	_tx, ok := vs.config.driver.(*txDriver)
	if !ok {
		panic("ent: VerySecret is not a transactional entity")
	}
	vs.config.driver = _tx.drv
	return vs
}

// String implements the fmt.Stringer.
func (vs *VerySecret) String() string {
	var builder strings.Builder
	builder.WriteString("VerySecret(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vs.ID))
	builder.WriteString("password=")
	builder.WriteString(vs.Password)
	builder.WriteByte(')')
	return builder.String()
}

// VerySecrets is a parsable slice of VerySecret.
type VerySecrets []*VerySecret
