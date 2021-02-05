// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/B6001186/Contagions/ent/area"
	"github.com/B6001186/Contagions/ent/disease"
	"github.com/B6001186/Contagions/ent/employee"
	"github.com/B6001186/Contagions/ent/level"
	"github.com/B6001186/Contagions/ent/statistic"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// AreaCreate is the builder for creating a Area entity.
type AreaCreate struct {
	config
	mutation *AreaMutation
	hooks    []Hook
}

// SetAreaName sets the AreaName field.
func (ac *AreaCreate) SetAreaName(s string) *AreaCreate {
	ac.mutation.SetAreaName(s)
	return ac
}

// SetAreaDistrict sets the AreaDistrict field.
func (ac *AreaCreate) SetAreaDistrict(s string) *AreaCreate {
	ac.mutation.SetAreaDistrict(s)
	return ac
}

// SetAreaSubDistrict sets the AreaSubDistrict field.
func (ac *AreaCreate) SetAreaSubDistrict(s string) *AreaCreate {
	ac.mutation.SetAreaSubDistrict(s)
	return ac
}

// SetDiseaseID sets the disease edge to Disease by id.
func (ac *AreaCreate) SetDiseaseID(id int) *AreaCreate {
	ac.mutation.SetDiseaseID(id)
	return ac
}

// SetNillableDiseaseID sets the disease edge to Disease by id if the given value is not nil.
func (ac *AreaCreate) SetNillableDiseaseID(id *int) *AreaCreate {
	if id != nil {
		ac = ac.SetDiseaseID(*id)
	}
	return ac
}

// SetDisease sets the disease edge to Disease.
func (ac *AreaCreate) SetDisease(d *Disease) *AreaCreate {
	return ac.SetDiseaseID(d.ID)
}

// SetStatisticID sets the statistic edge to Statistic by id.
func (ac *AreaCreate) SetStatisticID(id int) *AreaCreate {
	ac.mutation.SetStatisticID(id)
	return ac
}

// SetNillableStatisticID sets the statistic edge to Statistic by id if the given value is not nil.
func (ac *AreaCreate) SetNillableStatisticID(id *int) *AreaCreate {
	if id != nil {
		ac = ac.SetStatisticID(*id)
	}
	return ac
}

// SetStatistic sets the statistic edge to Statistic.
func (ac *AreaCreate) SetStatistic(s *Statistic) *AreaCreate {
	return ac.SetStatisticID(s.ID)
}

// SetLevelID sets the level edge to Level by id.
func (ac *AreaCreate) SetLevelID(id int) *AreaCreate {
	ac.mutation.SetLevelID(id)
	return ac
}

// SetNillableLevelID sets the level edge to Level by id if the given value is not nil.
func (ac *AreaCreate) SetNillableLevelID(id *int) *AreaCreate {
	if id != nil {
		ac = ac.SetLevelID(*id)
	}
	return ac
}

// SetLevel sets the level edge to Level.
func (ac *AreaCreate) SetLevel(l *Level) *AreaCreate {
	return ac.SetLevelID(l.ID)
}

// SetEmployeeID sets the employee edge to Employee by id.
func (ac *AreaCreate) SetEmployeeID(id int) *AreaCreate {
	ac.mutation.SetEmployeeID(id)
	return ac
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (ac *AreaCreate) SetNillableEmployeeID(id *int) *AreaCreate {
	if id != nil {
		ac = ac.SetEmployeeID(*id)
	}
	return ac
}

// SetEmployee sets the employee edge to Employee.
func (ac *AreaCreate) SetEmployee(e *Employee) *AreaCreate {
	return ac.SetEmployeeID(e.ID)
}

// Mutation returns the AreaMutation object of the builder.
func (ac *AreaCreate) Mutation() *AreaMutation {
	return ac.mutation
}

// Save creates the Area in the database.
func (ac *AreaCreate) Save(ctx context.Context) (*Area, error) {
	if _, ok := ac.mutation.AreaName(); !ok {
		return nil, &ValidationError{Name: "AreaName", err: errors.New("ent: missing required field \"AreaName\"")}
	}
	if v, ok := ac.mutation.AreaName(); ok {
		if err := area.AreaNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "AreaName", err: fmt.Errorf("ent: validator failed for field \"AreaName\": %w", err)}
		}
	}
	if _, ok := ac.mutation.AreaDistrict(); !ok {
		return nil, &ValidationError{Name: "AreaDistrict", err: errors.New("ent: missing required field \"AreaDistrict\"")}
	}
	if v, ok := ac.mutation.AreaDistrict(); ok {
		if err := area.AreaDistrictValidator(v); err != nil {
			return nil, &ValidationError{Name: "AreaDistrict", err: fmt.Errorf("ent: validator failed for field \"AreaDistrict\": %w", err)}
		}
	}
	if _, ok := ac.mutation.AreaSubDistrict(); !ok {
		return nil, &ValidationError{Name: "AreaSubDistrict", err: errors.New("ent: missing required field \"AreaSubDistrict\"")}
	}
	if v, ok := ac.mutation.AreaSubDistrict(); ok {
		if err := area.AreaSubDistrictValidator(v); err != nil {
			return nil, &ValidationError{Name: "AreaSubDistrict", err: fmt.Errorf("ent: validator failed for field \"AreaSubDistrict\": %w", err)}
		}
	}
	var (
		err  error
		node *Area
	)
	if len(ac.hooks) == 0 {
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AreaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AreaCreate) SaveX(ctx context.Context) *Area {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *AreaCreate) sqlSave(ctx context.Context) (*Area, error) {
	a, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	a.ID = int(id)
	return a, nil
}

func (ac *AreaCreate) createSpec() (*Area, *sqlgraph.CreateSpec) {
	var (
		a     = &Area{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: area.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: area.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.AreaName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: area.FieldAreaName,
		})
		a.AreaName = value
	}
	if value, ok := ac.mutation.AreaDistrict(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: area.FieldAreaDistrict,
		})
		a.AreaDistrict = value
	}
	if value, ok := ac.mutation.AreaSubDistrict(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: area.FieldAreaSubDistrict,
		})
		a.AreaSubDistrict = value
	}
	if nodes := ac.mutation.DiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   area.DiseaseTable,
			Columns: []string{area.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.StatisticIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   area.StatisticTable,
			Columns: []string{area.StatisticColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statistic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.LevelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   area.LevelTable,
			Columns: []string{area.LevelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: level.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   area.EmployeeTable,
			Columns: []string{area.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return a, _spec
}
