package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BoardLike holds the schema definition for the BoardLike entity.
type BoardLike struct {
	ent.Schema
}

// Fields of the BoardLike.
func (BoardLike) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.Int("board_id").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the BoardLike.
func (BoardLike) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("board_like").Unique().Field("user_id"),
		edge.From("board", Board.Type).Ref("board_like").Unique().Field("board_id"),
	}
}
