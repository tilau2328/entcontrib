// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/schemast/internal/mutatetest/ent/predicate"
	"github.com/tilau2328/entcontrib/schemast/internal/mutatetest/ent/withoutfields"
)

// WithoutFieldsQuery is the builder for querying WithoutFields entities.
type WithoutFieldsQuery struct {
	config
	ctx        *QueryContext
	order      []withoutfields.Order
	inters     []Interceptor
	predicates []predicate.WithoutFields
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WithoutFieldsQuery builder.
func (wfq *WithoutFieldsQuery) Where(ps ...predicate.WithoutFields) *WithoutFieldsQuery {
	wfq.predicates = append(wfq.predicates, ps...)
	return wfq
}

// Limit the number of records to be returned by this query.
func (wfq *WithoutFieldsQuery) Limit(limit int) *WithoutFieldsQuery {
	wfq.ctx.Limit = &limit
	return wfq
}

// Offset to start from.
func (wfq *WithoutFieldsQuery) Offset(offset int) *WithoutFieldsQuery {
	wfq.ctx.Offset = &offset
	return wfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wfq *WithoutFieldsQuery) Unique(unique bool) *WithoutFieldsQuery {
	wfq.ctx.Unique = &unique
	return wfq
}

// Order specifies how the records should be ordered.
func (wfq *WithoutFieldsQuery) Order(o ...withoutfields.Order) *WithoutFieldsQuery {
	wfq.order = append(wfq.order, o...)
	return wfq
}

// First returns the first WithoutFields entity from the query.
// Returns a *NotFoundError when no WithoutFields was found.
func (wfq *WithoutFieldsQuery) First(ctx context.Context) (*WithoutFields, error) {
	nodes, err := wfq.Limit(1).All(setContextOp(ctx, wfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{withoutfields.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) FirstX(ctx context.Context) *WithoutFields {
	node, err := wfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WithoutFields ID from the query.
// Returns a *NotFoundError when no WithoutFields ID was found.
func (wfq *WithoutFieldsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wfq.Limit(1).IDs(setContextOp(ctx, wfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{withoutfields.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) FirstIDX(ctx context.Context) int {
	id, err := wfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WithoutFields entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WithoutFields entity is found.
// Returns a *NotFoundError when no WithoutFields entities are found.
func (wfq *WithoutFieldsQuery) Only(ctx context.Context) (*WithoutFields, error) {
	nodes, err := wfq.Limit(2).All(setContextOp(ctx, wfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{withoutfields.Label}
	default:
		return nil, &NotSingularError{withoutfields.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) OnlyX(ctx context.Context) *WithoutFields {
	node, err := wfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WithoutFields ID in the query.
// Returns a *NotSingularError when more than one WithoutFields ID is found.
// Returns a *NotFoundError when no entities are found.
func (wfq *WithoutFieldsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wfq.Limit(2).IDs(setContextOp(ctx, wfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{withoutfields.Label}
	default:
		err = &NotSingularError{withoutfields.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) OnlyIDX(ctx context.Context) int {
	id, err := wfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WithoutFieldsSlice.
func (wfq *WithoutFieldsQuery) All(ctx context.Context) ([]*WithoutFields, error) {
	ctx = setContextOp(ctx, wfq.ctx, "All")
	if err := wfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WithoutFields, *WithoutFieldsQuery]()
	return withInterceptors[[]*WithoutFields](ctx, wfq, qr, wfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) AllX(ctx context.Context) []*WithoutFields {
	nodes, err := wfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WithoutFields IDs.
func (wfq *WithoutFieldsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if wfq.ctx.Unique == nil && wfq.path != nil {
		wfq.Unique(true)
	}
	ctx = setContextOp(ctx, wfq.ctx, "IDs")
	if err = wfq.Select(withoutfields.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) IDsX(ctx context.Context) []int {
	ids, err := wfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wfq *WithoutFieldsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wfq.ctx, "Count")
	if err := wfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wfq, querierCount[*WithoutFieldsQuery](), wfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) CountX(ctx context.Context) int {
	count, err := wfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wfq *WithoutFieldsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wfq.ctx, "Exist")
	switch _, err := wfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) ExistX(ctx context.Context) bool {
	exist, err := wfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WithoutFieldsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wfq *WithoutFieldsQuery) Clone() *WithoutFieldsQuery {
	if wfq == nil {
		return nil
	}
	return &WithoutFieldsQuery{
		config:     wfq.config,
		ctx:        wfq.ctx.Clone(),
		order:      append([]withoutfields.Order{}, wfq.order...),
		inters:     append([]Interceptor{}, wfq.inters...),
		predicates: append([]predicate.WithoutFields{}, wfq.predicates...),
		// clone intermediate query.
		sql:  wfq.sql.Clone(),
		path: wfq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (wfq *WithoutFieldsQuery) GroupBy(field string, fields ...string) *WithoutFieldsGroupBy {
	wfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WithoutFieldsGroupBy{build: wfq}
	grbuild.flds = &wfq.ctx.Fields
	grbuild.label = withoutfields.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (wfq *WithoutFieldsQuery) Select(fields ...string) *WithoutFieldsSelect {
	wfq.ctx.Fields = append(wfq.ctx.Fields, fields...)
	sbuild := &WithoutFieldsSelect{WithoutFieldsQuery: wfq}
	sbuild.label = withoutfields.Label
	sbuild.flds, sbuild.scan = &wfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WithoutFieldsSelect configured with the given aggregations.
func (wfq *WithoutFieldsQuery) Aggregate(fns ...AggregateFunc) *WithoutFieldsSelect {
	return wfq.Select().Aggregate(fns...)
}

func (wfq *WithoutFieldsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wfq); err != nil {
				return err
			}
		}
	}
	for _, f := range wfq.ctx.Fields {
		if !withoutfields.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wfq.path != nil {
		prev, err := wfq.path(ctx)
		if err != nil {
			return err
		}
		wfq.sql = prev
	}
	return nil
}

func (wfq *WithoutFieldsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WithoutFields, error) {
	var (
		nodes = []*WithoutFields{}
		_spec = wfq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WithoutFields).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WithoutFields{config: wfq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wfq *WithoutFieldsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wfq.querySpec()
	_spec.Node.Columns = wfq.ctx.Fields
	if len(wfq.ctx.Fields) > 0 {
		_spec.Unique = wfq.ctx.Unique != nil && *wfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wfq.driver, _spec)
}

func (wfq *WithoutFieldsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(withoutfields.Table, withoutfields.Columns, sqlgraph.NewFieldSpec(withoutfields.FieldID, field.TypeInt))
	_spec.From = wfq.sql
	if unique := wfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wfq.path != nil {
		_spec.Unique = true
	}
	if fields := wfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withoutfields.FieldID)
		for i := range fields {
			if fields[i] != withoutfields.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wfq *WithoutFieldsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wfq.driver.Dialect())
	t1 := builder.Table(withoutfields.Table)
	columns := wfq.ctx.Fields
	if len(columns) == 0 {
		columns = withoutfields.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wfq.sql != nil {
		selector = wfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wfq.ctx.Unique != nil && *wfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wfq.predicates {
		p(selector)
	}
	for _, p := range wfq.order {
		p(selector)
	}
	if offset := wfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithoutFieldsGroupBy is the group-by builder for WithoutFields entities.
type WithoutFieldsGroupBy struct {
	selector
	build *WithoutFieldsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wfgb *WithoutFieldsGroupBy) Aggregate(fns ...AggregateFunc) *WithoutFieldsGroupBy {
	wfgb.fns = append(wfgb.fns, fns...)
	return wfgb
}

// Scan applies the selector query and scans the result into the given value.
func (wfgb *WithoutFieldsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wfgb.build.ctx, "GroupBy")
	if err := wfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WithoutFieldsQuery, *WithoutFieldsGroupBy](ctx, wfgb.build, wfgb, wfgb.build.inters, v)
}

func (wfgb *WithoutFieldsGroupBy) sqlScan(ctx context.Context, root *WithoutFieldsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wfgb.fns))
	for _, fn := range wfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wfgb.flds)+len(wfgb.fns))
		for _, f := range *wfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WithoutFieldsSelect is the builder for selecting fields of WithoutFields entities.
type WithoutFieldsSelect struct {
	*WithoutFieldsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wfs *WithoutFieldsSelect) Aggregate(fns ...AggregateFunc) *WithoutFieldsSelect {
	wfs.fns = append(wfs.fns, fns...)
	return wfs
}

// Scan applies the selector query and scans the result into the given value.
func (wfs *WithoutFieldsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wfs.ctx, "Select")
	if err := wfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WithoutFieldsQuery, *WithoutFieldsSelect](ctx, wfs.WithoutFieldsQuery, wfs, wfs.inters, v)
}

func (wfs *WithoutFieldsSelect) sqlScan(ctx context.Context, root *WithoutFieldsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wfs.fns))
	for _, fn := range wfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
