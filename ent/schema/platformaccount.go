package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PlatformAccount holds the schema definition for the PlatformAccount entity.
type PlatformAccount struct {
	ent.Schema
}

// Fields of the PlatformAccount.
func (PlatformAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique(),
		field.String("password"),
		field.String("email").Optional().Default(""),
	}
}

// Edges of the PlatformAccount.
func (PlatformAccount) Edges() []ent.Edge {
	return nil
}
