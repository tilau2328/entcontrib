// Code generated by ent, DO NOT EDIT.

package task

import (
	"entgo.io/ent/dialect/sql"
	"github.com/tilau2328/entcontrib/entproto/cmd/protoc-gen-ent/internal/todo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDescription, v))
}

// Complete applies equality check predicate on the "complete" field. It's identical to CompleteEQ.
func Complete(v bool) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldComplete, v))
}

// Signature applies equality check predicate on the "signature" field. It's identical to SignatureEQ.
func Signature(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldSignature, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleIsNil applies the IsNil predicate on the "title" field.
func TitleIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldTitle))
}

// TitleNotNil applies the NotNil predicate on the "title" field.
func TitleNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldTitle))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldDescription, v))
}

// CompleteEQ applies the EQ predicate on the "complete" field.
func CompleteEQ(v bool) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldComplete, v))
}

// CompleteNEQ applies the NEQ predicate on the "complete" field.
func CompleteNEQ(v bool) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldComplete, v))
}

// SignatureEQ applies the EQ predicate on the "signature" field.
func SignatureEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldSignature, v))
}

// SignatureNEQ applies the NEQ predicate on the "signature" field.
func SignatureNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldSignature, v))
}

// SignatureIn applies the In predicate on the "signature" field.
func SignatureIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldSignature, vs...))
}

// SignatureNotIn applies the NotIn predicate on the "signature" field.
func SignatureNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldSignature, vs...))
}

// SignatureGT applies the GT predicate on the "signature" field.
func SignatureGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldSignature, v))
}

// SignatureGTE applies the GTE predicate on the "signature" field.
func SignatureGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldSignature, v))
}

// SignatureLT applies the LT predicate on the "signature" field.
func SignatureLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldSignature, v))
}

// SignatureLTE applies the LTE predicate on the "signature" field.
func SignatureLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldSignature, v))
}

// SignatureContains applies the Contains predicate on the "signature" field.
func SignatureContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldSignature, v))
}

// SignatureHasPrefix applies the HasPrefix predicate on the "signature" field.
func SignatureHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldSignature, v))
}

// SignatureHasSuffix applies the HasSuffix predicate on the "signature" field.
func SignatureHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldSignature, v))
}

// SignatureEqualFold applies the EqualFold predicate on the "signature" field.
func SignatureEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldSignature, v))
}

// SignatureContainsFold applies the ContainsFold predicate on the "signature" field.
func SignatureContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldSignature, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
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
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		p(s.Not())
	})
}
