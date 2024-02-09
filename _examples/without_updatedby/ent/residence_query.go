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
	"github.com/google/uuid"

	"github.com/flume/enthistory/_examples/without_updatedby/ent/character"
	"github.com/flume/enthistory/_examples/without_updatedby/ent/predicate"
	"github.com/flume/enthistory/_examples/without_updatedby/ent/residence"
)

// ResidenceQuery is the builder for querying Residence entities.
type ResidenceQuery struct {
	config
	ctx           *QueryContext
	order         []residence.OrderOption
	inters        []Interceptor
	predicates    []predicate.Residence
	withOccupants *CharacterQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ResidenceQuery builder.
func (rq *ResidenceQuery) Where(ps ...predicate.Residence) *ResidenceQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *ResidenceQuery) Limit(limit int) *ResidenceQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *ResidenceQuery) Offset(offset int) *ResidenceQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *ResidenceQuery) Unique(unique bool) *ResidenceQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *ResidenceQuery) Order(o ...residence.OrderOption) *ResidenceQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryOccupants chains the current query on the "occupants" edge.
func (rq *ResidenceQuery) QueryOccupants() *CharacterQuery {
	query := (&CharacterClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(residence.Table, residence.FieldID, selector),
			sqlgraph.To(character.Table, character.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, residence.OccupantsTable, residence.OccupantsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Residence entity from the query.
// Returns a *NotFoundError when no Residence was found.
func (rq *ResidenceQuery) First(ctx context.Context) (*Residence, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{residence.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ResidenceQuery) FirstX(ctx context.Context) *Residence {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Residence ID from the query.
// Returns a *NotFoundError when no Residence ID was found.
func (rq *ResidenceQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{residence.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *ResidenceQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Residence entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Residence entity is found.
// Returns a *NotFoundError when no Residence entities are found.
func (rq *ResidenceQuery) Only(ctx context.Context) (*Residence, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{residence.Label}
	default:
		return nil, &NotSingularError{residence.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ResidenceQuery) OnlyX(ctx context.Context) *Residence {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Residence ID in the query.
// Returns a *NotSingularError when more than one Residence ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *ResidenceQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{residence.Label}
	default:
		err = &NotSingularError{residence.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ResidenceQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Residences.
func (rq *ResidenceQuery) All(ctx context.Context) ([]*Residence, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Residence, *ResidenceQuery]()
	return withInterceptors[[]*Residence](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *ResidenceQuery) AllX(ctx context.Context) []*Residence {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Residence IDs.
func (rq *ResidenceQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(residence.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ResidenceQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ResidenceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*ResidenceQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ResidenceQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ResidenceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ResidenceQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ResidenceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ResidenceQuery) Clone() *ResidenceQuery {
	if rq == nil {
		return nil
	}
	return &ResidenceQuery{
		config:        rq.config,
		ctx:           rq.ctx.Clone(),
		order:         append([]residence.OrderOption{}, rq.order...),
		inters:        append([]Interceptor{}, rq.inters...),
		predicates:    append([]predicate.Residence{}, rq.predicates...),
		withOccupants: rq.withOccupants.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithOccupants tells the query-builder to eager-load the nodes that are connected to
// the "occupants" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ResidenceQuery) WithOccupants(opts ...func(*CharacterQuery)) *ResidenceQuery {
	query := (&CharacterClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withOccupants = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Residence.Query().
//		GroupBy(residence.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *ResidenceQuery) GroupBy(field string, fields ...string) *ResidenceGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ResidenceGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = residence.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Residence.Query().
//		Select(residence.FieldCreatedAt).
//		Scan(ctx, &v)
func (rq *ResidenceQuery) Select(fields ...string) *ResidenceSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &ResidenceSelect{ResidenceQuery: rq}
	sbuild.label = residence.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ResidenceSelect configured with the given aggregations.
func (rq *ResidenceQuery) Aggregate(fns ...AggregateFunc) *ResidenceSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *ResidenceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !residence.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ResidenceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Residence, error) {
	var (
		nodes       = []*Residence{}
		_spec       = rq.querySpec()
		loadedTypes = [1]bool{
			rq.withOccupants != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Residence).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Residence{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withOccupants; query != nil {
		if err := rq.loadOccupants(ctx, query, nodes,
			func(n *Residence) { n.Edges.Occupants = []*Character{} },
			func(n *Residence, e *Character) { n.Edges.Occupants = append(n.Edges.Occupants, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *ResidenceQuery) loadOccupants(ctx context.Context, query *CharacterQuery, nodes []*Residence, init func(*Residence), assign func(*Residence, *Character)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Residence)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Character(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(residence.OccupantsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.residence_occupants
		if fk == nil {
			return fmt.Errorf(`foreign-key "residence_occupants" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "residence_occupants" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *ResidenceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ResidenceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(residence.Table, residence.Columns, sqlgraph.NewFieldSpec(residence.FieldID, field.TypeUUID))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, residence.FieldID)
		for i := range fields {
			if fields[i] != residence.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *ResidenceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(residence.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = residence.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ResidenceGroupBy is the group-by builder for Residence entities.
type ResidenceGroupBy struct {
	selector
	build *ResidenceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ResidenceGroupBy) Aggregate(fns ...AggregateFunc) *ResidenceGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *ResidenceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ResidenceQuery, *ResidenceGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *ResidenceGroupBy) sqlScan(ctx context.Context, root *ResidenceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ResidenceSelect is the builder for selecting fields of Residence entities.
type ResidenceSelect struct {
	*ResidenceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *ResidenceSelect) Aggregate(fns ...AggregateFunc) *ResidenceSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ResidenceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ResidenceQuery, *ResidenceSelect](ctx, rs.ResidenceQuery, rs, rs.inters, v)
}

func (rs *ResidenceSelect) sqlScan(ctx context.Context, root *ResidenceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
