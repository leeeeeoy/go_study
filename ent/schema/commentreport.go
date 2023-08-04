package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CommentReport holds the schema definition for the CommentReport entity.
type CommentReport struct {
	ent.Schema
}

// Fields of the CommentReport.
func (CommentReport) Fields() []ent.Field {
	return []ent.Field{
		field.Int("comment_id").Optional(),
		field.Int("report_type_id").Optional(),
		field.Int("reporter_id").Optional(),
		field.String("desc").Optional(),
		field.Enum("status").Values("activate", "deleted"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CommentReport.
func (CommentReport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("comment_report").Unique().Field("reporter_id"),
		edge.From("comment", Comment.Type).Ref("comment_report").Unique().Field("comment_id"),
		edge.From("report_type", ReportType.Type).Ref("comment_report").Unique().Field("report_type_id"),
	}
}
