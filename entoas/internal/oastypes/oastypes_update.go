// Code generated by ent, DO NOT EDIT.

package oastypes

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tilau2328/entcontrib/entoas/internal/oastypes/oastypes"
	"github.com/tilau2328/entcontrib/entoas/internal/oastypes/predicate"
	"github.com/tilau2328/entcontrib/entoas/internal/oastypes/schema"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OASTypesUpdate is the builder for updating OASTypes entities.
type OASTypesUpdate struct {
	config
	hooks    []Hook
	mutation *OASTypesMutation
}

// Where appends a list predicates to the OASTypesUpdate builder.
func (otu *OASTypesUpdate) Where(ps ...predicate.OASTypes) *OASTypesUpdate {
	otu.mutation.Where(ps...)
	return otu
}

// SetInt sets the "int" field.
func (otu *OASTypesUpdate) SetInt(i int) *OASTypesUpdate {
	otu.mutation.ResetInt()
	otu.mutation.SetInt(i)
	return otu
}

// AddInt adds i to the "int" field.
func (otu *OASTypesUpdate) AddInt(i int) *OASTypesUpdate {
	otu.mutation.AddInt(i)
	return otu
}

// SetInt8 sets the "int8" field.
func (otu *OASTypesUpdate) SetInt8(i int8) *OASTypesUpdate {
	otu.mutation.ResetInt8()
	otu.mutation.SetInt8(i)
	return otu
}

// AddInt8 adds i to the "int8" field.
func (otu *OASTypesUpdate) AddInt8(i int8) *OASTypesUpdate {
	otu.mutation.AddInt8(i)
	return otu
}

// SetInt16 sets the "int16" field.
func (otu *OASTypesUpdate) SetInt16(i int16) *OASTypesUpdate {
	otu.mutation.ResetInt16()
	otu.mutation.SetInt16(i)
	return otu
}

// AddInt16 adds i to the "int16" field.
func (otu *OASTypesUpdate) AddInt16(i int16) *OASTypesUpdate {
	otu.mutation.AddInt16(i)
	return otu
}

// SetInt32 sets the "int32" field.
func (otu *OASTypesUpdate) SetInt32(i int32) *OASTypesUpdate {
	otu.mutation.ResetInt32()
	otu.mutation.SetInt32(i)
	return otu
}

// AddInt32 adds i to the "int32" field.
func (otu *OASTypesUpdate) AddInt32(i int32) *OASTypesUpdate {
	otu.mutation.AddInt32(i)
	return otu
}

// SetInt64 sets the "int64" field.
func (otu *OASTypesUpdate) SetInt64(i int64) *OASTypesUpdate {
	otu.mutation.ResetInt64()
	otu.mutation.SetInt64(i)
	return otu
}

// AddInt64 adds i to the "int64" field.
func (otu *OASTypesUpdate) AddInt64(i int64) *OASTypesUpdate {
	otu.mutation.AddInt64(i)
	return otu
}

// SetUint sets the "uint" field.
func (otu *OASTypesUpdate) SetUint(u uint) *OASTypesUpdate {
	otu.mutation.ResetUint()
	otu.mutation.SetUint(u)
	return otu
}

// AddUint adds u to the "uint" field.
func (otu *OASTypesUpdate) AddUint(u int) *OASTypesUpdate {
	otu.mutation.AddUint(u)
	return otu
}

// SetUint8 sets the "uint8" field.
func (otu *OASTypesUpdate) SetUint8(u uint8) *OASTypesUpdate {
	otu.mutation.ResetUint8()
	otu.mutation.SetUint8(u)
	return otu
}

// AddUint8 adds u to the "uint8" field.
func (otu *OASTypesUpdate) AddUint8(u int8) *OASTypesUpdate {
	otu.mutation.AddUint8(u)
	return otu
}

// SetUint16 sets the "uint16" field.
func (otu *OASTypesUpdate) SetUint16(u uint16) *OASTypesUpdate {
	otu.mutation.ResetUint16()
	otu.mutation.SetUint16(u)
	return otu
}

// AddUint16 adds u to the "uint16" field.
func (otu *OASTypesUpdate) AddUint16(u int16) *OASTypesUpdate {
	otu.mutation.AddUint16(u)
	return otu
}

// SetUint32 sets the "uint32" field.
func (otu *OASTypesUpdate) SetUint32(u uint32) *OASTypesUpdate {
	otu.mutation.ResetUint32()
	otu.mutation.SetUint32(u)
	return otu
}

// AddUint32 adds u to the "uint32" field.
func (otu *OASTypesUpdate) AddUint32(u int32) *OASTypesUpdate {
	otu.mutation.AddUint32(u)
	return otu
}

// SetUint64 sets the "uint64" field.
func (otu *OASTypesUpdate) SetUint64(u uint64) *OASTypesUpdate {
	otu.mutation.ResetUint64()
	otu.mutation.SetUint64(u)
	return otu
}

// AddUint64 adds u to the "uint64" field.
func (otu *OASTypesUpdate) AddUint64(u int64) *OASTypesUpdate {
	otu.mutation.AddUint64(u)
	return otu
}

// SetFloat32 sets the "float32" field.
func (otu *OASTypesUpdate) SetFloat32(f float32) *OASTypesUpdate {
	otu.mutation.ResetFloat32()
	otu.mutation.SetFloat32(f)
	return otu
}

