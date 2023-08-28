// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/leeeeeoy/go_study/ent/commentreport"
	"github.com/leeeeeoy/go_study/ent/predicate"
)

// CommentReportDelete is the builder for deleting a CommentReport entity.
type CommentReportDelete struct {
	config
	hooks    []Hook
	mutation *CommentReportMutation
}

// Where appends a list predicates to the CommentReportDelete builder.
func (crd *CommentReportDelete) Where(ps ...predicate.CommentReport) *CommentReportDelete {
	crd.mutation.Where(ps...)
	return crd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crd *CommentReportDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, crd.sqlExec, crd.mutation, crd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (crd *CommentReportDelete) ExecX(ctx context.Context) int {
	n, err := crd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crd *CommentReportDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(commentreport.Table, sqlgraph.NewFieldSpec(commentreport.FieldID, field.TypeInt))
	if ps := crd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, crd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	crd.mutation.done = true
	return affected, err
}

// CommentReportDeleteOne is the builder for deleting a single CommentReport entity.
type CommentReportDeleteOne struct {
	crd *CommentReportDelete
}

// Where appends a list predicates to the CommentReportDelete builder.
func (crdo *CommentReportDeleteOne) Where(ps ...predicate.CommentReport) *CommentReportDeleteOne {
	crdo.crd.mutation.Where(ps...)
	return crdo
}

// Exec executes the deletion query.
func (crdo *CommentReportDeleteOne) Exec(ctx context.Context) error {
	n, err := crdo.crd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{commentreport.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crdo *CommentReportDeleteOne) ExecX(ctx context.Context) {
	if err := crdo.Exec(ctx); err != nil {
		panic(err)
	}
}