package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BoardReport holds the schema definition for the BoardReport entity.
type BoardReport struct {
	ent.Schema
}

// Fields of the BoardReport.
func (BoardReport) Fields() []ent.Field {
	return []ent.Field{
		field.Int("board_id").Optional(),
		field.Int("report_type_id").Optional(),
		field.Int("reporter_id").Optional(),
		field.String("comment").Optional(),
		field.Enum("status").Values("0", "1").Comment("0 is deleted, 1 is activate"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the BoardReport.
func (BoardReport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("board_report").Unique().Field("reporter_id"),
		edge.From("board", Board.Type).Ref("board_report").Unique().Field("board_id"),
		edge.From("report_type", ReportType.Type).Ref("board_report").Unique().Field("report_type_id"),
	}
}
