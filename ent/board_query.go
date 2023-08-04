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
	"github.com/leeeeeoy/go_study/ent/board"
	"github.com/leeeeeoy/go_study/ent/boardhashtag"
	"github.com/leeeeeoy/go_study/ent/boardlike"
	"github.com/leeeeeoy/go_study/ent/boardreport"
	"github.com/leeeeeoy/go_study/ent/comment"
	"github.com/leeeeeoy/go_study/ent/predicate"
	"github.com/leeeeeoy/go_study/ent/user"
)

// BoardQuery is the builder for querying Board entities.
type BoardQuery struct {
	config
	ctx              *QueryContext
	order            []board.OrderOption
	inters           []Interceptor
	predicates       []predicate.Board
	withUser         *UserQuery
	withComments     *CommentQuery
	withBoardLike    *BoardLikeQuery
	withBoardHashtag *BoardHashtagQuery
	withBoardReport  *BoardReportQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BoardQuery builder.
func (bq *BoardQuery) Where(ps ...predicate.Board) *BoardQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BoardQuery) Limit(limit int) *BoardQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BoardQuery) Offset(offset int) *BoardQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BoardQuery) Unique(unique bool) *BoardQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BoardQuery) Order(o ...board.OrderOption) *BoardQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryUser chains the current query on the "user" edge.
func (bq *BoardQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(board.Table, board.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, board.UserTable, board.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryComments chains the current query on the "comments" edge.
func (bq *BoardQuery) QueryComments() *CommentQuery {
	query := (&CommentClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(board.Table, board.FieldID, selector),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, board.CommentsTable, board.CommentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBoardLike chains the current query on the "board_like" edge.
func (bq *BoardQuery) QueryBoardLike() *BoardLikeQuery {
	query := (&BoardLikeClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(board.Table, board.FieldID, selector),
			sqlgraph.To(boardlike.Table, boardlike.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, board.BoardLikeTable, board.BoardLikeColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBoardHashtag chains the current query on the "board_hashtag" edge.
func (bq *BoardQuery) QueryBoardHashtag() *BoardHashtagQuery {
	query := (&BoardHashtagClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(board.Table, board.FieldID, selector),
			sqlgraph.To(boardhashtag.Table, boardhashtag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, board.BoardHashtagTable, board.BoardHashtagColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBoardReport chains the current query on the "board_report" edge.
func (bq *BoardQuery) QueryBoardReport() *BoardReportQuery {
	query := (&BoardReportClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(board.Table, board.FieldID, selector),
			sqlgraph.To(boardreport.Table, boardreport.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, board.BoardReportTable, board.BoardReportColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Board entity from the query.
// Returns a *NotFoundError when no Board was found.
func (bq *BoardQuery) First(ctx context.Context) (*Board, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{board.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BoardQuery) FirstX(ctx context.Context) *Board {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Board ID from the query.
// Returns a *NotFoundError when no Board ID was found.
func (bq *BoardQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{board.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BoardQuery) FirstIDX(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Board entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Board entity is found.
// Returns a *NotFoundError when no Board entities are found.
func (bq *BoardQuery) Only(ctx context.Context) (*Board, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{board.Label}
	default:
		return nil, &NotSingularError{board.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BoardQuery) OnlyX(ctx context.Context) *Board {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Board ID in the query.
// Returns a *NotSingularError when more than one Board ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BoardQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{board.Label}
	default:
		err = &NotSingularError{board.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BoardQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Boards.
func (bq *BoardQuery) All(ctx context.Context) ([]*Board, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Board, *BoardQuery]()
	return withInterceptors[[]*Board](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BoardQuery) AllX(ctx context.Context) []*Board {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Board IDs.
func (bq *BoardQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(board.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BoardQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BoardQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BoardQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BoardQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BoardQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BoardQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BoardQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BoardQuery) Clone() *BoardQuery {
	if bq == nil {
		return nil
	}
	return &BoardQuery{
		config:           bq.config,
		ctx:              bq.ctx.Clone(),
		order:            append([]board.OrderOption{}, bq.order...),
		inters:           append([]Interceptor{}, bq.inters...),
		predicates:       append([]predicate.Board{}, bq.predicates...),
		withUser:         bq.withUser.Clone(),
		withComments:     bq.withComments.Clone(),
		withBoardLike:    bq.withBoardLike.Clone(),
		withBoardHashtag: bq.withBoardHashtag.Clone(),
		withBoardReport:  bq.withBoardReport.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BoardQuery) WithUser(opts ...func(*UserQuery)) *BoardQuery {
	query := (&UserClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withUser = query
	return bq
}

// WithComments tells the query-builder to eager-load the nodes that are connected to
// the "comments" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BoardQuery) WithComments(opts ...func(*CommentQuery)) *BoardQuery {
	query := (&CommentClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withComments = query
	return bq
}

// WithBoardLike tells the query-builder to eager-load the nodes that are connected to
// the "board_like" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BoardQuery) WithBoardLike(opts ...func(*BoardLikeQuery)) *BoardQuery {
	query := (&BoardLikeClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withBoardLike = query
	return bq
}

// WithBoardHashtag tells the query-builder to eager-load the nodes that are connected to
// the "board_hashtag" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BoardQuery) WithBoardHashtag(opts ...func(*BoardHashtagQuery)) *BoardQuery {
	query := (&BoardHashtagClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withBoardHashtag = query
	return bq
}

// WithBoardReport tells the query-builder to eager-load the nodes that are connected to
// the "board_report" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BoardQuery) WithBoardReport(opts ...func(*BoardReportQuery)) *BoardQuery {
	query := (&BoardReportClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withBoardReport = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Board.Query().
//		GroupBy(board.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BoardQuery) GroupBy(field string, fields ...string) *BoardGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BoardGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = board.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Board.Query().
//		Select(board.FieldTitle).
//		Scan(ctx, &v)
func (bq *BoardQuery) Select(fields ...string) *BoardSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BoardSelect{BoardQuery: bq}
	sbuild.label = board.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BoardSelect configured with the given aggregations.
func (bq *BoardQuery) Aggregate(fns ...AggregateFunc) *BoardSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BoardQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !board.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BoardQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Board, error) {
	var (
		nodes       = []*Board{}
		_spec       = bq.querySpec()
		loadedTypes = [5]bool{
			bq.withUser != nil,
			bq.withComments != nil,
			bq.withBoardLike != nil,
			bq.withBoardHashtag != nil,
			bq.withBoardReport != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Board).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Board{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withUser; query != nil {
		if err := bq.loadUser(ctx, query, nodes, nil,
			func(n *Board, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := bq.withComments; query != nil {
		if err := bq.loadComments(ctx, query, nodes,
			func(n *Board) { n.Edges.Comments = []*Comment{} },
			func(n *Board, e *Comment) { n.Edges.Comments = append(n.Edges.Comments, e) }); err != nil {
			return nil, err
		}
	}
	if query := bq.withBoardLike; query != nil {
		if err := bq.loadBoardLike(ctx, query, nodes,
			func(n *Board) { n.Edges.BoardLike = []*BoardLike{} },
			func(n *Board, e *BoardLike) { n.Edges.BoardLike = append(n.Edges.BoardLike, e) }); err != nil {
			return nil, err
		}
	}
	if query := bq.withBoardHashtag; query != nil {
		if err := bq.loadBoardHashtag(ctx, query, nodes,
			func(n *Board) { n.Edges.BoardHashtag = []*BoardHashtag{} },
			func(n *Board, e *BoardHashtag) { n.Edges.BoardHashtag = append(n.Edges.BoardHashtag, e) }); err != nil {
			return nil, err
		}
	}
	if query := bq.withBoardReport; query != nil {
		if err := bq.loadBoardReport(ctx, query, nodes,
			func(n *Board) { n.Edges.BoardReport = []*BoardReport{} },
			func(n *Board, e *BoardReport) { n.Edges.BoardReport = append(n.Edges.BoardReport, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BoardQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Board, init func(*Board), assign func(*Board, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Board)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bq *BoardQuery) loadComments(ctx context.Context, query *CommentQuery, nodes []*Board, init func(*Board), assign func(*Board, *Comment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Board)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(comment.FieldBoardID)
	}
	query.Where(predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(board.CommentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BoardID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "board_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (bq *BoardQuery) loadBoardLike(ctx context.Context, query *BoardLikeQuery, nodes []*Board, init func(*Board), assign func(*Board, *BoardLike)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Board)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(boardlike.FieldBoardID)
	}
	query.Where(predicate.BoardLike(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(board.BoardLikeColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BoardID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "board_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (bq *BoardQuery) loadBoardHashtag(ctx context.Context, query *BoardHashtagQuery, nodes []*Board, init func(*Board), assign func(*Board, *BoardHashtag)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Board)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(boardhashtag.FieldBoardID)
	}
	query.Where(predicate.BoardHashtag(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(board.BoardHashtagColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BoardID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "board_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (bq *BoardQuery) loadBoardReport(ctx context.Context, query *BoardReportQuery, nodes []*Board, init func(*Board), assign func(*Board, *BoardReport)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Board)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(boardreport.FieldBoardID)
	}
	query.Where(predicate.BoardReport(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(board.BoardReportColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BoardID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "board_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bq *BoardQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BoardQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(board.Table, board.Columns, sqlgraph.NewFieldSpec(board.FieldID, field.TypeInt))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, board.FieldID)
		for i := range fields {
			if fields[i] != board.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if bq.withUser != nil {
			_spec.Node.AddColumnOnce(board.FieldUserID)
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BoardQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(board.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = board.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BoardGroupBy is the group-by builder for Board entities.
type BoardGroupBy struct {
	selector
	build *BoardQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BoardGroupBy) Aggregate(fns ...AggregateFunc) *BoardGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BoardGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BoardQuery, *BoardGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BoardGroupBy) sqlScan(ctx context.Context, root *BoardQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BoardSelect is the builder for selecting fields of Board entities.
type BoardSelect struct {
	*BoardQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BoardSelect) Aggregate(fns ...AggregateFunc) *BoardSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BoardSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BoardQuery, *BoardSelect](ctx, bs.BoardQuery, bs, bs.inters, v)
}

func (bs *BoardSelect) sqlScan(ctx context.Context, root *BoardQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
