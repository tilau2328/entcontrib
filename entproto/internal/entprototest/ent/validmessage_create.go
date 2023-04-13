// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/validmessage"
)

// ValidMessageCreate is the builder for creating a ValidMessage entity.
type ValidMessageCreate struct {
	config
	mutation *ValidMessageMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (vmc *ValidMessageCreate) SetName(s string) *ValidMessageCreate {
	vmc.mutation.SetName(s)
	return vmc
}

// SetTs sets the "ts" field.
func (vmc *ValidMessageCreate) SetTs(t time.Time) *ValidMessageCreate {
	vmc.mutation.SetTs(t)
	return vmc
}

// SetUUID sets the "uuid" field.
func (vmc *ValidMessageCreate) SetUUID(u uuid.UUID) *ValidMessageCreate {
	vmc.mutation.SetUUID(u)
	return vmc
}

// SetU8 sets the "u8" field.
func (vmc *ValidMessageCreate) SetU8(u uint8) *ValidMessageCreate {
	vmc.mutation.SetU8(u)
	return vmc
}

// SetOpti8 sets the "opti8" field.
func (vmc *ValidMessageCreate) SetOpti8(i int8) *ValidMessageCreate {
	vmc.mutation.SetOpti8(i)
	return vmc
}

// SetNillableOpti8 sets the "opti8" field if the given value is not nil.
func (vmc *ValidMessageCreate) SetNillableOpti8(i *int8) *ValidMessageCreate {
	if i != nil {
		vmc.SetOpti8(*i)
	}
	return vmc
}

// Mutation returns the ValidMessageMutation object of the builder.
func (vmc *ValidMessageCreate) Mutation() *ValidMessageMutation {
	return vmc.mutation
}

// Save creates the ValidMessage in the database.
func (vmc *ValidMessageCreate) Save(ctx context.Context) (*ValidMessage, error) {
	return withHooks[*ValidMessage, ValidMessageMutation](ctx, vmc.sqlSave, vmc.mutation, vmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vmc *ValidMessageCreate) SaveX(ctx context.Context) *ValidMessage {
	v, err := vmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vmc *ValidMessageCreate) Exec(ctx context.Context) error {
	_, err := vmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmc *ValidMessageCreate) ExecX(ctx context.Context) {
	if err := vmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vmc *ValidMessageCreate) check() error {
	if _, ok := vmc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ValidMessage.name"`)}
	}
	if _, ok := vmc.mutation.Ts(); !ok {
		return &ValidationError{Name: "ts", err: errors.New(`ent: missing required field "ValidMessage.ts"`)}
	}
	if _, ok := vmc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "ValidMessage.uuid"`)}
	}
	if _, ok := vmc.mutation.U8(); !ok {
		return &ValidationError{Name: "u8", err: errors.New(`ent: missing required field "ValidMessage.u8"`)}
	}
	return nil
}

func (vmc *ValidMessageCreate) sqlSave(ctx context.Context) (*ValidMessage, error) {
	if err := vmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vmc.mutation.id = &_node.ID
	vmc.mutation.done = true
	return _node, nil
}

func (vmc *ValidMessageCreate) createSpec() (*ValidMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &ValidMessage{config: vmc.config}
		_spec = sqlgraph.NewCreateSpec(validmessage.Table, sqlgraph.NewFieldSpec(validmessage.FieldID, field.TypeInt))
	)
	if value, ok := vmc.mutation.Name(); ok {
		_spec.SetField(validmessage.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := vmc.mutation.Ts(); ok {
		_spec.SetField(validmessage.FieldTs, field.TypeTime, value)
		_node.Ts = value
	}
	if value, ok := vmc.mutation.UUID(); ok {
		_spec.SetField(validmessage.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if value, ok := vmc.mutation.U8(); ok {
		_spec.SetField(validmessage.FieldU8, field.TypeUint8, value)
		_node.U8 = value
	}
	if value, ok := vmc.mutation.Opti8(); ok {
		_spec.SetField(validmessage.FieldOpti8, field.TypeInt8, value)
		_node.Opti8 = &value
	}
	return _node, _spec
}

// ValidMessageCreateBulk is the builder for creating many ValidMessage entities in bulk.
type ValidMessageCreateBulk struct {
	config
	builders []*ValidMessageCreate
}

// Save creates the ValidMessage entities in the database.
func (vmcb *ValidMessageCreateBulk) Save(ctx context.Context) ([]*ValidMessage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vmcb.builders))
	nodes := make([]*ValidMessage, len(vmcb.builders))
	mutators := make([]Mutator, len(vmcb.builders))
	for i := range vmcb.builders {
		func(i int, root context.Context) {
			builder := vmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ValidMessageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vmcb *ValidMessageCreateBulk) SaveX(ctx context.Context) []*ValidMessage {
	v, err := vmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vmcb *ValidMessageCreateBulk) Exec(ctx context.Context) error {
	_, err := vmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmcb *ValidMessageCreateBulk) ExecX(ctx context.Context) {
	if err := vmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