// AddFloat32 adds f to the "float32" field.
func (otu *OASTypesUpdate) AddFloat32(f float32) *OASTypesUpdate {
	otu.mutation.AddFloat32(f)
	return otu
}

// SetFloat64 sets the "float64" field.
func (otu *OASTypesUpdate) SetFloat64(f float64) *OASTypesUpdate {
	otu.mutation.ResetFloat64()
	otu.mutation.SetFloat64(f)
	return otu
}

// AddFloat64 adds f to the "float64" field.
func (otu *OASTypesUpdate) AddFloat64(f float64) *OASTypesUpdate {
	otu.mutation.AddFloat64(f)
	return otu
}

// SetStringField sets the "string_field" field.
func (otu *OASTypesUpdate) SetStringField(s string) *OASTypesUpdate {
	otu.mutation.SetStringField(s)
	return otu
}

// SetBool sets the "bool" field.
func (otu *OASTypesUpdate) SetBool(b bool) *OASTypesUpdate {
	otu.mutation.SetBool(b)
	return otu
}

// SetUUID sets the "uuid" field.
func (otu *OASTypesUpdate) SetUUID(u uuid.UUID) *OASTypesUpdate {
	otu.mutation.SetUUID(u)
	return otu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (otu *OASTypesUpdate) SetNillableUUID(u *uuid.UUID) *OASTypesUpdate {
	if u != nil {
		otu.SetUUID(*u)
	}
	return otu
}

// SetTime sets the "time" field.
func (otu *OASTypesUpdate) SetTime(t time.Time) *OASTypesUpdate {
	otu.mutation.SetTime(t)
	return otu
}

// SetText sets the "text" field.
func (otu *OASTypesUpdate) SetText(s string) *OASTypesUpdate {
	otu.mutation.SetText(s)
	return otu
}

// SetState sets the "state" field.
func (otu *OASTypesUpdate) SetState(o oastypes.State) *OASTypesUpdate {
	otu.mutation.SetState(o)
	return otu
}

// SetStrings sets the "strings" field.
func (otu *OASTypesUpdate) SetStrings(s []string) *OASTypesUpdate {
	otu.mutation.SetStrings(s)
	return otu
}

// AppendStrings appends s to the "strings" field.
func (otu *OASTypesUpdate) AppendStrings(s []string) *OASTypesUpdate {
	otu.mutation.AppendStrings(s)
	return otu
}

// SetInts sets the "ints" field.
func (otu *OASTypesUpdate) SetInts(i []int) *OASTypesUpdate {
	otu.mutation.SetInts(i)
	return otu
}

// AppendInts appends i to the "ints" field.
func (otu *OASTypesUpdate) AppendInts(i []int) *OASTypesUpdate {
	otu.mutation.AppendInts(i)
	return otu
}

// SetFloats sets the "floats" field.
func (otu *OASTypesUpdate) SetFloats(f []float64) *OASTypesUpdate {
	otu.mutation.SetFloats(f)
	return otu
}

// AppendFloats appends f to the "floats" field.
func (otu *OASTypesUpdate) AppendFloats(f []float64) *OASTypesUpdate {
	otu.mutation.AppendFloats(f)
	return otu
}

// SetBytes sets the "bytes" field.
func (otu *OASTypesUpdate) SetBytes(b []byte) *OASTypesUpdate {
	otu.mutation.SetBytes(b)
	return otu
}

// SetNicknames sets the "nicknames" field.
func (otu *OASTypesUpdate) SetNicknames(s []string) *OASTypesUpdate {
	otu.mutation.SetNicknames(s)
	return otu
}

// AppendNicknames appends s to the "nicknames" field.
func (otu *OASTypesUpdate) AppendNicknames(s []string) *OASTypesUpdate {
	otu.mutation.AppendNicknames(s)
	return otu
}

// SetJSONSlice sets the "json_slice" field.
func (otu *OASTypesUpdate) SetJSONSlice(h []http.Dir) *OASTypesUpdate {
	otu.mutation.SetJSONSlice(h)
	return otu
}

// AppendJSONSlice appends h to the "json_slice" field.
func (otu *OASTypesUpdate) AppendJSONSlice(h []http.Dir) *OASTypesUpdate {
	otu.mutation.AppendJSONSlice(h)
	return otu
}

// SetJSONObj sets the "json_obj" field.
func (otu *OASTypesUpdate) SetJSONObj(u url.URL) *OASTypesUpdate {
	otu.mutation.SetJSONObj(u)
	return otu
}

// SetOther sets the "other" field.
func (otu *OASTypesUpdate) SetOther(s *schema.Link) *OASTypesUpdate {
	otu.mutation.SetOther(s)
	return otu
}

// SetOptional sets the "optional" field.
func (otu *OASTypesUpdate) SetOptional(i int) *OASTypesUpdate {
	otu.mutation.ResetOptional()
	otu.mutation.SetOptional(i)
	return otu
}

// SetNillableOptional sets the "optional" field if the given value is not nil.
func (otu *OASTypesUpdate) SetNillableOptional(i *int) *OASTypesUpdate {
	if i != nil {
		otu.SetOptional(*i)
	}
	return otu
}

// AddOptional adds i to the "optional" field.
func (otu *OASTypesUpdate) AddOptional(i int) *OASTypesUpdate {
	otu.mutation.AddOptional(i)
	return otu
}

// ClearOptional clears the value of the "optional" field.
func (otu *OASTypesUpdate) ClearOptional() *OASTypesUpdate {
	otu.mutation.ClearOptional()
	return otu
}

// SetNillable sets the "nillable" field.
func (otu *OASTypesUpdate) SetNillable(i int) *OASTypesUpdate {
	otu.mutation.ResetNillable()
	otu.mutation.SetNillable(i)
	return otu
}

// AddNillable adds i to the "nillable" field.
func (otu *OASTypesUpdate) AddNillable(i int) *OASTypesUpdate {
	otu.mutation.AddNillable(i)
	return otu
}

// SetOptionalAndNillable sets the "optional_and_nillable" field.
func (otu *OASTypesUpdate) SetOptionalAndNillable(i int) *OASTypesUpdate {
	otu.mutation.ResetOptionalAndNillable()
	otu.mutation.SetOptionalAndNillable(i)
	return otu
}

// SetNillableOptionalAndNillable sets the "optional_and_nillable" field if the given value is not nil.
func (otu *OASTypesUpdate) SetNillableOptionalAndNillable(i *int) *OASTypesUpdate {
	if i != nil {
		otu.SetOptionalAndNillable(*i)
	}
	return otu
}

// AddOptionalAndNillable adds i to the "optional_and_nillable" field.
func (otu *OASTypesUpdate) AddOptionalAndNillable(i int) *OASTypesUpdate {
	otu.mutation.AddOptionalAndNillable(i)
	return otu
}

// ClearOptionalAndNillable clears the value of the "optional_and_nillable" field.
func (otu *OASTypesUpdate) ClearOptionalAndNillable() *OASTypesUpdate {
	otu.mutation.ClearOptionalAndNillable()
	return otu
}

// Mutation returns the OASTypesMutation object of the builder.
func (otu *OASTypesUpdate) Mutation() *OASTypesMutation {
	return otu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (otu *OASTypesUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OASTypesMutation](ctx, otu.sqlSave, otu.mutation, otu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (otu *OASTypesUpdate) SaveX(ctx context.Context) int {
	affected, err := otu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (otu *OASTypesUpdate) Exec(ctx context.Context) error {
	_, err := otu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otu *OASTypesUpdate) ExecX(ctx context.Context) {
	if err := otu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (otu *OASTypesUpdate) check() error {
	if v, ok := otu.mutation.State(); ok {
		if err := oastypes.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`oastypes: validator failed for field "OASTypes.state": %w`, err)}
		}
	}
	return nil
}

func (otu *OASTypesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := otu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(oastypes.Table, oastypes.Columns, sqlgraph.NewFieldSpec(oastypes.FieldID, field.TypeInt))
	if ps := otu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := otu.mutation.Int(); ok {
		_spec.SetField(oastypes.FieldInt, field.TypeInt, value)
	}
	if value, ok := otu.mutation.AddedInt(); ok {
		_spec.AddField(oastypes.FieldInt, field.TypeInt, value)
	}
	if value, ok := otu.mutation.Int8(); ok {
		_spec.SetField(oastypes.FieldInt8, field.TypeInt8, value)
	}
	if value, ok := otu.mutation.AddedInt8(); ok {
		_spec.AddField(oastypes.FieldInt8, field.TypeInt8, value)
	}
	if value, ok := otu.mutation.Int16(); ok {
		_spec.SetField(oastypes.FieldInt16, field.TypeInt16, value)
	}
	if value, ok := otu.mutation.AddedInt16(); ok {
		_spec.AddField(oastypes.FieldInt16, field.TypeInt16, value)
	}
	if value, ok := otu.mutation.Int32(); ok {
		_spec.SetField(oastypes.FieldInt32, field.TypeInt32, value)
	}
	if value, ok := otu.mutation.AddedInt32(); ok {
		_spec.AddField(oastypes.FieldInt32, field.TypeInt32, value)
	}
	if value, ok := otu.mutation.Int64(); ok {
		_spec.SetField(oastypes.FieldInt64, field.TypeInt64, value)
	}
	if value, ok := otu.mutation.AddedInt64(); ok {
		_spec.AddField(oastypes.FieldInt64, field.TypeInt64, value)
	}
	if value, ok := otu.mutation.Uint(); ok {
		_spec.SetField(oastypes.FieldUint, field.TypeUint, value)
	}
	if value, ok := otu.mutation.AddedUint(); ok {
		_spec.AddField(oastypes.FieldUint, field.TypeUint, value)
	}
	if value, ok := otu.mutation.Uint8(); ok {
		_spec.SetField(oastypes.FieldUint8, field.TypeUint8, value)
	}
	if value, ok := otu.mutation.AddedUint8(); ok {
		_spec.AddField(oastypes.FieldUint8, field.TypeUint8, value)
	}
	if value, ok := otu.mutation.Uint16(); ok {
		_spec.SetField(oastypes.FieldUint16, field.TypeUint16, value)
	}
	if value, ok := otu.mutation.AddedUint16(); ok {
		_spec.AddField(oastypes.FieldUint16, field.TypeUint16, value)
	}
	if value, ok := otu.mutation.Uint32(); ok {
		_spec.SetField(oastypes.FieldUint32, field.TypeUint32, value)
	}
	if value, ok := otu.mutation.AddedUint32(); ok {
		_spec.AddField(oastypes.FieldUint32, field.TypeUint32, value)
	}
	if value, ok := otu.mutation.Uint64(); ok {
		_spec.SetField(oastypes.FieldUint64, field.TypeUint64, value)
	}
	if value, ok := otu.mutation.AddedUint64(); ok {
		_spec.AddField(oastypes.FieldUint64, field.TypeUint64, value)
	}
	if value, ok := otu.mutation.Float32(); ok {
		_spec.SetField(oastypes.FieldFloat32, field.TypeFloat32, value)
	}
	if value, ok := otu.mutation.AddedFloat32(); ok {
		_spec.AddField(oastypes.FieldFloat32, field.TypeFloat32, value)
	}
	if value, ok := otu.mutation.Float64(); ok {
		_spec.SetField(oastypes.FieldFloat64, field.TypeFloat64, value)
	}
	if value, ok := otu.mutation.AddedFloat64(); ok {
		_spec.AddField(oastypes.FieldFloat64, field.TypeFloat64, value)
	}
	if value, ok := otu.mutation.StringField(); ok {
		_spec.SetField(oastypes.FieldStringField, field.TypeString, value)
	}
	if value, ok := otu.mutation.Bool(); ok {
		_spec.SetField(oastypes.FieldBool, field.TypeBool, value)
	}
	if value, ok := otu.mutation.UUID(); ok {
		_spec.SetField(oastypes.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := otu.mutation.Time(); ok {
		_spec.SetField(oastypes.FieldTime, field.TypeTime, value)
	}
	if value, ok := otu.mutation.Text(); ok {
		_spec.SetField(oastypes.FieldText, field.TypeString, value)
	}
	if value, ok := otu.mutation.State(); ok {
		_spec.SetField(oastypes.FieldState, field.TypeEnum, value)
	}
	if value, ok := otu.mutation.Strings(); ok {
		_spec.SetField(oastypes.FieldStrings, field.TypeJSON, value)
	}
	if value, ok := otu.mutation.AppendedStrings(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oastypes.FieldStrings, value)
		})
	}
	if value, ok := otu.mutation.Ints(); ok {
		_spec.SetField(oastypes.FieldInts, field.TypeJSON, value)
	}
	if value, ok := otu.mutation.AppendedInts(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oastypes.FieldInts, value)
		})
	}
	if value, ok := otu.mutation.Floats(); ok {
		_spec.SetField(oastypes.FieldFloats, field.TypeJSON, value)
	}
	if value, ok := otu.mutation.AppendedFloats(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oastypes.FieldFloats, value)
		})
	}
	if value, ok := otu.mutation.Bytes(); ok {
		_spec.SetField(oastypes.FieldBytes, field.TypeBytes, value)
	}
	if value, ok := otu.mutation.Nicknames(); ok {
		_spec.SetField(oastypes.FieldNicknames, field.TypeJSON, value)
	}
	if value, ok := otu.mutation.AppendedNicknames(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oastypes.FieldNicknames, value)
		})
	}
	if value, ok := otu.mutation.JSONSlice(); ok {
		_spec.SetField(oastypes.FieldJSONSlice, field.TypeJSON, value)
	}
	if value, ok := otu.mutation.AppendedJSONSlice(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oastypes.FieldJSONSlice, value)
		})
	}
	if value, ok := otu.mutation.JSONObj(); ok {
		_spec.SetField(oastypes.FieldJSONObj, field.TypeJSON, value)
	}
	if value, ok := otu.mutation.Other(); ok {
		_spec.SetField(oastypes.FieldOther, field.TypeOther, value)
	}
	if value, ok := otu.mutation.Optional(); ok {
		_spec.SetField(oastypes.FieldOptional, field.TypeInt, value)
	}
	if value, ok := otu.mutation.AddedOptional(); ok {
		_spec.AddField(oastypes.FieldOptional, field.TypeInt, value)
	}
	if otu.mutation.OptionalCleared() {
		_spec.ClearField(oastypes.FieldOptional, field.TypeInt)
	}
	if value, ok := otu.mutation.Nillable(); ok {
		_spec.SetField(oastypes.FieldNillable, field.TypeInt, value)
	}
	if value, ok := otu.mutation.AddedNillable(); ok {
		_spec.AddField(oastypes.FieldNillable, field.TypeInt, value)
	}
	if value, ok := otu.mutation.OptionalAndNillable(); ok {
		_spec.SetField(oastypes.FieldOptionalAndNillable, field.TypeInt, value)
	}
	if value, ok := otu.mutation.AddedOptionalAndNillable(); ok {
		_spec.AddField(oastypes.FieldOptionalAndNillable, field.TypeInt, value)
	}
	if otu.mutation.OptionalAndNillableCleared() {
		_spec.ClearField(oastypes.FieldOptionalAndNillable, field.TypeInt)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, otu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oastypes.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	otu.mutation.done = true
	return n, nil
}

