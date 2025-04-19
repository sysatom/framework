// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sysatom/framework/ent/merchant"
	"github.com/sysatom/framework/ent/merchantaccount"
	"github.com/sysatom/framework/ent/predicate"
)

// MerchantUpdate is the builder for updating Merchant entities.
type MerchantUpdate struct {
	config
	hooks    []Hook
	mutation *MerchantMutation
}

// Where appends a list predicates to the MerchantUpdate builder.
func (mu *MerchantUpdate) Where(ps ...predicate.Merchant) *MerchantUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MerchantUpdate) SetUpdatedAt(t time.Time) *MerchantUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetDeletedAt sets the "deleted_at" field.
func (mu *MerchantUpdate) SetDeletedAt(t time.Time) *MerchantUpdate {
	mu.mutation.SetDeletedAt(t)
	return mu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableDeletedAt(t *time.Time) *MerchantUpdate {
	if t != nil {
		mu.SetDeletedAt(*t)
	}
	return mu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (mu *MerchantUpdate) ClearDeletedAt() *MerchantUpdate {
	mu.mutation.ClearDeletedAt()
	return mu
}

// SetMerchantName sets the "merchant_name" field.
func (mu *MerchantUpdate) SetMerchantName(s string) *MerchantUpdate {
	mu.mutation.SetMerchantName(s)
	return mu
}

// SetNillableMerchantName sets the "merchant_name" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableMerchantName(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetMerchantName(*s)
	}
	return mu
}

// SetContactPerson sets the "contact_person" field.
func (mu *MerchantUpdate) SetContactPerson(s string) *MerchantUpdate {
	mu.mutation.SetContactPerson(s)
	return mu
}

// SetNillableContactPerson sets the "contact_person" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableContactPerson(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetContactPerson(*s)
	}
	return mu
}

// ClearContactPerson clears the value of the "contact_person" field.
func (mu *MerchantUpdate) ClearContactPerson() *MerchantUpdate {
	mu.mutation.ClearContactPerson()
	return mu
}

// SetContactPhone sets the "contact_phone" field.
func (mu *MerchantUpdate) SetContactPhone(s string) *MerchantUpdate {
	mu.mutation.SetContactPhone(s)
	return mu
}

// SetNillableContactPhone sets the "contact_phone" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableContactPhone(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetContactPhone(*s)
	}
	return mu
}

// ClearContactPhone clears the value of the "contact_phone" field.
func (mu *MerchantUpdate) ClearContactPhone() *MerchantUpdate {
	mu.mutation.ClearContactPhone()
	return mu
}

// SetCountry sets the "country" field.
func (mu *MerchantUpdate) SetCountry(s string) *MerchantUpdate {
	mu.mutation.SetCountry(s)
	return mu
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableCountry(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetCountry(*s)
	}
	return mu
}

// ClearCountry clears the value of the "country" field.
func (mu *MerchantUpdate) ClearCountry() *MerchantUpdate {
	mu.mutation.ClearCountry()
	return mu
}

// SetProvince sets the "province" field.
func (mu *MerchantUpdate) SetProvince(s string) *MerchantUpdate {
	mu.mutation.SetProvince(s)
	return mu
}

// SetNillableProvince sets the "province" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableProvince(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetProvince(*s)
	}
	return mu
}

// ClearProvince clears the value of the "province" field.
func (mu *MerchantUpdate) ClearProvince() *MerchantUpdate {
	mu.mutation.ClearProvince()
	return mu
}

// SetCity sets the "city" field.
func (mu *MerchantUpdate) SetCity(s string) *MerchantUpdate {
	mu.mutation.SetCity(s)
	return mu
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableCity(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetCity(*s)
	}
	return mu
}

// ClearCity clears the value of the "city" field.
func (mu *MerchantUpdate) ClearCity() *MerchantUpdate {
	mu.mutation.ClearCity()
	return mu
}

// SetDistrict sets the "district" field.
func (mu *MerchantUpdate) SetDistrict(s string) *MerchantUpdate {
	mu.mutation.SetDistrict(s)
	return mu
}

// SetNillableDistrict sets the "district" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableDistrict(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetDistrict(*s)
	}
	return mu
}

// ClearDistrict clears the value of the "district" field.
func (mu *MerchantUpdate) ClearDistrict() *MerchantUpdate {
	mu.mutation.ClearDistrict()
	return mu
}

