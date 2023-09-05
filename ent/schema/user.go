package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").MaxLen(20),
		field.String("password").Sensitive().MinLen(8),
		field.String("name").MaxLen(10),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("boards", Board.Type),
		edge.To("board_like", BoardLike.Type),
		edge.To("book_marks", BookMark.Type),
		edge.To("comment_like", CommentLike.Type),
		edge.To("comments", Comment.Type),
		edge.To("comment_mention", CommentMention.Type),
		edge.To("board_report", BoardReport.Type),
		edge.To("comment_report", CommentReport.Type),
	}
}
