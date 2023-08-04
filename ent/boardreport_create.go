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
	"github.com/leeeeeoy/go_study/ent/boardreport"
	"github.com/leeeeeoy/go_study/ent/reporttype"
	"github.com/leeeeeoy/go_study/ent/user"
)

// BoardReportCreate is the builder for creating a BoardReport entity.
type BoardReportCreate struct {
	config
	mutation *BoardReportMutation
	hooks    []Hook
}

// SetBoardID sets the "board_id" field.
func (brc *BoardReportCreate) SetBoardID(i int) *BoardReportCreate {
	brc.mutation.SetBoardID(i)
	return brc
}

// SetNillableBoardID sets the "board_id" field if the given value is not nil.
func (brc *BoardReportCreate) SetNillableBoardID(i *int) *BoardReportCreate {
	if i != nil {
		brc.SetBoardID(*i)
	}
	return brc
}

// SetReportTypeID sets the "report_type_id" field.
func (brc *BoardReportCreate) SetReportTypeID(i int) *BoardReportCreate {
	brc.mutation.SetReportTypeID(i)
	return brc
}

// SetNillableReportTypeID sets the "report_type_id" field if the given value is not nil.
func (brc *BoardReportCreate) SetNillableReportTypeID(i *int) *BoardReportCreate {
	if i != nil {
		brc.SetReportTypeID(*i)
	}
	return brc
}

// SetReporterID sets the "reporter_id" field.
func (brc *BoardReportCreate) SetReporterID(i int) *BoardReportCreate {
	brc.mutation.SetReporterID(i)
	return brc
}

// SetNillableReporterID sets the "reporter_id" field if the given value is not nil.
func (brc *BoardReportCreate) SetNillableReporterID(i *int) *BoardReportCreate {
	if i != nil {
		brc.SetReporterID(*i)
	}
	return brc
}

