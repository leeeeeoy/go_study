package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CommentLike holds the schema definition for the CommentLike entity.
type CommentLike struct {
	ent.Schema
}

// Fields of the CommentLike.
func (CommentLike) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.Int("comment_id").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the CommentLike.
func (CommentLike) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("comment_like").Unique().Field("user_id"),
		edge.From("comment", Comment.Type).Ref("comment_like").Unique().Field("comment_id"),
	}
}
