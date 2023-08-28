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
	"github.com/leeeeeoy/go_study/ent/bookmark"
	"github.com/leeeeeoy/go_study/ent/user"
)

// BookMarkCreate is the builder for creating a BookMark entity.
type BookMarkCreate struct {
	config
	mutation *BookMarkMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (bmc *BookMarkCreate) SetUserID(i int) *BookMarkCreate {
	bmc.mutation.SetUserID(i)
	return bmc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (bmc *BookMarkCreate) SetNillableUserID(i *int) *BookMarkCreate {
	if i != nil {
		bmc.SetUserID(*i)
	}
	return bmc
}

// SetBoardID sets the "board_id" field.
func (bmc *BookMarkCreate) SetBoardID(i int) *BookMarkCreate {
	bmc.mutation.SetBoardID(i)
	return bmc
}

// SetNillableBoardID sets the "board_id" field if the given value is not nil.
func (bmc *BookMarkCreate) SetNillableBoardID(i *int) *BookMarkCreate {
	if i != nil {
		bmc.SetBoardID(*i)
	}
	return bmc
}

// SetCreatedAt sets the "created_at" field.
func (bmc *BookMarkCreate) SetCreatedAt(t time.Time) *BookMarkCreate {
	bmc.mutation.SetCreatedAt(t)
	return bmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bmc *BookMarkCreate) SetNillableCreatedAt(t *time.Time) *BookMarkCreate {
	if t != nil {
		bmc.SetCreatedAt(*t)
	}
	return bmc
}

// SetBoard sets the "board" edge to the Board entity.
func (bmc *BookMarkCreate) SetBoard(b *Board) *BookMarkCreate {
	return bmc.SetBoardID(b.ID)
}

// SetUser sets the "user" edge to the User entity.
func (bmc *BookMarkCreate) SetUser(u *User) *BookMarkCreate {
	return bmc.SetUserID(u.ID)
}

// Mutation returns the BookMarkMutation object of the builder.
func (bmc *BookMarkCreate) Mutation() *BookMarkMutation {
	return bmc.mutation
}

// Save creates the BookMark in the database.
func (bmc *BookMarkCreate) Save(ctx context.Context) (*BookMark, error) {
	bmc.defaults()
	return withHooks(ctx, bmc.sqlSave, bmc.mutation, bmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bmc *BookMarkCreate) SaveX(ctx context.Context) *BookMark {
	v, err := bmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bmc *BookMarkCreate) Exec(ctx context.Context) error {
	_, err := bmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmc *BookMarkCreate) ExecX(ctx context.Context) {
	if err := bmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bmc *BookMarkCreate) defaults() {
	if _, ok := bmc.mutation.CreatedAt(); !ok {
		v := bookmark.DefaultCreatedAt()
		bmc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bmc *BookMarkCreate) check() error {
	if _, ok := bmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BookMark.created_at"`)}
	}
	return nil
}

func (bmc *BookMarkCreate) sqlSave(ctx context.Context) (*BookMark, error) {
	if err := bmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bmc.mutation.id = &_node.ID
	bmc.mutation.done = true
	return _node, nil
}

func (bmc *BookMarkCreate) createSpec() (*BookMark, *sqlgraph.CreateSpec) {
	var (
		_node = &BookMark{config: bmc.config}
		_spec = sqlgraph.NewCreateSpec(bookmark.Table, sqlgraph.NewFieldSpec(bookmark.FieldID, field.TypeInt))
	)
	if value, ok := bmc.mutation.CreatedAt(); ok {
		_spec.SetField(bookmark.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := bmc.mutation.BoardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.BoardTable,
			Columns: []string{bookmark.BoardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BoardID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bmc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookmark.UserTable,
			Columns: []string{bookmark.UserColumn},
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
	return _node, _spec
}

// BookMarkCreateBulk is the builder for creating many BookMark entities in bulk.
type BookMarkCreateBulk struct {
	config
	builders []*BookMarkCreate
}

// Save creates the BookMark entities in the database.
func (bmcb *BookMarkCreateBulk) Save(ctx context.Context) ([]*BookMark, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bmcb.builders))
	nodes := make([]*BookMark, len(bmcb.builders))
	mutators := make([]Mutator, len(bmcb.builders))
	for i := range bmcb.builders {
		func(i int, root context.Context) {
			builder := bmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookMarkMutation)
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
					_, err = mutators[i+1].Mutate(root, bmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bmcb *BookMarkCreateBulk) SaveX(ctx context.Context) []*BookMark {
	v, err := bmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bmcb *BookMarkCreateBulk) Exec(ctx context.Context) error {
	_, err := bmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmcb *BookMarkCreateBulk) ExecX(ctx context.Context) {
	if err := bmcb.Exec(ctx); err != nil {
		panic(err)
	}
}