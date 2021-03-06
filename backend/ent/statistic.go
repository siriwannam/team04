// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/B6001186/Contagions/ent/statistic"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Statistic is the model entity for the Statistic schema.
type Statistic struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StatisticName holds the value of the "StatisticName" field.
	StatisticName string `json:"StatisticName,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StatisticQuery when eager-loading is set.
	Edges StatisticEdges `json:"edges"`
}

// StatisticEdges holds the relations/edges for other nodes in the graph.
type StatisticEdges struct {
	// Area holds the value of the area edge.
	Area []*Area
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AreaOrErr returns the Area value or an error if the edge
// was not loaded in eager-loading.
func (e StatisticEdges) AreaOrErr() ([]*Area, error) {
	if e.loadedTypes[0] {
		return e.Area, nil
	}
	return nil, &NotLoadedError{edge: "area"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Statistic) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // StatisticName
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Statistic fields.
func (s *Statistic) assignValues(values ...interface{}) error {
	if m, n := len(values), len(statistic.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field StatisticName", values[0])
	} else if value.Valid {
		s.StatisticName = value.String
	}
	return nil
}

// QueryArea queries the area edge of the Statistic.
func (s *Statistic) QueryArea() *AreaQuery {
	return (&StatisticClient{config: s.config}).QueryArea(s)
}

// Update returns a builder for updating this Statistic.
// Note that, you need to call Statistic.Unwrap() before calling this method, if this Statistic
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Statistic) Update() *StatisticUpdateOne {
	return (&StatisticClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Statistic) Unwrap() *Statistic {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Statistic is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Statistic) String() string {
	var builder strings.Builder
	builder.WriteString("Statistic(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", StatisticName=")
	builder.WriteString(s.StatisticName)
	builder.WriteByte(')')
	return builder.String()
}

// Statistics is a parsable slice of Statistic.
type Statistics []*Statistic

func (s Statistics) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
