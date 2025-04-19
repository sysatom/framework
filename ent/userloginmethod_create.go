// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sysatom/framework/ent/userloginmethod"
)

// UserLoginMethodCreate is the builder for creating a UserLoginMethod entity.
type UserLoginMethodCreate struct {
	config
	mutation *UserLoginMethodMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ulmc *UserLoginMethodCreate) SetCreatedAt(t time.Time) *UserLoginMethodCreate {
	ulmc.mutation.SetCreatedAt(t)
	return ulmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ulmc *UserLoginMethodCreate) SetNillableCreatedAt(t *time.Time) *UserLoginMethodCreate {
	if t != nil {
		ulmc.SetCreatedAt(*t)
	}
	return ulmc
}

// SetUpdatedAt sets the "updated_at" field.
func (ulmc *UserLoginMethodCreate) SetUpdatedAt(t time.Time) *UserLoginMethodCreate {
	ulmc.mutation.SetUpdatedAt(t)
	return ulmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ulmc *UserLoginMethodCreate) SetNillableUpdatedAt(t *time.Time) *UserLoginMethodCreate {
	if t != nil {
		ulmc.SetUpdatedAt(*t)
	}
	return ulmc
}

// SetDeletedAt sets the "deleted_at" field.
func (ulmc *UserLoginMethodCreate) SetDeletedAt(t time.Time) *UserLoginMethodCreate {
	ulmc.mutation.SetDeletedAt(t)
	return ulmc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ulmc *UserLoginMethodCreate) SetNillableDeletedAt(t *time.Time) *UserLoginMethodCreate {
	if t != nil {
		ulmc.SetDeletedAt(*t)
	}
	return ulmc
}

// SetLoginType sets the "login_type" field.
func (ulmc *UserLoginMethodCreate) SetLoginType(s string) *UserLoginMethodCreate {
	ulmc.mutation.SetLoginType(s)
	return ulmc
}

// SetIdentifier sets the "identifier" field.
func (ulmc *UserLoginMethodCreate) SetIdentifier(s string) *UserLoginMethodCreate {
	ulmc.mutation.SetIdentifier(s)
	return ulmc
}

// SetID sets the "id" field.
func (ulmc *UserLoginMethodCreate) SetID(u uint64) *UserLoginMethodCreate {
	ulmc.mutation.SetID(u)
	return ulmc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ulmc *UserLoginMethodCreate) SetNillableID(u *uint64) *UserLoginMethodCreate {
	if u != nil {
		ulmc.SetID(*u)
	}
	return ulmc
}

// Mutation returns the UserLoginMethodMutation object of the builder.
func (ulmc *UserLoginMethodCreate) Mutation() *UserLoginMethodMutation {
	return ulmc.mutation
}

// Save creates the UserLoginMethod in the database.
func (ulmc *UserLoginMethodCreate) Save(ctx context.Context) (*UserLoginMethod, error) {
	ulmc.defaults()
	return withHooks(ctx, ulmc.sqlSave, ulmc.mutation, ulmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ulmc *UserLoginMethodCreate) SaveX(ctx context.Context) *UserLoginMethod {
	v, err := ulmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ulmc *UserLoginMethodCreate) Exec(ctx context.Context) error {
	_, err := ulmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulmc *UserLoginMethodCreate) ExecX(ctx context.Context) {
	if err := ulmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ulmc *UserLoginMethodCreate) defaults() {
	if _, ok := ulmc.mutation.CreatedAt(); !ok {
		v := userloginmethod.DefaultCreatedAt()
		ulmc.mutation.SetCreatedAt(v)
	}
	if _, ok := ulmc.mutation.UpdatedAt(); !ok {
		v := userloginmethod.DefaultUpdatedAt()
		ulmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ulmc.mutation.ID(); !ok {
		v := userloginmethod.DefaultID()
		ulmc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ulmc *UserLoginMethodCreate) check() error {
	if _, ok := ulmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserLoginMethod.created_at"`)}
	}
	if _, ok := ulmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserLoginMethod.updated_at"`)}
	}
	if _, ok := ulmc.mutation.LoginType(); !ok {
		return &ValidationError{Name: "login_type", err: errors.New(`ent: missing required field "UserLoginMethod.login_type"`)}
	}
	if _, ok := ulmc.mutation.Identifier(); !ok {
		return &ValidationError{Name: "identifier", err: errors.New(`ent: missing required field "UserLoginMethod.identifier"`)}
	}
	return nil
}

func (ulmc *UserLoginMethodCreate) sqlSave(ctx context.Context) (*UserLoginMethod, error) {
	if err := ulmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ulmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ulmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ulmc.mutation.id = &_node.ID
	ulmc.mutation.done = true
	return _node, nil
}

func (ulmc *UserLoginMethodCreate) createSpec() (*UserLoginMethod, *sqlgraph.CreateSpec) {
	var (
		_node = &UserLoginMethod{config: ulmc.config}
		_spec = sqlgraph.NewCreateSpec(userloginmethod.Table, sqlgraph.NewFieldSpec(userloginmethod.FieldID, field.TypeUint64))
	)
	if id, ok := ulmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ulmc.mutation.CreatedAt(); ok {
		_spec.SetField(userloginmethod.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ulmc.mutation.UpdatedAt(); ok {
		_spec.SetField(userloginmethod.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ulmc.mutation.DeletedAt(); ok {
		_spec.SetField(userloginmethod.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ulmc.mutation.LoginType(); ok {
		_spec.SetField(userloginmethod.FieldLoginType, field.TypeString, value)
		_node.LoginType = value
	}
	if value, ok := ulmc.mutation.Identifier(); ok {
		_spec.SetField(userloginmethod.FieldIdentifier, field.TypeString, value)
		_node.Identifier = value
	}
	return _node, _spec
}

// UserLoginMethodCreateBulk is the builder for creating many UserLoginMethod entities in bulk.
type UserLoginMethodCreateBulk struct {
	config
	err      error
	builders []*UserLoginMethodCreate
}

// Save creates the UserLoginMethod entities in the database.
func (ulmcb *UserLoginMethodCreateBulk) Save(ctx context.Context) ([]*UserLoginMethod, error) {
	if ulmcb.err != nil {
		return nil, ulmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ulmcb.builders))
	nodes := make([]*UserLoginMethod, len(ulmcb.builders))
	mutators := make([]Mutator, len(ulmcb.builders))
	for i := range ulmcb.builders {
		func(i int, root context.Context) {
			builder := ulmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserLoginMethodMutation)
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
					_, err = mutators[i+1].Mutate(root, ulmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ulmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ulmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ulmcb *UserLoginMethodCreateBulk) SaveX(ctx context.Context) []*UserLoginMethod {
	v, err := ulmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ulmcb *UserLoginMethodCreateBulk) Exec(ctx context.Context) error {
	_, err := ulmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulmcb *UserLoginMethodCreateBulk) ExecX(ctx context.Context) {
	if err := ulmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
