// Code generated by ent, DO NOT EDIT.

package boardlike

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/leeeeeoy/go_study/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldEQ(FieldUserID, v))
}

// BoardID applies equality check predicate on the "board_id" field. It's identical to BoardIDEQ.
func BoardID(v int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldEQ(FieldBoardID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldEQ(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.BoardLike {
	return predicate.BoardLike(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.BoardLike {
	return predicate.BoardLike(sql.FieldNotNull(FieldUserID))
}

// BoardIDEQ applies the EQ predicate on the "board_id" field.
func BoardIDEQ(v int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldEQ(FieldBoardID, v))
}

// BoardIDNEQ applies the NEQ predicate on the "board_id" field.
func BoardIDNEQ(v int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldNEQ(FieldBoardID, v))
}

// BoardIDIn applies the In predicate on the "board_id" field.
func BoardIDIn(vs ...int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldIn(FieldBoardID, vs...))
}

// BoardIDNotIn applies the NotIn predicate on the "board_id" field.
func BoardIDNotIn(vs ...int) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldNotIn(FieldBoardID, vs...))
}

// BoardIDIsNil applies the IsNil predicate on the "board_id" field.
func BoardIDIsNil() predicate.BoardLike {
	return predicate.BoardLike(sql.FieldIsNull(FieldBoardID))
}

// BoardIDNotNil applies the NotNil predicate on the "board_id" field.
func BoardIDNotNil() predicate.BoardLike {
	return predicate.BoardLike(sql.FieldNotNull(FieldBoardID))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BoardLike {
	return predicate.BoardLike(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.BoardLike {
	return predicate.BoardLike(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.BoardLike {
	return predicate.BoardLike(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBoard applies the HasEdge predicate on the "board" edge.
func HasBoard() predicate.BoardLike {
	return predicate.BoardLike(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BoardTable, BoardColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBoardWith applies the HasEdge predicate on the "board" edge with a given conditions (other predicates).
func HasBoardWith(preds ...predicate.Board) predicate.BoardLike {
	return predicate.BoardLike(func(s *sql.Selector) {
		step := newBoardStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BoardLike) predicate.BoardLike {
	return predicate.BoardLike(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BoardLike) predicate.BoardLike {
	return predicate.BoardLike(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BoardLike) predicate.BoardLike {
	return predicate.BoardLike(func(s *sql.Selector) {
		p(s.Not())
	})
}
