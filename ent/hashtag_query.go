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
	"github.com/leeeeeoy/go_study/ent/boardhashtag"
	"github.com/leeeeeoy/go_study/ent/hashtag"
	"github.com/leeeeeoy/go_study/ent/predicate"
)

// HashtagQuery is the builder for querying Hashtag entities.
type HashtagQuery struct {
	config
	ctx              *QueryContext
	order            []hashtag.OrderOption
	inters           []Interceptor
	predicates       []predicate.Hashtag
	withBoardHashtag *BoardHashtagQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HashtagQuery builder.
func (hq *HashtagQuery) Where(ps ...predicate.Hashtag) *HashtagQuery {
	hq.predicates = append(hq.predicates, ps...)
	return hq
}

// Limit the number of records to be returned by this query.
func (hq *HashtagQuery) Limit(limit int) *HashtagQuery {
	hq.ctx.Limit = &limit
	return hq
}

// Offset to start from.
func (hq *HashtagQuery) Offset(offset int) *HashtagQuery {
	hq.ctx.Offset = &offset
	return hq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hq *HashtagQuery) Unique(unique bool) *HashtagQuery {
	hq.ctx.Unique = &unique
	return hq
}

// Order specifies how the records should be ordered.
func (hq *HashtagQuery) Order(o ...hashtag.OrderOption) *HashtagQuery {
	hq.order = append(hq.order, o...)
	return hq
}

// QueryBoardHashtag chains the current query on the "board_hashtag" edge.
func (hq *HashtagQuery) QueryBoardHashtag() *BoardHashtagQuery {
	query := (&BoardHashtagClient{config: hq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hashtag.Table, hashtag.FieldID, selector),
			sqlgraph.To(boardhashtag.Table, boardhashtag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, hashtag.BoardHashtagTable, hashtag.BoardHashtagColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Hashtag entity from the query.
// Returns a *NotFoundError when no Hashtag was found.
func (hq *HashtagQuery) First(ctx context.Context) (*Hashtag, error) {
	nodes, err := hq.Limit(1).All(setContextOp(ctx, hq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{hashtag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hq *HashtagQuery) FirstX(ctx context.Context) *Hashtag {
	node, err := hq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Hashtag ID from the query.
// Returns a *NotFoundError when no Hashtag ID was found.
func (hq *HashtagQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(1).IDs(setContextOp(ctx, hq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{hashtag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hq *HashtagQuery) FirstIDX(ctx context.Context) int {
	id, err := hq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Hashtag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Hashtag entity is found.
// Returns a *NotFoundError when no Hashtag entities are found.
func (hq *HashtagQuery) Only(ctx context.Context) (*Hashtag, error) {
	nodes, err := hq.Limit(2).All(setContextOp(ctx, hq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{hashtag.Label}
	default:
		return nil, &NotSingularError{hashtag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hq *HashtagQuery) OnlyX(ctx context.Context) *Hashtag {
	node, err := hq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Hashtag ID in the query.
// Returns a *NotSingularError when more than one Hashtag ID is found.
// Returns a *NotFoundError when no entities are found.
func (hq *HashtagQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(2).IDs(setContextOp(ctx, hq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{hashtag.Label}
	default:
		err = &NotSingularError{hashtag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hq *HashtagQuery) OnlyIDX(ctx context.Context) int {
	id, err := hq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Hashtags.
func (hq *HashtagQuery) All(ctx context.Context) ([]*Hashtag, error) {
	ctx = setContextOp(ctx, hq.ctx, "All")
	if err := hq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Hashtag, *HashtagQuery]()
	return withInterceptors[[]*Hashtag](ctx, hq, qr, hq.inters)
}

// AllX is like All, but panics if an error occurs.
func (hq *HashtagQuery) AllX(ctx context.Context) []*Hashtag {
	nodes, err := hq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Hashtag IDs.
func (hq *HashtagQuery) IDs(ctx context.Context) (ids []int, err error) {
	if hq.ctx.Unique == nil && hq.path != nil {
		hq.Unique(true)
	}
	ctx = setContextOp(ctx, hq.ctx, "IDs")
	if err = hq.Select(hashtag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hq *HashtagQuery) IDsX(ctx context.Context) []int {
	ids, err := hq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hq *HashtagQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, hq.ctx, "Count")
	if err := hq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, hq, querierCount[*HashtagQuery](), hq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (hq *HashtagQuery) CountX(ctx context.Context) int {
	count, err := hq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hq *HashtagQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, hq.ctx, "Exist")
	switch _, err := hq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (hq *HashtagQuery) ExistX(ctx context.Context) bool {
	exist, err := hq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HashtagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hq *HashtagQuery) Clone() *HashtagQuery {
	if hq == nil {
		return nil
	}
	return &HashtagQuery{
		config:           hq.config,
		ctx:              hq.ctx.Clone(),
		order:            append([]hashtag.OrderOption{}, hq.order...),
		inters:           append([]Interceptor{}, hq.inters...),
		predicates:       append([]predicate.Hashtag{}, hq.predicates...),
		withBoardHashtag: hq.withBoardHashtag.Clone(),
		// clone intermediate query.
		sql:  hq.sql.Clone(),
		path: hq.path,
	}
}

// WithBoardHashtag tells the query-builder to eager-load the nodes that are connected to
// the "board_hashtag" edge. The optional arguments are used to configure the query builder of the edge.
func (hq *HashtagQuery) WithBoardHashtag(opts ...func(*BoardHashtagQuery)) *HashtagQuery {
	query := (&BoardHashtagClient{config: hq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hq.withBoardHashtag = query
	return hq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value string `json:"value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Hashtag.Query().
//		GroupBy(hashtag.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (hq *HashtagQuery) GroupBy(field string, fields ...string) *HashtagGroupBy {
	hq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HashtagGroupBy{build: hq}
	grbuild.flds = &hq.ctx.Fields
	grbuild.label = hashtag.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Value string `json:"value,omitempty"`
//	}
//
//	client.Hashtag.Query().
//		Select(hashtag.FieldValue).
//		Scan(ctx, &v)
func (hq *HashtagQuery) Select(fields ...string) *HashtagSelect {
	hq.ctx.Fields = append(hq.ctx.Fields, fields...)
	sbuild := &HashtagSelect{HashtagQuery: hq}
	sbuild.label = hashtag.Label
	sbuild.flds, sbuild.scan = &hq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HashtagSelect configured with the given aggregations.
func (hq *HashtagQuery) Aggregate(fns ...AggregateFunc) *HashtagSelect {
	return hq.Select().Aggregate(fns...)
}

func (hq *HashtagQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range hq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, hq); err != nil {
				return err
			}
		}
	}
	for _, f := range hq.ctx.Fields {
		if !hashtag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hq.path != nil {
		prev, err := hq.path(ctx)
		if err != nil {
			return err
		}
		hq.sql = prev
	}
	return nil
}

func (hq *HashtagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Hashtag, error) {
	var (
		nodes       = []*Hashtag{}
		_spec       = hq.querySpec()
		loadedTypes = [1]bool{
			hq.withBoardHashtag != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Hashtag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Hashtag{config: hq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := hq.withBoardHashtag; query != nil {
		if err := hq.loadBoardHashtag(ctx, query, nodes,
			func(n *Hashtag) { n.Edges.BoardHashtag = []*BoardHashtag{} },
			func(n *Hashtag, e *BoardHashtag) { n.Edges.BoardHashtag = append(n.Edges.BoardHashtag, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (hq *HashtagQuery) loadBoardHashtag(ctx context.Context, query *BoardHashtagQuery, nodes []*Hashtag, init func(*Hashtag), assign func(*Hashtag, *BoardHashtag)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Hashtag)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(boardhashtag.FieldHashtagID)
	}
	query.Where(predicate.BoardHashtag(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(hashtag.BoardHashtagColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.HashtagID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "hashtag_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (hq *HashtagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hq.querySpec()
	_spec.Node.Columns = hq.ctx.Fields
	if len(hq.ctx.Fields) > 0 {
		_spec.Unique = hq.ctx.Unique != nil && *hq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, hq.driver, _spec)
}

func (hq *HashtagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(hashtag.Table, hashtag.Columns, sqlgraph.NewFieldSpec(hashtag.FieldID, field.TypeInt))
	_spec.From = hq.sql
	if unique := hq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if hq.path != nil {
		_spec.Unique = true
	}
	if fields := hq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hashtag.FieldID)
		for i := range fields {
			if fields[i] != hashtag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hq *HashtagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hq.driver.Dialect())
	t1 := builder.Table(hashtag.Table)
	columns := hq.ctx.Fields
	if len(columns) == 0 {
		columns = hashtag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hq.sql != nil {
		selector = hq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hq.ctx.Unique != nil && *hq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range hq.predicates {
		p(selector)
	}
	for _, p := range hq.order {
		p(selector)
	}
	if offset := hq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HashtagGroupBy is the group-by builder for Hashtag entities.
type HashtagGroupBy struct {
	selector
	build *HashtagQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hgb *HashtagGroupBy) Aggregate(fns ...AggregateFunc) *HashtagGroupBy {
	hgb.fns = append(hgb.fns, fns...)
	return hgb
}

// Scan applies the selector query and scans the result into the given value.
func (hgb *HashtagGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hgb.build.ctx, "GroupBy")
	if err := hgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HashtagQuery, *HashtagGroupBy](ctx, hgb.build, hgb, hgb.build.inters, v)
}

func (hgb *HashtagGroupBy) sqlScan(ctx context.Context, root *HashtagQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hgb.fns))
	for _, fn := range hgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hgb.flds)+len(hgb.fns))
		for _, f := range *hgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HashtagSelect is the builder for selecting fields of Hashtag entities.
type HashtagSelect struct {
	*HashtagQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hs *HashtagSelect) Aggregate(fns ...AggregateFunc) *HashtagSelect {
	hs.fns = append(hs.fns, fns...)
	return hs
}

// Scan applies the selector query and scans the result into the given value.
func (hs *HashtagSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hs.ctx, "Select")
	if err := hs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HashtagQuery, *HashtagSelect](ctx, hs.HashtagQuery, hs, hs.inters, v)
}

func (hs *HashtagSelect) sqlScan(ctx context.Context, root *HashtagQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hs.fns))
	for _, fn := range hs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
