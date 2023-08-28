package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Topic holds the schema definition for the Topic entity.
type Topic struct {
	ent.Schema
}

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.Int("category_id").Optional(),
		field.String("name"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("boards", Board.Type),
		edge.From("category", Category.Type).Unique().Ref("topics").Field("category_id"),
	}
}
