// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sysatom/framework/ent/platformaccount"
)

// PlatformAccountCreate is the builder for creating a PlatformAccount entity.
type PlatformAccountCreate struct {
	config
	mutation *PlatformAccountMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (pac *PlatformAccountCreate) SetUsername(s string) *PlatformAccountCreate {
	pac.mutation.SetUsername(s)
	return pac
}

// SetPassword sets the "password" field.
func (pac *PlatformAccountCreate) SetPassword(s string) *PlatformAccountCreate {
	pac.mutation.SetPassword(s)
	return pac
}

// SetEmail sets the "email" field.
func (pac *PlatformAccountCreate) SetEmail(s string) *PlatformAccountCreate {
	pac.mutation.SetEmail(s)
	return pac
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (pac *PlatformAccountCreate) SetNillableEmail(s *string) *PlatformAccountCreate {
	if s != nil {
		pac.SetEmail(*s)
	}
	return pac
}

// SetID sets the "id" field.
func (pac *PlatformAccountCreate) SetID(u uint64) *PlatformAccountCreate {
	pac.mutation.SetID(u)
	return pac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pac *PlatformAccountCreate) SetNillableID(u *uint64) *PlatformAccountCreate {
	if u != nil {
		pac.SetID(*u)
	}
	return pac
}

// Mutation returns the PlatformAccountMutation object of the builder.
func (pac *PlatformAccountCreate) Mutation() *PlatformAccountMutation {
	return pac.mutation
}

// Save creates the PlatformAccount in the database.
func (pac *PlatformAccountCreate) Save(ctx context.Context) (*PlatformAccount, error) {
	pac.defaults()
	return withHooks(ctx, pac.sqlSave, pac.mutation, pac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pac *PlatformAccountCreate) SaveX(ctx context.Context) *PlatformAccount {
	v, err := pac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pac *PlatformAccountCreate) Exec(ctx context.Context) error {
	_, err := pac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pac *PlatformAccountCreate) ExecX(ctx context.Context) {
	if err := pac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pac *PlatformAccountCreate) defaults() {
	if _, ok := pac.mutation.Email(); !ok {
		v := platformaccount.DefaultEmail
		pac.mutation.SetEmail(v)
	}
	if _, ok := pac.mutation.ID(); !ok {
		v := platformaccount.DefaultID()
		pac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pac *PlatformAccountCreate) check() error {
	if _, ok := pac.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "PlatformAccount.username"`)}
	}
	if _, ok := pac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "PlatformAccount.password"`)}
	}
	return nil
}

func (pac *PlatformAccountCreate) sqlSave(ctx context.Context) (*PlatformAccount, error) {
	if err := pac.check(); err != nil {
		return nil, err
	}
	_node, _spec := pac.createSpec()
	if err := sqlgraph.CreateNode(ctx, pac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	pac.mutation.id = &_node.ID
	pac.mutation.done = true
	return _node, nil
}

func (pac *PlatformAccountCreate) createSpec() (*PlatformAccount, *sqlgraph.CreateSpec) {
	var (
		_node = &PlatformAccount{config: pac.config}
		_spec = sqlgraph.NewCreateSpec(platformaccount.Table, sqlgraph.NewFieldSpec(platformaccount.FieldID, field.TypeUint64))
	)
	if id, ok := pac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pac.mutation.Username(); ok {
		_spec.SetField(platformaccount.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := pac.mutation.Password(); ok {
		_spec.SetField(platformaccount.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := pac.mutation.Email(); ok {
		_spec.SetField(platformaccount.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	return _node, _spec
}

// PlatformAccountCreateBulk is the builder for creating many PlatformAccount entities in bulk.
type PlatformAccountCreateBulk struct {
	config
	err      error
	builders []*PlatformAccountCreate
}

// Save creates the PlatformAccount entities in the database.
func (pacb *PlatformAccountCreateBulk) Save(ctx context.Context) ([]*PlatformAccount, error) {
	if pacb.err != nil {
		return nil, pacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pacb.builders))
	nodes := make([]*PlatformAccount, len(pacb.builders))
	mutators := make([]Mutator, len(pacb.builders))
	for i := range pacb.builders {
		func(i int, root context.Context) {
			builder := pacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlatformAccountMutation)
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
					_, err = mutators[i+1].Mutate(root, pacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, pacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pacb *PlatformAccountCreateBulk) SaveX(ctx context.Context) []*PlatformAccount {
	v, err := pacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pacb *PlatformAccountCreateBulk) Exec(ctx context.Context) error {
	_, err := pacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pacb *PlatformAccountCreateBulk) ExecX(ctx context.Context) {
	if err := pacb.Exec(ctx); err != nil {
		panic(err)
	}
}
