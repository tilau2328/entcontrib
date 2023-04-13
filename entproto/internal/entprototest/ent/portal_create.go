// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/category"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/portal"
)

// PortalCreate is the builder for creating a Portal entity.
type PortalCreate struct {
	config
	mutation *PortalMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PortalCreate) SetName(s string) *PortalCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PortalCreate) SetDescription(s string) *PortalCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (pc *PortalCreate) SetCategoryID(id int) *PortalCreate {
	pc.mutation.SetCategoryID(id)
	return pc
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (pc *PortalCreate) SetNillableCategoryID(id *int) *PortalCreate {
	if id != nil {
		pc = pc.SetCategoryID(*id)
	}
	return pc
}

// SetCategory sets the "category" edge to the Category entity.
func (pc *PortalCreate) SetCategory(c *Category) *PortalCreate {
	return pc.SetCategoryID(c.ID)
}

// Mutation returns the PortalMutation object of the builder.
func (pc *PortalCreate) Mutation() *PortalMutation {
	return pc.mutation
}

// Save creates the Portal in the database.
func (pc *PortalCreate) Save(ctx context.Context) (*Portal, error) {
	return withHooks[*Portal, PortalMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PortalCreate) SaveX(ctx context.Context) *Portal {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PortalCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PortalCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PortalCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Portal.name"`)}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Portal.description"`)}
	}
	return nil
}

func (pc *PortalCreate) sqlSave(ctx context.Context) (*Portal, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PortalCreate) createSpec() (*Portal, *sqlgraph.CreateSpec) {
	var (
		_node = &Portal{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(portal.Table, sqlgraph.NewFieldSpec(portal.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(portal.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(portal.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := pc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   portal.CategoryTable,
			Columns: []string{portal.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.portal_category = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PortalCreateBulk is the builder for creating many Portal entities in bulk.
type PortalCreateBulk struct {
	config
	builders []*PortalCreate
}

// Save creates the Portal entities in the database.
func (pcb *PortalCreateBulk) Save(ctx context.Context) ([]*Portal, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Portal, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PortalMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PortalCreateBulk) SaveX(ctx context.Context) []*Portal {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PortalCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PortalCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
