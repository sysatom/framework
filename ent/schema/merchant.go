package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/sysatom/framework/pkg/types"
)

// Merchant holds the schema definition for the Merchant entity.
type Merchant struct {
	ent.Schema
}

// Fields of the Merchant.
func (Merchant) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").DefaultFunc(types.Id),
		field.String("merchant_name"),
		field.String("contact_person").Optional(),
		field.String("contact_phone").Optional(),
		field.String("country").Optional(),
		field.String("province").Optional(),
		field.String("city").Optional(),
		field.String("district").Optional(),
		field.String("address").Optional(),
	}
}

// Edges of the Merchant.
func (Merchant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", MerchantAccount.Type),
	}
}

func (Merchant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
