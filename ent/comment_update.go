// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/leeeeeoy/go_study/ent/board"
	"github.com/leeeeeoy/go_study/ent/comment"
	"github.com/leeeeeoy/go_study/ent/commentlike"
	"github.com/leeeeeoy/go_study/ent/commentmention"
	"github.com/leeeeeoy/go_study/ent/commentreport"
	"github.com/leeeeeoy/go_study/ent/predicate"
	"github.com/leeeeeoy/go_study/ent/user"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetText sets the "text" field.
func (cu *CommentUpdate) SetText(s string) *CommentUpdate {
	cu.mutation.SetText(s)
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CommentUpdate) SetUserID(i int) *CommentUpdate {
	cu.mutation.SetUserID(i)
	return cu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableUserID(i *int) *CommentUpdate {
	if i != nil {
		cu.SetUserID(*i)
	}
	return cu
}

// ClearUserID clears the value of the "user_id" field.
func (cu *CommentUpdate) ClearUserID() *CommentUpdate {
	cu.mutation.ClearUserID()
	return cu
}

// SetBoardID sets the "board_id" field.
func (cu *CommentUpdate) SetBoardID(i int) *CommentUpdate {
	cu.mutation.SetBoardID(i)
	return cu
}

// SetNillableBoardID sets the "board_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableBoardID(i *int) *CommentUpdate {
	if i != nil {
		cu.SetBoardID(*i)
	}
	return cu
}

// ClearBoardID clears the value of the "board_id" field.
func (cu *CommentUpdate) ClearBoardID() *CommentUpdate {
	cu.mutation.ClearBoardID()
	return cu
}

// SetLikeCount sets the "like_count" field.
func (cu *CommentUpdate) SetLikeCount(i int) *CommentUpdate {
	cu.mutation.ResetLikeCount()
	cu.mutation.SetLikeCount(i)
	return cu
}

// AddLikeCount adds i to the "like_count" field.
func (cu *CommentUpdate) AddLikeCount(i int) *CommentUpdate {
	cu.mutation.AddLikeCount(i)
	return cu
}

// SetStatus sets the "status" field.
func (cu *CommentUpdate) SetStatus(c comment.Status) *CommentUpdate {
	cu.mutation.SetStatus(c)
	return cu
}

// SetReportCount sets the "report_count" field.
func (cu *CommentUpdate) SetReportCount(i int) *CommentUpdate {
	cu.mutation.ResetReportCount()
	cu.mutation.SetReportCount(i)
	return cu
}

// AddReportCount adds i to the "report_count" field.
func (cu *CommentUpdate) AddReportCount(i int) *CommentUpdate {
	cu.mutation.AddReportCount(i)
	return cu
}

// SetLanguageType sets the "language_type" field.
func (cu *CommentUpdate) SetLanguageType(s string) *CommentUpdate {
	cu.mutation.SetLanguageType(s)
	return cu
}

// SetAuthorHeart sets the "author_heart" field.
func (cu *CommentUpdate) SetAuthorHeart(b bool) *CommentUpdate {
	cu.mutation.SetAuthorHeart(b)
	return cu
}

// SetNillableAuthorHeart sets the "author_heart" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableAuthorHeart(b *bool) *CommentUpdate {
	if b != nil {
		cu.SetAuthorHeart(*b)
	}
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CommentUpdate) SetCreatedAt(t time.Time) *CommentUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableCreatedAt(t *time.Time) *CommentUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommentUpdate) SetUpdatedAt(t time.Time) *CommentUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetBoard sets the "board" edge to the Board entity.
func (cu *CommentUpdate) SetBoard(b *Board) *CommentUpdate {
	return cu.SetBoardID(b.ID)
}

// SetUser sets the "user" edge to the User entity.
func (cu *CommentUpdate) SetUser(u *User) *CommentUpdate {
	return cu.SetUserID(u.ID)
}

