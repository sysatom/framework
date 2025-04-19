package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MerchantAccount holds the schema definition for the MerchantAccount entity.
type MerchantAccount struct {
	ent.Schema
}

// Fields of the MerchantAccount.
func (MerchantAccount) Fields() []ent.Field {
	return []ent.Field{
		//field.Int64("merchant_id"),
		field.String("username").Unique(),
		field.String("password"),
		field.String("email").Optional().Default(""),
		field.String("phone").Optional().Default(""),
		field.Bool("is_main_account").Default(false),
	}
}

// Edges of the MerchantAccount.
func (MerchantAccount) Edges() []ent.Edge {
	return nil
}
