// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sysatom/framework/ent/merchantaccount"
	"github.com/sysatom/framework/ent/predicate"
)

// MerchantAccountQuery is the builder for querying MerchantAccount entities.
type MerchantAccountQuery struct {
	config
	ctx        *QueryContext
	order      []merchantaccount.OrderOption
	inters     []Interceptor
	predicates []predicate.MerchantAccount
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MerchantAccountQuery builder.
func (maq *MerchantAccountQuery) Where(ps ...predicate.MerchantAccount) *MerchantAccountQuery {
	maq.predicates = append(maq.predicates, ps...)
	return maq
}

// Limit the number of records to be returned by this query.
func (maq *MerchantAccountQuery) Limit(limit int) *MerchantAccountQuery {
	maq.ctx.Limit = &limit
	return maq
}

// Offset to start from.
func (maq *MerchantAccountQuery) Offset(offset int) *MerchantAccountQuery {
	maq.ctx.Offset = &offset
	return maq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (maq *MerchantAccountQuery) Unique(unique bool) *MerchantAccountQuery {
	maq.ctx.Unique = &unique
	return maq
}

// Order specifies how the records should be ordered.
func (maq *MerchantAccountQuery) Order(o ...merchantaccount.OrderOption) *MerchantAccountQuery {
	maq.order = append(maq.order, o...)
	return maq
}

// First returns the first MerchantAccount entity from the query.
// Returns a *NotFoundError when no MerchantAccount was found.
func (maq *MerchantAccountQuery) First(ctx context.Context) (*MerchantAccount, error) {
	nodes, err := maq.Limit(1).All(setContextOp(ctx, maq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{merchantaccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (maq *MerchantAccountQuery) FirstX(ctx context.Context) *MerchantAccount {
	node, err := maq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MerchantAccount ID from the query.
// Returns a *NotFoundError when no MerchantAccount ID was found.
func (maq *MerchantAccountQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = maq.Limit(1).IDs(setContextOp(ctx, maq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{merchantaccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (maq *MerchantAccountQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := maq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MerchantAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MerchantAccount entity is found.
// Returns a *NotFoundError when no MerchantAccount entities are found.
func (maq *MerchantAccountQuery) Only(ctx context.Context) (*MerchantAccount, error) {
	nodes, err := maq.Limit(2).All(setContextOp(ctx, maq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{merchantaccount.Label}
	default:
		return nil, &NotSingularError{merchantaccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (maq *MerchantAccountQuery) OnlyX(ctx context.Context) *MerchantAccount {
	node, err := maq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MerchantAccount ID in the query.
// Returns a *NotSingularError when more than one MerchantAccount ID is found.
// Returns a *NotFoundError when no entities are found.
func (maq *MerchantAccountQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = maq.Limit(2).IDs(setContextOp(ctx, maq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{merchantaccount.Label}
	default:
		err = &NotSingularError{merchantaccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (maq *MerchantAccountQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := maq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MerchantAccounts.
func (maq *MerchantAccountQuery) All(ctx context.Context) ([]*MerchantAccount, error) {
	ctx = setContextOp(ctx, maq.ctx, ent.OpQueryAll)
	if err := maq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MerchantAccount, *MerchantAccountQuery]()
	return withInterceptors[[]*MerchantAccount](ctx, maq, qr, maq.inters)
}

// AllX is like All, but panics if an error occurs.
func (maq *MerchantAccountQuery) AllX(ctx context.Context) []*MerchantAccount {
	nodes, err := maq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MerchantAccount IDs.
func (maq *MerchantAccountQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if maq.ctx.Unique == nil && maq.path != nil {
		maq.Unique(true)
	}
	ctx = setContextOp(ctx, maq.ctx, ent.OpQueryIDs)
	if err = maq.Select(merchantaccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (maq *MerchantAccountQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := maq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (maq *MerchantAccountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, maq.ctx, ent.OpQueryCount)
	if err := maq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, maq, querierCount[*MerchantAccountQuery](), maq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (maq *MerchantAccountQuery) CountX(ctx context.Context) int {
	count, err := maq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (maq *MerchantAccountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, maq.ctx, ent.OpQueryExist)
	switch _, err := maq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (maq *MerchantAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := maq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MerchantAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (maq *MerchantAccountQuery) Clone() *MerchantAccountQuery {
	if maq == nil {
		return nil
	}
	return &MerchantAccountQuery{
		config:     maq.config,
		ctx:        maq.ctx.Clone(),
		order:      append([]merchantaccount.OrderOption{}, maq.order...),
		inters:     append([]Interceptor{}, maq.inters...),
		predicates: append([]predicate.MerchantAccount{}, maq.predicates...),
		// clone intermediate query.
		sql:  maq.sql.Clone(),
		path: maq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MerchantAccount.Query().
//		GroupBy(merchantaccount.FieldUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (maq *MerchantAccountQuery) GroupBy(field string, fields ...string) *MerchantAccountGroupBy {
	maq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MerchantAccountGroupBy{build: maq}
	grbuild.flds = &maq.ctx.Fields
	grbuild.label = merchantaccount.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.MerchantAccount.Query().
//		Select(merchantaccount.FieldUsername).
//		Scan(ctx, &v)
func (maq *MerchantAccountQuery) Select(fields ...string) *MerchantAccountSelect {
	maq.ctx.Fields = append(maq.ctx.Fields, fields...)
	sbuild := &MerchantAccountSelect{MerchantAccountQuery: maq}
	sbuild.label = merchantaccount.Label
	sbuild.flds, sbuild.scan = &maq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MerchantAccountSelect configured with the given aggregations.
func (maq *MerchantAccountQuery) Aggregate(fns ...AggregateFunc) *MerchantAccountSelect {
	return maq.Select().Aggregate(fns...)
}

func (maq *MerchantAccountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range maq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, maq); err != nil {
				return err
			}
		}
	}
	for _, f := range maq.ctx.Fields {
		if !merchantaccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if maq.path != nil {
		prev, err := maq.path(ctx)
		if err != nil {
			return err
		}
		maq.sql = prev
	}
	return nil
}

func (maq *MerchantAccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MerchantAccount, error) {
	var (
		nodes   = []*MerchantAccount{}
		withFKs = maq.withFKs
		_spec   = maq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, merchantaccount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MerchantAccount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MerchantAccount{config: maq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, maq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (maq *MerchantAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := maq.querySpec()
	_spec.Node.Columns = maq.ctx.Fields
	if len(maq.ctx.Fields) > 0 {
		_spec.Unique = maq.ctx.Unique != nil && *maq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, maq.driver, _spec)
}

func (maq *MerchantAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(merchantaccount.Table, merchantaccount.Columns, sqlgraph.NewFieldSpec(merchantaccount.FieldID, field.TypeUint64))
	_spec.From = maq.sql
	if unique := maq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if maq.path != nil {
		_spec.Unique = true
	}
	if fields := maq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, merchantaccount.FieldID)
		for i := range fields {
			if fields[i] != merchantaccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := maq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := maq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := maq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := maq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (maq *MerchantAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(maq.driver.Dialect())
	t1 := builder.Table(merchantaccount.Table)
	columns := maq.ctx.Fields
	if len(columns) == 0 {
		columns = merchantaccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if maq.sql != nil {
		selector = maq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if maq.ctx.Unique != nil && *maq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range maq.predicates {
		p(selector)
	}
	for _, p := range maq.order {
		p(selector)
	}
	if offset := maq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := maq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MerchantAccountGroupBy is the group-by builder for MerchantAccount entities.
type MerchantAccountGroupBy struct {
	selector
	build *MerchantAccountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (magb *MerchantAccountGroupBy) Aggregate(fns ...AggregateFunc) *MerchantAccountGroupBy {
	magb.fns = append(magb.fns, fns...)
	return magb
}

// Scan applies the selector query and scans the result into the given value.
func (magb *MerchantAccountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, magb.build.ctx, ent.OpQueryGroupBy)
	if err := magb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MerchantAccountQuery, *MerchantAccountGroupBy](ctx, magb.build, magb, magb.build.inters, v)
}

func (magb *MerchantAccountGroupBy) sqlScan(ctx context.Context, root *MerchantAccountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(magb.fns))
	for _, fn := range magb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*magb.flds)+len(magb.fns))
		for _, f := range *magb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*magb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := magb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MerchantAccountSelect is the builder for selecting fields of MerchantAccount entities.
type MerchantAccountSelect struct {
	*MerchantAccountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mas *MerchantAccountSelect) Aggregate(fns ...AggregateFunc) *MerchantAccountSelect {
	mas.fns = append(mas.fns, fns...)
	return mas
}

// Scan applies the selector query and scans the result into the given value.
func (mas *MerchantAccountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mas.ctx, ent.OpQuerySelect)
	if err := mas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MerchantAccountQuery, *MerchantAccountSelect](ctx, mas.MerchantAccountQuery, mas, mas.inters, v)
}

func (mas *MerchantAccountSelect) sqlScan(ctx context.Context, root *MerchantAccountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mas.fns))
	for _, fn := range mas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
