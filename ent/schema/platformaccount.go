package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/sysatom/framework/pkg/types"
)

// PlatformAccount holds the schema definition for the PlatformAccount entity.
type PlatformAccount struct {
	ent.Schema
}

// Fields of the PlatformAccount.
func (PlatformAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").DefaultFunc(types.Id),
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

func (PlatformAccount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
