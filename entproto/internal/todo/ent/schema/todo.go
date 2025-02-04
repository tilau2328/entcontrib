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

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/tilau2328/entcontrib/entproto"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("task").
			Annotations(entproto.Field(2)),
		field.Enum("status").
			Values("pending", "in_progress", "done").
			Default("pending").
			Annotations(
				entproto.Field(3),
				entproto.Enum(map[string]int32{
					"pending":     0,
					"in_progress": 1,
					"done":        2,
				}),
			),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Annotations(entproto.Field(4)),
	}
}