// SetAddress sets the "address" field.
func (mu *MerchantUpdate) SetAddress(s string) *MerchantUpdate {
	mu.mutation.SetAddress(s)
	return mu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (mu *MerchantUpdate) SetNillableAddress(s *string) *MerchantUpdate {
	if s != nil {
		mu.SetAddress(*s)
	}
	return mu
}

// ClearAddress clears the value of the "address" field.
func (mu *MerchantUpdate) ClearAddress() *MerchantUpdate {
	mu.mutation.ClearAddress()
	return mu
}

// AddAccountIDs adds the "accounts" edge to the MerchantAccount entity by IDs.
func (mu *MerchantUpdate) AddAccountIDs(ids ...uint64) *MerchantUpdate {
	mu.mutation.AddAccountIDs(ids...)
	return mu
}

// AddAccounts adds the "accounts" edges to the MerchantAccount entity.
func (mu *MerchantUpdate) AddAccounts(m ...*MerchantAccount) *MerchantUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddAccountIDs(ids...)
}

// Mutation returns the MerchantMutation object of the builder.
func (mu *MerchantUpdate) Mutation() *MerchantMutation {
	return mu.mutation
}

// ClearAccounts clears all "accounts" edges to the MerchantAccount entity.
func (mu *MerchantUpdate) ClearAccounts() *MerchantUpdate {
	mu.mutation.ClearAccounts()
	return mu
}

// RemoveAccountIDs removes the "accounts" edge to MerchantAccount entities by IDs.
func (mu *MerchantUpdate) RemoveAccountIDs(ids ...uint64) *MerchantUpdate {
	mu.mutation.RemoveAccountIDs(ids...)
	return mu
}

