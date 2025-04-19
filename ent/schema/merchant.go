package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Merchant holds the schema definition for the Merchant entity.
type Merchant struct {
	ent.Schema
}

// Fields of the Merchant.
func (Merchant) Fields() []ent.Field {
	return []ent.Field{
		field.String("merchant_name"),
		field.String("contact_person"),
		field.String("contact_phone"),
		field.String("country"),
		field.String("province"),
		field.String("city"),
		field.String("district"),
		field.String("address"),
	}
}

// Edges of the Merchant.
func (Merchant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", MerchantAccount.Type),
	}
}