// OASTypesUpdateOne is the builder for updating a single OASTypes entity.
type OASTypesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OASTypesMutation
}

// SetInt sets the "int" field.
func (otuo *OASTypesUpdateOne) SetInt(i int) *OASTypesUpdateOne {
	otuo.mutation.ResetInt()
	otuo.mutation.SetInt(i)
	return otuo
}

// AddInt adds i to the "int" field.
func (otuo *OASTypesUpdateOne) AddInt(i int) *OASTypesUpdateOne {
	otuo.mutation.AddInt(i)
	return otuo
}

// SetInt8 sets the "int8" field.
func (otuo *OASTypesUpdateOne) SetInt8(i int8) *OASTypesUpdateOne {
	otuo.mutation.ResetInt8()
	otuo.mutation.SetInt8(i)
	return otuo
}

// AddInt8 adds i to the "int8" field.
func (otuo *OASTypesUpdateOne) AddInt8(i int8) *OASTypesUpdateOne {
	otuo.mutation.AddInt8(i)
	return otuo
}

// SetInt16 sets the "int16" field.
func (otuo *OASTypesUpdateOne) SetInt16(i int16) *OASTypesUpdateOne {
	otuo.mutation.ResetInt16()
	otuo.mutation.SetInt16(i)
	return otuo
}

