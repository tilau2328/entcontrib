// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/schemast/internal/mutatetest/ent/predicate"
	"github.com/tilau2328/entcontrib/schemast/internal/mutatetest/ent/withoutfields"
)

// WithoutFieldsUpdate is the builder for updating WithoutFields entities.
type WithoutFieldsUpdate struct {
	config
	hooks    []Hook
	mutation *WithoutFieldsMutation
}

// Where appends a list predicates to the WithoutFieldsUpdate builder.
func (wfu *WithoutFieldsUpdate) Where(ps ...predicate.WithoutFields) *WithoutFieldsUpdate {
	wfu.mutation.Where(ps...)
	return wfu
}

// Mutation returns the WithoutFieldsMutation object of the builder.
func (wfu *WithoutFieldsUpdate) Mutation() *WithoutFieldsMutation {
	return wfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wfu *WithoutFieldsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, WithoutFieldsMutation](ctx, wfu.sqlSave, wfu.mutation, wfu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wfu *WithoutFieldsUpdate) SaveX(ctx context.Context) int {
	affected, err := wfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wfu *WithoutFieldsUpdate) Exec(ctx context.Context) error {
	_, err := wfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wfu *WithoutFieldsUpdate) ExecX(ctx context.Context) {
	if err := wfu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wfu *WithoutFieldsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(withoutfields.Table, withoutfields.Columns, sqlgraph.NewFieldSpec(withoutfields.FieldID, field.TypeInt))
	if ps := wfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withoutfields.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wfu.mutation.done = true
	return n, nil
}

// WithoutFieldsUpdateOne is the builder for updating a single WithoutFields entity.
type WithoutFieldsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WithoutFieldsMutation
}

// Mutation returns the WithoutFieldsMutation object of the builder.
func (wfuo *WithoutFieldsUpdateOne) Mutation() *WithoutFieldsMutation {
	return wfuo.mutation
}

// Where appends a list predicates to the WithoutFieldsUpdate builder.
func (wfuo *WithoutFieldsUpdateOne) Where(ps ...predicate.WithoutFields) *WithoutFieldsUpdateOne {
	wfuo.mutation.Where(ps...)
	return wfuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wfuo *WithoutFieldsUpdateOne) Select(field string, fields ...string) *WithoutFieldsUpdateOne {
	wfuo.fields = append([]string{field}, fields...)
	return wfuo
}

// Save executes the query and returns the updated WithoutFields entity.
func (wfuo *WithoutFieldsUpdateOne) Save(ctx context.Context) (*WithoutFields, error) {
	return withHooks[*WithoutFields, WithoutFieldsMutation](ctx, wfuo.sqlSave, wfuo.mutation, wfuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wfuo *WithoutFieldsUpdateOne) SaveX(ctx context.Context) *WithoutFields {
	node, err := wfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wfuo *WithoutFieldsUpdateOne) Exec(ctx context.Context) error {
	_, err := wfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wfuo *WithoutFieldsUpdateOne) ExecX(ctx context.Context) {
	if err := wfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wfuo *WithoutFieldsUpdateOne) sqlSave(ctx context.Context) (_node *WithoutFields, err error) {
	_spec := sqlgraph.NewUpdateSpec(withoutfields.Table, withoutfields.Columns, sqlgraph.NewFieldSpec(withoutfields.FieldID, field.TypeInt))
	id, ok := wfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WithoutFields.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withoutfields.FieldID)
		for _, f := range fields {
			if !withoutfields.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != withoutfields.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &WithoutFields{config: wfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withoutfields.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wfuo.mutation.done = true
	return _node, nil
}
