package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("username").
			Annotations(entproto.Field(2)),
		field.String("password").
			Annotations(entproto.Field(3)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("machines", Machine.Type).
			Annotations(entproto.Field(4)),
		edge.From("keys", Key.Type).Ref("user").
			Annotations(entproto.Field(5)),
	}
}
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Service(),
		entproto.Message(),
	}
}
