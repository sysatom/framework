// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sysatom/framework/ent/merchantaccount"
	"github.com/sysatom/framework/ent/predicate"
)

// MerchantAccountDelete is the builder for deleting a MerchantAccount entity.
type MerchantAccountDelete struct {
	config
	hooks    []Hook
	mutation *MerchantAccountMutation
}

// Where appends a list predicates to the MerchantAccountDelete builder.
func (mad *MerchantAccountDelete) Where(ps ...predicate.MerchantAccount) *MerchantAccountDelete {
	mad.mutation.Where(ps...)
	return mad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mad *MerchantAccountDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mad.sqlExec, mad.mutation, mad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mad *MerchantAccountDelete) ExecX(ctx context.Context) int {
	n, err := mad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mad *MerchantAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(merchantaccount.Table, sqlgraph.NewFieldSpec(merchantaccount.FieldID, field.TypeInt))
	if ps := mad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mad.mutation.done = true
	return affected, err
}

// MerchantAccountDeleteOne is the builder for deleting a single MerchantAccount entity.
type MerchantAccountDeleteOne struct {
	mad *MerchantAccountDelete
}

// Where appends a list predicates to the MerchantAccountDelete builder.
func (mado *MerchantAccountDeleteOne) Where(ps ...predicate.MerchantAccount) *MerchantAccountDeleteOne {
	mado.mad.mutation.Where(ps...)
	return mado
}

// Exec executes the deletion query.
func (mado *MerchantAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := mado.mad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{merchantaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mado *MerchantAccountDeleteOne) ExecX(ctx context.Context) {
	if err := mado.Exec(ctx); err != nil {
		panic(err)
	}
}