// AddInt16 adds i to the "int16" field.
func (otuo *OASTypesUpdateOne) AddInt16(i int16) *OASTypesUpdateOne {
	otuo.mutation.AddInt16(i)
	return otuo
}

// SetInt32 sets the "int32" field.
func (otuo *OASTypesUpdateOne) SetInt32(i int32) *OASTypesUpdateOne {
	otuo.mutation.ResetInt32()
	otuo.mutation.SetInt32(i)
	return otuo
}

// AddInt32 adds i to the "int32" field.
func (otuo *OASTypesUpdateOne) AddInt32(i int32) *OASTypesUpdateOne {
	otuo.mutation.AddInt32(i)
	return otuo
}

// SetInt64 sets the "int64" field.
func (otuo *OASTypesUpdateOne) SetInt64(i int64) *OASTypesUpdateOne {
	otuo.mutation.ResetInt64()
	otuo.mutation.SetInt64(i)
	return otuo
}

// AddInt64 adds i to the "int64" field.
func (otuo *OASTypesUpdateOne) AddInt64(i int64) *OASTypesUpdateOne {
	otuo.mutation.AddInt64(i)
	return otuo
}

// SetUint sets the "uint" field.
func (otuo *OASTypesUpdateOne) SetUint(u uint) *OASTypesUpdateOne {
	otuo.mutation.ResetUint()
	otuo.mutation.SetUint(u)
	return otuo
}

