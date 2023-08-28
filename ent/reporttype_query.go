// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/leeeeeoy/go_study/ent/boardreport"
	"github.com/leeeeeoy/go_study/ent/commentreport"
	"github.com/leeeeeoy/go_study/ent/predicate"
	"github.com/leeeeeoy/go_study/ent/reporttype"
)

// ReportTypeQuery is the builder for querying ReportType entities.
type ReportTypeQuery struct {
	config
	ctx               *QueryContext
	order             []reporttype.OrderOption
	inters            []Interceptor
	predicates        []predicate.ReportType
	withCommentReport *CommentReportQuery
	withBoardReport   *BoardReportQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReportTypeQuery builder.
func (rtq *ReportTypeQuery) Where(ps ...predicate.ReportType) *ReportTypeQuery {
	rtq.predicates = append(rtq.predicates, ps...)
	return rtq
}

// Limit the number of records to be returned by this query.
func (rtq *ReportTypeQuery) Limit(limit int) *ReportTypeQuery {
	rtq.ctx.Limit = &limit
	return rtq
}

// Offset to start from.
func (rtq *ReportTypeQuery) Offset(offset int) *ReportTypeQuery {
	rtq.ctx.Offset = &offset
	return rtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rtq *ReportTypeQuery) Unique(unique bool) *ReportTypeQuery {
	rtq.ctx.Unique = &unique
	return rtq
}

// Order specifies how the records should be ordered.
func (rtq *ReportTypeQuery) Order(o ...reporttype.OrderOption) *ReportTypeQuery {
	rtq.order = append(rtq.order, o...)
	return rtq
}

// QueryCommentReport chains the current query on the "comment_report" edge.
func (rtq *ReportTypeQuery) QueryCommentReport() *CommentReportQuery {
	query := (&CommentReportClient{config: rtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reporttype.Table, reporttype.FieldID, selector),
			sqlgraph.To(commentreport.Table, commentreport.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, reporttype.CommentReportTable, reporttype.CommentReportColumn),
		)
		fromU = sqlgraph.SetNeighbors(rtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBoardReport chains the current query on the "board_report" edge.
func (rtq *ReportTypeQuery) QueryBoardReport() *BoardReportQuery {
	query := (&BoardReportClient{config: rtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reporttype.Table, reporttype.FieldID, selector),
			sqlgraph.To(boardreport.Table, boardreport.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, reporttype.BoardReportTable, reporttype.BoardReportColumn),
		)
		fromU = sqlgraph.SetNeighbors(rtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ReportType entity from the query.
// Returns a *NotFoundError when no ReportType was found.
func (rtq *ReportTypeQuery) First(ctx context.Context) (*ReportType, error) {
	nodes, err := rtq.Limit(1).All(setContextOp(ctx, rtq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{reporttype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rtq *ReportTypeQuery) FirstX(ctx context.Context) *ReportType {
	node, err := rtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ReportType ID from the query.
// Returns a *NotFoundError when no ReportType ID was found.
func (rtq *ReportTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rtq.Limit(1).IDs(setContextOp(ctx, rtq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{reporttype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rtq *ReportTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := rtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ReportType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ReportType entity is found.
// Returns a *NotFoundError when no ReportType entities are found.
func (rtq *ReportTypeQuery) Only(ctx context.Context) (*ReportType, error) {
	nodes, err := rtq.Limit(2).All(setContextOp(ctx, rtq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{reporttype.Label}
	default:
		return nil, &NotSingularError{reporttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rtq *ReportTypeQuery) OnlyX(ctx context.Context) *ReportType {
	node, err := rtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ReportType ID in the query.
// Returns a *NotSingularError when more than one ReportType ID is found.
// Returns a *NotFoundError when no entities are found.
func (rtq *ReportTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rtq.Limit(2).IDs(setContextOp(ctx, rtq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{reporttype.Label}
	default:
		err = &NotSingularError{reporttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rtq *ReportTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := rtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ReportTypes.
func (rtq *ReportTypeQuery) All(ctx context.Context) ([]*ReportType, error) {
	ctx = setContextOp(ctx, rtq.ctx, "All")
	if err := rtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ReportType, *ReportTypeQuery]()
	return withInterceptors[[]*ReportType](ctx, rtq, qr, rtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rtq *ReportTypeQuery) AllX(ctx context.Context) []*ReportType {
	nodes, err := rtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ReportType IDs.
func (rtq *ReportTypeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rtq.ctx.Unique == nil && rtq.path != nil {
		rtq.Unique(true)
	}
	ctx = setContextOp(ctx, rtq.ctx, "IDs")
	if err = rtq.Select(reporttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rtq *ReportTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := rtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rtq *ReportTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rtq.ctx, "Count")
	if err := rtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rtq, querierCount[*ReportTypeQuery](), rtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rtq *ReportTypeQuery) CountX(ctx context.Context) int {
	count, err := rtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rtq *ReportTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rtq.ctx, "Exist")
	switch _, err := rtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rtq *ReportTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := rtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReportTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rtq *ReportTypeQuery) Clone() *ReportTypeQuery {
	if rtq == nil {
		return nil
	}
	return &ReportTypeQuery{
		config:            rtq.config,
		ctx:               rtq.ctx.Clone(),
		order:             append([]reporttype.OrderOption{}, rtq.order...),
		inters:            append([]Interceptor{}, rtq.inters...),
		predicates:        append([]predicate.ReportType{}, rtq.predicates...),
		withCommentReport: rtq.withCommentReport.Clone(),
		withBoardReport:   rtq.withBoardReport.Clone(),
		// clone intermediate query.
		sql:  rtq.sql.Clone(),
		path: rtq.path,
	}
}

// WithCommentReport tells the query-builder to eager-load the nodes that are connected to
// the "comment_report" edge. The optional arguments are used to configure the query builder of the edge.
func (rtq *ReportTypeQuery) WithCommentReport(opts ...func(*CommentReportQuery)) *ReportTypeQuery {
	query := (&CommentReportClient{config: rtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rtq.withCommentReport = query
	return rtq
}

// WithBoardReport tells the query-builder to eager-load the nodes that are connected to
// the "board_report" edge. The optional arguments are used to configure the query builder of the edge.
func (rtq *ReportTypeQuery) WithBoardReport(opts ...func(*BoardReportQuery)) *ReportTypeQuery {
	query := (&BoardReportClient{config: rtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rtq.withBoardReport = query
	return rtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ReportType.Query().
//		GroupBy(reporttype.FieldDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rtq *ReportTypeQuery) GroupBy(field string, fields ...string) *ReportTypeGroupBy {
	rtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReportTypeGroupBy{build: rtq}
	grbuild.flds = &rtq.ctx.Fields
	grbuild.label = reporttype.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//	}
//
//	client.ReportType.Query().
//		Select(reporttype.FieldDescription).
//		Scan(ctx, &v)
func (rtq *ReportTypeQuery) Select(fields ...string) *ReportTypeSelect {
	rtq.ctx.Fields = append(rtq.ctx.Fields, fields...)
	sbuild := &ReportTypeSelect{ReportTypeQuery: rtq}
	sbuild.label = reporttype.Label
	sbuild.flds, sbuild.scan = &rtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReportTypeSelect configured with the given aggregations.
func (rtq *ReportTypeQuery) Aggregate(fns ...AggregateFunc) *ReportTypeSelect {
	return rtq.Select().Aggregate(fns...)
}

func (rtq *ReportTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rtq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rtq); err != nil {
				return err
			}
		}
	}
	for _, f := range rtq.ctx.Fields {
		if !reporttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rtq.path != nil {
		prev, err := rtq.path(ctx)
		if err != nil {
			return err
		}
		rtq.sql = prev
	}
	return nil
}

func (rtq *ReportTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ReportType, error) {
	var (
		nodes       = []*ReportType{}
		_spec       = rtq.querySpec()
		loadedTypes = [2]bool{
			rtq.withCommentReport != nil,
			rtq.withBoardReport != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ReportType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ReportType{config: rtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rtq.withCommentReport; query != nil {
		if err := rtq.loadCommentReport(ctx, query, nodes, nil,
			func(n *ReportType, e *CommentReport) { n.Edges.CommentReport = e }); err != nil {
			return nil, err
		}
	}
	if query := rtq.withBoardReport; query != nil {
		if err := rtq.loadBoardReport(ctx, query, nodes, nil,
			func(n *ReportType, e *BoardReport) { n.Edges.BoardReport = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rtq *ReportTypeQuery) loadCommentReport(ctx context.Context, query *CommentReportQuery, nodes []*ReportType, init func(*ReportType), assign func(*ReportType, *CommentReport)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ReportType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(commentreport.FieldReportTypeID)
	}
	query.Where(predicate.CommentReport(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(reporttype.CommentReportColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ReportTypeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "report_type_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rtq *ReportTypeQuery) loadBoardReport(ctx context.Context, query *BoardReportQuery, nodes []*ReportType, init func(*ReportType), assign func(*ReportType, *BoardReport)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ReportType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(boardreport.FieldReportTypeID)
	}
	query.Where(predicate.BoardReport(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(reporttype.BoardReportColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ReportTypeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "report_type_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rtq *ReportTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rtq.querySpec()
	_spec.Node.Columns = rtq.ctx.Fields
	if len(rtq.ctx.Fields) > 0 {
		_spec.Unique = rtq.ctx.Unique != nil && *rtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rtq.driver, _spec)
}

func (rtq *ReportTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(reporttype.Table, reporttype.Columns, sqlgraph.NewFieldSpec(reporttype.FieldID, field.TypeInt))
	_spec.From = rtq.sql
	if unique := rtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rtq.path != nil {
		_spec.Unique = true
	}
	if fields := rtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reporttype.FieldID)
		for i := range fields {
			if fields[i] != reporttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rtq *ReportTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rtq.driver.Dialect())
	t1 := builder.Table(reporttype.Table)
	columns := rtq.ctx.Fields
	if len(columns) == 0 {
		columns = reporttype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rtq.sql != nil {
		selector = rtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rtq.ctx.Unique != nil && *rtq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rtq.predicates {
		p(selector)
	}
	for _, p := range rtq.order {
		p(selector)
	}
	if offset := rtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReportTypeGroupBy is the group-by builder for ReportType entities.
type ReportTypeGroupBy struct {
	selector
	build *ReportTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rtgb *ReportTypeGroupBy) Aggregate(fns ...AggregateFunc) *ReportTypeGroupBy {
	rtgb.fns = append(rtgb.fns, fns...)
	return rtgb
}

// Scan applies the selector query and scans the result into the given value.
func (rtgb *ReportTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rtgb.build.ctx, "GroupBy")
	if err := rtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReportTypeQuery, *ReportTypeGroupBy](ctx, rtgb.build, rtgb, rtgb.build.inters, v)
}

func (rtgb *ReportTypeGroupBy) sqlScan(ctx context.Context, root *ReportTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rtgb.fns))
	for _, fn := range rtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rtgb.flds)+len(rtgb.fns))
		for _, f := range *rtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReportTypeSelect is the builder for selecting fields of ReportType entities.
type ReportTypeSelect struct {
	*ReportTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rts *ReportTypeSelect) Aggregate(fns ...AggregateFunc) *ReportTypeSelect {
	rts.fns = append(rts.fns, fns...)
	return rts
}

// Scan applies the selector query and scans the result into the given value.
func (rts *ReportTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rts.ctx, "Select")
	if err := rts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReportTypeQuery, *ReportTypeSelect](ctx, rts.ReportTypeQuery, rts, rts.inters, v)
}

func (rts *ReportTypeSelect) sqlScan(ctx context.Context, root *ReportTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rts.fns))
	for _, fn := range rts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}