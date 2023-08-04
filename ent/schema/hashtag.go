package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Hashtag holds the schema definition for the Hashtag entity.
type Hashtag struct {
	ent.Schema
}

// Fields of the Hashtag.
func (Hashtag) Fields() []ent.Field {
	return []ent.Field{
		field.String("value"),
		field.Int("used_count").Default(0),
	}
}

// Edges of the Hashtag.
func (Hashtag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("board_hashtag", BoardHashtag.Type),
	}
}
