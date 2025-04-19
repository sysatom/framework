package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/sysatom/framework/pkg/types"
)

// MerchantAccount holds the schema definition for the MerchantAccount entity.
type MerchantAccount struct {
	ent.Schema
}

// Fields of the MerchantAccount.
func (MerchantAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").DefaultFunc(types.Id),
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

func (MerchantAccount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_main_account"),
	}
}

func (MerchantAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
