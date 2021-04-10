package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Machine holds the schema definition for the Machine entity.
type Machine struct {
	ent.Schema
}

// Fields of the Machine.
func (Machine) Fields() []ent.Field {
	return []ent.Field{
		field.String("hwid").NotEmpty().
			Annotations(entproto.Field(2)),
		field.String("hostname").NotEmpty().
			Annotations(entproto.Field(3)),
		field.String("fingerprint").NotEmpty().
			Annotations(entproto.Field(4)),
	}
}

// Edges of the Machine.
func (Machine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ips", IP.Type).Ref("machine").Annotations(entproto.Field(5)),
		edge.From("users", User.Type).Ref("machines").Annotations(entproto.Field(6)),
		edge.From("keys", Key.Type).Ref("machine").Annotations(entproto.Field(7)),
	}
}

func (Machine) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Service(),
		entproto.Message(),
	}
}
