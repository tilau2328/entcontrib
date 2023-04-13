// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tilau2328/entcontrib/entproto/internal/todo/ent/attachment"
	"github.com/tilau2328/entcontrib/entproto/internal/todo/ent/pet"
	"github.com/tilau2328/entcontrib/entproto/internal/todo/ent/predicate"
	"github.com/tilau2328/entcontrib/entproto/internal/todo/ent/user"
)

// PetUpdate is the builder for updating Pet entities.
type PetUpdate struct {
	config
	hooks    []Hook
	mutation *PetMutation
}

// Where appends a list predicates to the PetUpdate builder.
func (pu *PetUpdate) Where(ps ...predicate.Pet) *PetUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pu *PetUpdate) SetOwnerID(id uint32) *PetUpdate {
	pu.mutation.SetOwnerID(id)
	return pu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pu *PetUpdate) SetNillableOwnerID(id *uint32) *PetUpdate {
	if id != nil {
		pu = pu.SetOwnerID(*id)
	}
	return pu
}

// SetOwner sets the "owner" edge to the User entity.
func (pu *PetUpdate) SetOwner(u *User) *PetUpdate {
	return pu.SetOwnerID(u.ID)
}

// AddAttachmentIDs adds the "attachment" edge to the Attachment entity by IDs.
func (pu *PetUpdate) AddAttachmentIDs(ids ...uuid.UUID) *PetUpdate {
	pu.mutation.AddAttachmentIDs(ids...)
	return pu
}

// AddAttachment adds the "attachment" edges to the Attachment entity.
func (pu *PetUpdate) AddAttachment(a ...*Attachment) *PetUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.AddAttachmentIDs(ids...)
}

// Mutation returns the PetMutation object of the builder.
func (pu *PetUpdate) Mutation() *PetMutation {
	return pu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (pu *PetUpdate) ClearOwner() *PetUpdate {
	pu.mutation.ClearOwner()
	return pu
}

// ClearAttachment clears all "attachment" edges to the Attachment entity.
func (pu *PetUpdate) ClearAttachment() *PetUpdate {
	pu.mutation.ClearAttachment()
	return pu
}

// RemoveAttachmentIDs removes the "attachment" edge to Attachment entities by IDs.
func (pu *PetUpdate) RemoveAttachmentIDs(ids ...uuid.UUID) *PetUpdate {
	pu.mutation.RemoveAttachmentIDs(ids...)
	return pu
}

// RemoveAttachment removes "attachment" edges to Attachment entities.
func (pu *PetUpdate) RemoveAttachment(a ...*Attachment) *PetUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.RemoveAttachmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PetMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PetUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PetUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PetUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(pet.Table, pet.Columns, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.AttachmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pet.AttachmentTable,
			Columns: []string{pet.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedAttachmentIDs(); len(nodes) > 0 && !pu.mutation.AttachmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pet.AttachmentTable,
			Columns: []string{pet.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AttachmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pet.AttachmentTable,
			Columns: []string{pet.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PetUpdateOne is the builder for updating a single Pet entity.
type PetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PetMutation
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (puo *PetUpdateOne) SetOwnerID(id uint32) *PetUpdateOne {
	puo.mutation.SetOwnerID(id)
	return puo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (puo *PetUpdateOne) SetNillableOwnerID(id *uint32) *PetUpdateOne {
	if id != nil {
		puo = puo.SetOwnerID(*id)
	}
	return puo
}

// SetOwner sets the "owner" edge to the User entity.
func (puo *PetUpdateOne) SetOwner(u *User) *PetUpdateOne {
	return puo.SetOwnerID(u.ID)
}

// AddAttachmentIDs adds the "attachment" edge to the Attachment entity by IDs.
func (puo *PetUpdateOne) AddAttachmentIDs(ids ...uuid.UUID) *PetUpdateOne {
	puo.mutation.AddAttachmentIDs(ids...)
	return puo
}

// AddAttachment adds the "attachment" edges to the Attachment entity.
func (puo *PetUpdateOne) AddAttachment(a ...*Attachment) *PetUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.AddAttachmentIDs(ids...)
}

// Mutation returns the PetMutation object of the builder.
func (puo *PetUpdateOne) Mutation() *PetMutation {
	return puo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (puo *PetUpdateOne) ClearOwner() *PetUpdateOne {
	puo.mutation.ClearOwner()
	return puo
}

// ClearAttachment clears all "attachment" edges to the Attachment entity.
func (puo *PetUpdateOne) ClearAttachment() *PetUpdateOne {
	puo.mutation.ClearAttachment()
	return puo
}

// RemoveAttachmentIDs removes the "attachment" edge to Attachment entities by IDs.
func (puo *PetUpdateOne) RemoveAttachmentIDs(ids ...uuid.UUID) *PetUpdateOne {
	puo.mutation.RemoveAttachmentIDs(ids...)
	return puo
}

// RemoveAttachment removes "attachment" edges to Attachment entities.
func (puo *PetUpdateOne) RemoveAttachment(a ...*Attachment) *PetUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.RemoveAttachmentIDs(ids...)
}

// Where appends a list predicates to the PetUpdate builder.
func (puo *PetUpdateOne) Where(ps ...predicate.Pet) *PetUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PetUpdateOne) Select(field string, fields ...string) *PetUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pet entity.
func (puo *PetUpdateOne) Save(ctx context.Context) (*Pet, error) {
	return withHooks[*Pet, PetMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PetUpdateOne) SaveX(ctx context.Context) *Pet {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PetUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PetUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PetUpdateOne) sqlSave(ctx context.Context) (_node *Pet, err error) {
	_spec := sqlgraph.NewUpdateSpec(pet.Table, pet.Columns, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pet.FieldID)
		for _, f := range fields {
			if !pet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if puo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.AttachmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pet.AttachmentTable,
			Columns: []string{pet.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedAttachmentIDs(); len(nodes) > 0 && !puo.mutation.AttachmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pet.AttachmentTable,
			Columns: []string{pet.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AttachmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pet.AttachmentTable,
			Columns: []string{pet.AttachmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Pet{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
