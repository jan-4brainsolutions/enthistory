// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/basic/ent/characterhistory"
	"_examples/basic/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CharacterHistoryQuery is the builder for querying CharacterHistory entities.
type CharacterHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []characterhistory.OrderOption
	inters     []Interceptor
	predicates []predicate.CharacterHistory
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CharacterHistoryQuery builder.
func (chq *CharacterHistoryQuery) Where(ps ...predicate.CharacterHistory) *CharacterHistoryQuery {
	chq.predicates = append(chq.predicates, ps...)
	return chq
}

// Limit the number of records to be returned by this query.
func (chq *CharacterHistoryQuery) Limit(limit int) *CharacterHistoryQuery {
	chq.ctx.Limit = &limit
	return chq
}

// Offset to start from.
func (chq *CharacterHistoryQuery) Offset(offset int) *CharacterHistoryQuery {
	chq.ctx.Offset = &offset
	return chq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (chq *CharacterHistoryQuery) Unique(unique bool) *CharacterHistoryQuery {
	chq.ctx.Unique = &unique
	return chq
}

// Order specifies how the records should be ordered.
func (chq *CharacterHistoryQuery) Order(o ...characterhistory.OrderOption) *CharacterHistoryQuery {
	chq.order = append(chq.order, o...)
	return chq
}

// First returns the first CharacterHistory entity from the query.
// Returns a *NotFoundError when no CharacterHistory was found.
func (chq *CharacterHistoryQuery) First(ctx context.Context) (*CharacterHistory, error) {
	nodes, err := chq.Limit(1).All(setContextOp(ctx, chq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{characterhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (chq *CharacterHistoryQuery) FirstX(ctx context.Context) *CharacterHistory {
	node, err := chq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CharacterHistory ID from the query.
// Returns a *NotFoundError when no CharacterHistory ID was found.
func (chq *CharacterHistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = chq.Limit(1).IDs(setContextOp(ctx, chq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{characterhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (chq *CharacterHistoryQuery) FirstIDX(ctx context.Context) int {
	id, err := chq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CharacterHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CharacterHistory entity is found.
// Returns a *NotFoundError when no CharacterHistory entities are found.
func (chq *CharacterHistoryQuery) Only(ctx context.Context) (*CharacterHistory, error) {
	nodes, err := chq.Limit(2).All(setContextOp(ctx, chq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{characterhistory.Label}
	default:
		return nil, &NotSingularError{characterhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (chq *CharacterHistoryQuery) OnlyX(ctx context.Context) *CharacterHistory {
	node, err := chq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CharacterHistory ID in the query.
// Returns a *NotSingularError when more than one CharacterHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (chq *CharacterHistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = chq.Limit(2).IDs(setContextOp(ctx, chq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{characterhistory.Label}
	default:
		err = &NotSingularError{characterhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (chq *CharacterHistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := chq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CharacterHistories.
func (chq *CharacterHistoryQuery) All(ctx context.Context) ([]*CharacterHistory, error) {
	ctx = setContextOp(ctx, chq.ctx, "All")
	if err := chq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CharacterHistory, *CharacterHistoryQuery]()
	return withInterceptors[[]*CharacterHistory](ctx, chq, qr, chq.inters)
}

// AllX is like All, but panics if an error occurs.
func (chq *CharacterHistoryQuery) AllX(ctx context.Context) []*CharacterHistory {
	nodes, err := chq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CharacterHistory IDs.
func (chq *CharacterHistoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if chq.ctx.Unique == nil && chq.path != nil {
		chq.Unique(true)
	}
	ctx = setContextOp(ctx, chq.ctx, "IDs")
	if err = chq.Select(characterhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (chq *CharacterHistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := chq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (chq *CharacterHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, chq.ctx, "Count")
	if err := chq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, chq, querierCount[*CharacterHistoryQuery](), chq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (chq *CharacterHistoryQuery) CountX(ctx context.Context) int {
	count, err := chq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (chq *CharacterHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, chq.ctx, "Exist")
	switch _, err := chq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (chq *CharacterHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := chq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CharacterHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (chq *CharacterHistoryQuery) Clone() *CharacterHistoryQuery {
	if chq == nil {
		return nil
	}
	return &CharacterHistoryQuery{
		config:     chq.config,
		ctx:        chq.ctx.Clone(),
		order:      append([]characterhistory.OrderOption{}, chq.order...),
		inters:     append([]Interceptor{}, chq.inters...),
		predicates: append([]predicate.CharacterHistory{}, chq.predicates...),
		// clone intermediate query.
		sql:  chq.sql.Clone(),
		path: chq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CharacterHistory.Query().
//		GroupBy(characterhistory.FieldHistoryTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (chq *CharacterHistoryQuery) GroupBy(field string, fields ...string) *CharacterHistoryGroupBy {
	chq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CharacterHistoryGroupBy{build: chq}
	grbuild.flds = &chq.ctx.Fields
	grbuild.label = characterhistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//	}
//
//	client.CharacterHistory.Query().
//		Select(characterhistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (chq *CharacterHistoryQuery) Select(fields ...string) *CharacterHistorySelect {
	chq.ctx.Fields = append(chq.ctx.Fields, fields...)
	sbuild := &CharacterHistorySelect{CharacterHistoryQuery: chq}
	sbuild.label = characterhistory.Label
	sbuild.flds, sbuild.scan = &chq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CharacterHistorySelect configured with the given aggregations.
func (chq *CharacterHistoryQuery) Aggregate(fns ...AggregateFunc) *CharacterHistorySelect {
	return chq.Select().Aggregate(fns...)
}

func (chq *CharacterHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range chq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, chq); err != nil {
				return err
			}
		}
	}
	for _, f := range chq.ctx.Fields {
		if !characterhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if chq.path != nil {
		prev, err := chq.path(ctx)
		if err != nil {
			return err
		}
		chq.sql = prev
	}
	return nil
}

func (chq *CharacterHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CharacterHistory, error) {
	var (
		nodes = []*CharacterHistory{}
		_spec = chq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CharacterHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CharacterHistory{config: chq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, chq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (chq *CharacterHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := chq.querySpec()
	_spec.Node.Columns = chq.ctx.Fields
	if len(chq.ctx.Fields) > 0 {
		_spec.Unique = chq.ctx.Unique != nil && *chq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, chq.driver, _spec)
}

func (chq *CharacterHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(characterhistory.Table, characterhistory.Columns, sqlgraph.NewFieldSpec(characterhistory.FieldID, field.TypeInt))
	_spec.From = chq.sql
	if unique := chq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if chq.path != nil {
		_spec.Unique = true
	}
	if fields := chq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, characterhistory.FieldID)
		for i := range fields {
			if fields[i] != characterhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := chq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := chq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := chq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := chq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (chq *CharacterHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(chq.driver.Dialect())
	t1 := builder.Table(characterhistory.Table)
	columns := chq.ctx.Fields
	if len(columns) == 0 {
		columns = characterhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if chq.sql != nil {
		selector = chq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if chq.ctx.Unique != nil && *chq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range chq.predicates {
		p(selector)
	}
	for _, p := range chq.order {
		p(selector)
	}
	if offset := chq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := chq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CharacterHistoryGroupBy is the group-by builder for CharacterHistory entities.
type CharacterHistoryGroupBy struct {
	selector
	build *CharacterHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (chgb *CharacterHistoryGroupBy) Aggregate(fns ...AggregateFunc) *CharacterHistoryGroupBy {
	chgb.fns = append(chgb.fns, fns...)
	return chgb
}

// Scan applies the selector query and scans the result into the given value.
func (chgb *CharacterHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, chgb.build.ctx, "GroupBy")
	if err := chgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CharacterHistoryQuery, *CharacterHistoryGroupBy](ctx, chgb.build, chgb, chgb.build.inters, v)
}

func (chgb *CharacterHistoryGroupBy) sqlScan(ctx context.Context, root *CharacterHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(chgb.fns))
	for _, fn := range chgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*chgb.flds)+len(chgb.fns))
		for _, f := range *chgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*chgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := chgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CharacterHistorySelect is the builder for selecting fields of CharacterHistory entities.
type CharacterHistorySelect struct {
	*CharacterHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (chs *CharacterHistorySelect) Aggregate(fns ...AggregateFunc) *CharacterHistorySelect {
	chs.fns = append(chs.fns, fns...)
	return chs
}

// Scan applies the selector query and scans the result into the given value.
func (chs *CharacterHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, chs.ctx, "Select")
	if err := chs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CharacterHistoryQuery, *CharacterHistorySelect](ctx, chs.CharacterHistoryQuery, chs, chs.inters, v)
}

func (chs *CharacterHistorySelect) sqlScan(ctx context.Context, root *CharacterHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(chs.fns))
	for _, fn := range chs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*chs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := chs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