// AddCommentLikeIDs adds the "comment_like" edge to the CommentLike entity by IDs.
func (cu *CommentUpdate) AddCommentLikeIDs(ids ...int) *CommentUpdate {
	cu.mutation.AddCommentLikeIDs(ids...)
	return cu
}

// AddCommentLike adds the "comment_like" edges to the CommentLike entity.
func (cu *CommentUpdate) AddCommentLike(c ...*CommentLike) *CommentUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCommentLikeIDs(ids...)
}

// AddCommentMentionIDs adds the "comment_mention" edge to the CommentMention entity by IDs.
func (cu *CommentUpdate) AddCommentMentionIDs(ids ...int) *CommentUpdate {
	cu.mutation.AddCommentMentionIDs(ids...)
	return cu
}

// AddCommentMention adds the "comment_mention" edges to the CommentMention entity.
func (cu *CommentUpdate) AddCommentMention(c ...*CommentMention) *CommentUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCommentMentionIDs(ids...)
}

// AddCommentReportIDs adds the "comment_report" edge to the CommentReport entity by IDs.
func (cu *CommentUpdate) AddCommentReportIDs(ids ...int) *CommentUpdate {
	cu.mutation.AddCommentReportIDs(ids...)
	return cu
}

// AddCommentReport adds the "comment_report" edges to the CommentReport entity.
func (cu *CommentUpdate) AddCommentReport(c ...*CommentReport) *CommentUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCommentReportIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// ClearBoard clears the "board" edge to the Board entity.
func (cu *CommentUpdate) ClearBoard() *CommentUpdate {
	cu.mutation.ClearBoard()
	return cu
}

// ClearUser clears the "user" edge to the User entity.
func (cu *CommentUpdate) ClearUser() *CommentUpdate {
	cu.mutation.ClearUser()
	return cu
}

// ClearCommentLike clears all "comment_like" edges to the CommentLike entity.
func (cu *CommentUpdate) ClearCommentLike() *CommentUpdate {
	cu.mutation.ClearCommentLike()
	return cu
}

// RemoveCommentLikeIDs removes the "comment_like" edge to CommentLike entities by IDs.
func (cu *CommentUpdate) RemoveCommentLikeIDs(ids ...int) *CommentUpdate {
	cu.mutation.RemoveCommentLikeIDs(ids...)
	return cu
}

// RemoveCommentLike removes "comment_like" edges to CommentLike entities.
func (cu *CommentUpdate) RemoveCommentLike(c ...*CommentLike) *CommentUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCommentLikeIDs(ids...)
}

// ClearCommentMention clears all "comment_mention" edges to the CommentMention entity.
func (cu *CommentUpdate) ClearCommentMention() *CommentUpdate {
	cu.mutation.ClearCommentMention()
	return cu
}

// RemoveCommentMentionIDs removes the "comment_mention" edge to CommentMention entities by IDs.
func (cu *CommentUpdate) RemoveCommentMentionIDs(ids ...int) *CommentUpdate {
	cu.mutation.RemoveCommentMentionIDs(ids...)
	return cu
}

// RemoveCommentMention removes "comment_mention" edges to CommentMention entities.
func (cu *CommentUpdate) RemoveCommentMention(c ...*CommentMention) *CommentUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCommentMentionIDs(ids...)
}

// ClearCommentReport clears all "comment_report" edges to the CommentReport entity.
func (cu *CommentUpdate) ClearCommentReport() *CommentUpdate {
	cu.mutation.ClearCommentReport()
	return cu
}

// RemoveCommentReportIDs removes the "comment_report" edge to CommentReport entities by IDs.
func (cu *CommentUpdate) RemoveCommentReportIDs(ids ...int) *CommentUpdate {
	cu.mutation.RemoveCommentReportIDs(ids...)
	return cu
}

