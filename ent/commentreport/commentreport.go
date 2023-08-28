// Code generated by ent, DO NOT EDIT.

package commentreport

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the commentreport type in the database.
	Label = "comment_report"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCommentID holds the string denoting the comment_id field in the database.
	FieldCommentID = "comment_id"
	// FieldReportTypeID holds the string denoting the report_type_id field in the database.
	FieldReportTypeID = "report_type_id"
	// FieldReporterID holds the string denoting the reporter_id field in the database.
	FieldReporterID = "reporter_id"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeComment holds the string denoting the comment edge name in mutations.
	EdgeComment = "comment"
	// EdgeReportType holds the string denoting the report_type edge name in mutations.
	EdgeReportType = "report_type"
	// Table holds the table name of the commentreport in the database.
	Table = "comment_reports"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "comment_reports"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "reporter_id"
	// CommentTable is the table that holds the comment relation/edge.
	CommentTable = "comment_reports"
	// CommentInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentInverseTable = "comments"
	// CommentColumn is the table column denoting the comment relation/edge.
	CommentColumn = "comment_id"
	// ReportTypeTable is the table that holds the report_type relation/edge.
	ReportTypeTable = "comment_reports"
	// ReportTypeInverseTable is the table name for the ReportType entity.
	// It exists in this package in order to avoid circular dependency with the "reporttype" package.
	ReportTypeInverseTable = "report_types"
	// ReportTypeColumn is the table column denoting the report_type relation/edge.
	ReportTypeColumn = "report_type_id"
)

// Columns holds all SQL columns for commentreport fields.
var Columns = []string{
	FieldID,
	FieldCommentID,
	FieldReportTypeID,
	FieldReporterID,
	FieldDesc,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	Status0 Status = "0"
	Status1 Status = "1"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case Status0, Status1:
		return nil
	default:
		return fmt.Errorf("commentreport: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the CommentReport queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCommentID orders the results by the comment_id field.
func ByCommentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommentID, opts...).ToFunc()
}

// ByReportTypeID orders the results by the report_type_id field.
func ByReportTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReportTypeID, opts...).ToFunc()
}

// ByReporterID orders the results by the reporter_id field.
func ByReporterID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReporterID, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByCommentField orders the results by comment field.
func ByCommentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentStep(), sql.OrderByField(field, opts...))
	}
}

// ByReportTypeField orders the results by report_type field.
func ByReportTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReportTypeStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newCommentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CommentTable, CommentColumn),
	)
}
func newReportTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReportTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ReportTypeTable, ReportTypeColumn),
	)
}