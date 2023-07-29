// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/leeeeeoy/go_study/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/leeeeeoy/go_study/ent/board"
	"github.com/leeeeeoy/go_study/ent/boardlike"
	"github.com/leeeeeoy/go_study/ent/comment"
	"github.com/leeeeeoy/go_study/ent/commentlike"
	"github.com/leeeeeoy/go_study/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Board is the client for interacting with the Board builders.
	Board *BoardClient
	// BoardLike is the client for interacting with the BoardLike builders.
	BoardLike *BoardLikeClient
	// Comment is the client for interacting with the Comment builders.
	Comment *CommentClient
	// CommentLike is the client for interacting with the CommentLike builders.
	CommentLike *CommentLikeClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Board = NewBoardClient(c.config)
	c.BoardLike = NewBoardLikeClient(c.config)
	c.Comment = NewCommentClient(c.config)
	c.CommentLike = NewCommentLikeClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Board:       NewBoardClient(cfg),
		BoardLike:   NewBoardLikeClient(cfg),
		Comment:     NewCommentClient(cfg),
		CommentLike: NewCommentLikeClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Board:       NewBoardClient(cfg),
		BoardLike:   NewBoardLikeClient(cfg),
		Comment:     NewCommentClient(cfg),
		CommentLike: NewCommentLikeClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Board.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Board.Use(hooks...)
	c.BoardLike.Use(hooks...)
	c.Comment.Use(hooks...)
	c.CommentLike.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Board.Intercept(interceptors...)
	c.BoardLike.Intercept(interceptors...)
	c.Comment.Intercept(interceptors...)
	c.CommentLike.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *BoardMutation:
		return c.Board.mutate(ctx, m)
	case *BoardLikeMutation:
		return c.BoardLike.mutate(ctx, m)
	case *CommentMutation:
		return c.Comment.mutate(ctx, m)
	case *CommentLikeMutation:
		return c.CommentLike.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// BoardClient is a client for the Board schema.
type BoardClient struct {
	config
}