// RemoveCommentReport removes "comment_report" edges to CommentReport entities.
func (cu *CommentUpdate) RemoveCommentReport(c ...*CommentReport) *CommentUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCommentReportIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CommentUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := comment.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CommentUpdate) check() error {
	if v, ok := cu.mutation.Status(); ok {
		if err := comment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Comment.status": %w`, err)}
		}
	}
	if v, ok := cu.mutation.ReportCount(); ok {
		if err := comment.ReportCountValidator(v); err != nil {
			return &ValidationError{Name: "report_count", err: fmt.Errorf(`ent: validator failed for field "Comment.report_count": %w`, err)}
		}
	}
	return nil
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Text(); ok {
		_spec.SetField(comment.FieldText, field.TypeString, value)
	}
	if value, ok := cu.mutation.LikeCount(); ok {
		_spec.SetField(comment.FieldLikeCount, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedLikeCount(); ok {
		_spec.AddField(comment.FieldLikeCount, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(comment.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.ReportCount(); ok {
		_spec.SetField(comment.FieldReportCount, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedReportCount(); ok {
		_spec.AddField(comment.FieldReportCount, field.TypeInt, value)
	}
	if value, ok := cu.mutation.LanguageType(); ok {
		_spec.SetField(comment.FieldLanguageType, field.TypeString, value)
	}
	if value, ok := cu.mutation.AuthorHeart(); ok {
		_spec.SetField(comment.FieldAuthorHeart, field.TypeBool, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.BoardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.BoardTable,
			Columns: []string{comment.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.BoardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.BoardTable,
			Columns: []string{comment.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CommentLikeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentLikeTable,
			Columns: []string{comment.CommentLikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentlike.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCommentLikeIDs(); len(nodes) > 0 && !cu.mutation.CommentLikeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentLikeTable,
			Columns: []string{comment.CommentLikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CommentLikeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentLikeTable,
			Columns: []string{comment.CommentLikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CommentMentionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentMentionTable,
			Columns: []string{comment.CommentMentionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentmention.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCommentMentionIDs(); len(nodes) > 0 && !cu.mutation.CommentMentionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentMentionTable,
			Columns: []string{comment.CommentMentionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentmention.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CommentMentionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentMentionTable,
			Columns: []string{comment.CommentMentionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentmention.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CommentReportCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentReportTable,
			Columns: []string{comment.CommentReportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentreport.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCommentReportIDs(); len(nodes) > 0 && !cu.mutation.CommentReportCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentReportTable,
			Columns: []string{comment.CommentReportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentreport.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CommentReportIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentReportTable,
			Columns: []string{comment.CommentReportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentreport.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetText sets the "text" field.
func (cuo *CommentUpdateOne) SetText(s string) *CommentUpdateOne {
	cuo.mutation.SetText(s)
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *CommentUpdateOne) SetUserID(i int) *CommentUpdateOne {
	cuo.mutation.SetUserID(i)
	return cuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableUserID(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetUserID(*i)
	}
	return cuo
}

// ClearUserID clears the value of the "user_id" field.
func (cuo *CommentUpdateOne) ClearUserID() *CommentUpdateOne {
	cuo.mutation.ClearUserID()
	return cuo
}

// SetBoardID sets the "board_id" field.
func (cuo *CommentUpdateOne) SetBoardID(i int) *CommentUpdateOne {
	cuo.mutation.SetBoardID(i)
	return cuo
}

// SetNillableBoardID sets the "board_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableBoardID(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetBoardID(*i)
	}
	return cuo
}

// ClearBoardID clears the value of the "board_id" field.
func (cuo *CommentUpdateOne) ClearBoardID() *CommentUpdateOne {
	cuo.mutation.ClearBoardID()
	return cuo
}

// SetLikeCount sets the "like_count" field.
func (cuo *CommentUpdateOne) SetLikeCount(i int) *CommentUpdateOne {
	cuo.mutation.ResetLikeCount()
	cuo.mutation.SetLikeCount(i)
	return cuo
}

// AddLikeCount adds i to the "like_count" field.
func (cuo *CommentUpdateOne) AddLikeCount(i int) *CommentUpdateOne {
	cuo.mutation.AddLikeCount(i)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CommentUpdateOne) SetStatus(c comment.Status) *CommentUpdateOne {
	cuo.mutation.SetStatus(c)
	return cuo
}

// SetReportCount sets the "report_count" field.
func (cuo *CommentUpdateOne) SetReportCount(i int) *CommentUpdateOne {
	cuo.mutation.ResetReportCount()
	cuo.mutation.SetReportCount(i)
	return cuo
}

// AddReportCount adds i to the "report_count" field.
func (cuo *CommentUpdateOne) AddReportCount(i int) *CommentUpdateOne {
	cuo.mutation.AddReportCount(i)
	return cuo
}

// SetLanguageType sets the "language_type" field.
func (cuo *CommentUpdateOne) SetLanguageType(s string) *CommentUpdateOne {
	cuo.mutation.SetLanguageType(s)
	return cuo
}

// SetAuthorHeart sets the "author_heart" field.
func (cuo *CommentUpdateOne) SetAuthorHeart(b bool) *CommentUpdateOne {
	cuo.mutation.SetAuthorHeart(b)
	return cuo
}

// SetNillableAuthorHeart sets the "author_heart" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableAuthorHeart(b *bool) *CommentUpdateOne {
	if b != nil {
		cuo.SetAuthorHeart(*b)
	}
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CommentUpdateOne) SetCreatedAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableCreatedAt(t *time.Time) *CommentUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommentUpdateOne) SetUpdatedAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetBoard sets the "board" edge to the Board entity.
func (cuo *CommentUpdateOne) SetBoard(b *Board) *CommentUpdateOne {
	return cuo.SetBoardID(b.ID)
}

// SetUser sets the "user" edge to the User entity.
func (cuo *CommentUpdateOne) SetUser(u *User) *CommentUpdateOne {
	return cuo.SetUserID(u.ID)
}

// AddCommentLikeIDs adds the "comment_like" edge to the CommentLike entity by IDs.
func (cuo *CommentUpdateOne) AddCommentLikeIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.AddCommentLikeIDs(ids...)
	return cuo
}

// AddCommentLike adds the "comment_like" edges to the CommentLike entity.
func (cuo *CommentUpdateOne) AddCommentLike(c ...*CommentLike) *CommentUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCommentLikeIDs(ids...)
}

// AddCommentMentionIDs adds the "comment_mention" edge to the CommentMention entity by IDs.
func (cuo *CommentUpdateOne) AddCommentMentionIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.AddCommentMentionIDs(ids...)
	return cuo
}

// AddCommentMention adds the "comment_mention" edges to the CommentMention entity.
func (cuo *CommentUpdateOne) AddCommentMention(c ...*CommentMention) *CommentUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCommentMentionIDs(ids...)
}

// AddCommentReportIDs adds the "comment_report" edge to the CommentReport entity by IDs.
func (cuo *CommentUpdateOne) AddCommentReportIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.AddCommentReportIDs(ids...)
	return cuo
}

// AddCommentReport adds the "comment_report" edges to the CommentReport entity.
func (cuo *CommentUpdateOne) AddCommentReport(c ...*CommentReport) *CommentUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCommentReportIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// ClearBoard clears the "board" edge to the Board entity.
func (cuo *CommentUpdateOne) ClearBoard() *CommentUpdateOne {
	cuo.mutation.ClearBoard()
	return cuo
}

// ClearUser clears the "user" edge to the User entity.
func (cuo *CommentUpdateOne) ClearUser() *CommentUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// ClearCommentLike clears all "comment_like" edges to the CommentLike entity.
func (cuo *CommentUpdateOne) ClearCommentLike() *CommentUpdateOne {
	cuo.mutation.ClearCommentLike()
	return cuo
}

// RemoveCommentLikeIDs removes the "comment_like" edge to CommentLike entities by IDs.
func (cuo *CommentUpdateOne) RemoveCommentLikeIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.RemoveCommentLikeIDs(ids...)
	return cuo
}

// RemoveCommentLike removes "comment_like" edges to CommentLike entities.
func (cuo *CommentUpdateOne) RemoveCommentLike(c ...*CommentLike) *CommentUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCommentLikeIDs(ids...)
}

// ClearCommentMention clears all "comment_mention" edges to the CommentMention entity.
func (cuo *CommentUpdateOne) ClearCommentMention() *CommentUpdateOne {
	cuo.mutation.ClearCommentMention()
	return cuo
}

// RemoveCommentMentionIDs removes the "comment_mention" edge to CommentMention entities by IDs.
func (cuo *CommentUpdateOne) RemoveCommentMentionIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.RemoveCommentMentionIDs(ids...)
	return cuo
}

// RemoveCommentMention removes "comment_mention" edges to CommentMention entities.
func (cuo *CommentUpdateOne) RemoveCommentMention(c ...*CommentMention) *CommentUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCommentMentionIDs(ids...)
}

// ClearCommentReport clears all "comment_report" edges to the CommentReport entity.
func (cuo *CommentUpdateOne) ClearCommentReport() *CommentUpdateOne {
	cuo.mutation.ClearCommentReport()
	return cuo
}

// RemoveCommentReportIDs removes the "comment_report" edge to CommentReport entities by IDs.
func (cuo *CommentUpdateOne) RemoveCommentReportIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.RemoveCommentReportIDs(ids...)
	return cuo
}

// RemoveCommentReport removes "comment_report" edges to CommentReport entities.
func (cuo *CommentUpdateOne) RemoveCommentReport(c ...*CommentReport) *CommentUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCommentReportIDs(ids...)
}

// Where appends a list predicates to the CommentUpdate builder.
func (cuo *CommentUpdateOne) Where(ps ...predicate.Comment) *CommentUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CommentUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := comment.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CommentUpdateOne) check() error {
	if v, ok := cuo.mutation.Status(); ok {
		if err := comment.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Comment.status": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.ReportCount(); ok {
		if err := comment.ReportCountValidator(v); err != nil {
			return &ValidationError{Name: "report_count", err: fmt.Errorf(`ent: validator failed for field "Comment.report_count": %w`, err)}
		}
	}
	return nil
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Text(); ok {
		_spec.SetField(comment.FieldText, field.TypeString, value)
	}
	if value, ok := cuo.mutation.LikeCount(); ok {
		_spec.SetField(comment.FieldLikeCount, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedLikeCount(); ok {
		_spec.AddField(comment.FieldLikeCount, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(comment.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.ReportCount(); ok {
		_spec.SetField(comment.FieldReportCount, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedReportCount(); ok {
		_spec.AddField(comment.FieldReportCount, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.LanguageType(); ok {
		_spec.SetField(comment.FieldLanguageType, field.TypeString, value)
	}
	if value, ok := cuo.mutation.AuthorHeart(); ok {
		_spec.SetField(comment.FieldAuthorHeart, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.BoardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.BoardTable,
			Columns: []string{comment.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.BoardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.BoardTable,
			Columns: []string{comment.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CommentLikeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentLikeTable,
			Columns: []string{comment.CommentLikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentlike.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCommentLikeIDs(); len(nodes) > 0 && !cuo.mutation.CommentLikeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentLikeTable,
			Columns: []string{comment.CommentLikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CommentLikeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentLikeTable,
			Columns: []string{comment.CommentLikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CommentMentionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentMentionTable,
			Columns: []string{comment.CommentMentionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentmention.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCommentMentionIDs(); len(nodes) > 0 && !cuo.mutation.CommentMentionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentMentionTable,
			Columns: []string{comment.CommentMentionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentmention.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CommentMentionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentMentionTable,
			Columns: []string{comment.CommentMentionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentmention.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CommentReportCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentReportTable,
			Columns: []string{comment.CommentReportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentreport.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCommentReportIDs(); len(nodes) > 0 && !cuo.mutation.CommentReportCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentReportTable,
			Columns: []string{comment.CommentReportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentreport.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CommentReportIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.CommentReportTable,
			Columns: []string{comment.CommentReportColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(commentreport.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
