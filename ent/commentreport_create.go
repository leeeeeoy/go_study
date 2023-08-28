// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/leeeeeoy/go_study/ent/comment"
	"github.com/leeeeeoy/go_study/ent/commentreport"
	"github.com/leeeeeoy/go_study/ent/reporttype"
	"github.com/leeeeeoy/go_study/ent/user"
)

// CommentReportCreate is the builder for creating a CommentReport entity.
type CommentReportCreate struct {
	config
	mutation *CommentReportMutation
	hooks    []Hook
}

// SetCommentID sets the "comment_id" field.
func (crc *CommentReportCreate) SetCommentID(i int) *CommentReportCreate {
	crc.mutation.SetCommentID(i)
	return crc
}

// SetNillableCommentID sets the "comment_id" field if the given value is not nil.
func (crc *CommentReportCreate) SetNillableCommentID(i *int) *CommentReportCreate {
	if i != nil {
		crc.SetCommentID(*i)
	}
	return crc
}

// SetReportTypeID sets the "report_type_id" field.
func (crc *CommentReportCreate) SetReportTypeID(i int) *CommentReportCreate {
	crc.mutation.SetReportTypeID(i)
	return crc
}

// SetNillableReportTypeID sets the "report_type_id" field if the given value is not nil.
func (crc *CommentReportCreate) SetNillableReportTypeID(i *int) *CommentReportCreate {
	if i != nil {
		crc.SetReportTypeID(*i)
	}
	return crc
}

// SetReporterID sets the "reporter_id" field.
func (crc *CommentReportCreate) SetReporterID(i int) *CommentReportCreate {
	crc.mutation.SetReporterID(i)
	return crc
}

// SetNillableReporterID sets the "reporter_id" field if the given value is not nil.
func (crc *CommentReportCreate) SetNillableReporterID(i *int) *CommentReportCreate {
	if i != nil {
		crc.SetReporterID(*i)
	}
	return crc
}

