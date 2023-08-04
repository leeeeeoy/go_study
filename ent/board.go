// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/leeeeeoy/go_study/ent/board"
	"github.com/leeeeeoy/go_study/ent/user"
)

// Board is the model entity for the Board schema.
type Board struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// LikeCount holds the value of the "like_count" field.
	LikeCount int `json:"like_count,omitempty"`
	// CommentCount holds the value of the "comment_count" field.
	CommentCount int `json:"comment_count,omitempty"`
	// ViewCount holds the value of the "view_count" field.
	ViewCount int `json:"view_count,omitempty"`
	// ReportCount holds the value of the "report_count" field.
	ReportCount int `json:"report_count,omitempty"`
	// Status holds the value of the "status" field.
	Status board.Status `json:"status,omitempty"`
	// LanguageType holds the value of the "language_type" field.
	LanguageType string `json:"language_type,omitempty"`
	// Attachments holds the value of the "attachments" field.
	Attachments string `json:"attachments,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BoardQuery when eager-loading is set.
	Edges        BoardEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BoardEdges holds the relations/edges for other nodes in the graph.
type BoardEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// BoardLike holds the value of the board_like edge.
	BoardLike []*BoardLike `json:"board_like,omitempty"`
	// BoardHashtag holds the value of the board_hashtag edge.
	BoardHashtag []*BoardHashtag `json:"board_hashtag,omitempty"`
	// BoardReport holds the value of the board_report edge.
	BoardReport []*BoardReport `json:"board_report,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BoardEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e BoardEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[1] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// BoardLikeOrErr returns the BoardLike value or an error if the edge
// was not loaded in eager-loading.
func (e BoardEdges) BoardLikeOrErr() ([]*BoardLike, error) {
	if e.loadedTypes[2] {
		return e.BoardLike, nil
	}
	return nil, &NotLoadedError{edge: "board_like"}
}

// BoardHashtagOrErr returns the BoardHashtag value or an error if the edge
// was not loaded in eager-loading.
func (e BoardEdges) BoardHashtagOrErr() ([]*BoardHashtag, error) {
	if e.loadedTypes[3] {
		return e.BoardHashtag, nil
	}
	return nil, &NotLoadedError{edge: "board_hashtag"}
}

// BoardReportOrErr returns the BoardReport value or an error if the edge
// was not loaded in eager-loading.
func (e BoardEdges) BoardReportOrErr() ([]*BoardReport, error) {
	if e.loadedTypes[4] {
		return e.BoardReport, nil
	}
	return nil, &NotLoadedError{edge: "board_report"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Board) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case board.FieldID, board.FieldUserID, board.FieldLikeCount, board.FieldCommentCount, board.FieldViewCount, board.FieldReportCount:
			values[i] = new(sql.NullInt64)
		case board.FieldTitle, board.FieldText, board.FieldStatus, board.FieldLanguageType, board.FieldAttachments:
			values[i] = new(sql.NullString)
		case board.FieldCreatedAt, board.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Board fields.
func (b *Board) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case board.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case board.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				b.Title = value.String
			}
		case board.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				b.Text = value.String
			}
		case board.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				b.UserID = int(value.Int64)
			}
		case board.FieldLikeCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field like_count", values[i])
			} else if value.Valid {
				b.LikeCount = int(value.Int64)
			}
		case board.FieldCommentCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field comment_count", values[i])
			} else if value.Valid {
				b.CommentCount = int(value.Int64)
			}
		case board.FieldViewCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field view_count", values[i])
			} else if value.Valid {
				b.ViewCount = int(value.Int64)
			}
		case board.FieldReportCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field report_count", values[i])
			} else if value.Valid {
				b.ReportCount = int(value.Int64)
			}
		case board.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = board.Status(value.String)
			}
		case board.FieldLanguageType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language_type", values[i])
			} else if value.Valid {
				b.LanguageType = value.String
			}
		case board.FieldAttachments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field attachments", values[i])
			} else if value.Valid {
				b.Attachments = value.String
			}
		case board.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case board.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Board.
// This includes values selected through modifiers, order, etc.
func (b *Board) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Board entity.
func (b *Board) QueryUser() *UserQuery {
	return NewBoardClient(b.config).QueryUser(b)
}

// QueryComments queries the "comments" edge of the Board entity.
func (b *Board) QueryComments() *CommentQuery {
	return NewBoardClient(b.config).QueryComments(b)
}

// QueryBoardLike queries the "board_like" edge of the Board entity.
func (b *Board) QueryBoardLike() *BoardLikeQuery {
	return NewBoardClient(b.config).QueryBoardLike(b)
}

// QueryBoardHashtag queries the "board_hashtag" edge of the Board entity.
func (b *Board) QueryBoardHashtag() *BoardHashtagQuery {
	return NewBoardClient(b.config).QueryBoardHashtag(b)
}

// QueryBoardReport queries the "board_report" edge of the Board entity.
func (b *Board) QueryBoardReport() *BoardReportQuery {
	return NewBoardClient(b.config).QueryBoardReport(b)
}

// Update returns a builder for updating this Board.
// Note that you need to call Board.Unwrap() before calling this method if this Board
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Board) Update() *BoardUpdateOne {
	return NewBoardClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Board entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Board) Unwrap() *Board {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Board is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Board) String() string {
	var builder strings.Builder
	builder.WriteString("Board(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("title=")
	builder.WriteString(b.Title)
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(b.Text)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", b.UserID))
	builder.WriteString(", ")
	builder.WriteString("like_count=")
	builder.WriteString(fmt.Sprintf("%v", b.LikeCount))
	builder.WriteString(", ")
	builder.WriteString("comment_count=")
	builder.WriteString(fmt.Sprintf("%v", b.CommentCount))
	builder.WriteString(", ")
	builder.WriteString("view_count=")
	builder.WriteString(fmt.Sprintf("%v", b.ViewCount))
	builder.WriteString(", ")
	builder.WriteString("report_count=")
	builder.WriteString(fmt.Sprintf("%v", b.ReportCount))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", ")
	builder.WriteString("language_type=")
	builder.WriteString(b.LanguageType)
	builder.WriteString(", ")
	builder.WriteString("attachments=")
	builder.WriteString(b.Attachments)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Boards is a parsable slice of Board.
type Boards []*Board