// AddUint adds u to the "uint" field.
func (otuo *OASTypesUpdateOne) AddUint(u int) *OASTypesUpdateOne {
	otuo.mutation.AddUint(u)
	return otuo
}

// SetUint8 sets the "uint8" field.
func (otuo *OASTypesUpdateOne) SetUint8(u uint8) *OASTypesUpdateOne {
	otuo.mutation.ResetUint8()
	otuo.mutation.SetUint8(u)
	return otuo
}

// AddUint8 adds u to the "uint8" field.
func (otuo *OASTypesUpdateOne) AddUint8(u int8) *OASTypesUpdateOne {
	otuo.mutation.AddUint8(u)
	return otuo
}

// SetUint16 sets the "uint16" field.
func (otuo *OASTypesUpdateOne) SetUint16(u uint16) *OASTypesUpdateOne {
	otuo.mutation.ResetUint16()
	otuo.mutation.SetUint16(u)
	return otuo
}

// AddUint16 adds u to the "uint16" field.
func (otuo *OASTypesUpdateOne) AddUint16(u int16) *OASTypesUpdateOne {
	otuo.mutation.AddUint16(u)
	return otuo
}

// SetUint32 sets the "uint32" field.
func (otuo *OASTypesUpdateOne) SetUint32(u uint32) *OASTypesUpdateOne {
	otuo.mutation.ResetUint32()
	otuo.mutation.SetUint32(u)
	return otuo
}

// AddUint32 adds u to the "uint32" field.
func (otuo *OASTypesUpdateOne) AddUint32(u int32) *OASTypesUpdateOne {
	otuo.mutation.AddUint32(u)
	return otuo
}

// SetUint64 sets the "uint64" field.
func (otuo *OASTypesUpdateOne) SetUint64(u uint64) *OASTypesUpdateOne {
	otuo.mutation.ResetUint64()
	otuo.mutation.SetUint64(u)
	return otuo
}

// AddUint64 adds u to the "uint64" field.
func (otuo *OASTypesUpdateOne) AddUint64(u int64) *OASTypesUpdateOne {
	otuo.mutation.AddUint64(u)
	return otuo
}

// SetFloat32 sets the "float32" field.
func (otuo *OASTypesUpdateOne) SetFloat32(f float32) *OASTypesUpdateOne {
	otuo.mutation.ResetFloat32()
	otuo.mutation.SetFloat32(f)
	return otuo
}

// AddFloat32 adds f to the "float32" field.
func (otuo *OASTypesUpdateOne) AddFloat32(f float32) *OASTypesUpdateOne {
	otuo.mutation.AddFloat32(f)
	return otuo
}

// SetFloat64 sets the "float64" field.
func (otuo *OASTypesUpdateOne) SetFloat64(f float64) *OASTypesUpdateOne {
	otuo.mutation.ResetFloat64()
	otuo.mutation.SetFloat64(f)
	return otuo
}

// AddFloat64 adds f to the "float64" field.
func (otuo *OASTypesUpdateOne) AddFloat64(f float64) *OASTypesUpdateOne {
	otuo.mutation.AddFloat64(f)
	return otuo
}

// SetStringField sets the "string_field" field.
func (otuo *OASTypesUpdateOne) SetStringField(s string) *OASTypesUpdateOne {
	otuo.mutation.SetStringField(s)
	return otuo
}

// SetBool sets the "bool" field.
func (otuo *OASTypesUpdateOne) SetBool(b bool) *OASTypesUpdateOne {
	otuo.mutation.SetBool(b)
	return otuo
}

