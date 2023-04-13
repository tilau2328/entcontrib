// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/predicate"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/skipedgeexample"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/user"
)

// SkipEdgeExampleQuery is the builder for querying SkipEdgeExample entities.
type SkipEdgeExampleQuery struct {
	config
	ctx        *QueryContext
	order      []skipedgeexample.Order
	inters     []Interceptor
	predicates []predicate.SkipEdgeExample
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SkipEdgeExampleQuery builder.
func (seeq *SkipEdgeExampleQuery) Where(ps ...predicate.SkipEdgeExample) *SkipEdgeExampleQuery {
	seeq.predicates = append(seeq.predicates, ps...)
	return seeq
}

// Limit the number of records to be returned by this query.
func (seeq *SkipEdgeExampleQuery) Limit(limit int) *SkipEdgeExampleQuery {
	seeq.ctx.Limit = &limit
	return seeq
}

// Offset to start from.
func (seeq *SkipEdgeExampleQuery) Offset(offset int) *SkipEdgeExampleQuery {
	seeq.ctx.Offset = &offset
	return seeq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (seeq *SkipEdgeExampleQuery) Unique(unique bool) *SkipEdgeExampleQuery {
	seeq.ctx.Unique = &unique
	return seeq
}

// Order specifies how the records should be ordered.
func (seeq *SkipEdgeExampleQuery) Order(o ...skipedgeexample.Order) *SkipEdgeExampleQuery {
	seeq.order = append(seeq.order, o...)
	return seeq
}

// QueryUser chains the current query on the "user" edge.
func (seeq *SkipEdgeExampleQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: seeq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := seeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := seeq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(skipedgeexample.Table, skipedgeexample.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, skipedgeexample.UserTable, skipedgeexample.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(seeq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SkipEdgeExample entity from the query.
// Returns a *NotFoundError when no SkipEdgeExample was found.
func (seeq *SkipEdgeExampleQuery) First(ctx context.Context) (*SkipEdgeExample, error) {
	nodes, err := seeq.Limit(1).All(setContextOp(ctx, seeq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{skipedgeexample.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (seeq *SkipEdgeExampleQuery) FirstX(ctx context.Context) *SkipEdgeExample {
	node, err := seeq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SkipEdgeExample ID from the query.
// Returns a *NotFoundError when no SkipEdgeExample ID was found.
func (seeq *SkipEdgeExampleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = seeq.Limit(1).IDs(setContextOp(ctx, seeq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{skipedgeexample.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (seeq *SkipEdgeExampleQuery) FirstIDX(ctx context.Context) int {
	id, err := seeq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SkipEdgeExample entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SkipEdgeExample entity is found.
// Returns a *NotFoundError when no SkipEdgeExample entities are found.
func (seeq *SkipEdgeExampleQuery) Only(ctx context.Context) (*SkipEdgeExample, error) {
	nodes, err := seeq.Limit(2).All(setContextOp(ctx, seeq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{skipedgeexample.Label}
	default:
		return nil, &NotSingularError{skipedgeexample.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (seeq *SkipEdgeExampleQuery) OnlyX(ctx context.Context) *SkipEdgeExample {
	node, err := seeq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SkipEdgeExample ID in the query.
// Returns a *NotSingularError when more than one SkipEdgeExample ID is found.
// Returns a *NotFoundError when no entities are found.
func (seeq *SkipEdgeExampleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = seeq.Limit(2).IDs(setContextOp(ctx, seeq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{skipedgeexample.Label}
	default:
		err = &NotSingularError{skipedgeexample.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (seeq *SkipEdgeExampleQuery) OnlyIDX(ctx context.Context) int {
	id, err := seeq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SkipEdgeExamples.
func (seeq *SkipEdgeExampleQuery) All(ctx context.Context) ([]*SkipEdgeExample, error) {
	ctx = setContextOp(ctx, seeq.ctx, "All")
	if err := seeq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SkipEdgeExample, *SkipEdgeExampleQuery]()
	return withInterceptors[[]*SkipEdgeExample](ctx, seeq, qr, seeq.inters)
}

// AllX is like All, but panics if an error occurs.
func (seeq *SkipEdgeExampleQuery) AllX(ctx context.Context) []*SkipEdgeExample {
	nodes, err := seeq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SkipEdgeExample IDs.
func (seeq *SkipEdgeExampleQuery) IDs(ctx context.Context) (ids []int, err error) {
	if seeq.ctx.Unique == nil && seeq.path != nil {
		seeq.Unique(true)
	}
	ctx = setContextOp(ctx, seeq.ctx, "IDs")
	if err = seeq.Select(skipedgeexample.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (seeq *SkipEdgeExampleQuery) IDsX(ctx context.Context) []int {
	ids, err := seeq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (seeq *SkipEdgeExampleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, seeq.ctx, "Count")
	if err := seeq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, seeq, querierCount[*SkipEdgeExampleQuery](), seeq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (seeq *SkipEdgeExampleQuery) CountX(ctx context.Context) int {
	count, err := seeq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (seeq *SkipEdgeExampleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, seeq.ctx, "Exist")
	switch _, err := seeq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (seeq *SkipEdgeExampleQuery) ExistX(ctx context.Context) bool {
	exist, err := seeq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SkipEdgeExampleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (seeq *SkipEdgeExampleQuery) Clone() *SkipEdgeExampleQuery {
	if seeq == nil {
		return nil
	}
	return &SkipEdgeExampleQuery{
		config:     seeq.config,
		ctx:        seeq.ctx.Clone(),
		order:      append([]skipedgeexample.Order{}, seeq.order...),
		inters:     append([]Interceptor{}, seeq.inters...),
		predicates: append([]predicate.SkipEdgeExample{}, seeq.predicates...),
		withUser:   seeq.withUser.Clone(),
		// clone intermediate query.
		sql:  seeq.sql.Clone(),
		path: seeq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (seeq *SkipEdgeExampleQuery) WithUser(opts ...func(*UserQuery)) *SkipEdgeExampleQuery {
	query := (&UserClient{config: seeq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	seeq.withUser = query
	return seeq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (seeq *SkipEdgeExampleQuery) GroupBy(field string, fields ...string) *SkipEdgeExampleGroupBy {
	seeq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SkipEdgeExampleGroupBy{build: seeq}
	grbuild.flds = &seeq.ctx.Fields
	grbuild.label = skipedgeexample.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (seeq *SkipEdgeExampleQuery) Select(fields ...string) *SkipEdgeExampleSelect {
	seeq.ctx.Fields = append(seeq.ctx.Fields, fields...)
	sbuild := &SkipEdgeExampleSelect{SkipEdgeExampleQuery: seeq}
	sbuild.label = skipedgeexample.Label
	sbuild.flds, sbuild.scan = &seeq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SkipEdgeExampleSelect configured with the given aggregations.
func (seeq *SkipEdgeExampleQuery) Aggregate(fns ...AggregateFunc) *SkipEdgeExampleSelect {
	return seeq.Select().Aggregate(fns...)
}

func (seeq *SkipEdgeExampleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range seeq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, seeq); err != nil {
				return err
			}
		}
	}
	for _, f := range seeq.ctx.Fields {
		if !skipedgeexample.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if seeq.path != nil {
		prev, err := seeq.path(ctx)
		if err != nil {
			return err
		}
		seeq.sql = prev
	}
	return nil
}

func (seeq *SkipEdgeExampleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SkipEdgeExample, error) {
	var (
		nodes       = []*SkipEdgeExample{}
		withFKs     = seeq.withFKs
		_spec       = seeq.querySpec()
		loadedTypes = [1]bool{
			seeq.withUser != nil,
		}
	)
	if seeq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, skipedgeexample.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SkipEdgeExample).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SkipEdgeExample{config: seeq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, seeq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := seeq.withUser; query != nil {
		if err := seeq.loadUser(ctx, query, nodes, nil,
			func(n *SkipEdgeExample, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (seeq *SkipEdgeExampleQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*SkipEdgeExample, init func(*SkipEdgeExample), assign func(*SkipEdgeExample, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SkipEdgeExample)
	for i := range nodes {
		if nodes[i].user_skip_edge == nil {
			continue
		}
		fk := *nodes[i].user_skip_edge
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
			return fmt.Errorf(`unexpected foreign-key "user_skip_edge" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (seeq *SkipEdgeExampleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := seeq.querySpec()
	_spec.Node.Columns = seeq.ctx.Fields
	if len(seeq.ctx.Fields) > 0 {
		_spec.Unique = seeq.ctx.Unique != nil && *seeq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, seeq.driver, _spec)
}

func (seeq *SkipEdgeExampleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(skipedgeexample.Table, skipedgeexample.Columns, sqlgraph.NewFieldSpec(skipedgeexample.FieldID, field.TypeInt))
	_spec.From = seeq.sql
	if unique := seeq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if seeq.path != nil {
		_spec.Unique = true
	}
	if fields := seeq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skipedgeexample.FieldID)
		for i := range fields {
			if fields[i] != skipedgeexample.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := seeq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := seeq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := seeq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := seeq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (seeq *SkipEdgeExampleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(seeq.driver.Dialect())
	t1 := builder.Table(skipedgeexample.Table)
	columns := seeq.ctx.Fields
	if len(columns) == 0 {
		columns = skipedgeexample.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if seeq.sql != nil {
		selector = seeq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if seeq.ctx.Unique != nil && *seeq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range seeq.predicates {
		p(selector)
	}
	for _, p := range seeq.order {
		p(selector)
	}
	if offset := seeq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := seeq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SkipEdgeExampleGroupBy is the group-by builder for SkipEdgeExample entities.
type SkipEdgeExampleGroupBy struct {
	selector
	build *SkipEdgeExampleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (seegb *SkipEdgeExampleGroupBy) Aggregate(fns ...AggregateFunc) *SkipEdgeExampleGroupBy {
	seegb.fns = append(seegb.fns, fns...)
	return seegb
}

// Scan applies the selector query and scans the result into the given value.
func (seegb *SkipEdgeExampleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, seegb.build.ctx, "GroupBy")
	if err := seegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SkipEdgeExampleQuery, *SkipEdgeExampleGroupBy](ctx, seegb.build, seegb, seegb.build.inters, v)
}

func (seegb *SkipEdgeExampleGroupBy) sqlScan(ctx context.Context, root *SkipEdgeExampleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(seegb.fns))
	for _, fn := range seegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*seegb.flds)+len(seegb.fns))
		for _, f := range *seegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*seegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := seegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SkipEdgeExampleSelect is the builder for selecting fields of SkipEdgeExample entities.
type SkipEdgeExampleSelect struct {
	*SkipEdgeExampleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sees *SkipEdgeExampleSelect) Aggregate(fns ...AggregateFunc) *SkipEdgeExampleSelect {
	sees.fns = append(sees.fns, fns...)
	return sees
}

// Scan applies the selector query and scans the result into the given value.
func (sees *SkipEdgeExampleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sees.ctx, "Select")
	if err := sees.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SkipEdgeExampleQuery, *SkipEdgeExampleSelect](ctx, sees.SkipEdgeExampleQuery, sees, sees.inters, v)
}

func (sees *SkipEdgeExampleSelect) sqlScan(ctx context.Context, root *SkipEdgeExampleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sees.fns))
	for _, fn := range sees.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sees.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sees.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