// SetDesc sets the "desc" field.
func (crc *CommentReportCreate) SetDesc(s string) *CommentReportCreate {
	crc.mutation.SetDesc(s)
	return crc
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (crc *CommentReportCreate) SetNillableDesc(s *string) *CommentReportCreate {
	if s != nil {
		crc.SetDesc(*s)
	}
	return crc
}

// SetStatus sets the "status" field.
func (crc *CommentReportCreate) SetStatus(c commentreport.Status) *CommentReportCreate {
	crc.mutation.SetStatus(c)
	return crc
}

// SetCreatedAt sets the "created_at" field.
func (crc *CommentReportCreate) SetCreatedAt(t time.Time) *CommentReportCreate {
	crc.mutation.SetCreatedAt(t)
	return crc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (crc *CommentReportCreate) SetNillableCreatedAt(t *time.Time) *CommentReportCreate {
	if t != nil {
		crc.SetCreatedAt(*t)
	}
	return crc
}

// SetUpdatedAt sets the "updated_at" field.
func (crc *CommentReportCreate) SetUpdatedAt(t time.Time) *CommentReportCreate {
	crc.mutation.SetUpdatedAt(t)
	return crc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (crc *CommentReportCreate) SetNillableUpdatedAt(t *time.Time) *CommentReportCreate {
	if t != nil {
		crc.SetUpdatedAt(*t)
	}
	return crc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (crc *CommentReportCreate) SetUserID(id int) *CommentReportCreate {
	crc.mutation.SetUserID(id)
	return crc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (crc *CommentReportCreate) SetNillableUserID(id *int) *CommentReportCreate {
	if id != nil {
		crc = crc.SetUserID(*id)
	}
	return crc
}

// SetUser sets the "user" edge to the User entity.
func (crc *CommentReportCreate) SetUser(u *User) *CommentReportCreate {
	return crc.SetUserID(u.ID)
}

// SetComment sets the "comment" edge to the Comment entity.
func (crc *CommentReportCreate) SetComment(c *Comment) *CommentReportCreate {
	return crc.SetCommentID(c.ID)
}

// SetReportType sets the "report_type" edge to the ReportType entity.
func (crc *CommentReportCreate) SetReportType(r *ReportType) *CommentReportCreate {
	return crc.SetReportTypeID(r.ID)
}

// Mutation returns the CommentReportMutation object of the builder.
func (crc *CommentReportCreate) Mutation() *CommentReportMutation {
	return crc.mutation
}

// Save creates the CommentReport in the database.
func (crc *CommentReportCreate) Save(ctx context.Context) (*CommentReport, error) {
	crc.defaults()
	return withHooks(ctx, crc.sqlSave, crc.mutation, crc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (crc *CommentReportCreate) SaveX(ctx context.Context) *CommentReport {
	v, err := crc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crc *CommentReportCreate) Exec(ctx context.Context) error {
	_, err := crc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crc *CommentReportCreate) ExecX(ctx context.Context) {
	if err := crc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (crc *CommentReportCreate) defaults() {
	if _, ok := crc.mutation.CreatedAt(); !ok {
		v := commentreport.DefaultCreatedAt()
		crc.mutation.SetCreatedAt(v)
	}
	if _, ok := crc.mutation.UpdatedAt(); !ok {
		v := commentreport.DefaultUpdatedAt()
		crc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (crc *CommentReportCreate) check() error {
	if _, ok := crc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "CommentReport.status"`)}
	}
	if v, ok := crc.mutation.Status(); ok {
		if err := commentreport.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "CommentReport.status": %w`, err)}
		}
	}
	if _, ok := crc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CommentReport.created_at"`)}
	}
	if _, ok := crc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CommentReport.updated_at"`)}
	}
	return nil
}

func (crc *CommentReportCreate) sqlSave(ctx context.Context) (*CommentReport, error) {
	if err := crc.check(); err != nil {
		return nil, err
	}
	_node, _spec := crc.createSpec()
	if err := sqlgraph.CreateNode(ctx, crc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	crc.mutation.id = &_node.ID
	crc.mutation.done = true
	return _node, nil
}

func (crc *CommentReportCreate) createSpec() (*CommentReport, *sqlgraph.CreateSpec) {
	var (
		_node = &CommentReport{config: crc.config}
		_spec = sqlgraph.NewCreateSpec(commentreport.Table, sqlgraph.NewFieldSpec(commentreport.FieldID, field.TypeInt))
	)
	if value, ok := crc.mutation.Desc(); ok {
		_spec.SetField(commentreport.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if value, ok := crc.mutation.Status(); ok {
		_spec.SetField(commentreport.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := crc.mutation.CreatedAt(); ok {
		_spec.SetField(commentreport.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := crc.mutation.UpdatedAt(); ok {
		_spec.SetField(commentreport.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := crc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commentreport.UserTable,
			Columns: []string{commentreport.UserColumn},
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
	if nodes := crc.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   commentreport.CommentTable,
			Columns: []string{commentreport.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CommentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := crc.mutation.ReportTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   commentreport.ReportTypeTable,
			Columns: []string{commentreport.ReportTypeColumn},
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

// CommentReportCreateBulk is the builder for creating many CommentReport entities in bulk.
type CommentReportCreateBulk struct {
	config
	builders []*CommentReportCreate
}

// Save creates the CommentReport entities in the database.
func (crcb *CommentReportCreateBulk) Save(ctx context.Context) ([]*CommentReport, error) {
	specs := make([]*sqlgraph.CreateSpec, len(crcb.builders))
	nodes := make([]*CommentReport, len(crcb.builders))
	mutators := make([]Mutator, len(crcb.builders))
	for i := range crcb.builders {
		func(i int, root context.Context) {
			builder := crcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentReportMutation)
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
					_, err = mutators[i+1].Mutate(root, crcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, crcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, crcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (crcb *CommentReportCreateBulk) SaveX(ctx context.Context) []*CommentReport {
	v, err := crcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crcb *CommentReportCreateBulk) Exec(ctx context.Context) error {
	_, err := crcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crcb *CommentReportCreateBulk) ExecX(ctx context.Context) {
	if err := crcb.Exec(ctx); err != nil {
		panic(err)
	}
}