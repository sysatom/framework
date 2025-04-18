// Code generated by ent, DO NOT EDIT.

package platformaccount

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/sysatom/framework/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldDeletedAt, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldPassword, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldEmail, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotNull(FieldDeletedAt))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldContainsFold(FieldPassword, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.FieldContainsFold(FieldEmail, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PlatformAccount) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PlatformAccount) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PlatformAccount) predicate.PlatformAccount {
	return predicate.PlatformAccount(sql.NotPredicates(p))
}
