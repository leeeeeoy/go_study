package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CommentMention holds the schema definition for the CommentMention entity.
type CommentMention struct {
	ent.Schema
}

// Fields of the CommentMention.
func (CommentMention) Fields() []ent.Field {
	return []ent.Field{
		field.Int("comment_id").Optional(),
		field.Int("user_id").Optional(),
	}
}

// Edges of the CommentMention.
func (CommentMention) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("comment_mention").Unique().Field("user_id"),
		edge.From("comment", Comment.Type).Ref("comment_mention").Unique().Field("comment_id"),
	}
}
