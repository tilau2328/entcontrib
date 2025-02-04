// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/predicate"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/twomethodservice"
)

// TwoMethodServiceUpdate is the builder for updating TwoMethodService entities.
type TwoMethodServiceUpdate struct {
	config
	hooks    []Hook
	mutation *TwoMethodServiceMutation
}

// Where appends a list predicates to the TwoMethodServiceUpdate builder.
func (tmsu *TwoMethodServiceUpdate) Where(ps ...predicate.TwoMethodService) *TwoMethodServiceUpdate {
	tmsu.mutation.Where(ps...)
	return tmsu
}

// Mutation returns the TwoMethodServiceMutation object of the builder.
func (tmsu *TwoMethodServiceUpdate) Mutation() *TwoMethodServiceMutation {
	return tmsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tmsu *TwoMethodServiceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TwoMethodServiceMutation](ctx, tmsu.sqlSave, tmsu.mutation, tmsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tmsu *TwoMethodServiceUpdate) SaveX(ctx context.Context) int {
	affected, err := tmsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tmsu *TwoMethodServiceUpdate) Exec(ctx context.Context) error {
	_, err := tmsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmsu *TwoMethodServiceUpdate) ExecX(ctx context.Context) {
	if err := tmsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tmsu *TwoMethodServiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(twomethodservice.Table, twomethodservice.Columns, sqlgraph.NewFieldSpec(twomethodservice.FieldID, field.TypeInt))
	if ps := tmsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tmsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{twomethodservice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tmsu.mutation.done = true
	return n, nil
}

// TwoMethodServiceUpdateOne is the builder for updating a single TwoMethodService entity.
type TwoMethodServiceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TwoMethodServiceMutation
}

// Mutation returns the TwoMethodServiceMutation object of the builder.
func (tmsuo *TwoMethodServiceUpdateOne) Mutation() *TwoMethodServiceMutation {
	return tmsuo.mutation
}

// Where appends a list predicates to the TwoMethodServiceUpdate builder.
func (tmsuo *TwoMethodServiceUpdateOne) Where(ps ...predicate.TwoMethodService) *TwoMethodServiceUpdateOne {
	tmsuo.mutation.Where(ps...)
	return tmsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tmsuo *TwoMethodServiceUpdateOne) Select(field string, fields ...string) *TwoMethodServiceUpdateOne {
	tmsuo.fields = append([]string{field}, fields...)
	return tmsuo
}

// Save executes the query and returns the updated TwoMethodService entity.
func (tmsuo *TwoMethodServiceUpdateOne) Save(ctx context.Context) (*TwoMethodService, error) {
	return withHooks[*TwoMethodService, TwoMethodServiceMutation](ctx, tmsuo.sqlSave, tmsuo.mutation, tmsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tmsuo *TwoMethodServiceUpdateOne) SaveX(ctx context.Context) *TwoMethodService {
	node, err := tmsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tmsuo *TwoMethodServiceUpdateOne) Exec(ctx context.Context) error {
	_, err := tmsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmsuo *TwoMethodServiceUpdateOne) ExecX(ctx context.Context) {
	if err := tmsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tmsuo *TwoMethodServiceUpdateOne) sqlSave(ctx context.Context) (_node *TwoMethodService, err error) {
	_spec := sqlgraph.NewUpdateSpec(twomethodservice.Table, twomethodservice.Columns, sqlgraph.NewFieldSpec(twomethodservice.FieldID, field.TypeInt))
	id, ok := tmsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TwoMethodService.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tmsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, twomethodservice.FieldID)
		for _, f := range fields {
			if !twomethodservice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != twomethodservice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tmsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &TwoMethodService{config: tmsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tmsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{twomethodservice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tmsuo.mutation.done = true
	return _node, nil
}
