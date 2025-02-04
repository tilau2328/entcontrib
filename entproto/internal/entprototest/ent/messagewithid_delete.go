// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/messagewithid"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/predicate"
)

// MessageWithIDDelete is the builder for deleting a MessageWithID entity.
type MessageWithIDDelete struct {
	config
	hooks    []Hook
	mutation *MessageWithIDMutation
}

// Where appends a list predicates to the MessageWithIDDelete builder.
func (mwid *MessageWithIDDelete) Where(ps ...predicate.MessageWithID) *MessageWithIDDelete {
	mwid.mutation.Where(ps...)
	return mwid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mwid *MessageWithIDDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, MessageWithIDMutation](ctx, mwid.sqlExec, mwid.mutation, mwid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mwid *MessageWithIDDelete) ExecX(ctx context.Context) int {
	n, err := mwid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mwid *MessageWithIDDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(messagewithid.Table, sqlgraph.NewFieldSpec(messagewithid.FieldID, field.TypeInt32))
	if ps := mwid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mwid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mwid.mutation.done = true
	return affected, err
}

// MessageWithIDDeleteOne is the builder for deleting a single MessageWithID entity.
type MessageWithIDDeleteOne struct {
	mwid *MessageWithIDDelete
}

// Where appends a list predicates to the MessageWithIDDelete builder.
func (mwido *MessageWithIDDeleteOne) Where(ps ...predicate.MessageWithID) *MessageWithIDDeleteOne {
	mwido.mwid.mutation.Where(ps...)
	return mwido
}

// Exec executes the deletion query.
func (mwido *MessageWithIDDeleteOne) Exec(ctx context.Context) error {
	n, err := mwido.mwid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{messagewithid.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mwido *MessageWithIDDeleteOne) ExecX(ctx context.Context) {
	if err := mwido.Exec(ctx); err != nil {
		panic(err)
	}
}