// SetUUID sets the "uuid" field.
func (otuo *OASTypesUpdateOne) SetUUID(u uuid.UUID) *OASTypesUpdateOne {
	otuo.mutation.SetUUID(u)
	return otuo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (otuo *OASTypesUpdateOne) SetNillableUUID(u *uuid.UUID) *OASTypesUpdateOne {
	if u != nil {
		otuo.SetUUID(*u)
	}
	return otuo
}

// SetTime sets the "time" field.
func (otuo *OASTypesUpdateOne) SetTime(t time.Time) *OASTypesUpdateOne {
	otuo.mutation.SetTime(t)
	return otuo
}

// SetText sets the "text" field.
func (otuo *OASTypesUpdateOne) SetText(s string) *OASTypesUpdateOne {
	otuo.mutation.SetText(s)
	return otuo
}

// SetState sets the "state" field.
func (otuo *OASTypesUpdateOne) SetState(o oastypes.State) *OASTypesUpdateOne {
	otuo.mutation.SetState(o)
	return otuo
}

// SetStrings sets the "strings" field.
func (otuo *OASTypesUpdateOne) SetStrings(s []string) *OASTypesUpdateOne {
	otuo.mutation.SetStrings(s)
	return otuo
}

// AppendStrings appends s to the "strings" field.
func (otuo *OASTypesUpdateOne) AppendStrings(s []string) *OASTypesUpdateOne {
	otuo.mutation.AppendStrings(s)
	return otuo
}

// SetInts sets the "ints" field.
func (otuo *OASTypesUpdateOne) SetInts(i []int) *OASTypesUpdateOne {
	otuo.mutation.SetInts(i)
	return otuo
}

// AppendInts appends i to the "ints" field.
func (otuo *OASTypesUpdateOne) AppendInts(i []int) *OASTypesUpdateOne {
	otuo.mutation.AppendInts(i)
	return otuo
}

// SetFloats sets the "floats" field.
func (otuo *OASTypesUpdateOne) SetFloats(f []float64) *OASTypesUpdateOne {
	otuo.mutation.SetFloats(f)
	return otuo
}

// AppendFloats appends f to the "floats" field.
func (otuo *OASTypesUpdateOne) AppendFloats(f []float64) *OASTypesUpdateOne {
	otuo.mutation.AppendFloats(f)
	return otuo
}

// SetBytes sets the "bytes" field.
func (otuo *OASTypesUpdateOne) SetBytes(b []byte) *OASTypesUpdateOne {
	otuo.mutation.SetBytes(b)
	return otuo
}

// SetNicknames sets the "nicknames" field.
func (otuo *OASTypesUpdateOne) SetNicknames(s []string) *OASTypesUpdateOne {
	otuo.mutation.SetNicknames(s)
	return otuo
}

// AppendNicknames appends s to the "nicknames" field.
func (otuo *OASTypesUpdateOne) AppendNicknames(s []string) *OASTypesUpdateOne {
	otuo.mutation.AppendNicknames(s)
	return otuo
}

// SetJSONSlice sets the "json_slice" field.
func (otuo *OASTypesUpdateOne) SetJSONSlice(h []http.Dir) *OASTypesUpdateOne {
	otuo.mutation.SetJSONSlice(h)
	return otuo
}

// AppendJSONSlice appends h to the "json_slice" field.
func (otuo *OASTypesUpdateOne) AppendJSONSlice(h []http.Dir) *OASTypesUpdateOne {
	otuo.mutation.AppendJSONSlice(h)
	return otuo
}

// SetJSONObj sets the "json_obj" field.
func (otuo *OASTypesUpdateOne) SetJSONObj(u url.URL) *OASTypesUpdateOne {
	otuo.mutation.SetJSONObj(u)
	return otuo
}

// SetOther sets the "other" field.
func (otuo *OASTypesUpdateOne) SetOther(s *schema.Link) *OASTypesUpdateOne {
	otuo.mutation.SetOther(s)
	return otuo
}

// SetOptional sets the "optional" field.
func (otuo *OASTypesUpdateOne) SetOptional(i int) *OASTypesUpdateOne {
	otuo.mutation.ResetOptional()
	otuo.mutation.SetOptional(i)
	return otuo
}

// SetNillableOptional sets the "optional" field if the given value is not nil.
func (otuo *OASTypesUpdateOne) SetNillableOptional(i *int) *OASTypesUpdateOne {
	if i != nil {
		otuo.SetOptional(*i)
	}
	return otuo
}

// AddOptional adds i to the "optional" field.
func (otuo *OASTypesUpdateOne) AddOptional(i int) *OASTypesUpdateOne {
	otuo.mutation.AddOptional(i)
	return otuo
}

// ClearOptional clears the value of the "optional" field.
func (otuo *OASTypesUpdateOne) ClearOptional() *OASTypesUpdateOne {
	otuo.mutation.ClearOptional()
	return otuo
}

// SetNillable sets the "nillable" field.
func (otuo *OASTypesUpdateOne) SetNillable(i int) *OASTypesUpdateOne {
	otuo.mutation.ResetNillable()
	otuo.mutation.SetNillable(i)
	return otuo
}

// AddNillable adds i to the "nillable" field.
func (otuo *OASTypesUpdateOne) AddNillable(i int) *OASTypesUpdateOne {
	otuo.mutation.AddNillable(i)
	return otuo
}

// SetOptionalAndNillable sets the "optional_and_nillable" field.
func (otuo *OASTypesUpdateOne) SetOptionalAndNillable(i int) *OASTypesUpdateOne {
	otuo.mutation.ResetOptionalAndNillable()
	otuo.mutation.SetOptionalAndNillable(i)
	return otuo
}

// SetNillableOptionalAndNillable sets the "optional_and_nillable" field if the given value is not nil.
func (otuo *OASTypesUpdateOne) SetNillableOptionalAndNillable(i *int) *OASTypesUpdateOne {
	if i != nil {
		otuo.SetOptionalAndNillable(*i)
	}
	return otuo
}

// AddOptionalAndNillable adds i to the "optional_and_nillable" field.
func (otuo *OASTypesUpdateOne) AddOptionalAndNillable(i int) *OASTypesUpdateOne {
	otuo.mutation.AddOptionalAndNillable(i)
	return otuo
}

// ClearOptionalAndNillable clears the value of the "optional_and_nillable" field.
func (otuo *OASTypesUpdateOne) ClearOptionalAndNillable() *OASTypesUpdateOne {
	otuo.mutation.ClearOptionalAndNillable()
	return otuo
}

// Mutation returns the OASTypesMutation object of the builder.
func (otuo *OASTypesUpdateOne) Mutation() *OASTypesMutation {
	return otuo.mutation
}

// Where appends a list predicates to the OASTypesUpdate builder.
func (otuo *OASTypesUpdateOne) Where(ps ...predicate.OASTypes) *OASTypesUpdateOne {
	otuo.mutation.Where(ps...)
	return otuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (otuo *OASTypesUpdateOne) Select(field string, fields ...string) *OASTypesUpdateOne {
	otuo.fields = append([]string{field}, fields...)
	return otuo
}

// Save executes the query and returns the updated OASTypes entity.
func (otuo *OASTypesUpdateOne) Save(ctx context.Context) (*OASTypes, error) {
	return withHooks[*OASTypes, OASTypesMutation](ctx, otuo.sqlSave, otuo.mutation, otuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (otuo *OASTypesUpdateOne) SaveX(ctx context.Context) *OASTypes {
	node, err := otuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (otuo *OASTypesUpdateOne) Exec(ctx context.Context) error {
	_, err := otuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (otuo *OASTypesUpdateOne) ExecX(ctx context.Context) {
	if err := otuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (otuo *OASTypesUpdateOne) check() error {
	if v, ok := otuo.mutation.State(); ok {
		if err := oastypes.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`oastypes: validator failed for field "OASTypes.state": %w`, err)}
		}
	}
	return nil
}

func (otuo *OASTypesUpdateOne) sqlSave(ctx context.Context) (_node *OASTypes, err error) {
	if err := otuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(oastypes.Table, oastypes.Columns, sqlgraph.NewFieldSpec(oastypes.FieldID, field.TypeInt))
	id, ok := otuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`oastypes: missing "OASTypes.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := otuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oastypes.FieldID)
		for _, f := range fields {
			if !oastypes.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("oastypes: invalid field %q for query", f)}
			}
			if f != oastypes.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := otuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := otuo.mutation.Int(); ok {
		_spec.SetField(oastypes.FieldInt, field.TypeInt, value)
	}
	if value, ok := otuo.mutation.AddedInt(); ok {
		_spec.AddField(oastypes.FieldInt, field.TypeInt, value)
	}
	if value, ok := otuo.mutation.Int8(); ok {
		_spec.SetField(oastypes.FieldInt8, field.TypeInt8, value)
	}
	if value, ok := otuo.mutation.AddedInt8(); ok {
		_spec.AddField(oastypes.FieldInt8, field.TypeInt8, value)
	}
	if value, ok := otuo.mutation.Int16(); ok {
		_spec.SetField(oastypes.FieldInt16, field.TypeInt16, value)
	}
	if value, ok := otuo.mutation.AddedInt16(); ok {
		_spec.AddField(oastypes.FieldInt16, field.TypeInt16, value)
	}
	if value, ok := otuo.mutation.Int32(); ok {
		_spec.SetField(oastypes.FieldInt32, field.TypeInt32, value)
	}
	if value, ok := otuo.mutation.AddedInt32(); ok {
		_spec.AddField(oastypes.FieldInt32, field.TypeInt32, value)
	}
	if value, ok := otuo.mutation.Int64(); ok {
		_spec.SetField(oastypes.FieldInt64, field.TypeInt64, value)
	}
	if value, ok := otuo.mutation.AddedInt64(); ok {
		_spec.AddField(oastypes.FieldInt64, field.TypeInt64, value)
	}
	if value, ok := otuo.mutation.Uint(); ok {
		_spec.SetField(oastypes.FieldUint, field.TypeUint, value)
	}
	if value, ok := otuo.mutation.AddedUint(); ok {
		_spec.AddField(oastypes.FieldUint, field.TypeUint, value)
	}
	if value, ok := otuo.mutation.Uint8(); ok {
		_spec.SetField(oastypes.FieldUint8, field.TypeUint8, value)
	}
	if value, ok := otuo.mutation.AddedUint8(); ok {
		_spec.AddField(oastypes.FieldUint8, field.TypeUint8, value)
	}
	if value, ok := otuo.mutation.Uint16(); ok {
		_spec.SetField(oastypes.FieldUint16, field.TypeUint16, value)
	}
	if value, ok := otuo.mutation.AddedUint16(); ok {
		_spec.AddField(oastypes.FieldUint16, field.TypeUint16, value)
	}
	if value, ok := otuo.mutation.Uint32(); ok {
		_spec.SetField(oastypes.FieldUint32, field.TypeUint32, value)
	}
	if value, ok := otuo.mutation.AddedUint32(); ok {
		_spec.AddField(oastypes.FieldUint32, field.TypeUint32, value)
	}
	if value, ok := otuo.mutation.Uint64(); ok {
		_spec.SetField(oastypes.FieldUint64, field.TypeUint64, value)
	}
	if value, ok := otuo.mutation.AddedUint64(); ok {
		_spec.AddField(oastypes.FieldUint64, field.TypeUint64, value)
	}
	if value, ok := otuo.mutation.Float32(); ok {
		_spec.SetField(oastypes.FieldFloat32, field.TypeFloat32, value)
	}
	if value, ok := otuo.mutation.AddedFloat32(); ok {
		_spec.AddField(oastypes.FieldFloat32, field.TypeFloat32, value)
	}
	if value, ok := otuo.mutation.Float64(); ok {
		_spec.SetField(oastypes.FieldFloat64, field.TypeFloat64, value)
	}
	if value, ok := otuo.mutation.AddedFloat64(); ok {
		_spec.AddField(oastypes.FieldFloat64, field.TypeFloat64, value)
	}
	if value, ok := otuo.mutation.StringField(); ok {
		_spec.SetField(oastypes.FieldStringField, field.TypeString, value)
	}
	if value, ok := otuo.mutation.Bool(); ok {
		_spec.SetField(oastypes.FieldBool, field.TypeBool, value)
	}
	if value, ok := otuo.mutation.UUID(); ok {
		_spec.SetField(oastypes.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := otuo.mutation.Time(); ok {
		_spec.SetField(oastypes.FieldTime, field.TypeTime, value)
	}
	if value, ok := otuo.mutation.Text(); ok {
		_spec.SetField(oastypes.FieldText, field.TypeString, value)
	}
	if value, ok := otuo.mutation.State(); ok {
		_spec.SetField(oastypes.FieldState, field.TypeEnum, value)
	}
	if value, ok := otuo.mutation.Strings(); ok {
		_spec.SetField(oastypes.FieldStrings, field.TypeJSON, value)
	}
	if value, ok := otuo.mutation.AppendedStrings(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oastypes.FieldStrings, value)
		})
	}
	if value, ok := otuo.mutation.Ints(); ok {
		_spec.SetField(oastypes.FieldInts, field.TypeJSON, value)
	}
	if value, ok := otuo.mutation.AppendedInts(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oastypes.FieldInts, value)
		})
	}
	if value, ok := otuo.mutation.Floats(); ok {
		_spec.SetField(oastypes.FieldFloats, field.TypeJSON, value)
	}
	if value, ok := otuo.mutation.AppendedFloats(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oastypes.FieldFloats, value)
		})
	}
	if value, ok := otuo.mutation.Bytes(); ok {
		_spec.SetField(oastypes.FieldBytes, field.TypeBytes, value)
	}
	if value, ok := otuo.mutation.Nicknames(); ok {
		_spec.SetField(oastypes.FieldNicknames, field.TypeJSON, value)
	}
	if value, ok := otuo.mutation.AppendedNicknames(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oastypes.FieldNicknames, value)
		})
	}
	if value, ok := otuo.mutation.JSONSlice(); ok {
		_spec.SetField(oastypes.FieldJSONSlice, field.TypeJSON, value)
	}
	if value, ok := otuo.mutation.AppendedJSONSlice(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, oastypes.FieldJSONSlice, value)
		})
	}
	if value, ok := otuo.mutation.JSONObj(); ok {
		_spec.SetField(oastypes.FieldJSONObj, field.TypeJSON, value)
	}
	if value, ok := otuo.mutation.Other(); ok {
		_spec.SetField(oastypes.FieldOther, field.TypeOther, value)
	}
	if value, ok := otuo.mutation.Optional(); ok {
		_spec.SetField(oastypes.FieldOptional, field.TypeInt, value)
	}
	if value, ok := otuo.mutation.AddedOptional(); ok {
		_spec.AddField(oastypes.FieldOptional, field.TypeInt, value)
	}
	if otuo.mutation.OptionalCleared() {
		_spec.ClearField(oastypes.FieldOptional, field.TypeInt)
	}
	if value, ok := otuo.mutation.Nillable(); ok {
		_spec.SetField(oastypes.FieldNillable, field.TypeInt, value)
	}
	if value, ok := otuo.mutation.AddedNillable(); ok {
		_spec.AddField(oastypes.FieldNillable, field.TypeInt, value)
	}
	if value, ok := otuo.mutation.OptionalAndNillable(); ok {
		_spec.SetField(oastypes.FieldOptionalAndNillable, field.TypeInt, value)
	}
	if value, ok := otuo.mutation.AddedOptionalAndNillable(); ok {
		_spec.AddField(oastypes.FieldOptionalAndNillable, field.TypeInt, value)
	}
	if otuo.mutation.OptionalAndNillableCleared() {
		_spec.ClearField(oastypes.FieldOptionalAndNillable, field.TypeInt)
	}
	_node = &OASTypes{config: otuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, otuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oastypes.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	otuo.mutation.done = true
	return _node, nil
}
