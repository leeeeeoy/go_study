package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BookMark holds the schema definition for the BookMark entity.
type BookMark struct {
	ent.Schema
}

// Fields of the BookMark.
func (BookMark) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.Int("board_id").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the BookMark.
func (BookMark) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("board", Board.Type).Unique().Ref("book_marks").Field("board_id"),
		edge.From("user", User.Type).Unique().Ref("book_marks").Field("user_id"),
	}
}
