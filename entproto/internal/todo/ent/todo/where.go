// Code generated by ent, DO NOT EDIT.

package todo

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tilau2328/entcontrib/entproto/internal/todo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldID, id))
}

// Task applies equality check predicate on the "task" field. It's identical to TaskEQ.
func Task(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldTask, v))
}

// TaskEQ applies the EQ predicate on the "task" field.
func TaskEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldTask, v))
}

// TaskNEQ applies the NEQ predicate on the "task" field.
func TaskNEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldTask, v))
}

// TaskIn applies the In predicate on the "task" field.
func TaskIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldTask, vs...))
}

// TaskNotIn applies the NotIn predicate on the "task" field.
func TaskNotIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldTask, vs...))
}

// TaskGT applies the GT predicate on the "task" field.
func TaskGT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldTask, v))
}

// TaskGTE applies the GTE predicate on the "task" field.
func TaskGTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldTask, v))
}

// TaskLT applies the LT predicate on the "task" field.
func TaskLT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldTask, v))
}

// TaskLTE applies the LTE predicate on the "task" field.
func TaskLTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldTask, v))
}

// TaskContains applies the Contains predicate on the "task" field.
func TaskContains(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContains(FieldTask, v))
}

// TaskHasPrefix applies the HasPrefix predicate on the "task" field.
func TaskHasPrefix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasPrefix(FieldTask, v))
}

// TaskHasSuffix applies the HasSuffix predicate on the "task" field.
func TaskHasSuffix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasSuffix(FieldTask, v))
}

// TaskEqualFold applies the EqualFold predicate on the "task" field.
func TaskEqualFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEqualFold(FieldTask, v))
}

// TaskContainsFold applies the ContainsFold predicate on the "task" field.
func TaskContainsFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContainsFold(FieldTask, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldStatus, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		p(s.Not())
	})
}
