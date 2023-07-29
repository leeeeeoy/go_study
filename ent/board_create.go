// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/leeeeeoy/go_study/ent/board"
	"github.com/leeeeeoy/go_study/ent/boardlike"
	"github.com/leeeeeoy/go_study/ent/comment"
	"github.com/leeeeeoy/go_study/ent/user"
)

// BoardCreate is the builder for creating a Board entity.
type BoardCreate struct {
	config
	mutation *BoardMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (bc *BoardCreate) SetTitle(s string) *BoardCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetContent sets the "content" field.
func (bc *BoardCreate) SetContent(s string) *BoardCreate {
	bc.mutation.SetContent(s)
	return bc
}

// SetUserID sets the "user_id" field.
func (bc *BoardCreate) SetUserID(i int) *BoardCreate {
	bc.mutation.SetUserID(i)
	return bc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (bc *BoardCreate) SetNillableUserID(i *int) *BoardCreate {
	if i != nil {
		bc.SetUserID(*i)
	}
	return bc
}

// SetLikeCount sets the "like_count" field.
func (bc *BoardCreate) SetLikeCount(i int) *BoardCreate {
	bc.mutation.SetLikeCount(i)
	return bc
}

// SetCommentCount sets the "comment_count" field.
func (bc *BoardCreate) SetCommentCount(i int) *BoardCreate {
	bc.mutation.SetCommentCount(i)
	return bc
}

// SetCreatedAt sets the "created_at" field.
func (bc *BoardCreate) SetCreatedAt(t time.Time) *BoardCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BoardCreate) SetNillableCreatedAt(t *time.Time) *BoardCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BoardCreate) SetUpdatedAt(t time.Time) *BoardCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BoardCreate) SetNillableUpdatedAt(t *time.Time) *BoardCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetUser sets the "user" edge to the User entity.
func (bc *BoardCreate) SetUser(u *User) *BoardCreate {
	return bc.SetUserID(u.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (bc *BoardCreate) AddCommentIDs(ids ...int) *BoardCreate {
	bc.mutation.AddCommentIDs(ids...)
	return bc
}

// AddComments adds the "comments" edges to the Comment entity.
func (bc *BoardCreate) AddComments(c ...*Comment) *BoardCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bc.AddCommentIDs(ids...)
}

// AddBoardLikeIDs adds the "board_like" edge to the BoardLike entity by IDs.
func (bc *BoardCreate) AddBoardLikeIDs(ids ...int) *BoardCreate {
	bc.mutation.AddBoardLikeIDs(ids...)
	return bc
}

// AddBoardLike adds the "board_like" edges to the BoardLike entity.
func (bc *BoardCreate) AddBoardLike(b ...*BoardLike) *BoardCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddBoardLikeIDs(ids...)
}

// Mutation returns the BoardMutation object of the builder.
func (bc *BoardCreate) Mutation() *BoardMutation {
	return bc.mutation
}

// Save creates the Board in the database.
func (bc *BoardCreate) Save(ctx context.Context) (*Board, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BoardCreate) SaveX(ctx context.Context) *Board {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BoardCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BoardCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BoardCreate) defaults() {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := board.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := board.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BoardCreate) check() error {
	if _, ok := bc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Board.title"`)}
	}
	if _, ok := bc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Board.content"`)}
	}
	if _, ok := bc.mutation.LikeCount(); !ok {
		return &ValidationError{Name: "like_count", err: errors.New(`ent: missing required field "Board.like_count"`)}
	}
	if _, ok := bc.mutation.CommentCount(); !ok {
		return &ValidationError{Name: "comment_count", err: errors.New(`ent: missing required field "Board.comment_count"`)}
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Board.created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Board.updated_at"`)}
	}
	return nil
}

func (bc *BoardCreate) sqlSave(ctx context.Context) (*Board, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BoardCreate) createSpec() (*Board, *sqlgraph.CreateSpec) {
	var (
		_node = &Board{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(board.Table, sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.Title(); ok {
		_spec.SetField(board.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bc.mutation.Content(); ok {
		_spec.SetField(board.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := bc.mutation.LikeCount(); ok {
		_spec.SetField(board.FieldLikeCount, field.TypeInt, value)
		_node.LikeCount = value
	}
	if value, ok := bc.mutation.CommentCount(); ok {
		_spec.SetField(board.FieldCommentCount, field.TypeInt, value)
		_node.CommentCount = value
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(board.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(board.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := bc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   board.UserTable,
			Columns: []string{board.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   board.CommentsTable,
			Columns: []string{board.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.BoardLikeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   board.BoardLikeTable,
			Columns: []string{board.BoardLikeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(boardlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BoardCreateBulk is the builder for creating many Board entities in bulk.
type BoardCreateBulk struct {
	config
	builders []*BoardCreate
}

// Save creates the Board entities in the database.
func (bcb *BoardCreateBulk) Save(ctx context.Context) ([]*Board, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Board, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BoardMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BoardCreateBulk) SaveX(ctx context.Context) []*Board {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BoardCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BoardCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
