package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ReportType holds the schema definition for the ReportType entity.
type ReportType struct {
	ent.Schema
}

// Fields of the ReportType.
func (ReportType) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Optional(),
		field.Bool("in_active").Default(false),
		field.Int("order_num"),
	}
}

// Edges of the ReportType.
func (ReportType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comment_report", CommentReport.Type).Unique(),
		edge.To("board_report", BoardReport.Type).Unique(),
	}
}
