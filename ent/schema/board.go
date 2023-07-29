package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Board holds the schema definition for the Board entity.
type Board struct {
	ent.Schema
}

// Fields of the Board.
func (Board) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("content"),
		field.Int("user_id").Optional(),
		field.Int("like_count"),
		field.Int("comment_count"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Board.
func (Board) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Unique().Ref("boards").Field("user_id"),
		edge.To("comments", Comment.Type),
		edge.To("board_like", BoardLike.Type),
	}
}
