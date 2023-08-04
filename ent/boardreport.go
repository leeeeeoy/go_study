// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/leeeeeoy/go_study/ent/board"
	"github.com/leeeeeoy/go_study/ent/boardreport"
	"github.com/leeeeeoy/go_study/ent/reporttype"
	"github.com/leeeeeoy/go_study/ent/user"
)

// BoardReport is the model entity for the BoardReport schema.
type BoardReport struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BoardID holds the value of the "board_id" field.
	BoardID int `json:"board_id,omitempty"`
	// ReportTypeID holds the value of the "report_type_id" field.
	ReportTypeID int `json:"report_type_id,omitempty"`
	// ReporterID holds the value of the "reporter_id" field.
	ReporterID int `json:"reporter_id,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitempty"`
	// Status holds the value of the "status" field.
	Status boardreport.Status `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BoardReportQuery when eager-loading is set.
	Edges        BoardReportEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BoardReportEdges holds the relations/edges for other nodes in the graph.
type BoardReportEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Board holds the value of the board edge.
	Board *Board `json:"board,omitempty"`
	// ReportType holds the value of the report_type edge.
	ReportType *ReportType `json:"report_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BoardReportEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// BoardOrErr returns the Board value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BoardReportEdges) BoardOrErr() (*Board, error) {
	if e.loadedTypes[1] {
		if e.Board == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: board.Label}
		}
		return e.Board, nil
	}
	return nil, &NotLoadedError{edge: "board"}
}

// ReportTypeOrErr returns the ReportType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BoardReportEdges) ReportTypeOrErr() (*ReportType, error) {
	if e.loadedTypes[2] {
		if e.ReportType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: reporttype.Label}
		}
		return e.ReportType, nil
	}
	return nil, &NotLoadedError{edge: "report_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BoardReport) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case boardreport.FieldID, boardreport.FieldBoardID, boardreport.FieldReportTypeID, boardreport.FieldReporterID:
			values[i] = new(sql.NullInt64)
		case boardreport.FieldComment, boardreport.FieldStatus:
			values[i] = new(sql.NullString)
		case boardreport.FieldCreatedAt, boardreport.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BoardReport fields.
func (br *BoardReport) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case boardreport.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			br.ID = int(value.Int64)
		case boardreport.FieldBoardID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field board_id", values[i])
			} else if value.Valid {
				br.BoardID = int(value.Int64)
			}
		case boardreport.FieldReportTypeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field report_type_id", values[i])
			} else if value.Valid {
				br.ReportTypeID = int(value.Int64)
			}
		case boardreport.FieldReporterID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reporter_id", values[i])
			} else if value.Valid {
				br.ReporterID = int(value.Int64)
			}
		case boardreport.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				br.Comment = value.String
			}
		case boardreport.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				br.Status = boardreport.Status(value.String)
			}
		case boardreport.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				br.CreatedAt = value.Time
			}
		case boardreport.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				br.UpdatedAt = value.Time
			}
		default:
			br.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BoardReport.
// This includes values selected through modifiers, order, etc.
func (br *BoardReport) Value(name string) (ent.Value, error) {
	return br.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the BoardReport entity.
func (br *BoardReport) QueryUser() *UserQuery {
	return NewBoardReportClient(br.config).QueryUser(br)
}

// QueryBoard queries the "board" edge of the BoardReport entity.
func (br *BoardReport) QueryBoard() *BoardQuery {
	return NewBoardReportClient(br.config).QueryBoard(br)
}

// QueryReportType queries the "report_type" edge of the BoardReport entity.
func (br *BoardReport) QueryReportType() *ReportTypeQuery {
	return NewBoardReportClient(br.config).QueryReportType(br)
}

// Update returns a builder for updating this BoardReport.
// Note that you need to call BoardReport.Unwrap() before calling this method if this BoardReport
// was returned from a transaction, and the transaction was committed or rolled back.
func (br *BoardReport) Update() *BoardReportUpdateOne {
	return NewBoardReportClient(br.config).UpdateOne(br)
}

// Unwrap unwraps the BoardReport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (br *BoardReport) Unwrap() *BoardReport {
	_tx, ok := br.config.driver.(*txDriver)
	if !ok {
		panic("ent: BoardReport is not a transactional entity")
	}
	br.config.driver = _tx.drv
	return br
}

// String implements the fmt.Stringer.
func (br *BoardReport) String() string {
	var builder strings.Builder
	builder.WriteString("BoardReport(")
	builder.WriteString(fmt.Sprintf("id=%v, ", br.ID))
	builder.WriteString("board_id=")
	builder.WriteString(fmt.Sprintf("%v", br.BoardID))
	builder.WriteString(", ")
	builder.WriteString("report_type_id=")
	builder.WriteString(fmt.Sprintf("%v", br.ReportTypeID))
	builder.WriteString(", ")
	builder.WriteString("reporter_id=")
	builder.WriteString(fmt.Sprintf("%v", br.ReporterID))
	builder.WriteString(", ")
	builder.WriteString("comment=")
	builder.WriteString(br.Comment)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", br.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(br.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(br.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BoardReports is a parsable slice of BoardReport.
type BoardReports []*BoardReport
