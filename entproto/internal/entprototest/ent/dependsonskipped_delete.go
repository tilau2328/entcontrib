// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/dependsonskipped"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/predicate"
)

// DependsOnSkippedDelete is the builder for deleting a DependsOnSkipped entity.
type DependsOnSkippedDelete struct {
	config
	hooks    []Hook
	mutation *DependsOnSkippedMutation
}

// Where appends a list predicates to the DependsOnSkippedDelete builder.
func (dosd *DependsOnSkippedDelete) Where(ps ...predicate.DependsOnSkipped) *DependsOnSkippedDelete {
	dosd.mutation.Where(ps...)
	return dosd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dosd *DependsOnSkippedDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, DependsOnSkippedMutation](ctx, dosd.sqlExec, dosd.mutation, dosd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dosd *DependsOnSkippedDelete) ExecX(ctx context.Context) int {
	n, err := dosd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dosd *DependsOnSkippedDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(dependsonskipped.Table, sqlgraph.NewFieldSpec(dependsonskipped.FieldID, field.TypeInt))
	if ps := dosd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dosd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dosd.mutation.done = true
	return affected, err
}

// DependsOnSkippedDeleteOne is the builder for deleting a single DependsOnSkipped entity.
type DependsOnSkippedDeleteOne struct {
	dosd *DependsOnSkippedDelete
}

// Where appends a list predicates to the DependsOnSkippedDelete builder.
func (dosdo *DependsOnSkippedDeleteOne) Where(ps ...predicate.DependsOnSkipped) *DependsOnSkippedDeleteOne {
	dosdo.dosd.mutation.Where(ps...)
	return dosdo
}

// Exec executes the deletion query.
func (dosdo *DependsOnSkippedDeleteOne) Exec(ctx context.Context) error {
	n, err := dosdo.dosd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dependsonskipped.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dosdo *DependsOnSkippedDeleteOne) ExecX(ctx context.Context) {
	if err := dosdo.Exec(ctx); err != nil {
		panic(err)
	}
}
