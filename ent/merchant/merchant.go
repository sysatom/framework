// Code generated by ent, DO NOT EDIT.

package merchant

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the merchant type in the database.
	Label = "merchant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldMerchantName holds the string denoting the merchant_name field in the database.
	FieldMerchantName = "merchant_name"
	// FieldContactPerson holds the string denoting the contact_person field in the database.
	FieldContactPerson = "contact_person"
	// FieldContactPhone holds the string denoting the contact_phone field in the database.
	FieldContactPhone = "contact_phone"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldProvince holds the string denoting the province field in the database.
	FieldProvince = "province"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldDistrict holds the string denoting the district field in the database.
	FieldDistrict = "district"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// EdgeAccounts holds the string denoting the accounts edge name in mutations.
	EdgeAccounts = "accounts"
	// Table holds the table name of the merchant in the database.
	Table = "merchants"
	// AccountsTable is the table that holds the accounts relation/edge.
	AccountsTable = "merchant_accounts"
	// AccountsInverseTable is the table name for the MerchantAccount entity.
	// It exists in this package in order to avoid circular dependency with the "merchantaccount" package.
	AccountsInverseTable = "merchant_accounts"
	// AccountsColumn is the table column denoting the accounts relation/edge.
	AccountsColumn = "merchant_accounts"
)

// Columns holds all SQL columns for merchant fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldMerchantName,
	FieldContactPerson,
	FieldContactPhone,
	FieldCountry,
	FieldProvince,
	FieldCity,
	FieldDistrict,
	FieldAddress,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uint64
)

// OrderOption defines the ordering options for the Merchant queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByMerchantName orders the results by the merchant_name field.
func ByMerchantName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMerchantName, opts...).ToFunc()
}

// ByContactPerson orders the results by the contact_person field.
func ByContactPerson(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContactPerson, opts...).ToFunc()
}

// ByContactPhone orders the results by the contact_phone field.
func ByContactPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContactPhone, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// ByProvince orders the results by the province field.
func ByProvince(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProvince, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByDistrict orders the results by the district field.
func ByDistrict(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDistrict, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByAccountsCount orders the results by accounts count.
func ByAccountsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAccountsStep(), opts...)
	}
}

// ByAccounts orders the results by accounts terms.
func ByAccounts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAccountsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AccountsTable, AccountsColumn),
	)
}
