// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/messagewithpackagename"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/predicate"
)

// MessageWithPackageNameUpdate is the builder for updating MessageWithPackageName entities.
type MessageWithPackageNameUpdate struct {
	config
	hooks    []Hook
	mutation *MessageWithPackageNameMutation
}

// Where appends a list predicates to the MessageWithPackageNameUpdate builder.
func (mwpnu *MessageWithPackageNameUpdate) Where(ps ...predicate.MessageWithPackageName) *MessageWithPackageNameUpdate {
	mwpnu.mutation.Where(ps...)
	return mwpnu
}

// SetName sets the "name" field.
func (mwpnu *MessageWithPackageNameUpdate) SetName(s string) *MessageWithPackageNameUpdate {
	mwpnu.mutation.SetName(s)
	return mwpnu
}

// Mutation returns the MessageWithPackageNameMutation object of the builder.
func (mwpnu *MessageWithPackageNameUpdate) Mutation() *MessageWithPackageNameMutation {
	return mwpnu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mwpnu *MessageWithPackageNameUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, MessageWithPackageNameMutation](ctx, mwpnu.sqlSave, mwpnu.mutation, mwpnu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mwpnu *MessageWithPackageNameUpdate) SaveX(ctx context.Context) int {
	affected, err := mwpnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mwpnu *MessageWithPackageNameUpdate) Exec(ctx context.Context) error {
	_, err := mwpnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwpnu *MessageWithPackageNameUpdate) ExecX(ctx context.Context) {
	if err := mwpnu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mwpnu *MessageWithPackageNameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(messagewithpackagename.Table, messagewithpackagename.Columns, sqlgraph.NewFieldSpec(messagewithpackagename.FieldID, field.TypeInt))
	if ps := mwpnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mwpnu.mutation.Name(); ok {
		_spec.SetField(messagewithpackagename.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mwpnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithpackagename.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mwpnu.mutation.done = true
	return n, nil
}

// MessageWithPackageNameUpdateOne is the builder for updating a single MessageWithPackageName entity.
type MessageWithPackageNameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageWithPackageNameMutation
}

// SetName sets the "name" field.
func (mwpnuo *MessageWithPackageNameUpdateOne) SetName(s string) *MessageWithPackageNameUpdateOne {
	mwpnuo.mutation.SetName(s)
	return mwpnuo
}

// Mutation returns the MessageWithPackageNameMutation object of the builder.
func (mwpnuo *MessageWithPackageNameUpdateOne) Mutation() *MessageWithPackageNameMutation {
	return mwpnuo.mutation
}

// Where appends a list predicates to the MessageWithPackageNameUpdate builder.
func (mwpnuo *MessageWithPackageNameUpdateOne) Where(ps ...predicate.MessageWithPackageName) *MessageWithPackageNameUpdateOne {
	mwpnuo.mutation.Where(ps...)
	return mwpnuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mwpnuo *MessageWithPackageNameUpdateOne) Select(field string, fields ...string) *MessageWithPackageNameUpdateOne {
	mwpnuo.fields = append([]string{field}, fields...)
	return mwpnuo
}

// Save executes the query and returns the updated MessageWithPackageName entity.
func (mwpnuo *MessageWithPackageNameUpdateOne) Save(ctx context.Context) (*MessageWithPackageName, error) {
	return withHooks[*MessageWithPackageName, MessageWithPackageNameMutation](ctx, mwpnuo.sqlSave, mwpnuo.mutation, mwpnuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mwpnuo *MessageWithPackageNameUpdateOne) SaveX(ctx context.Context) *MessageWithPackageName {
	node, err := mwpnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mwpnuo *MessageWithPackageNameUpdateOne) Exec(ctx context.Context) error {
	_, err := mwpnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwpnuo *MessageWithPackageNameUpdateOne) ExecX(ctx context.Context) {
	if err := mwpnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mwpnuo *MessageWithPackageNameUpdateOne) sqlSave(ctx context.Context) (_node *MessageWithPackageName, err error) {
	_spec := sqlgraph.NewUpdateSpec(messagewithpackagename.Table, messagewithpackagename.Columns, sqlgraph.NewFieldSpec(messagewithpackagename.FieldID, field.TypeInt))
	id, ok := mwpnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MessageWithPackageName.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mwpnuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagewithpackagename.FieldID)
		for _, f := range fields {
			if !messagewithpackagename.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != messagewithpackagename.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mwpnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mwpnuo.mutation.Name(); ok {
		_spec.SetField(messagewithpackagename.FieldName, field.TypeString, value)
	}
	_node = &MessageWithPackageName{config: mwpnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mwpnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithpackagename.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mwpnuo.mutation.done = true
	return _node, nil
}
