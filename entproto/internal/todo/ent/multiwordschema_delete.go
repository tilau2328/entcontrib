// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/entproto/internal/todo/ent/multiwordschema"
	"github.com/tilau2328/entcontrib/entproto/internal/todo/ent/predicate"
)

// MultiWordSchemaDelete is the builder for deleting a MultiWordSchema entity.
type MultiWordSchemaDelete struct {
	config
	hooks    []Hook
	mutation *MultiWordSchemaMutation
}

// Where appends a list predicates to the MultiWordSchemaDelete builder.
func (mwsd *MultiWordSchemaDelete) Where(ps ...predicate.MultiWordSchema) *MultiWordSchemaDelete {
	mwsd.mutation.Where(ps...)
	return mwsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mwsd *MultiWordSchemaDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, MultiWordSchemaMutation](ctx, mwsd.sqlExec, mwsd.mutation, mwsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mwsd *MultiWordSchemaDelete) ExecX(ctx context.Context) int {
	n, err := mwsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mwsd *MultiWordSchemaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(multiwordschema.Table, sqlgraph.NewFieldSpec(multiwordschema.FieldID, field.TypeInt))
	if ps := mwsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mwsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mwsd.mutation.done = true
	return affected, err
}

// MultiWordSchemaDeleteOne is the builder for deleting a single MultiWordSchema entity.
type MultiWordSchemaDeleteOne struct {
	mwsd *MultiWordSchemaDelete
}

// Where appends a list predicates to the MultiWordSchemaDelete builder.
func (mwsdo *MultiWordSchemaDeleteOne) Where(ps ...predicate.MultiWordSchema) *MultiWordSchemaDeleteOne {
	mwsdo.mwsd.mutation.Where(ps...)
	return mwsdo
}

// Exec executes the deletion query.
func (mwsdo *MultiWordSchemaDeleteOne) Exec(ctx context.Context) error {
	n, err := mwsdo.mwsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{multiwordschema.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mwsdo *MultiWordSchemaDeleteOne) ExecX(ctx context.Context) {
	if err := mwsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
