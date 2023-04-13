// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/schemast/internal/mutatetest/ent/predicate"
	"github.com/tilau2328/entcontrib/schemast/internal/mutatetest/ent/withoutfields"
)

// WithoutFieldsDelete is the builder for deleting a WithoutFields entity.
type WithoutFieldsDelete struct {
	config
	hooks    []Hook
	mutation *WithoutFieldsMutation
}

// Where appends a list predicates to the WithoutFieldsDelete builder.
func (wfd *WithoutFieldsDelete) Where(ps ...predicate.WithoutFields) *WithoutFieldsDelete {
	wfd.mutation.Where(ps...)
	return wfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wfd *WithoutFieldsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, WithoutFieldsMutation](ctx, wfd.sqlExec, wfd.mutation, wfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wfd *WithoutFieldsDelete) ExecX(ctx context.Context) int {
	n, err := wfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wfd *WithoutFieldsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(withoutfields.Table, sqlgraph.NewFieldSpec(withoutfields.FieldID, field.TypeInt))
	if ps := wfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wfd.mutation.done = true
	return affected, err
}

// WithoutFieldsDeleteOne is the builder for deleting a single WithoutFields entity.
type WithoutFieldsDeleteOne struct {
	wfd *WithoutFieldsDelete
}

// Where appends a list predicates to the WithoutFieldsDelete builder.
func (wfdo *WithoutFieldsDeleteOne) Where(ps ...predicate.WithoutFields) *WithoutFieldsDeleteOne {
	wfdo.wfd.mutation.Where(ps...)
	return wfdo
}

// Exec executes the deletion query.
func (wfdo *WithoutFieldsDeleteOne) Exec(ctx context.Context) error {
	n, err := wfdo.wfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{withoutfields.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wfdo *WithoutFieldsDeleteOne) ExecX(ctx context.Context) {
	if err := wfdo.Exec(ctx); err != nil {
		panic(err)
	}
}
