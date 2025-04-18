// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sysatom/framework/ent/merchant"
)

// Merchant is the model entity for the Merchant schema.
type Merchant struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// MerchantName holds the value of the "merchant_name" field.
	MerchantName string `json:"merchant_name,omitempty"`
	// ContactPerson holds the value of the "contact_person" field.
	ContactPerson string `json:"contact_person,omitempty"`
	// ContactPhone holds the value of the "contact_phone" field.
	ContactPhone string `json:"contact_phone,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Province holds the value of the "province" field.
	Province string `json:"province,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// District holds the value of the "district" field.
	District string `json:"district,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MerchantQuery when eager-loading is set.
	Edges        MerchantEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MerchantEdges holds the relations/edges for other nodes in the graph.
type MerchantEdges struct {
	// Accounts holds the value of the accounts edge.
	Accounts []*MerchantAccount `json:"accounts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AccountsOrErr returns the Accounts value or an error if the edge
// was not loaded in eager-loading.
func (e MerchantEdges) AccountsOrErr() ([]*MerchantAccount, error) {
	if e.loadedTypes[0] {
		return e.Accounts, nil
	}
	return nil, &NotLoadedError{edge: "accounts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Merchant) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case merchant.FieldID:
			values[i] = new(sql.NullInt64)
		case merchant.FieldMerchantName, merchant.FieldContactPerson, merchant.FieldContactPhone, merchant.FieldCountry, merchant.FieldProvince, merchant.FieldCity, merchant.FieldDistrict, merchant.FieldAddress:
			values[i] = new(sql.NullString)
		case merchant.FieldCreatedAt, merchant.FieldUpdatedAt, merchant.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Merchant fields.
func (m *Merchant) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case merchant.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = uint64(value.Int64)
		case merchant.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case merchant.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case merchant.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				m.DeletedAt = value.Time
			}
		case merchant.FieldMerchantName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field merchant_name", values[i])
			} else if value.Valid {
				m.MerchantName = value.String
			}
		case merchant.FieldContactPerson:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact_person", values[i])
			} else if value.Valid {
				m.ContactPerson = value.String
			}
		case merchant.FieldContactPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact_phone", values[i])
			} else if value.Valid {
				m.ContactPhone = value.String
			}
		case merchant.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				m.Country = value.String
			}
		case merchant.FieldProvince:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province", values[i])
			} else if value.Valid {
				m.Province = value.String
			}
		case merchant.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				m.City = value.String
			}
		case merchant.FieldDistrict:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field district", values[i])
			} else if value.Valid {
				m.District = value.String
			}
		case merchant.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				m.Address = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Merchant.
// This includes values selected through modifiers, order, etc.
func (m *Merchant) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryAccounts queries the "accounts" edge of the Merchant entity.
func (m *Merchant) QueryAccounts() *MerchantAccountQuery {
	return NewMerchantClient(m.config).QueryAccounts(m)
}

// Update returns a builder for updating this Merchant.
// Note that you need to call Merchant.Unwrap() before calling this method if this Merchant
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Merchant) Update() *MerchantUpdateOne {
	return NewMerchantClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Merchant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Merchant) Unwrap() *Merchant {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Merchant is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Merchant) String() string {
	var builder strings.Builder
	builder.WriteString("Merchant(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(m.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("merchant_name=")
	builder.WriteString(m.MerchantName)
	builder.WriteString(", ")
	builder.WriteString("contact_person=")
	builder.WriteString(m.ContactPerson)
	builder.WriteString(", ")
	builder.WriteString("contact_phone=")
	builder.WriteString(m.ContactPhone)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(m.Country)
	builder.WriteString(", ")
	builder.WriteString("province=")
	builder.WriteString(m.Province)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(m.City)
	builder.WriteString(", ")
	builder.WriteString("district=")
	builder.WriteString(m.District)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(m.Address)
	builder.WriteByte(')')
	return builder.String()
}

// Merchants is a parsable slice of Merchant.
type Merchants []*Merchant
