// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/entproto/internal/todo/ent/attachment"
	"github.com/tilau2328/entcontrib/entproto/internal/todo/ent/predicate"
)

// AttachmentDelete is the builder for deleting a Attachment entity.
type AttachmentDelete struct {
	config
	hooks    []Hook
	mutation *AttachmentMutation
}

// Where appends a list predicates to the AttachmentDelete builder.
func (ad *AttachmentDelete) Where(ps ...predicate.Attachment) *AttachmentDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AttachmentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, AttachmentMutation](ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AttachmentDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AttachmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(attachment.Table, sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeUUID))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AttachmentDeleteOne is the builder for deleting a single Attachment entity.
type AttachmentDeleteOne struct {
	ad *AttachmentDelete
}

// Where appends a list predicates to the AttachmentDelete builder.
func (ado *AttachmentDeleteOne) Where(ps ...predicate.Attachment) *AttachmentDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *AttachmentDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{attachment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AttachmentDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
