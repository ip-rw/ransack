package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IP holds the schema definition for the IP entity.
type IP struct {
	ent.Schema
}

// Fields of the IP.
func (IP) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip").Unique().
			Annotations(entproto.Field(2)),
		field.String("net").
			Annotations(entproto.Field(3)),
		field.Int("mask").Default(24).
			Annotations(entproto.Field(4)),
	}
}

// Edges of the IP.
func (IP) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("machine", Machine.Type).Unique().
			Annotations(entproto.Field(5)),
	}
}

func (IP) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Service(),
		entproto.Message(),
	}
}
