// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/messagewithpackagename"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/predicate"
)

// MessageWithPackageNameQuery is the builder for querying MessageWithPackageName entities.
type MessageWithPackageNameQuery struct {
	config
	ctx        *QueryContext
	order      []messagewithpackagename.Order
	inters     []Interceptor
	predicates []predicate.MessageWithPackageName
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MessageWithPackageNameQuery builder.
func (mwpnq *MessageWithPackageNameQuery) Where(ps ...predicate.MessageWithPackageName) *MessageWithPackageNameQuery {
	mwpnq.predicates = append(mwpnq.predicates, ps...)
	return mwpnq
}

// Limit the number of records to be returned by this query.
func (mwpnq *MessageWithPackageNameQuery) Limit(limit int) *MessageWithPackageNameQuery {
	mwpnq.ctx.Limit = &limit
	return mwpnq
}

// Offset to start from.
func (mwpnq *MessageWithPackageNameQuery) Offset(offset int) *MessageWithPackageNameQuery {
	mwpnq.ctx.Offset = &offset
	return mwpnq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mwpnq *MessageWithPackageNameQuery) Unique(unique bool) *MessageWithPackageNameQuery {
	mwpnq.ctx.Unique = &unique
	return mwpnq
}

// Order specifies how the records should be ordered.
func (mwpnq *MessageWithPackageNameQuery) Order(o ...messagewithpackagename.Order) *MessageWithPackageNameQuery {
	mwpnq.order = append(mwpnq.order, o...)
	return mwpnq
}

// First returns the first MessageWithPackageName entity from the query.
// Returns a *NotFoundError when no MessageWithPackageName was found.
func (mwpnq *MessageWithPackageNameQuery) First(ctx context.Context) (*MessageWithPackageName, error) {
	nodes, err := mwpnq.Limit(1).All(setContextOp(ctx, mwpnq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{messagewithpackagename.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) FirstX(ctx context.Context) *MessageWithPackageName {
	node, err := mwpnq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MessageWithPackageName ID from the query.
// Returns a *NotFoundError when no MessageWithPackageName ID was found.
func (mwpnq *MessageWithPackageNameQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mwpnq.Limit(1).IDs(setContextOp(ctx, mwpnq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{messagewithpackagename.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) FirstIDX(ctx context.Context) int {
	id, err := mwpnq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MessageWithPackageName entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MessageWithPackageName entity is found.
// Returns a *NotFoundError when no MessageWithPackageName entities are found.
func (mwpnq *MessageWithPackageNameQuery) Only(ctx context.Context) (*MessageWithPackageName, error) {
	nodes, err := mwpnq.Limit(2).All(setContextOp(ctx, mwpnq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{messagewithpackagename.Label}
	default:
		return nil, &NotSingularError{messagewithpackagename.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) OnlyX(ctx context.Context) *MessageWithPackageName {
	node, err := mwpnq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MessageWithPackageName ID in the query.
// Returns a *NotSingularError when more than one MessageWithPackageName ID is found.
// Returns a *NotFoundError when no entities are found.
func (mwpnq *MessageWithPackageNameQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mwpnq.Limit(2).IDs(setContextOp(ctx, mwpnq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{messagewithpackagename.Label}
	default:
		err = &NotSingularError{messagewithpackagename.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) OnlyIDX(ctx context.Context) int {
	id, err := mwpnq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MessageWithPackageNames.
func (mwpnq *MessageWithPackageNameQuery) All(ctx context.Context) ([]*MessageWithPackageName, error) {
	ctx = setContextOp(ctx, mwpnq.ctx, "All")
	if err := mwpnq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MessageWithPackageName, *MessageWithPackageNameQuery]()
	return withInterceptors[[]*MessageWithPackageName](ctx, mwpnq, qr, mwpnq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) AllX(ctx context.Context) []*MessageWithPackageName {
	nodes, err := mwpnq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MessageWithPackageName IDs.
func (mwpnq *MessageWithPackageNameQuery) IDs(ctx context.Context) (ids []int, err error) {
	if mwpnq.ctx.Unique == nil && mwpnq.path != nil {
		mwpnq.Unique(true)
	}
	ctx = setContextOp(ctx, mwpnq.ctx, "IDs")
	if err = mwpnq.Select(messagewithpackagename.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) IDsX(ctx context.Context) []int {
	ids, err := mwpnq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mwpnq *MessageWithPackageNameQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mwpnq.ctx, "Count")
	if err := mwpnq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mwpnq, querierCount[*MessageWithPackageNameQuery](), mwpnq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) CountX(ctx context.Context) int {
	count, err := mwpnq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mwpnq *MessageWithPackageNameQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mwpnq.ctx, "Exist")
	switch _, err := mwpnq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mwpnq *MessageWithPackageNameQuery) ExistX(ctx context.Context) bool {
	exist, err := mwpnq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MessageWithPackageNameQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mwpnq *MessageWithPackageNameQuery) Clone() *MessageWithPackageNameQuery {
	if mwpnq == nil {
		return nil
	}
	return &MessageWithPackageNameQuery{
		config:     mwpnq.config,
		ctx:        mwpnq.ctx.Clone(),
		order:      append([]messagewithpackagename.Order{}, mwpnq.order...),
		inters:     append([]Interceptor{}, mwpnq.inters...),
		predicates: append([]predicate.MessageWithPackageName{}, mwpnq.predicates...),
		// clone intermediate query.
		sql:  mwpnq.sql.Clone(),
		path: mwpnq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MessageWithPackageName.Query().
//		GroupBy(messagewithpackagename.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mwpnq *MessageWithPackageNameQuery) GroupBy(field string, fields ...string) *MessageWithPackageNameGroupBy {
	mwpnq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MessageWithPackageNameGroupBy{build: mwpnq}
	grbuild.flds = &mwpnq.ctx.Fields
	grbuild.label = messagewithpackagename.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.MessageWithPackageName.Query().
//		Select(messagewithpackagename.FieldName).
//		Scan(ctx, &v)
func (mwpnq *MessageWithPackageNameQuery) Select(fields ...string) *MessageWithPackageNameSelect {
	mwpnq.ctx.Fields = append(mwpnq.ctx.Fields, fields...)
	sbuild := &MessageWithPackageNameSelect{MessageWithPackageNameQuery: mwpnq}
	sbuild.label = messagewithpackagename.Label
	sbuild.flds, sbuild.scan = &mwpnq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MessageWithPackageNameSelect configured with the given aggregations.
func (mwpnq *MessageWithPackageNameQuery) Aggregate(fns ...AggregateFunc) *MessageWithPackageNameSelect {
	return mwpnq.Select().Aggregate(fns...)
}

func (mwpnq *MessageWithPackageNameQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mwpnq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mwpnq); err != nil {
				return err
			}
		}
	}
	for _, f := range mwpnq.ctx.Fields {
		if !messagewithpackagename.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mwpnq.path != nil {
		prev, err := mwpnq.path(ctx)
		if err != nil {
			return err
		}
		mwpnq.sql = prev
	}
	return nil
}

func (mwpnq *MessageWithPackageNameQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MessageWithPackageName, error) {
	var (
		nodes = []*MessageWithPackageName{}
		_spec = mwpnq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MessageWithPackageName).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MessageWithPackageName{config: mwpnq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mwpnq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mwpnq *MessageWithPackageNameQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mwpnq.querySpec()
	_spec.Node.Columns = mwpnq.ctx.Fields
	if len(mwpnq.ctx.Fields) > 0 {
		_spec.Unique = mwpnq.ctx.Unique != nil && *mwpnq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mwpnq.driver, _spec)
}

func (mwpnq *MessageWithPackageNameQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(messagewithpackagename.Table, messagewithpackagename.Columns, sqlgraph.NewFieldSpec(messagewithpackagename.FieldID, field.TypeInt))
	_spec.From = mwpnq.sql
	if unique := mwpnq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mwpnq.path != nil {
		_spec.Unique = true
	}
	if fields := mwpnq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagewithpackagename.FieldID)
		for i := range fields {
			if fields[i] != messagewithpackagename.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mwpnq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mwpnq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mwpnq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mwpnq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mwpnq *MessageWithPackageNameQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mwpnq.driver.Dialect())
	t1 := builder.Table(messagewithpackagename.Table)
	columns := mwpnq.ctx.Fields
	if len(columns) == 0 {
		columns = messagewithpackagename.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mwpnq.sql != nil {
		selector = mwpnq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mwpnq.ctx.Unique != nil && *mwpnq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mwpnq.predicates {
		p(selector)
	}
	for _, p := range mwpnq.order {
		p(selector)
	}
	if offset := mwpnq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mwpnq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MessageWithPackageNameGroupBy is the group-by builder for MessageWithPackageName entities.
type MessageWithPackageNameGroupBy struct {
	selector
	build *MessageWithPackageNameQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mwpngb *MessageWithPackageNameGroupBy) Aggregate(fns ...AggregateFunc) *MessageWithPackageNameGroupBy {
	mwpngb.fns = append(mwpngb.fns, fns...)
	return mwpngb
}

// Scan applies the selector query and scans the result into the given value.
func (mwpngb *MessageWithPackageNameGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mwpngb.build.ctx, "GroupBy")
	if err := mwpngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MessageWithPackageNameQuery, *MessageWithPackageNameGroupBy](ctx, mwpngb.build, mwpngb, mwpngb.build.inters, v)
}

func (mwpngb *MessageWithPackageNameGroupBy) sqlScan(ctx context.Context, root *MessageWithPackageNameQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mwpngb.fns))
	for _, fn := range mwpngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mwpngb.flds)+len(mwpngb.fns))
		for _, f := range *mwpngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mwpngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mwpngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MessageWithPackageNameSelect is the builder for selecting fields of MessageWithPackageName entities.
type MessageWithPackageNameSelect struct {
	*MessageWithPackageNameQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mwpns *MessageWithPackageNameSelect) Aggregate(fns ...AggregateFunc) *MessageWithPackageNameSelect {
	mwpns.fns = append(mwpns.fns, fns...)
	return mwpns
}

// Scan applies the selector query and scans the result into the given value.
func (mwpns *MessageWithPackageNameSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mwpns.ctx, "Select")
	if err := mwpns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MessageWithPackageNameQuery, *MessageWithPackageNameSelect](ctx, mwpns.MessageWithPackageNameQuery, mwpns, mwpns.inters, v)
}

func (mwpns *MessageWithPackageNameSelect) sqlScan(ctx context.Context, root *MessageWithPackageNameQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mwpns.fns))
	for _, fn := range mwpns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mwpns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mwpns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
