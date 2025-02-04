// Code generated by ent, DO NOT EDIT.

package allmethodsservice

import (
	"entgo.io/ent/dialect/sql"
	"github.com/tilau2328/entcontrib/entproto/internal/entprototest/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AllMethodsService {
	return predicate.AllMethodsService(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AllMethodsService {
	return predicate.AllMethodsService(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AllMethodsService {
	return predicate.AllMethodsService(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AllMethodsService {
	return predicate.AllMethodsService(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AllMethodsService {
	return predicate.AllMethodsService(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AllMethodsService {
	return predicate.AllMethodsService(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AllMethodsService {
	return predicate.AllMethodsService(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AllMethodsService {
	return predicate.AllMethodsService(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AllMethodsService {
	return predicate.AllMethodsService(sql.FieldLTE(FieldID, id))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AllMethodsService) predicate.AllMethodsService {
	return predicate.AllMethodsService(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AllMethodsService) predicate.AllMethodsService {
	return predicate.AllMethodsService(func(s *sql.Selector) {
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
func Not(p predicate.AllMethodsService) predicate.AllMethodsService {
	return predicate.AllMethodsService(func(s *sql.Selector) {
		p(s.Not())
	})
}
