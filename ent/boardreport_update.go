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
	"github.com/leeeeeoy/go_study/ent/boardreport"
	"github.com/leeeeeoy/go_study/ent/predicate"
	"github.com/leeeeeoy/go_study/ent/reporttype"
	"github.com/leeeeeoy/go_study/ent/user"
)

// BoardReportUpdate is the builder for updating BoardReport entities.
type BoardReportUpdate struct {
	config
	hooks    []Hook
	mutation *BoardReportMutation
}

// Where appends a list predicates to the BoardReportUpdate builder.
func (bru *BoardReportUpdate) Where(ps ...predicate.BoardReport) *BoardReportUpdate {
	bru.mutation.Where(ps...)
	return bru
}

// SetBoardID sets the "board_id" field.
func (bru *BoardReportUpdate) SetBoardID(i int) *BoardReportUpdate {
	bru.mutation.SetBoardID(i)
	return bru
}

// SetNillableBoardID sets the "board_id" field if the given value is not nil.
func (bru *BoardReportUpdate) SetNillableBoardID(i *int) *BoardReportUpdate {
	if i != nil {
		bru.SetBoardID(*i)
	}
	return bru
}

// ClearBoardID clears the value of the "board_id" field.
func (bru *BoardReportUpdate) ClearBoardID() *BoardReportUpdate {
	bru.mutation.ClearBoardID()
	return bru
}

// SetReportTypeID sets the "report_type_id" field.
func (bru *BoardReportUpdate) SetReportTypeID(i int) *BoardReportUpdate {
	bru.mutation.SetReportTypeID(i)
	return bru
}

// SetNillableReportTypeID sets the "report_type_id" field if the given value is not nil.
func (bru *BoardReportUpdate) SetNillableReportTypeID(i *int) *BoardReportUpdate {
	if i != nil {
		bru.SetReportTypeID(*i)
	}
	return bru
}

// ClearReportTypeID clears the value of the "report_type_id" field.
func (bru *BoardReportUpdate) ClearReportTypeID() *BoardReportUpdate {
	bru.mutation.ClearReportTypeID()
	return bru
}

// SetReporterID sets the "reporter_id" field.
func (bru *BoardReportUpdate) SetReporterID(i int) *BoardReportUpdate {
	bru.mutation.SetReporterID(i)
	return bru
}

// SetNillableReporterID sets the "reporter_id" field if the given value is not nil.
func (bru *BoardReportUpdate) SetNillableReporterID(i *int) *BoardReportUpdate {
	if i != nil {
		bru.SetReporterID(*i)
	}
	return bru
}

// ClearReporterID clears the value of the "reporter_id" field.
func (bru *BoardReportUpdate) ClearReporterID() *BoardReportUpdate {
	bru.mutation.ClearReporterID()
	return bru
}