// SetComment sets the "comment" field.
func (brc *BoardReportCreate) SetComment(s string) *BoardReportCreate {
	brc.mutation.SetComment(s)
	return brc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (brc *BoardReportCreate) SetNillableComment(s *string) *BoardReportCreate {
	if s != nil {
		brc.SetComment(*s)
	}
	return brc
}

// SetStatus sets the "status" field.
func (brc *BoardReportCreate) SetStatus(b boardreport.Status) *BoardReportCreate {
	brc.mutation.SetStatus(b)
	return brc
}

// SetCreatedAt sets the "created_at" field.
func (brc *BoardReportCreate) SetCreatedAt(t time.Time) *BoardReportCreate {
	brc.mutation.SetCreatedAt(t)
	return brc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (brc *BoardReportCreate) SetNillableCreatedAt(t *time.Time) *BoardReportCreate {
	if t != nil {
		brc.SetCreatedAt(*t)
	}
	return brc
}

// SetUpdatedAt sets the "updated_at" field.
func (brc *BoardReportCreate) SetUpdatedAt(t time.Time) *BoardReportCreate {
	brc.mutation.SetUpdatedAt(t)
	return brc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (brc *BoardReportCreate) SetNillableUpdatedAt(t *time.Time) *BoardReportCreate {
	if t != nil {
		brc.SetUpdatedAt(*t)
	}
	return brc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (brc *BoardReportCreate) SetUserID(id int) *BoardReportCreate {
	brc.mutation.SetUserID(id)
	return brc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (brc *BoardReportCreate) SetNillableUserID(id *int) *BoardReportCreate {
	if id != nil {
		brc = brc.SetUserID(*id)
	}
	return brc
}

// SetUser sets the "user" edge to the User entity.
func (brc *BoardReportCreate) SetUser(u *User) *BoardReportCreate {
	return brc.SetUserID(u.ID)
}

// SetBoard sets the "board" edge to the Board entity.
func (brc *BoardReportCreate) SetBoard(b *Board) *BoardReportCreate {
	return brc.SetBoardID(b.ID)
}

// SetReportType sets the "report_type" edge to the ReportType entity.
func (brc *BoardReportCreate) SetReportType(r *ReportType) *BoardReportCreate {
	return brc.SetReportTypeID(r.ID)
}

// Mutation returns the BoardReportMutation object of the builder.
func (brc *BoardReportCreate) Mutation() *BoardReportMutation {
	return brc.mutation
}

// Save creates the BoardReport in the database.
func (brc *BoardReportCreate) Save(ctx context.Context) (*BoardReport, error) {
	brc.defaults()
	return withHooks(ctx, brc.sqlSave, brc.mutation, brc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (brc *BoardReportCreate) SaveX(ctx context.Context) *BoardReport {
	v, err := brc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (brc *BoardReportCreate) Exec(ctx context.Context) error {
	_, err := brc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (brc *BoardReportCreate) ExecX(ctx context.Context) {
	if err := brc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (brc *BoardReportCreate) defaults() {
	if _, ok := brc.mutation.CreatedAt(); !ok {
		v := boardreport.DefaultCreatedAt()
		brc.mutation.SetCreatedAt(v)
	}
	if _, ok := brc.mutation.UpdatedAt(); !ok {
		v := boardreport.DefaultUpdatedAt()
		brc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (brc *BoardReportCreate) check() error {
	if _, ok := brc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "BoardReport.status"`)}
	}
	if v, ok := brc.mutation.Status(); ok {
		if err := boardreport.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "BoardReport.status": %w`, err)}
		}
	}
	if _, ok := brc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BoardReport.created_at"`)}
	}
	if _, ok := brc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "BoardReport.updated_at"`)}
	}
	return nil
}

func (brc *BoardReportCreate) sqlSave(ctx context.Context) (*BoardReport, error) {
	if err := brc.check(); err != nil {
		return nil, err
	}
	_node, _spec := brc.createSpec()
	if err := sqlgraph.CreateNode(ctx, brc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	brc.mutation.id = &_node.ID
	brc.mutation.done = true
	return _node, nil
}

func (brc *BoardReportCreate) createSpec() (*BoardReport, *sqlgraph.CreateSpec) {
	var (
		_node = &BoardReport{config: brc.config}
		_spec = sqlgraph.NewCreateSpec(boardreport.Table, sqlgraph.NewFieldSpec(boardreport.FieldID, field.TypeInt))
	)
	if value, ok := brc.mutation.Comment(); ok {
		_spec.SetField(boardreport.FieldComment, field.TypeString, value)
		_node.Comment = value
	}
	if value, ok := brc.mutation.Status(); ok {
		_spec.SetField(boardreport.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := brc.mutation.CreatedAt(); ok {
		_spec.SetField(boardreport.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := brc.mutation.UpdatedAt(); ok {
		_spec.SetField(boardreport.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := brc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   boardreport.UserTable,
			Columns: []string{boardreport.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ReporterID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := brc.mutation.BoardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   boardreport.BoardTable,
			Columns: []string{boardreport.BoardColumn},
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
	if nodes := brc.mutation.ReportTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   boardreport.ReportTypeTable,
			Columns: []string{boardreport.ReportTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reporttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ReportTypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BoardReportCreateBulk is the builder for creating many BoardReport entities in bulk.
type BoardReportCreateBulk struct {
	config
	builders []*BoardReportCreate
}

// Save creates the BoardReport entities in the database.
func (brcb *BoardReportCreateBulk) Save(ctx context.Context) ([]*BoardReport, error) {
	specs := make([]*sqlgraph.CreateSpec, len(brcb.builders))
	nodes := make([]*BoardReport, len(brcb.builders))
	mutators := make([]Mutator, len(brcb.builders))
	for i := range brcb.builders {
		func(i int, root context.Context) {
			builder := brcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BoardReportMutation)
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
					_, err = mutators[i+1].Mutate(root, brcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, brcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, brcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (brcb *BoardReportCreateBulk) SaveX(ctx context.Context) []*BoardReport {
	v, err := brcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (brcb *BoardReportCreateBulk) Exec(ctx context.Context) error {
	_, err := brcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (brcb *BoardReportCreateBulk) ExecX(ctx context.Context) {
	if err := brcb.Exec(ctx); err != nil {
		panic(err)
	}
}
