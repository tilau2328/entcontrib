// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/schemast/internal/mutatetest/ent/withoutfields"
)

// WithoutFieldsCreate is the builder for creating a WithoutFields entity.
type WithoutFieldsCreate struct {
	config
	mutation *WithoutFieldsMutation
	hooks    []Hook
}

// Mutation returns the WithoutFieldsMutation object of the builder.
func (wfc *WithoutFieldsCreate) Mutation() *WithoutFieldsMutation {
	return wfc.mutation
}

// Save creates the WithoutFields in the database.
func (wfc *WithoutFieldsCreate) Save(ctx context.Context) (*WithoutFields, error) {
	return withHooks[*WithoutFields, WithoutFieldsMutation](ctx, wfc.sqlSave, wfc.mutation, wfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wfc *WithoutFieldsCreate) SaveX(ctx context.Context) *WithoutFields {
	v, err := wfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wfc *WithoutFieldsCreate) Exec(ctx context.Context) error {
	_, err := wfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wfc *WithoutFieldsCreate) ExecX(ctx context.Context) {
	if err := wfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wfc *WithoutFieldsCreate) check() error {
	return nil
}

func (wfc *WithoutFieldsCreate) sqlSave(ctx context.Context) (*WithoutFields, error) {
	if err := wfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wfc.mutation.id = &_node.ID
	wfc.mutation.done = true
	return _node, nil
}

func (wfc *WithoutFieldsCreate) createSpec() (*WithoutFields, *sqlgraph.CreateSpec) {
	var (
		_node = &WithoutFields{config: wfc.config}
		_spec = sqlgraph.NewCreateSpec(withoutfields.Table, sqlgraph.NewFieldSpec(withoutfields.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// WithoutFieldsCreateBulk is the builder for creating many WithoutFields entities in bulk.
type WithoutFieldsCreateBulk struct {
	config
	builders []*WithoutFieldsCreate
}

// Save creates the WithoutFields entities in the database.
func (wfcb *WithoutFieldsCreateBulk) Save(ctx context.Context) ([]*WithoutFields, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wfcb.builders))
	nodes := make([]*WithoutFields, len(wfcb.builders))
	mutators := make([]Mutator, len(wfcb.builders))
	for i := range wfcb.builders {
		func(i int, root context.Context) {
			builder := wfcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WithoutFieldsMutation)
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
					_, err = mutators[i+1].Mutate(root, wfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wfcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wfcb *WithoutFieldsCreateBulk) SaveX(ctx context.Context) []*WithoutFields {
	v, err := wfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wfcb *WithoutFieldsCreateBulk) Exec(ctx context.Context) error {
	_, err := wfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wfcb *WithoutFieldsCreateBulk) ExecX(ctx context.Context) {
	if err := wfcb.Exec(ctx); err != nil {
		panic(err)
	}
}
