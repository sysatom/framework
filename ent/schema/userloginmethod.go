package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/sysatom/framework/pkg/types"
)

// UserLoginMethod holds the schema definition for the UserLoginMethod entity.
type UserLoginMethod struct {
	ent.Schema
}

// Fields of the UserLoginMethod.
func (UserLoginMethod) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").DefaultFunc(types.Id),
		//field.Int64("user_id"),
		field.String("login_type"), // password, wechat
		field.String("identifier"),
	}
}

// Edges of the UserLoginMethod.
func (UserLoginMethod) Edges() []ent.Edge {
	return nil
}
