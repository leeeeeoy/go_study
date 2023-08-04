// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/leeeeeoy/go_study/ent/comment"
	"github.com/leeeeeoy/go_study/ent/commentmention"
	"github.com/leeeeeoy/go_study/ent/predicate"
	"github.com/leeeeeoy/go_study/ent/user"
)

// CommentMentionUpdate is the builder for updating CommentMention entities.
type CommentMentionUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMentionMutation
}

// Where appends a list predicates to the CommentMentionUpdate builder.
func (cmu *CommentMentionUpdate) Where(ps ...predicate.CommentMention) *CommentMentionUpdate {
	cmu.mutation.Where(ps...)
	return cmu
}

// SetCommentID sets the "comment_id" field.
func (cmu *CommentMentionUpdate) SetCommentID(i int) *CommentMentionUpdate {
	cmu.mutation.SetCommentID(i)
	return cmu
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (cmu *CommentMentionUpdate) SetNillableCommentID(i *int) *CommentMentionUpdate {
	if i != nil {
		cmu.SetCommentID(*i)
	}
	return cmu
}

// ClearCommentID clears the value of the "comment_id" field.
func (cmu *CommentMentionUpdate) ClearCommentID() *CommentMentionUpdate {
	cmu.mutation.ClearCommentID()
	return cmu
}

// SetUserID sets the "user_id" field.
func (cmu *CommentMentionUpdate) SetUserID(i int) *CommentMentionUpdate {
	cmu.mutation.SetUserID(i)
	return cmu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cmu *CommentMentionUpdate) SetNillableUserID(i *int) *CommentMentionUpdate {
	if i != nil {
		cmu.SetUserID(*i)
	}
	return cmu
}

// ClearUserID clears the value of the "user_id" field.
func (cmu *CommentMentionUpdate) ClearUserID() *CommentMentionUpdate {
	cmu.mutation.ClearUserID()
	return cmu
}

// SetUser sets the "user" edge to the User entity.
func (cmu *CommentMentionUpdate) SetUser(u *User) *CommentMentionUpdate {
	return cmu.SetUserID(u.ID)
}

// SetComment sets the "comment" edge to the Comment entity.
func (cmu *CommentMentionUpdate) SetComment(c *Comment) *CommentMentionUpdate {
	return cmu.SetCommentID(c.ID)
}

// Mutation returns the CommentMentionMutation object of the builder.
func (cmu *CommentMentionUpdate) Mutation() *CommentMentionMutation {
	return cmu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cmu *CommentMentionUpdate) ClearUser() *CommentMentionUpdate {
	cmu.mutation.ClearUser()
	return cmu
}

// ClearComment clears the "comment" edge to the Comment entity.
func (cmu *CommentMentionUpdate) ClearComment() *CommentMentionUpdate {
	cmu.mutation.ClearComment()
	return cmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cmu *CommentMentionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cmu.sqlSave, cmu.mutation, cmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cmu *CommentMentionUpdate) SaveX(ctx context.Context) int {
	affected, err := cmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cmu *CommentMentionUpdate) Exec(ctx context.Context) error {
	_, err := cmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmu *CommentMentionUpdate) ExecX(ctx context.Context) {
	if err := cmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cmu *CommentMentionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(commentmention.Table, commentmention.Columns, sqlgraph.NewFieldSpec(commentmention.FieldID, field.TypeInt))
	if ps := cmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cmu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commentmention.UserTable,
			Columns: []string{commentmention.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commentmention.UserTable,
			Columns: []string{commentmention.UserColumn},
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
	if cmu.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commentmention.CommentTable,
			Columns: []string{commentmention.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmu.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commentmention.CommentTable,
			Columns: []string{commentmention.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commentmention.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cmu.mutation.done = true
	return n, nil
}

// CommentMentionUpdateOne is the builder for updating a single CommentMention entity.
type CommentMentionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMentionMutation
}

// SetCommentID sets the "comment_id" field.
func (cmuo *CommentMentionUpdateOne) SetCommentID(i int) *CommentMentionUpdateOne {
	cmuo.mutation.SetCommentID(i)
	return cmuo
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (cmuo *CommentMentionUpdateOne) SetNillableCommentID(i *int) *CommentMentionUpdateOne {
	if i != nil {
		cmuo.SetCommentID(*i)
	}
	return cmuo
}

// ClearCommentID clears the value of the "comment_id" field.
func (cmuo *CommentMentionUpdateOne) ClearCommentID() *CommentMentionUpdateOne {
	cmuo.mutation.ClearCommentID()
	return cmuo
}

// SetUserID sets the "user_id" field.
func (cmuo *CommentMentionUpdateOne) SetUserID(i int) *CommentMentionUpdateOne {
	cmuo.mutation.SetUserID(i)
	return cmuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cmuo *CommentMentionUpdateOne) SetNillableUserID(i *int) *CommentMentionUpdateOne {
	if i != nil {
		cmuo.SetUserID(*i)
	}
	return cmuo
}

// ClearUserID clears the value of the "user_id" field.
func (cmuo *CommentMentionUpdateOne) ClearUserID() *CommentMentionUpdateOne {
	cmuo.mutation.ClearUserID()
	return cmuo
}

// SetUser sets the "user" edge to the User entity.
func (cmuo *CommentMentionUpdateOne) SetUser(u *User) *CommentMentionUpdateOne {
	return cmuo.SetUserID(u.ID)
}

// SetComment sets the "comment" edge to the Comment entity.
func (cmuo *CommentMentionUpdateOne) SetComment(c *Comment) *CommentMentionUpdateOne {
	return cmuo.SetCommentID(c.ID)
}

// Mutation returns the CommentMentionMutation object of the builder.
func (cmuo *CommentMentionUpdateOne) Mutation() *CommentMentionMutation {
	return cmuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cmuo *CommentMentionUpdateOne) ClearUser() *CommentMentionUpdateOne {
	cmuo.mutation.ClearUser()
	return cmuo
}

// ClearComment clears the "comment" edge to the Comment entity.
func (cmuo *CommentMentionUpdateOne) ClearComment() *CommentMentionUpdateOne {
	cmuo.mutation.ClearComment()
	return cmuo
}

// Where appends a list predicates to the CommentMentionUpdate builder.
func (cmuo *CommentMentionUpdateOne) Where(ps ...predicate.CommentMention) *CommentMentionUpdateOne {
	cmuo.mutation.Where(ps...)
	return cmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cmuo *CommentMentionUpdateOne) Select(field string, fields ...string) *CommentMentionUpdateOne {
	cmuo.fields = append([]string{field}, fields...)
	return cmuo
}

// Save executes the query and returns the updated CommentMention entity.
func (cmuo *CommentMentionUpdateOne) Save(ctx context.Context) (*CommentMention, error) {
	return withHooks(ctx, cmuo.sqlSave, cmuo.mutation, cmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cmuo *CommentMentionUpdateOne) SaveX(ctx context.Context) *CommentMention {
	node, err := cmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cmuo *CommentMentionUpdateOne) Exec(ctx context.Context) error {
	_, err := cmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmuo *CommentMentionUpdateOne) ExecX(ctx context.Context) {
	if err := cmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cmuo *CommentMentionUpdateOne) sqlSave(ctx context.Context) (_node *CommentMention, err error) {
	_spec := sqlgraph.NewUpdateSpec(commentmention.Table, commentmention.Columns, sqlgraph.NewFieldSpec(commentmention.FieldID, field.TypeInt))
	id, ok := cmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CommentMention.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, commentmention.FieldID)
		for _, f := range fields {
			if !commentmention.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != commentmention.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cmuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commentmention.UserTable,
			Columns: []string{commentmention.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commentmention.UserTable,
			Columns: []string{commentmention.UserColumn},
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
	if cmuo.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commentmention.CommentTable,
			Columns: []string{commentmention.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmuo.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commentmention.CommentTable,
			Columns: []string{commentmention.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CommentMention{config: cmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{commentmention.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cmuo.mutation.done = true
	return _node, nil
}