// NewBoardClient returns a client for the Board from the given config.
func NewBoardClient(c config) *BoardClient {
	return &BoardClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `board.Hooks(f(g(h())))`.
func (c *BoardClient) Use(hooks ...Hook) {
	c.hooks.Board = append(c.hooks.Board, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `board.Intercept(f(g(h())))`.
func (c *BoardClient) Intercept(interceptors ...Interceptor) {
	c.inters.Board = append(c.inters.Board, interceptors...)
}

// Create returns a builder for creating a Board entity.
func (c *BoardClient) Create() *BoardCreate {
	mutation := newBoardMutation(c.config, OpCreate)
	return &BoardCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Board entities.
func (c *BoardClient) CreateBulk(builders ...*BoardCreate) *BoardCreateBulk {
	return &BoardCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Board.
func (c *BoardClient) Update() *BoardUpdate {
	mutation := newBoardMutation(c.config, OpUpdate)
	return &BoardUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BoardClient) UpdateOne(b *Board) *BoardUpdateOne {
	mutation := newBoardMutation(c.config, OpUpdateOne, withBoard(b))
	return &BoardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BoardClient) UpdateOneID(id int) *BoardUpdateOne {
	mutation := newBoardMutation(c.config, OpUpdateOne, withBoardID(id))
	return &BoardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Board.
func (c *BoardClient) Delete() *BoardDelete {
	mutation := newBoardMutation(c.config, OpDelete)
	return &BoardDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BoardClient) DeleteOne(b *Board) *BoardDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BoardClient) DeleteOneID(id int) *BoardDeleteOne {
	builder := c.Delete().Where(board.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BoardDeleteOne{builder}
}

// Query returns a query builder for Board.
func (c *BoardClient) Query() *BoardQuery {
	return &BoardQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBoard},
		inters: c.Interceptors(),
	}
}

// Get returns a Board entity by its id.
func (c *BoardClient) Get(ctx context.Context, id int) (*Board, error) {
	return c.Query().Where(board.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BoardClient) GetX(ctx context.Context, id int) *Board {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Board.
func (c *BoardClient) QueryUser(b *Board) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(board.Table, board.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, board.UserTable, board.UserColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComments queries the comments edge of a Board.
func (c *BoardClient) QueryComments(b *Board) *CommentQuery {
	query := (&CommentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(board.Table, board.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, board.CommentsTable, board.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBoardLike queries the board_like edge of a Board.
func (c *BoardClient) QueryBoardLike(b *Board) *BoardLikeQuery {
	query := (&BoardLikeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(board.Table, board.FieldID, id),
			sqlgraph.To(boardlike.Table, boardlike.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, board.BoardLikeTable, board.BoardLikeColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BoardClient) Hooks() []Hook {
	return c.hooks.Board
}

// Interceptors returns the client interceptors.
func (c *BoardClient) Interceptors() []Interceptor {
	return c.inters.Board
}

func (c *BoardClient) mutate(ctx context.Context, m *BoardMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BoardCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BoardUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BoardUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BoardDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Board mutation op: %q", m.Op())
	}
}

// BoardLikeClient is a client for the BoardLike schema.
type BoardLikeClient struct {
	config
}

// NewBoardLikeClient returns a client for the BoardLike from the given config.
func NewBoardLikeClient(c config) *BoardLikeClient {
	return &BoardLikeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `boardlike.Hooks(f(g(h())))`.
func (c *BoardLikeClient) Use(hooks ...Hook) {
	c.hooks.BoardLike = append(c.hooks.BoardLike, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `boardlike.Intercept(f(g(h())))`.
func (c *BoardLikeClient) Intercept(interceptors ...Interceptor) {
	c.inters.BoardLike = append(c.inters.BoardLike, interceptors...)
}

// Create returns a builder for creating a BoardLike entity.
func (c *BoardLikeClient) Create() *BoardLikeCreate {
	mutation := newBoardLikeMutation(c.config, OpCreate)
	return &BoardLikeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BoardLike entities.
func (c *BoardLikeClient) CreateBulk(builders ...*BoardLikeCreate) *BoardLikeCreateBulk {
	return &BoardLikeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BoardLike.
func (c *BoardLikeClient) Update() *BoardLikeUpdate {
	mutation := newBoardLikeMutation(c.config, OpUpdate)
	return &BoardLikeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BoardLikeClient) UpdateOne(bl *BoardLike) *BoardLikeUpdateOne {
	mutation := newBoardLikeMutation(c.config, OpUpdateOne, withBoardLike(bl))
	return &BoardLikeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BoardLikeClient) UpdateOneID(id int) *BoardLikeUpdateOne {
	mutation := newBoardLikeMutation(c.config, OpUpdateOne, withBoardLikeID(id))
	return &BoardLikeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BoardLike.
func (c *BoardLikeClient) Delete() *BoardLikeDelete {
	mutation := newBoardLikeMutation(c.config, OpDelete)
	return &BoardLikeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BoardLikeClient) DeleteOne(bl *BoardLike) *BoardLikeDeleteOne {
	return c.DeleteOneID(bl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BoardLikeClient) DeleteOneID(id int) *BoardLikeDeleteOne {
	builder := c.Delete().Where(boardlike.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BoardLikeDeleteOne{builder}
}

// Query returns a query builder for BoardLike.
func (c *BoardLikeClient) Query() *BoardLikeQuery {
	return &BoardLikeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBoardLike},
		inters: c.Interceptors(),
	}
}

// Get returns a BoardLike entity by its id.
func (c *BoardLikeClient) Get(ctx context.Context, id int) (*BoardLike, error) {
	return c.Query().Where(boardlike.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BoardLikeClient) GetX(ctx context.Context, id int) *BoardLike {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a BoardLike.
func (c *BoardLikeClient) QueryUser(bl *BoardLike) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := bl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(boardlike.Table, boardlike.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, boardlike.UserTable, boardlike.UserColumn),
		)
		fromV = sqlgraph.Neighbors(bl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBoard queries the board edge of a BoardLike.
func (c *BoardLikeClient) QueryBoard(bl *BoardLike) *BoardQuery {
	query := (&BoardClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := bl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(boardlike.Table, boardlike.FieldID, id),
			sqlgraph.To(board.Table, board.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, boardlike.BoardTable, boardlike.BoardColumn),
		)
		fromV = sqlgraph.Neighbors(bl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BoardLikeClient) Hooks() []Hook {
	return c.hooks.BoardLike
}

// Interceptors returns the client interceptors.
func (c *BoardLikeClient) Interceptors() []Interceptor {
	return c.inters.BoardLike
}

func (c *BoardLikeClient) mutate(ctx context.Context, m *BoardLikeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BoardLikeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BoardLikeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BoardLikeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BoardLikeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown BoardLike mutation op: %q", m.Op())
	}
}

// CommentClient is a client for the Comment schema.
type CommentClient struct {
	config
}

// NewCommentClient returns a client for the Comment from the given config.
func NewCommentClient(c config) *CommentClient {
	return &CommentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `comment.Hooks(f(g(h())))`.
func (c *CommentClient) Use(hooks ...Hook) {
	c.hooks.Comment = append(c.hooks.Comment, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `comment.Intercept(f(g(h())))`.
func (c *CommentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Comment = append(c.inters.Comment, interceptors...)
}

// Create returns a builder for creating a Comment entity.
func (c *CommentClient) Create() *CommentCreate {
	mutation := newCommentMutation(c.config, OpCreate)
	return &CommentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Comment entities.
func (c *CommentClient) CreateBulk(builders ...*CommentCreate) *CommentCreateBulk {
	return &CommentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Comment.
func (c *CommentClient) Update() *CommentUpdate {
	mutation := newCommentMutation(c.config, OpUpdate)
	return &CommentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentClient) UpdateOne(co *Comment) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withComment(co))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentClient) UpdateOneID(id int) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withCommentID(id))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Comment.
func (c *CommentClient) Delete() *CommentDelete {
	mutation := newCommentMutation(c.config, OpDelete)
	return &CommentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CommentClient) DeleteOne(co *Comment) *CommentDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CommentClient) DeleteOneID(id int) *CommentDeleteOne {
	builder := c.Delete().Where(comment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommentDeleteOne{builder}
}

// Query returns a query builder for Comment.
func (c *CommentClient) Query() *CommentQuery {
	return &CommentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeComment},
		inters: c.Interceptors(),
	}
}

// Get returns a Comment entity by its id.
func (c *CommentClient) Get(ctx context.Context, id int) (*Comment, error) {
	return c.Query().Where(comment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentClient) GetX(ctx context.Context, id int) *Comment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBoard queries the board edge of a Comment.
func (c *CommentClient) QueryBoard(co *Comment) *BoardQuery {
	query := (&BoardClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(board.Table, board.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.BoardTable, comment.BoardColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Comment.
func (c *CommentClient) QueryUser(co *Comment) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.UserTable, comment.UserColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCommentLike queries the comment_like edge of a Comment.
func (c *CommentClient) QueryCommentLike(co *Comment) *CommentLikeQuery {
	query := (&CommentLikeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(commentlike.Table, commentlike.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, comment.CommentLikeTable, comment.CommentLikeColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CommentClient) Hooks() []Hook {
	return c.hooks.Comment
}

// Interceptors returns the client interceptors.
func (c *CommentClient) Interceptors() []Interceptor {
	return c.inters.Comment
}

func (c *CommentClient) mutate(ctx context.Context, m *CommentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CommentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CommentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CommentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Comment mutation op: %q", m.Op())
	}
}

// CommentLikeClient is a client for the CommentLike schema.
type CommentLikeClient struct {
	config
}

// NewCommentLikeClient returns a client for the CommentLike from the given config.
func NewCommentLikeClient(c config) *CommentLikeClient {
	return &CommentLikeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `commentlike.Hooks(f(g(h())))`.
func (c *CommentLikeClient) Use(hooks ...Hook) {
	c.hooks.CommentLike = append(c.hooks.CommentLike, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `commentlike.Intercept(f(g(h())))`.
func (c *CommentLikeClient) Intercept(interceptors ...Interceptor) {
	c.inters.CommentLike = append(c.inters.CommentLike, interceptors...)
}

// Create returns a builder for creating a CommentLike entity.
func (c *CommentLikeClient) Create() *CommentLikeCreate {
	mutation := newCommentLikeMutation(c.config, OpCreate)
	return &CommentLikeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CommentLike entities.
func (c *CommentLikeClient) CreateBulk(builders ...*CommentLikeCreate) *CommentLikeCreateBulk {
	return &CommentLikeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CommentLike.
func (c *CommentLikeClient) Update() *CommentLikeUpdate {
	mutation := newCommentLikeMutation(c.config, OpUpdate)
	return &CommentLikeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentLikeClient) UpdateOne(cl *CommentLike) *CommentLikeUpdateOne {
	mutation := newCommentLikeMutation(c.config, OpUpdateOne, withCommentLike(cl))
	return &CommentLikeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentLikeClient) UpdateOneID(id int) *CommentLikeUpdateOne {
	mutation := newCommentLikeMutation(c.config, OpUpdateOne, withCommentLikeID(id))
	return &CommentLikeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CommentLike.
func (c *CommentLikeClient) Delete() *CommentLikeDelete {
	mutation := newCommentLikeMutation(c.config, OpDelete)
	return &CommentLikeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CommentLikeClient) DeleteOne(cl *CommentLike) *CommentLikeDeleteOne {
	return c.DeleteOneID(cl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CommentLikeClient) DeleteOneID(id int) *CommentLikeDeleteOne {
	builder := c.Delete().Where(commentlike.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommentLikeDeleteOne{builder}
}

// Query returns a query builder for CommentLike.
func (c *CommentLikeClient) Query() *CommentLikeQuery {
	return &CommentLikeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCommentLike},
		inters: c.Interceptors(),
	}
}

// Get returns a CommentLike entity by its id.
func (c *CommentLikeClient) Get(ctx context.Context, id int) (*CommentLike, error) {
	return c.Query().Where(commentlike.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentLikeClient) GetX(ctx context.Context, id int) *CommentLike {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a CommentLike.
func (c *CommentLikeClient) QueryUser(cl *CommentLike) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(commentlike.Table, commentlike.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, commentlike.UserTable, commentlike.UserColumn),
		)
		fromV = sqlgraph.Neighbors(cl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComment queries the comment edge of a CommentLike.
func (c *CommentLikeClient) QueryComment(cl *CommentLike) *CommentQuery {
	query := (&CommentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(commentlike.Table, commentlike.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, commentlike.CommentTable, commentlike.CommentColumn),
		)
		fromV = sqlgraph.Neighbors(cl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CommentLikeClient) Hooks() []Hook {
	return c.hooks.CommentLike
}

// Interceptors returns the client interceptors.
func (c *CommentLikeClient) Interceptors() []Interceptor {
	return c.inters.CommentLike
}

func (c *CommentLikeClient) mutate(ctx context.Context, m *CommentLikeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CommentLikeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CommentLikeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CommentLikeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CommentLikeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CommentLike mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBoards queries the boards edge of a User.
func (c *UserClient) QueryBoards(u *User) *BoardQuery {
	query := (&BoardClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(board.Table, board.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.BoardsTable, user.BoardsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBoardLike queries the board_like edge of a User.
func (c *UserClient) QueryBoardLike(u *User) *BoardLikeQuery {
	query := (&BoardLikeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(boardlike.Table, boardlike.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.BoardLikeTable, user.BoardLikeColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCommentLike queries the comment_like edge of a User.
func (c *UserClient) QueryCommentLike(u *User) *CommentLikeQuery {
	query := (&CommentLikeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(commentlike.Table, commentlike.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CommentLikeTable, user.CommentLikeColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryComments queries the comments edge of a User.
func (c *UserClient) QueryComments(u *User) *CommentQuery {
	query := (&CommentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CommentsTable, user.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Board, BoardLike, Comment, CommentLike, User []ent.Hook
	}
	inters struct {
		Board, BoardLike, Comment, CommentLike, User []ent.Interceptor
	}
)