package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Key holds the schema definition for the Key entity.
type Key struct {
	ent.Schema
}

// Fields of the Key.
func (Key) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("keyfile").
			Annotations(entproto.Field(2)),
		//field.String("username").Optional(),
		field.String("password").Optional().
			Annotations(entproto.Field(3)),
	}
}

// Edges of the Key.
func (Key) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("machine", Machine.Type).Unique().
			Annotations(entproto.Field(4)),
		edge.To("user", User.Type).Unique().
			Annotations(entproto.Field(5)),
	}
}

func (Key) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Service(),
		entproto.Message(),
	}
}