// SetComment sets the "comment" field.
func (bru *BoardReportUpdate) SetComment(s string) *BoardReportUpdate {
	bru.mutation.SetComment(s)
	return bru
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (bru *BoardReportUpdate) SetNillableComment(s *string) *BoardReportUpdate {
	if s != nil {
		bru.SetComment(*s)
	}
	return bru
}

// ClearComment clears the value of the "comment" field.
func (bru *BoardReportUpdate) ClearComment() *BoardReportUpdate {
	bru.mutation.ClearComment()
	return bru
}

// SetStatus sets the "status" field.
func (bru *BoardReportUpdate) SetStatus(b boardreport.Status) *BoardReportUpdate {
	bru.mutation.SetStatus(b)
	return bru
}

// SetCreatedAt sets the "created_at" field.
func (bru *BoardReportUpdate) SetCreatedAt(t time.Time) *BoardReportUpdate {
	bru.mutation.SetCreatedAt(t)
	return bru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bru *BoardReportUpdate) SetNillableCreatedAt(t *time.Time) *BoardReportUpdate {
	if t != nil {
		bru.SetCreatedAt(*t)
	}
	return bru
}

// SetUpdatedAt sets the "updated_at" field.
func (bru *BoardReportUpdate) SetUpdatedAt(t time.Time) *BoardReportUpdate {
	bru.mutation.SetUpdatedAt(t)
	return bru
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bru *BoardReportUpdate) SetUserID(id int) *BoardReportUpdate {
	bru.mutation.SetUserID(id)
	return bru
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (bru *BoardReportUpdate) SetNillableUserID(id *int) *BoardReportUpdate {
	if id != nil {
		bru = bru.SetUserID(*id)
	}
	return bru
}

// SetUser sets the "user" edge to the User entity.
func (bru *BoardReportUpdate) SetUser(u *User) *BoardReportUpdate {
	return bru.SetUserID(u.ID)
}

// SetBoard sets the "board" edge to the Board entity.
func (bru *BoardReportUpdate) SetBoard(b *Board) *BoardReportUpdate {
	return bru.SetBoardID(b.ID)
}

// SetReportType sets the "report_type" edge to the ReportType entity.
func (bru *BoardReportUpdate) SetReportType(r *ReportType) *BoardReportUpdate {
	return bru.SetReportTypeID(r.ID)
}

// Mutation returns the BoardReportMutation object of the builder.
func (bru *BoardReportUpdate) Mutation() *BoardReportMutation {
	return bru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (bru *BoardReportUpdate) ClearUser() *BoardReportUpdate {
	bru.mutation.ClearUser()
	return bru
}

// ClearBoard clears the "board" edge to the Board entity.
func (bru *BoardReportUpdate) ClearBoard() *BoardReportUpdate {
	bru.mutation.ClearBoard()
	return bru
}

// ClearReportType clears the "report_type" edge to the ReportType entity.
func (bru *BoardReportUpdate) ClearReportType() *BoardReportUpdate {
	bru.mutation.ClearReportType()
	return bru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bru *BoardReportUpdate) Save(ctx context.Context) (int, error) {
	bru.defaults()
	return withHooks(ctx, bru.sqlSave, bru.mutation, bru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bru *BoardReportUpdate) SaveX(ctx context.Context) int {
	affected, err := bru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bru *BoardReportUpdate) Exec(ctx context.Context) error {
	_, err := bru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bru *BoardReportUpdate) ExecX(ctx context.Context) {
	if err := bru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bru *BoardReportUpdate) defaults() {
	if _, ok := bru.mutation.UpdatedAt(); !ok {
		v := boardreport.UpdateDefaultUpdatedAt()
		bru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bru *BoardReportUpdate) check() error {
	if v, ok := bru.mutation.Status(); ok {
		if err := boardreport.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "BoardReport.status": %w`, err)}
		}
	}
	return nil
}

func (bru *BoardReportUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(boardreport.Table, boardreport.Columns, sqlgraph.NewFieldSpec(boardreport.FieldID, field.TypeInt))
	if ps := bru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bru.mutation.Comment(); ok {
		_spec.SetField(boardreport.FieldComment, field.TypeString, value)
	}
	if bru.mutation.CommentCleared() {
		_spec.ClearField(boardreport.FieldComment, field.TypeString)
	}
	if value, ok := bru.mutation.Status(); ok {
		_spec.SetField(boardreport.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := bru.mutation.CreatedAt(); ok {
		_spec.SetField(boardreport.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bru.mutation.UpdatedAt(); ok {
		_spec.SetField(boardreport.FieldUpdatedAt, field.TypeTime, value)
	}
	if bru.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bru.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bru.mutation.BoardCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bru.mutation.BoardIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bru.mutation.ReportTypeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bru.mutation.ReportTypeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{boardreport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bru.mutation.done = true
	return n, nil
}

// BoardReportUpdateOne is the builder for updating a single BoardReport entity.
type BoardReportUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BoardReportMutation
}

// SetBoardID sets the "board_id" field.
func (bruo *BoardReportUpdateOne) SetBoardID(i int) *BoardReportUpdateOne {
	bruo.mutation.SetBoardID(i)
	return bruo
}

// SetNillableBoardID sets the "board_id" field if the given value is not nil.
func (bruo *BoardReportUpdateOne) SetNillableBoardID(i *int) *BoardReportUpdateOne {
	if i != nil {
		bruo.SetBoardID(*i)
	}
	return bruo
}

// ClearBoardID clears the value of the "board_id" field.
func (bruo *BoardReportUpdateOne) ClearBoardID() *BoardReportUpdateOne {
	bruo.mutation.ClearBoardID()
	return bruo
}

// SetReportTypeID sets the "report_type_id" field.
func (bruo *BoardReportUpdateOne) SetReportTypeID(i int) *BoardReportUpdateOne {
	bruo.mutation.SetReportTypeID(i)
	return bruo
}

// SetNillableReportTypeID sets the "report_type_id" field if the given value is not nil.
func (bruo *BoardReportUpdateOne) SetNillableReportTypeID(i *int) *BoardReportUpdateOne {
	if i != nil {
		bruo.SetReportTypeID(*i)
	}
	return bruo
}

// ClearReportTypeID clears the value of the "report_type_id" field.
func (bruo *BoardReportUpdateOne) ClearReportTypeID() *BoardReportUpdateOne {
	bruo.mutation.ClearReportTypeID()
	return bruo
}

// SetReporterID sets the "reporter_id" field.
func (bruo *BoardReportUpdateOne) SetReporterID(i int) *BoardReportUpdateOne {
	bruo.mutation.SetReporterID(i)
	return bruo
}

// SetNillableReporterID sets the "reporter_id" field if the given value is not nil.
func (bruo *BoardReportUpdateOne) SetNillableReporterID(i *int) *BoardReportUpdateOne {
	if i != nil {
		bruo.SetReporterID(*i)
	}
	return bruo
}

// ClearReporterID clears the value of the "reporter_id" field.
func (bruo *BoardReportUpdateOne) ClearReporterID() *BoardReportUpdateOne {
	bruo.mutation.ClearReporterID()
	return bruo
}

// SetComment sets the "comment" field.
func (bruo *BoardReportUpdateOne) SetComment(s string) *BoardReportUpdateOne {
	bruo.mutation.SetComment(s)
	return bruo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (bruo *BoardReportUpdateOne) SetNillableComment(s *string) *BoardReportUpdateOne {
	if s != nil {
		bruo.SetComment(*s)
	}
	return bruo
}

// ClearComment clears the value of the "comment" field.
func (bruo *BoardReportUpdateOne) ClearComment() *BoardReportUpdateOne {
	bruo.mutation.ClearComment()
	return bruo
}

// SetStatus sets the "status" field.
func (bruo *BoardReportUpdateOne) SetStatus(b boardreport.Status) *BoardReportUpdateOne {
	bruo.mutation.SetStatus(b)
	return bruo
}

// SetCreatedAt sets the "created_at" field.
func (bruo *BoardReportUpdateOne) SetCreatedAt(t time.Time) *BoardReportUpdateOne {
	bruo.mutation.SetCreatedAt(t)
	return bruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bruo *BoardReportUpdateOne) SetNillableCreatedAt(t *time.Time) *BoardReportUpdateOne {
	if t != nil {
		bruo.SetCreatedAt(*t)
	}
	return bruo
}

// SetUpdatedAt sets the "updated_at" field.
func (bruo *BoardReportUpdateOne) SetUpdatedAt(t time.Time) *BoardReportUpdateOne {
	bruo.mutation.SetUpdatedAt(t)
	return bruo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bruo *BoardReportUpdateOne) SetUserID(id int) *BoardReportUpdateOne {
	bruo.mutation.SetUserID(id)
	return bruo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (bruo *BoardReportUpdateOne) SetNillableUserID(id *int) *BoardReportUpdateOne {
	if id != nil {
		bruo = bruo.SetUserID(*id)
	}
	return bruo
}

// SetUser sets the "user" edge to the User entity.
func (bruo *BoardReportUpdateOne) SetUser(u *User) *BoardReportUpdateOne {
	return bruo.SetUserID(u.ID)
}

// SetBoard sets the "board" edge to the Board entity.
func (bruo *BoardReportUpdateOne) SetBoard(b *Board) *BoardReportUpdateOne {
	return bruo.SetBoardID(b.ID)
}

// SetReportType sets the "report_type" edge to the ReportType entity.
func (bruo *BoardReportUpdateOne) SetReportType(r *ReportType) *BoardReportUpdateOne {
	return bruo.SetReportTypeID(r.ID)
}

// Mutation returns the BoardReportMutation object of the builder.
func (bruo *BoardReportUpdateOne) Mutation() *BoardReportMutation {
	return bruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (bruo *BoardReportUpdateOne) ClearUser() *BoardReportUpdateOne {
	bruo.mutation.ClearUser()
	return bruo
}

// ClearBoard clears the "board" edge to the Board entity.
func (bruo *BoardReportUpdateOne) ClearBoard() *BoardReportUpdateOne {
	bruo.mutation.ClearBoard()
	return bruo
}

// ClearReportType clears the "report_type" edge to the ReportType entity.
func (bruo *BoardReportUpdateOne) ClearReportType() *BoardReportUpdateOne {
	bruo.mutation.ClearReportType()
	return bruo
}

// Where appends a list predicates to the BoardReportUpdate builder.
func (bruo *BoardReportUpdateOne) Where(ps ...predicate.BoardReport) *BoardReportUpdateOne {
	bruo.mutation.Where(ps...)
	return bruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bruo *BoardReportUpdateOne) Select(field string, fields ...string) *BoardReportUpdateOne {
	bruo.fields = append([]string{field}, fields...)
	return bruo
}

// Save executes the query and returns the updated BoardReport entity.
func (bruo *BoardReportUpdateOne) Save(ctx context.Context) (*BoardReport, error) {
	bruo.defaults()
	return withHooks(ctx, bruo.sqlSave, bruo.mutation, bruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bruo *BoardReportUpdateOne) SaveX(ctx context.Context) *BoardReport {
	node, err := bruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bruo *BoardReportUpdateOne) Exec(ctx context.Context) error {
	_, err := bruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bruo *BoardReportUpdateOne) ExecX(ctx context.Context) {
	if err := bruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bruo *BoardReportUpdateOne) defaults() {
	if _, ok := bruo.mutation.UpdatedAt(); !ok {
		v := boardreport.UpdateDefaultUpdatedAt()
		bruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bruo *BoardReportUpdateOne) check() error {
	if v, ok := bruo.mutation.Status(); ok {
		if err := boardreport.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "BoardReport.status": %w`, err)}
		}
	}
	return nil
}

func (bruo *BoardReportUpdateOne) sqlSave(ctx context.Context) (_node *BoardReport, err error) {
	if err := bruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(boardreport.Table, boardreport.Columns, sqlgraph.NewFieldSpec(boardreport.FieldID, field.TypeInt))
	id, ok := bruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BoardReport.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, boardreport.FieldID)
		for _, f := range fields {
			if !boardreport.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != boardreport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bruo.mutation.Comment(); ok {
		_spec.SetField(boardreport.FieldComment, field.TypeString, value)
	}
	if bruo.mutation.CommentCleared() {
		_spec.ClearField(boardreport.FieldComment, field.TypeString)
	}
	if value, ok := bruo.mutation.Status(); ok {
		_spec.SetField(boardreport.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := bruo.mutation.CreatedAt(); ok {
		_spec.SetField(boardreport.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bruo.mutation.UpdatedAt(); ok {
		_spec.SetField(boardreport.FieldUpdatedAt, field.TypeTime, value)
	}
	if bruo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bruo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bruo.mutation.BoardCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bruo.mutation.BoardIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bruo.mutation.ReportTypeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bruo.mutation.ReportTypeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BoardReport{config: bruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{boardreport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bruo.mutation.done = true
	return _node, nil
}