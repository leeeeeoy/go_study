// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/leeeeeoy/go_study/ent/board"
	"github.com/leeeeeoy/go_study/ent/bookmark"
	"github.com/leeeeeoy/go_study/ent/user"
)

// BookMark is the model entity for the BookMark schema.
type BookMark struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// BoardID holds the value of the "board_id" field.
	BoardID int `json:"board_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookMarkQuery when eager-loading is set.
	Edges        BookMarkEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BookMarkEdges holds the relations/edges for other nodes in the graph.
type BookMarkEdges struct {
	// Board holds the value of the board edge.
	Board *Board `json:"board,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BoardOrErr returns the Board value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookMarkEdges) BoardOrErr() (*Board, error) {
	if e.loadedTypes[0] {
		if e.Board == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: board.Label}
		}
		return e.Board, nil
	}
	return nil, &NotLoadedError{edge: "board"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookMarkEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BookMark) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bookmark.FieldID, bookmark.FieldUserID, bookmark.FieldBoardID:
			values[i] = new(sql.NullInt64)
		case bookmark.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BookMark fields.
func (bm *BookMark) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bookmark.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bm.ID = int(value.Int64)
		case bookmark.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				bm.UserID = int(value.Int64)
			}
		case bookmark.FieldBoardID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field board_id", values[i])
			} else if value.Valid {
				bm.BoardID = int(value.Int64)
			}
		case bookmark.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bm.CreatedAt = value.Time
			}
		default:
			bm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BookMark.
// This includes values selected through modifiers, order, etc.
func (bm *BookMark) Value(name string) (ent.Value, error) {
	return bm.selectValues.Get(name)
}

// QueryBoard queries the "board" edge of the BookMark entity.
func (bm *BookMark) QueryBoard() *BoardQuery {
	return NewBookMarkClient(bm.config).QueryBoard(bm)
}

// QueryUser queries the "user" edge of the BookMark entity.
func (bm *BookMark) QueryUser() *UserQuery {
	return NewBookMarkClient(bm.config).QueryUser(bm)
}

// Update returns a builder for updating this BookMark.
// Note that you need to call BookMark.Unwrap() before calling this method if this BookMark
// was returned from a transaction, and the transaction was committed or rolled back.
func (bm *BookMark) Update() *BookMarkUpdateOne {
	return NewBookMarkClient(bm.config).UpdateOne(bm)
}

// Unwrap unwraps the BookMark entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bm *BookMark) Unwrap() *BookMark {
	_tx, ok := bm.config.driver.(*txDriver)
	if !ok {
		panic("ent: BookMark is not a transactional entity")
	}
	bm.config.driver = _tx.drv
	return bm
}

// String implements the fmt.Stringer.
func (bm *BookMark) String() string {
	var builder strings.Builder
	builder.WriteString("BookMark(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bm.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", bm.UserID))
	builder.WriteString(", ")
	builder.WriteString("board_id=")
	builder.WriteString(fmt.Sprintf("%v", bm.BoardID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(bm.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BookMarks is a parsable slice of BookMark.
type BookMarks []*BookMark