// RemoveAccounts removes "accounts" edges to MerchantAccount entities.
func (mu *MerchantUpdate) RemoveAccounts(m ...*MerchantAccount) *MerchantUpdate {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveAccountIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MerchantUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MerchantUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MerchantUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MerchantUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MerchantUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := merchant.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

func (mu *MerchantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(merchant.Table, merchant.Columns, sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(merchant.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.DeletedAt(); ok {
		_spec.SetField(merchant.FieldDeletedAt, field.TypeTime, value)
	}
	if mu.mutation.DeletedAtCleared() {
		_spec.ClearField(merchant.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := mu.mutation.MerchantName(); ok {
		_spec.SetField(merchant.FieldMerchantName, field.TypeString, value)
	}
	if value, ok := mu.mutation.ContactPerson(); ok {
		_spec.SetField(merchant.FieldContactPerson, field.TypeString, value)
	}
	if mu.mutation.ContactPersonCleared() {
		_spec.ClearField(merchant.FieldContactPerson, field.TypeString)
	}
	if value, ok := mu.mutation.ContactPhone(); ok {
		_spec.SetField(merchant.FieldContactPhone, field.TypeString, value)
	}
	if mu.mutation.ContactPhoneCleared() {
		_spec.ClearField(merchant.FieldContactPhone, field.TypeString)
	}
	if value, ok := mu.mutation.Country(); ok {
		_spec.SetField(merchant.FieldCountry, field.TypeString, value)
	}
	if mu.mutation.CountryCleared() {
		_spec.ClearField(merchant.FieldCountry, field.TypeString)
	}
	if value, ok := mu.mutation.Province(); ok {
		_spec.SetField(merchant.FieldProvince, field.TypeString, value)
	}
	if mu.mutation.ProvinceCleared() {
		_spec.ClearField(merchant.FieldProvince, field.TypeString)
	}
	if value, ok := mu.mutation.City(); ok {
		_spec.SetField(merchant.FieldCity, field.TypeString, value)
	}
	if mu.mutation.CityCleared() {
		_spec.ClearField(merchant.FieldCity, field.TypeString)
	}
	if value, ok := mu.mutation.District(); ok {
		_spec.SetField(merchant.FieldDistrict, field.TypeString, value)
	}
	if mu.mutation.DistrictCleared() {
		_spec.ClearField(merchant.FieldDistrict, field.TypeString)
	}
	if value, ok := mu.mutation.Address(); ok {
		_spec.SetField(merchant.FieldAddress, field.TypeString, value)
	}
	if mu.mutation.AddressCleared() {
		_spec.ClearField(merchant.FieldAddress, field.TypeString)
	}
	if mu.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.AccountsTable,
			Columns: []string{merchant.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantaccount.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !mu.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.AccountsTable,
			Columns: []string{merchant.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantaccount.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.AccountsTable,
			Columns: []string{merchant.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantaccount.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MerchantUpdateOne is the builder for updating a single Merchant entity.
type MerchantUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MerchantMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MerchantUpdateOne) SetUpdatedAt(t time.Time) *MerchantUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetDeletedAt sets the "deleted_at" field.
func (muo *MerchantUpdateOne) SetDeletedAt(t time.Time) *MerchantUpdateOne {
	muo.mutation.SetDeletedAt(t)
	return muo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableDeletedAt(t *time.Time) *MerchantUpdateOne {
	if t != nil {
		muo.SetDeletedAt(*t)
	}
	return muo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (muo *MerchantUpdateOne) ClearDeletedAt() *MerchantUpdateOne {
	muo.mutation.ClearDeletedAt()
	return muo
}

// SetMerchantName sets the "merchant_name" field.
func (muo *MerchantUpdateOne) SetMerchantName(s string) *MerchantUpdateOne {
	muo.mutation.SetMerchantName(s)
	return muo
}

// SetNillableMerchantName sets the "merchant_name" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableMerchantName(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetMerchantName(*s)
	}
	return muo
}

// SetContactPerson sets the "contact_person" field.
func (muo *MerchantUpdateOne) SetContactPerson(s string) *MerchantUpdateOne {
	muo.mutation.SetContactPerson(s)
	return muo
}

// SetNillableContactPerson sets the "contact_person" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableContactPerson(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetContactPerson(*s)
	}
	return muo
}

// ClearContactPerson clears the value of the "contact_person" field.
func (muo *MerchantUpdateOne) ClearContactPerson() *MerchantUpdateOne {
	muo.mutation.ClearContactPerson()
	return muo
}

// SetContactPhone sets the "contact_phone" field.
func (muo *MerchantUpdateOne) SetContactPhone(s string) *MerchantUpdateOne {
	muo.mutation.SetContactPhone(s)
	return muo
}

// SetNillableContactPhone sets the "contact_phone" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableContactPhone(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetContactPhone(*s)
	}
	return muo
}

// ClearContactPhone clears the value of the "contact_phone" field.
func (muo *MerchantUpdateOne) ClearContactPhone() *MerchantUpdateOne {
	muo.mutation.ClearContactPhone()
	return muo
}

// SetCountry sets the "country" field.
func (muo *MerchantUpdateOne) SetCountry(s string) *MerchantUpdateOne {
	muo.mutation.SetCountry(s)
	return muo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableCountry(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetCountry(*s)
	}
	return muo
}

// ClearCountry clears the value of the "country" field.
func (muo *MerchantUpdateOne) ClearCountry() *MerchantUpdateOne {
	muo.mutation.ClearCountry()
	return muo
}

// SetProvince sets the "province" field.
func (muo *MerchantUpdateOne) SetProvince(s string) *MerchantUpdateOne {
	muo.mutation.SetProvince(s)
	return muo
}

// SetNillableProvince sets the "province" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableProvince(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetProvince(*s)
	}
	return muo
}

// ClearProvince clears the value of the "province" field.
func (muo *MerchantUpdateOne) ClearProvince() *MerchantUpdateOne {
	muo.mutation.ClearProvince()
	return muo
}

// SetCity sets the "city" field.
func (muo *MerchantUpdateOne) SetCity(s string) *MerchantUpdateOne {
	muo.mutation.SetCity(s)
	return muo
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableCity(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetCity(*s)
	}
	return muo
}

// ClearCity clears the value of the "city" field.
func (muo *MerchantUpdateOne) ClearCity() *MerchantUpdateOne {
	muo.mutation.ClearCity()
	return muo
}

// SetDistrict sets the "district" field.
func (muo *MerchantUpdateOne) SetDistrict(s string) *MerchantUpdateOne {
	muo.mutation.SetDistrict(s)
	return muo
}

// SetNillableDistrict sets the "district" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableDistrict(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetDistrict(*s)
	}
	return muo
}

// ClearDistrict clears the value of the "district" field.
func (muo *MerchantUpdateOne) ClearDistrict() *MerchantUpdateOne {
	muo.mutation.ClearDistrict()
	return muo
}

// SetAddress sets the "address" field.
func (muo *MerchantUpdateOne) SetAddress(s string) *MerchantUpdateOne {
	muo.mutation.SetAddress(s)
	return muo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (muo *MerchantUpdateOne) SetNillableAddress(s *string) *MerchantUpdateOne {
	if s != nil {
		muo.SetAddress(*s)
	}
	return muo
}

// ClearAddress clears the value of the "address" field.
func (muo *MerchantUpdateOne) ClearAddress() *MerchantUpdateOne {
	muo.mutation.ClearAddress()
	return muo
}

// AddAccountIDs adds the "accounts" edge to the MerchantAccount entity by IDs.
func (muo *MerchantUpdateOne) AddAccountIDs(ids ...uint64) *MerchantUpdateOne {
	muo.mutation.AddAccountIDs(ids...)
	return muo
}

// AddAccounts adds the "accounts" edges to the MerchantAccount entity.
func (muo *MerchantUpdateOne) AddAccounts(m ...*MerchantAccount) *MerchantUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddAccountIDs(ids...)
}

