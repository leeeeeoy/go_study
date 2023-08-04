package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BoardHashtag holds the schema definition for the BoardHashtag entity.
type BoardHashtag struct {
	ent.Schema
}

// Fields of the BoardHashtag.
func (BoardHashtag) Fields() []ent.Field {
	return []ent.Field{
		field.Int("board_id").Optional(),
		field.Int("hashtag_id").Optional(),
	}
}

// Edges of the BoardHashtag.
func (BoardHashtag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("board", Board.Type).Ref("board_hashtag").Unique().Field("board_id"),
		edge.From("hashtag", Hashtag.Type).Ref("board_hashtag").Unique().Field("hashtag_id"),
	}
}
