package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Beer holds the schema definition for the Beer entity.
type Beer struct {
	ent.Schema
}

// Fields of the Beer.
func (Beer) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
		field.String("description"),
		field.Int64("count"),
		field.Int64("price"),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the Beer.
func (Beer) Edges() []ent.Edge {
	return []ent.Edge{}
}