// Mutation returns the MerchantMutation object of the builder.
func (muo *MerchantUpdateOne) Mutation() *MerchantMutation {
	return muo.mutation
}

// ClearAccounts clears all "accounts" edges to the MerchantAccount entity.
func (muo *MerchantUpdateOne) ClearAccounts() *MerchantUpdateOne {
	muo.mutation.ClearAccounts()
	return muo
}

// RemoveAccountIDs removes the "accounts" edge to MerchantAccount entities by IDs.
func (muo *MerchantUpdateOne) RemoveAccountIDs(ids ...uint64) *MerchantUpdateOne {
	muo.mutation.RemoveAccountIDs(ids...)
	return muo
}

// RemoveAccounts removes "accounts" edges to MerchantAccount entities.
func (muo *MerchantUpdateOne) RemoveAccounts(m ...*MerchantAccount) *MerchantUpdateOne {
	ids := make([]uint64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveAccountIDs(ids...)
}

// Where appends a list predicates to the MerchantUpdate builder.
func (muo *MerchantUpdateOne) Where(ps ...predicate.Merchant) *MerchantUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MerchantUpdateOne) Select(field string, fields ...string) *MerchantUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Merchant entity.
func (muo *MerchantUpdateOne) Save(ctx context.Context) (*Merchant, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MerchantUpdateOne) SaveX(ctx context.Context) *Merchant {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MerchantUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MerchantUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MerchantUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := merchant.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

func (muo *MerchantUpdateOne) sqlSave(ctx context.Context) (_node *Merchant, err error) {
	_spec := sqlgraph.NewUpdateSpec(merchant.Table, merchant.Columns, sqlgraph.NewFieldSpec(merchant.FieldID, field.TypeUint64))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Merchant.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, merchant.FieldID)
		for _, f := range fields {
			if !merchant.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != merchant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(merchant.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.DeletedAt(); ok {
		_spec.SetField(merchant.FieldDeletedAt, field.TypeTime, value)
	}
	if muo.mutation.DeletedAtCleared() {
		_spec.ClearField(merchant.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := muo.mutation.MerchantName(); ok {
		_spec.SetField(merchant.FieldMerchantName, field.TypeString, value)
	}
	if value, ok := muo.mutation.ContactPerson(); ok {
		_spec.SetField(merchant.FieldContactPerson, field.TypeString, value)
	}
	if muo.mutation.ContactPersonCleared() {
		_spec.ClearField(merchant.FieldContactPerson, field.TypeString)
	}
	if value, ok := muo.mutation.ContactPhone(); ok {
		_spec.SetField(merchant.FieldContactPhone, field.TypeString, value)
	}
	if muo.mutation.ContactPhoneCleared() {
		_spec.ClearField(merchant.FieldContactPhone, field.TypeString)
	}
	if value, ok := muo.mutation.Country(); ok {
		_spec.SetField(merchant.FieldCountry, field.TypeString, value)
	}
	if muo.mutation.CountryCleared() {
		_spec.ClearField(merchant.FieldCountry, field.TypeString)
	}
	if value, ok := muo.mutation.Province(); ok {
		_spec.SetField(merchant.FieldProvince, field.TypeString, value)
	}
	if muo.mutation.ProvinceCleared() {
		_spec.ClearField(merchant.FieldProvince, field.TypeString)
	}
	if value, ok := muo.mutation.City(); ok {
		_spec.SetField(merchant.FieldCity, field.TypeString, value)
	}
	if muo.mutation.CityCleared() {
		_spec.ClearField(merchant.FieldCity, field.TypeString)
	}
	if value, ok := muo.mutation.District(); ok {
		_spec.SetField(merchant.FieldDistrict, field.TypeString, value)
	}
	if muo.mutation.DistrictCleared() {
		_spec.ClearField(merchant.FieldDistrict, field.TypeString)
	}
	if value, ok := muo.mutation.Address(); ok {
		_spec.SetField(merchant.FieldAddress, field.TypeString, value)
	}
	if muo.mutation.AddressCleared() {
		_spec.ClearField(merchant.FieldAddress, field.TypeString)
	}
	if muo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.AccountsTable,
			Columns: []string{merchant.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantaccount.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !muo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.AccountsTable,
			Columns: []string{merchant.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantaccount.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   merchant.AccountsTable,
			Columns: []string{merchant.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(merchantaccount.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Merchant{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{merchant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
