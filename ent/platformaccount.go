// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sysatom/framework/ent/platformaccount"
)

// PlatformAccount is the model entity for the PlatformAccount schema.
type PlatformAccount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Email holds the value of the "email" field.
	Email        string `json:"email,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PlatformAccount) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case platformaccount.FieldID:
			values[i] = new(sql.NullInt64)
		case platformaccount.FieldUsername, platformaccount.FieldPassword, platformaccount.FieldEmail:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PlatformAccount fields.
func (pa *PlatformAccount) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case platformaccount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case platformaccount.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				pa.Username = value.String
			}
		case platformaccount.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				pa.Password = value.String
			}
		case platformaccount.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				pa.Email = value.String
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PlatformAccount.
// This includes values selected through modifiers, order, etc.
func (pa *PlatformAccount) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// Update returns a builder for updating this PlatformAccount.
// Note that you need to call PlatformAccount.Unwrap() before calling this method if this PlatformAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *PlatformAccount) Update() *PlatformAccountUpdateOne {
	return NewPlatformAccountClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the PlatformAccount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *PlatformAccount) Unwrap() *PlatformAccount {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: PlatformAccount is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *PlatformAccount) String() string {
	var builder strings.Builder
	builder.WriteString("PlatformAccount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("username=")
	builder.WriteString(pa.Username)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(pa.Password)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(pa.Email)
	builder.WriteByte(')')
	return builder.String()
}

// PlatformAccounts is a parsable slice of PlatformAccount.
type PlatformAccounts []*PlatformAccount
