// Code generated by entc, DO NOT EDIT.

package machine

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ip-rw/ransack/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Hwid applies equality check predicate on the "hwid" field. It's identical to HwidEQ.
func Hwid(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHwid), v))
	})
}

// Hostname applies equality check predicate on the "hostname" field. It's identical to HostnameEQ.
func Hostname(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostname), v))
	})
}

// Fingerprint applies equality check predicate on the "fingerprint" field. It's identical to FingerprintEQ.
func Fingerprint(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFingerprint), v))
	})
}

// HwidEQ applies the EQ predicate on the "hwid" field.
func HwidEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHwid), v))
	})
}

// HwidNEQ applies the NEQ predicate on the "hwid" field.
func HwidNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHwid), v))
	})
}

// HwidIn applies the In predicate on the "hwid" field.
func HwidIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHwid), v...))
	})
}

// HwidNotIn applies the NotIn predicate on the "hwid" field.
func HwidNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHwid), v...))
	})
}

// HwidGT applies the GT predicate on the "hwid" field.
func HwidGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHwid), v))
	})
}

// HwidGTE applies the GTE predicate on the "hwid" field.
func HwidGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHwid), v))
	})
}

// HwidLT applies the LT predicate on the "hwid" field.
func HwidLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHwid), v))
	})
}

// HwidLTE applies the LTE predicate on the "hwid" field.
func HwidLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHwid), v))
	})
}

// HwidContains applies the Contains predicate on the "hwid" field.
func HwidContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHwid), v))
	})
}

// HwidHasPrefix applies the HasPrefix predicate on the "hwid" field.
func HwidHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHwid), v))
	})
}

// HwidHasSuffix applies the HasSuffix predicate on the "hwid" field.
func HwidHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHwid), v))
	})
}

// HwidEqualFold applies the EqualFold predicate on the "hwid" field.
func HwidEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHwid), v))
	})
}

// HwidContainsFold applies the ContainsFold predicate on the "hwid" field.
func HwidContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHwid), v))
	})
}

// HostnameEQ applies the EQ predicate on the "hostname" field.
func HostnameEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostname), v))
	})
}

// HostnameNEQ applies the NEQ predicate on the "hostname" field.
func HostnameNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHostname), v))
	})
}

// HostnameIn applies the In predicate on the "hostname" field.
func HostnameIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHostname), v...))
	})
}

// HostnameNotIn applies the NotIn predicate on the "hostname" field.
func HostnameNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHostname), v...))
	})
}

// HostnameGT applies the GT predicate on the "hostname" field.
func HostnameGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHostname), v))
	})
}

// HostnameGTE applies the GTE predicate on the "hostname" field.
func HostnameGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHostname), v))
	})
}

// HostnameLT applies the LT predicate on the "hostname" field.
func HostnameLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHostname), v))
	})
}

// HostnameLTE applies the LTE predicate on the "hostname" field.
func HostnameLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHostname), v))
	})
}

// HostnameContains applies the Contains predicate on the "hostname" field.
func HostnameContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHostname), v))
	})
}

// HostnameHasPrefix applies the HasPrefix predicate on the "hostname" field.
func HostnameHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHostname), v))
	})
}

// HostnameHasSuffix applies the HasSuffix predicate on the "hostname" field.
func HostnameHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHostname), v))
	})
}

// HostnameEqualFold applies the EqualFold predicate on the "hostname" field.
func HostnameEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHostname), v))
	})
}

// HostnameContainsFold applies the ContainsFold predicate on the "hostname" field.
func HostnameContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHostname), v))
	})
}

// FingerprintEQ applies the EQ predicate on the "fingerprint" field.
func FingerprintEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFingerprint), v))
	})
}

// FingerprintNEQ applies the NEQ predicate on the "fingerprint" field.
func FingerprintNEQ(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFingerprint), v))
	})
}

// FingerprintIn applies the In predicate on the "fingerprint" field.
func FingerprintIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFingerprint), v...))
	})
}

// FingerprintNotIn applies the NotIn predicate on the "fingerprint" field.
func FingerprintNotIn(vs ...string) predicate.Machine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Machine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFingerprint), v...))
	})
}

// FingerprintGT applies the GT predicate on the "fingerprint" field.
func FingerprintGT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFingerprint), v))
	})
}

// FingerprintGTE applies the GTE predicate on the "fingerprint" field.
func FingerprintGTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFingerprint), v))
	})
}

// FingerprintLT applies the LT predicate on the "fingerprint" field.
func FingerprintLT(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFingerprint), v))
	})
}

// FingerprintLTE applies the LTE predicate on the "fingerprint" field.
func FingerprintLTE(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFingerprint), v))
	})
}

// FingerprintContains applies the Contains predicate on the "fingerprint" field.
func FingerprintContains(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFingerprint), v))
	})
}

// FingerprintHasPrefix applies the HasPrefix predicate on the "fingerprint" field.
func FingerprintHasPrefix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFingerprint), v))
	})
}

// FingerprintHasSuffix applies the HasSuffix predicate on the "fingerprint" field.
func FingerprintHasSuffix(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFingerprint), v))
	})
}

// FingerprintEqualFold applies the EqualFold predicate on the "fingerprint" field.
func FingerprintEqualFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFingerprint), v))
	})
}

// FingerprintContainsFold applies the ContainsFold predicate on the "fingerprint" field.
func FingerprintContainsFold(v string) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFingerprint), v))
	})
}

// HasIps applies the HasEdge predicate on the "ips" edge.
func HasIps() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IpsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, IpsTable, IpsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIpsWith applies the HasEdge predicate on the "ips" edge with a given conditions (other predicates).
func HasIpsWith(preds ...predicate.IP) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IpsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, IpsTable, IpsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasKeys applies the HasEdge predicate on the "keys" edge.
func HasKeys() predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(KeysTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, KeysTable, KeysColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasKeysWith applies the HasEdge predicate on the "keys" edge with a given conditions (other predicates).
func HasKeysWith(preds ...predicate.Key) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(KeysInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, KeysTable, KeysColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Machine) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Machine) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
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
func Not(p predicate.Machine) predicate.Machine {
	return predicate.Machine(func(s *sql.Selector) {
		p(s.Not())
	})
}