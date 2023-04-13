// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/messagewithfieldone"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/predicate"
)

// MessageWithFieldOneDelete is the builder for deleting a MessageWithFieldOne entity.
type MessageWithFieldOneDelete struct {
	config
	hooks    []Hook
	mutation *MessageWithFieldOneMutation
}

// Where appends a list predicates to the MessageWithFieldOneDelete builder.
func (mwfod *MessageWithFieldOneDelete) Where(ps ...predicate.MessageWithFieldOne) *MessageWithFieldOneDelete {
	mwfod.mutation.Where(ps...)
	return mwfod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mwfod *MessageWithFieldOneDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, MessageWithFieldOneMutation](ctx, mwfod.sqlExec, mwfod.mutation, mwfod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mwfod *MessageWithFieldOneDelete) ExecX(ctx context.Context) int {
	n, err := mwfod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mwfod *MessageWithFieldOneDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(messagewithfieldone.Table, sqlgraph.NewFieldSpec(messagewithfieldone.FieldID, field.TypeInt))
	if ps := mwfod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mwfod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mwfod.mutation.done = true
	return affected, err
}

// MessageWithFieldOneDeleteOne is the builder for deleting a single MessageWithFieldOne entity.
type MessageWithFieldOneDeleteOne struct {
	mwfod *MessageWithFieldOneDelete
}

// Where appends a list predicates to the MessageWithFieldOneDelete builder.
func (mwfodo *MessageWithFieldOneDeleteOne) Where(ps ...predicate.MessageWithFieldOne) *MessageWithFieldOneDeleteOne {
	mwfodo.mwfod.mutation.Where(ps...)
	return mwfodo
}

// Exec executes the deletion query.
func (mwfodo *MessageWithFieldOneDeleteOne) Exec(ctx context.Context) error {
	n, err := mwfodo.mwfod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{messagewithfieldone.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mwfodo *MessageWithFieldOneDeleteOne) ExecX(ctx context.Context) {
	if err := mwfodo.Exec(ctx); err != nil {
		panic(err)
	}
}
