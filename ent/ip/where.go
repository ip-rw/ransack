// Code generated by entc, DO NOT EDIT.

package ip

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ip-rw/ransack/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// Net applies equality check predicate on the "net" field. It's identical to NetEQ.
func Net(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNet), v))
	})
}

// Mask applies equality check predicate on the "mask" field. It's identical to MaskEQ.
func Mask(v int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMask), v))
	})
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIP), v))
	})
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIP), v))
	})
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.IP {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IP(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIP), v...))
	})
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.IP {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IP(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIP), v...))
	})
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIP), v))
	})
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIP), v))
	})
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIP), v))
	})
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIP), v))
	})
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIP), v))
	})
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIP), v))
	})
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIP), v))
	})
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIP), v))
	})
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIP), v))
	})
}

// NetEQ applies the EQ predicate on the "net" field.
func NetEQ(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNet), v))
	})
}

// NetNEQ applies the NEQ predicate on the "net" field.
func NetNEQ(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNet), v))
	})
}

// NetIn applies the In predicate on the "net" field.
func NetIn(vs ...string) predicate.IP {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IP(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNet), v...))
	})
}

// NetNotIn applies the NotIn predicate on the "net" field.
func NetNotIn(vs ...string) predicate.IP {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IP(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNet), v...))
	})
}

// NetGT applies the GT predicate on the "net" field.
func NetGT(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNet), v))
	})
}

// NetGTE applies the GTE predicate on the "net" field.
func NetGTE(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNet), v))
	})
}

// NetLT applies the LT predicate on the "net" field.
func NetLT(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNet), v))
	})
}

// NetLTE applies the LTE predicate on the "net" field.
func NetLTE(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNet), v))
	})
}

// NetContains applies the Contains predicate on the "net" field.
func NetContains(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNet), v))
	})
}

// NetHasPrefix applies the HasPrefix predicate on the "net" field.
func NetHasPrefix(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNet), v))
	})
}

// NetHasSuffix applies the HasSuffix predicate on the "net" field.
func NetHasSuffix(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNet), v))
	})
}

// NetEqualFold applies the EqualFold predicate on the "net" field.
func NetEqualFold(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNet), v))
	})
}

// NetContainsFold applies the ContainsFold predicate on the "net" field.
func NetContainsFold(v string) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNet), v))
	})
}

// MaskEQ applies the EQ predicate on the "mask" field.
func MaskEQ(v int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMask), v))
	})
}

// MaskNEQ applies the NEQ predicate on the "mask" field.
func MaskNEQ(v int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMask), v))
	})
}

// MaskIn applies the In predicate on the "mask" field.
func MaskIn(vs ...int) predicate.IP {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IP(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMask), v...))
	})
}

// MaskNotIn applies the NotIn predicate on the "mask" field.
func MaskNotIn(vs ...int) predicate.IP {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.IP(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMask), v...))
	})
}

// MaskGT applies the GT predicate on the "mask" field.
func MaskGT(v int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMask), v))
	})
}

// MaskGTE applies the GTE predicate on the "mask" field.
func MaskGTE(v int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMask), v))
	})
}

// MaskLT applies the LT predicate on the "mask" field.
func MaskLT(v int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMask), v))
	})
}

// MaskLTE applies the LTE predicate on the "mask" field.
func MaskLTE(v int) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMask), v))
	})
}

// HasMachine applies the HasEdge predicate on the "machine" edge.
func HasMachine() predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MachineTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MachineTable, MachineColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMachineWith applies the HasEdge predicate on the "machine" edge with a given conditions (other predicates).
func HasMachineWith(preds ...predicate.Machine) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MachineInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MachineTable, MachineColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.IP) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.IP) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
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
func Not(p predicate.IP) predicate.IP {
	return predicate.IP(func(s *sql.Selector) {
		p(s.Not())
	})
}
