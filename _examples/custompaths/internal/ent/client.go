// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/flume/enthistory"
	"github.com/flume/enthistory/_examples/custompaths/internal/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/flume/enthistory/_examples/custompaths/internal/ent/character"
	"github.com/flume/enthistory/_examples/custompaths/internal/ent/characterhistory"
	"github.com/flume/enthistory/_examples/custompaths/internal/ent/friendship"
	"github.com/flume/enthistory/_examples/custompaths/internal/ent/friendshiphistory"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Character is the client for interacting with the Character builders.
	Character *CharacterClient
	// CharacterHistory is the client for interacting with the CharacterHistory builders.
	CharacterHistory *CharacterHistoryClient
	// Friendship is the client for interacting with the Friendship builders.
	Friendship *FriendshipClient
	// FriendshipHistory is the client for interacting with the FriendshipHistory builders.
	FriendshipHistory *FriendshipHistoryClient
	// historyActivated determines if the history hooks have already been activated
	historyActivated bool
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
	c.Character = NewCharacterClient(c.config)
	c.CharacterHistory = NewCharacterHistoryClient(c.config)
	c.Friendship = NewFriendshipClient(c.config)
	c.FriendshipHistory = NewFriendshipHistoryClient(c.config)
}

// withHistory adds the history hooks to the appropriate schemas - generated by enthistory
func (c *Client) WithHistory() {
	if !c.historyActivated {
		for _, hook := range enthistory.HistoryHooks[*CharacterMutation]() {
			c.Character.Use(hook)
		}
		for _, hook := range enthistory.HistoryHooks[*FriendshipMutation]() {
			c.Friendship.Use(hook)
		}
		c.historyActivated = true
	}
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

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		Character:         NewCharacterClient(cfg),
		CharacterHistory:  NewCharacterHistoryClient(cfg),
		Friendship:        NewFriendshipClient(cfg),
		FriendshipHistory: NewFriendshipHistoryClient(cfg),
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
		ctx:               ctx,
		config:            cfg,
		Character:         NewCharacterClient(cfg),
		CharacterHistory:  NewCharacterHistoryClient(cfg),
		Friendship:        NewFriendshipClient(cfg),
		FriendshipHistory: NewFriendshipHistoryClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Character.
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
	c.Character.Use(hooks...)
	c.CharacterHistory.Use(hooks...)
	c.Friendship.Use(hooks...)
	c.FriendshipHistory.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Character.Intercept(interceptors...)
	c.CharacterHistory.Intercept(interceptors...)
	c.Friendship.Intercept(interceptors...)
	c.FriendshipHistory.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CharacterMutation:
		return c.Character.mutate(ctx, m)
	case *CharacterHistoryMutation:
		return c.CharacterHistory.mutate(ctx, m)
	case *FriendshipMutation:
		return c.Friendship.mutate(ctx, m)
	case *FriendshipHistoryMutation:
		return c.FriendshipHistory.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CharacterClient is a client for the Character schema.
type CharacterClient struct {
	config
}

// NewCharacterClient returns a client for the Character from the given config.
func NewCharacterClient(c config) *CharacterClient {
	return &CharacterClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `character.Hooks(f(g(h())))`.
func (c *CharacterClient) Use(hooks ...Hook) {
	c.hooks.Character = append(c.hooks.Character, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `character.Intercept(f(g(h())))`.
func (c *CharacterClient) Intercept(interceptors ...Interceptor) {
	c.inters.Character = append(c.inters.Character, interceptors...)
}

// Create returns a builder for creating a Character entity.
func (c *CharacterClient) Create() *CharacterCreate {
	mutation := newCharacterMutation(c.config, OpCreate)
	return &CharacterCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Character entities.
func (c *CharacterClient) CreateBulk(builders ...*CharacterCreate) *CharacterCreateBulk {
	return &CharacterCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CharacterClient) MapCreateBulk(slice any, setFunc func(*CharacterCreate, int)) *CharacterCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CharacterCreateBulk{err: fmt.Errorf("calling to CharacterClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CharacterCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CharacterCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Character.
func (c *CharacterClient) Update() *CharacterUpdate {
	mutation := newCharacterMutation(c.config, OpUpdate)
	return &CharacterUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CharacterClient) UpdateOne(ch *Character) *CharacterUpdateOne {
	mutation := newCharacterMutation(c.config, OpUpdateOne, withCharacter(ch))
	return &CharacterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CharacterClient) UpdateOneID(id int) *CharacterUpdateOne {
	mutation := newCharacterMutation(c.config, OpUpdateOne, withCharacterID(id))
	return &CharacterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Character.
func (c *CharacterClient) Delete() *CharacterDelete {
	mutation := newCharacterMutation(c.config, OpDelete)
	return &CharacterDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CharacterClient) DeleteOne(ch *Character) *CharacterDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CharacterClient) DeleteOneID(id int) *CharacterDeleteOne {
	builder := c.Delete().Where(character.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CharacterDeleteOne{builder}
}

// Query returns a query builder for Character.
func (c *CharacterClient) Query() *CharacterQuery {
	return &CharacterQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCharacter},
		inters: c.Interceptors(),
	}
}

// Get returns a Character entity by its id.
func (c *CharacterClient) Get(ctx context.Context, id int) (*Character, error) {
	return c.Query().Where(character.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CharacterClient) GetX(ctx context.Context, id int) *Character {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFriends queries the friends edge of a Character.
func (c *CharacterClient) QueryFriends(ch *Character) *CharacterQuery {
	query := (&CharacterClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(character.Table, character.FieldID, id),
			sqlgraph.To(character.Table, character.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, character.FriendsTable, character.FriendsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFriendships queries the friendships edge of a Character.
func (c *CharacterClient) QueryFriendships(ch *Character) *FriendshipQuery {
	query := (&FriendshipClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ch.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(character.Table, character.FieldID, id),
			sqlgraph.To(friendship.Table, friendship.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, character.FriendshipsTable, character.FriendshipsColumn),
		)
		fromV = sqlgraph.Neighbors(ch.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CharacterClient) Hooks() []Hook {
	return c.hooks.Character
}

// Interceptors returns the client interceptors.
func (c *CharacterClient) Interceptors() []Interceptor {
	return c.inters.Character
}

func (c *CharacterClient) mutate(ctx context.Context, m *CharacterMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CharacterCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CharacterUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CharacterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CharacterDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Character mutation op: %q", m.Op())
	}
}

// CharacterHistoryClient is a client for the CharacterHistory schema.
type CharacterHistoryClient struct {
	config
}

// NewCharacterHistoryClient returns a client for the CharacterHistory from the given config.
func NewCharacterHistoryClient(c config) *CharacterHistoryClient {
	return &CharacterHistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `characterhistory.Hooks(f(g(h())))`.
func (c *CharacterHistoryClient) Use(hooks ...Hook) {
	c.hooks.CharacterHistory = append(c.hooks.CharacterHistory, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `characterhistory.Intercept(f(g(h())))`.
func (c *CharacterHistoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.CharacterHistory = append(c.inters.CharacterHistory, interceptors...)
}

// Create returns a builder for creating a CharacterHistory entity.
func (c *CharacterHistoryClient) Create() *CharacterHistoryCreate {
	mutation := newCharacterHistoryMutation(c.config, OpCreate)
	return &CharacterHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CharacterHistory entities.
func (c *CharacterHistoryClient) CreateBulk(builders ...*CharacterHistoryCreate) *CharacterHistoryCreateBulk {
	return &CharacterHistoryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CharacterHistoryClient) MapCreateBulk(slice any, setFunc func(*CharacterHistoryCreate, int)) *CharacterHistoryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CharacterHistoryCreateBulk{err: fmt.Errorf("calling to CharacterHistoryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CharacterHistoryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CharacterHistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CharacterHistory.
func (c *CharacterHistoryClient) Update() *CharacterHistoryUpdate {
	mutation := newCharacterHistoryMutation(c.config, OpUpdate)
	return &CharacterHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CharacterHistoryClient) UpdateOne(ch *CharacterHistory) *CharacterHistoryUpdateOne {
	mutation := newCharacterHistoryMutation(c.config, OpUpdateOne, withCharacterHistory(ch))
	return &CharacterHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CharacterHistoryClient) UpdateOneID(id int) *CharacterHistoryUpdateOne {
	mutation := newCharacterHistoryMutation(c.config, OpUpdateOne, withCharacterHistoryID(id))
	return &CharacterHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CharacterHistory.
func (c *CharacterHistoryClient) Delete() *CharacterHistoryDelete {
	mutation := newCharacterHistoryMutation(c.config, OpDelete)
	return &CharacterHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CharacterHistoryClient) DeleteOne(ch *CharacterHistory) *CharacterHistoryDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CharacterHistoryClient) DeleteOneID(id int) *CharacterHistoryDeleteOne {
	builder := c.Delete().Where(characterhistory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CharacterHistoryDeleteOne{builder}
}

// Query returns a query builder for CharacterHistory.
func (c *CharacterHistoryClient) Query() *CharacterHistoryQuery {
	return &CharacterHistoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCharacterHistory},
		inters: c.Interceptors(),
	}
}

// Get returns a CharacterHistory entity by its id.
func (c *CharacterHistoryClient) Get(ctx context.Context, id int) (*CharacterHistory, error) {
	return c.Query().Where(characterhistory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CharacterHistoryClient) GetX(ctx context.Context, id int) *CharacterHistory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CharacterHistoryClient) Hooks() []Hook {
	return c.hooks.CharacterHistory
}

// Interceptors returns the client interceptors.
func (c *CharacterHistoryClient) Interceptors() []Interceptor {
	return c.inters.CharacterHistory
}

func (c *CharacterHistoryClient) mutate(ctx context.Context, m *CharacterHistoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CharacterHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CharacterHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CharacterHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CharacterHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CharacterHistory mutation op: %q", m.Op())
	}
}

// FriendshipClient is a client for the Friendship schema.
type FriendshipClient struct {
	config
}

// NewFriendshipClient returns a client for the Friendship from the given config.
func NewFriendshipClient(c config) *FriendshipClient {
	return &FriendshipClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `friendship.Hooks(f(g(h())))`.
func (c *FriendshipClient) Use(hooks ...Hook) {
	c.hooks.Friendship = append(c.hooks.Friendship, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `friendship.Intercept(f(g(h())))`.
func (c *FriendshipClient) Intercept(interceptors ...Interceptor) {
	c.inters.Friendship = append(c.inters.Friendship, interceptors...)
}

// Create returns a builder for creating a Friendship entity.
func (c *FriendshipClient) Create() *FriendshipCreate {
	mutation := newFriendshipMutation(c.config, OpCreate)
	return &FriendshipCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Friendship entities.
func (c *FriendshipClient) CreateBulk(builders ...*FriendshipCreate) *FriendshipCreateBulk {
	return &FriendshipCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *FriendshipClient) MapCreateBulk(slice any, setFunc func(*FriendshipCreate, int)) *FriendshipCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &FriendshipCreateBulk{err: fmt.Errorf("calling to FriendshipClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*FriendshipCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &FriendshipCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Friendship.
func (c *FriendshipClient) Update() *FriendshipUpdate {
	mutation := newFriendshipMutation(c.config, OpUpdate)
	return &FriendshipUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FriendshipClient) UpdateOne(f *Friendship) *FriendshipUpdateOne {
	mutation := newFriendshipMutation(c.config, OpUpdateOne, withFriendship(f))
	return &FriendshipUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FriendshipClient) UpdateOneID(id int) *FriendshipUpdateOne {
	mutation := newFriendshipMutation(c.config, OpUpdateOne, withFriendshipID(id))
	return &FriendshipUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Friendship.
func (c *FriendshipClient) Delete() *FriendshipDelete {
	mutation := newFriendshipMutation(c.config, OpDelete)
	return &FriendshipDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FriendshipClient) DeleteOne(f *Friendship) *FriendshipDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FriendshipClient) DeleteOneID(id int) *FriendshipDeleteOne {
	builder := c.Delete().Where(friendship.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FriendshipDeleteOne{builder}
}

// Query returns a query builder for Friendship.
func (c *FriendshipClient) Query() *FriendshipQuery {
	return &FriendshipQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFriendship},
		inters: c.Interceptors(),
	}
}

// Get returns a Friendship entity by its id.
func (c *FriendshipClient) Get(ctx context.Context, id int) (*Friendship, error) {
	return c.Query().Where(friendship.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FriendshipClient) GetX(ctx context.Context, id int) *Friendship {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCharacter queries the character edge of a Friendship.
func (c *FriendshipClient) QueryCharacter(f *Friendship) *CharacterQuery {
	query := (&CharacterClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(friendship.Table, friendship.FieldID, id),
			sqlgraph.To(character.Table, character.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, friendship.CharacterTable, friendship.CharacterColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFriend queries the friend edge of a Friendship.
func (c *FriendshipClient) QueryFriend(f *Friendship) *CharacterQuery {
	query := (&CharacterClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(friendship.Table, friendship.FieldID, id),
			sqlgraph.To(character.Table, character.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, friendship.FriendTable, friendship.FriendColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FriendshipClient) Hooks() []Hook {
	return c.hooks.Friendship
}

// Interceptors returns the client interceptors.
func (c *FriendshipClient) Interceptors() []Interceptor {
	return c.inters.Friendship
}

func (c *FriendshipClient) mutate(ctx context.Context, m *FriendshipMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FriendshipCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FriendshipUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FriendshipUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FriendshipDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Friendship mutation op: %q", m.Op())
	}
}

// FriendshipHistoryClient is a client for the FriendshipHistory schema.
type FriendshipHistoryClient struct {
	config
}

// NewFriendshipHistoryClient returns a client for the FriendshipHistory from the given config.
func NewFriendshipHistoryClient(c config) *FriendshipHistoryClient {
	return &FriendshipHistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `friendshiphistory.Hooks(f(g(h())))`.
func (c *FriendshipHistoryClient) Use(hooks ...Hook) {
	c.hooks.FriendshipHistory = append(c.hooks.FriendshipHistory, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `friendshiphistory.Intercept(f(g(h())))`.
func (c *FriendshipHistoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.FriendshipHistory = append(c.inters.FriendshipHistory, interceptors...)
}

// Create returns a builder for creating a FriendshipHistory entity.
func (c *FriendshipHistoryClient) Create() *FriendshipHistoryCreate {
	mutation := newFriendshipHistoryMutation(c.config, OpCreate)
	return &FriendshipHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FriendshipHistory entities.
func (c *FriendshipHistoryClient) CreateBulk(builders ...*FriendshipHistoryCreate) *FriendshipHistoryCreateBulk {
	return &FriendshipHistoryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *FriendshipHistoryClient) MapCreateBulk(slice any, setFunc func(*FriendshipHistoryCreate, int)) *FriendshipHistoryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &FriendshipHistoryCreateBulk{err: fmt.Errorf("calling to FriendshipHistoryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*FriendshipHistoryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &FriendshipHistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FriendshipHistory.
func (c *FriendshipHistoryClient) Update() *FriendshipHistoryUpdate {
	mutation := newFriendshipHistoryMutation(c.config, OpUpdate)
	return &FriendshipHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FriendshipHistoryClient) UpdateOne(fh *FriendshipHistory) *FriendshipHistoryUpdateOne {
	mutation := newFriendshipHistoryMutation(c.config, OpUpdateOne, withFriendshipHistory(fh))
	return &FriendshipHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FriendshipHistoryClient) UpdateOneID(id int) *FriendshipHistoryUpdateOne {
	mutation := newFriendshipHistoryMutation(c.config, OpUpdateOne, withFriendshipHistoryID(id))
	return &FriendshipHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FriendshipHistory.
func (c *FriendshipHistoryClient) Delete() *FriendshipHistoryDelete {
	mutation := newFriendshipHistoryMutation(c.config, OpDelete)
	return &FriendshipHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FriendshipHistoryClient) DeleteOne(fh *FriendshipHistory) *FriendshipHistoryDeleteOne {
	return c.DeleteOneID(fh.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FriendshipHistoryClient) DeleteOneID(id int) *FriendshipHistoryDeleteOne {
	builder := c.Delete().Where(friendshiphistory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FriendshipHistoryDeleteOne{builder}
}

// Query returns a query builder for FriendshipHistory.
func (c *FriendshipHistoryClient) Query() *FriendshipHistoryQuery {
	return &FriendshipHistoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFriendshipHistory},
		inters: c.Interceptors(),
	}
}

// Get returns a FriendshipHistory entity by its id.
func (c *FriendshipHistoryClient) Get(ctx context.Context, id int) (*FriendshipHistory, error) {
	return c.Query().Where(friendshiphistory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FriendshipHistoryClient) GetX(ctx context.Context, id int) *FriendshipHistory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FriendshipHistoryClient) Hooks() []Hook {
	return c.hooks.FriendshipHistory
}

// Interceptors returns the client interceptors.
func (c *FriendshipHistoryClient) Interceptors() []Interceptor {
	return c.inters.FriendshipHistory
}

func (c *FriendshipHistoryClient) mutate(ctx context.Context, m *FriendshipHistoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FriendshipHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FriendshipHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FriendshipHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FriendshipHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown FriendshipHistory mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Character, CharacterHistory, Friendship, FriendshipHistory []ent.Hook
	}
	inters struct {
		Character, CharacterHistory, Friendship, FriendshipHistory []ent.Interceptor
	}
)
