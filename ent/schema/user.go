package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/sysatom/framework/pkg/types"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").DefaultFunc(types.Id),
		field.String("username").Comment("user name"),
		field.String("phone").Optional().Default(""),
		field.String("email").Optional().Default(""),
		//field.Int64("introducer_id").Default(0),
		//field.Int64("default_merchant_id").Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("login_methods", UserLoginMethod.Type),
		edge.To("introducer", UserLoginMethod.Type).Unique(),
		edge.To("default_merchant", Merchant.Type).Unique(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
