package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
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
