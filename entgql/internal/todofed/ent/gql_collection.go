// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/tilau2328/entcontrib/entgql"
	"github.com/tilau2328/entcontrib/entgql/internal/todofed/ent/category"
	"github.com/tilau2328/entcontrib/entgql/internal/todofed/ent/todo"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CategoryQuery) CollectFields(ctx context.Context, satisfies ...string) (*CategoryQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CategoryQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(category.Columns))
		selectedFields = []string{category.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "todos":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TodoClient{config: c.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			c.WithNamedTodos(alias, func(wq *TodoQuery) {
				*wq = *query
			})
		case "text":
			if _, ok := fieldSeen[category.FieldText]; !ok {
				selectedFields = append(selectedFields, category.FieldText)
				fieldSeen[category.FieldText] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[category.FieldStatus]; !ok {
				selectedFields = append(selectedFields, category.FieldStatus)
				fieldSeen[category.FieldStatus] = struct{}{}
			}
		case "config":
			if _, ok := fieldSeen[category.FieldConfig]; !ok {
				selectedFields = append(selectedFields, category.FieldConfig)
				fieldSeen[category.FieldConfig] = struct{}{}
			}
		case "duration":
			if _, ok := fieldSeen[category.FieldDuration]; !ok {
				selectedFields = append(selectedFields, category.FieldDuration)
				fieldSeen[category.FieldDuration] = struct{}{}
			}
		case "count":
			if _, ok := fieldSeen[category.FieldCount]; !ok {
				selectedFields = append(selectedFields, category.FieldCount)
				fieldSeen[category.FieldCount] = struct{}{}
			}
		case "strings":
			if _, ok := fieldSeen[category.FieldStrings]; !ok {
				selectedFields = append(selectedFields, category.FieldStrings)
				fieldSeen[category.FieldStrings] = struct{}{}
			}
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		c.Select(selectedFields...)
	}
	return nil
}

type categoryPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CategoryPaginateOption
}

func newCategoryPaginateArgs(rv map[string]interface{}) *categoryPaginateArgs {
	args := &categoryPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &CategoryOrder{Field: &CategoryOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCategoryOrder(order))
			}
		case *CategoryOrder:
			if v != nil {
				args.opts = append(args.opts, WithCategoryOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TodoQuery) CollectFields(ctx context.Context, satisfies ...string) (*TodoQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TodoQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(todo.Columns))
		selectedFields = []string{todo.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "parent":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TodoClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.withParent = query
		case "children":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TodoClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.WithNamedChildren(alias, func(wq *TodoQuery) {
				*wq = *query
			})
		case "category":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CategoryClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.withCategory = query
		case "createdAt":
			if _, ok := fieldSeen[todo.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, todo.FieldCreatedAt)
				fieldSeen[todo.FieldCreatedAt] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[todo.FieldStatus]; !ok {
				selectedFields = append(selectedFields, todo.FieldStatus)
				fieldSeen[todo.FieldStatus] = struct{}{}
			}
		case "priority":
			if _, ok := fieldSeen[todo.FieldPriority]; !ok {
				selectedFields = append(selectedFields, todo.FieldPriority)
				fieldSeen[todo.FieldPriority] = struct{}{}
			}
		case "text":
			if _, ok := fieldSeen[todo.FieldText]; !ok {
				selectedFields = append(selectedFields, todo.FieldText)
				fieldSeen[todo.FieldText] = struct{}{}
			}
		case "blob":
			if _, ok := fieldSeen[todo.FieldBlob]; !ok {
				selectedFields = append(selectedFields, todo.FieldBlob)
				fieldSeen[todo.FieldBlob] = struct{}{}
			}
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type todoPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TodoPaginateOption
}

func newTodoPaginateArgs(rv map[string]interface{}) *todoPaginateArgs {
	args := &todoPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &TodoOrder{Field: &TodoOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTodoOrder(order))
			}
		case *TodoOrder:
			if v != nil {
				args.opts = append(args.opts, WithTodoOrder(v))
			}
		}
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Alias == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}
