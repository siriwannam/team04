// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/B6001186/Contagions/ent/bloodtype"
	"github.com/B6001186/Contagions/ent/category"
	"github.com/B6001186/Contagions/ent/diagnosis"
	"github.com/B6001186/Contagions/ent/employee"
	"github.com/B6001186/Contagions/ent/gender"
	"github.com/B6001186/Contagions/ent/nametitle"
	"github.com/B6001186/Contagions/ent/patient"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PatientCreate is the builder for creating a Patient entity.
type PatientCreate struct {
	config
	mutation *PatientMutation
	hooks    []Hook
}

// SetIdcard sets the Idcard field.
func (pc *PatientCreate) SetIdcard(s string) *PatientCreate {
	pc.mutation.SetIdcard(s)
	return pc
}

// SetPatientName sets the PatientName field.
func (pc *PatientCreate) SetPatientName(s string) *PatientCreate {
	pc.mutation.SetPatientName(s)
	return pc
}

// SetAddress sets the Address field.
func (pc *PatientCreate) SetAddress(s string) *PatientCreate {
	pc.mutation.SetAddress(s)
	return pc
}

// SetCongenital sets the Congenital field.
func (pc *PatientCreate) SetCongenital(s string) *PatientCreate {
	pc.mutation.SetCongenital(s)
	return pc
}

// SetAllergic sets the Allergic field.
func (pc *PatientCreate) SetAllergic(s string) *PatientCreate {
	pc.mutation.SetAllergic(s)
	return pc
}

// SetEmployeeID sets the employee edge to Employee by id.
func (pc *PatientCreate) SetEmployeeID(id int) *PatientCreate {
	pc.mutation.SetEmployeeID(id)
	return pc
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (pc *PatientCreate) SetNillableEmployeeID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetEmployeeID(*id)
	}
	return pc
}

// SetEmployee sets the employee edge to Employee.
func (pc *PatientCreate) SetEmployee(e *Employee) *PatientCreate {
	return pc.SetEmployeeID(e.ID)
}

// SetCategoryID sets the category edge to Category by id.
func (pc *PatientCreate) SetCategoryID(id int) *PatientCreate {
	pc.mutation.SetCategoryID(id)
	return pc
}

// SetNillableCategoryID sets the category edge to Category by id if the given value is not nil.
func (pc *PatientCreate) SetNillableCategoryID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetCategoryID(*id)
	}
	return pc
}

// SetCategory sets the category edge to Category.
func (pc *PatientCreate) SetCategory(c *Category) *PatientCreate {
	return pc.SetCategoryID(c.ID)
}

// SetBloodtypeID sets the bloodtype edge to Bloodtype by id.
func (pc *PatientCreate) SetBloodtypeID(id int) *PatientCreate {
	pc.mutation.SetBloodtypeID(id)
	return pc
}

// SetNillableBloodtypeID sets the bloodtype edge to Bloodtype by id if the given value is not nil.
func (pc *PatientCreate) SetNillableBloodtypeID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetBloodtypeID(*id)
	}
	return pc
}

// SetBloodtype sets the bloodtype edge to Bloodtype.
func (pc *PatientCreate) SetBloodtype(b *Bloodtype) *PatientCreate {
	return pc.SetBloodtypeID(b.ID)
}

// SetGenderID sets the gender edge to Gender by id.
func (pc *PatientCreate) SetGenderID(id int) *PatientCreate {
	pc.mutation.SetGenderID(id)
	return pc
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (pc *PatientCreate) SetNillableGenderID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetGenderID(*id)
	}
	return pc
}

// SetGender sets the gender edge to Gender.
func (pc *PatientCreate) SetGender(g *Gender) *PatientCreate {
	return pc.SetGenderID(g.ID)
}

// SetNametitleID sets the nametitle edge to Nametitle by id.
func (pc *PatientCreate) SetNametitleID(id int) *PatientCreate {
	pc.mutation.SetNametitleID(id)
	return pc
}

// SetNillableNametitleID sets the nametitle edge to Nametitle by id if the given value is not nil.
func (pc *PatientCreate) SetNillableNametitleID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetNametitleID(*id)
	}
	return pc
}

// SetNametitle sets the nametitle edge to Nametitle.
func (pc *PatientCreate) SetNametitle(n *Nametitle) *PatientCreate {
	return pc.SetNametitleID(n.ID)
}

// AddDiagnosiIDs adds the diagnosis edge to Diagnosis by ids.
func (pc *PatientCreate) AddDiagnosiIDs(ids ...int) *PatientCreate {
	pc.mutation.AddDiagnosiIDs(ids...)
	return pc
}

// AddDiagnosis adds the diagnosis edges to Diagnosis.
func (pc *PatientCreate) AddDiagnosis(d ...*Diagnosis) *PatientCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pc.AddDiagnosiIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (pc *PatientCreate) Mutation() *PatientMutation {
	return pc.mutation
}

// Save creates the Patient in the database.
func (pc *PatientCreate) Save(ctx context.Context) (*Patient, error) {
	if _, ok := pc.mutation.Idcard(); !ok {
		return nil, &ValidationError{Name: "Idcard", err: errors.New("ent: missing required field \"Idcard\"")}
	}
	if v, ok := pc.mutation.Idcard(); ok {
		if err := patient.IdcardValidator(v); err != nil {
			return nil, &ValidationError{Name: "Idcard", err: fmt.Errorf("ent: validator failed for field \"Idcard\": %w", err)}
		}
	}
	if _, ok := pc.mutation.PatientName(); !ok {
		return nil, &ValidationError{Name: "PatientName", err: errors.New("ent: missing required field \"PatientName\"")}
	}
	if v, ok := pc.mutation.PatientName(); ok {
		if err := patient.PatientNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "PatientName", err: fmt.Errorf("ent: validator failed for field \"PatientName\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Address(); !ok {
		return nil, &ValidationError{Name: "Address", err: errors.New("ent: missing required field \"Address\"")}
	}
	if v, ok := pc.mutation.Address(); ok {
		if err := patient.AddressValidator(v); err != nil {
			return nil, &ValidationError{Name: "Address", err: fmt.Errorf("ent: validator failed for field \"Address\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Congenital(); !ok {
		return nil, &ValidationError{Name: "Congenital", err: errors.New("ent: missing required field \"Congenital\"")}
	}
	if v, ok := pc.mutation.Congenital(); ok {
		if err := patient.CongenitalValidator(v); err != nil {
			return nil, &ValidationError{Name: "Congenital", err: fmt.Errorf("ent: validator failed for field \"Congenital\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Allergic(); !ok {
		return nil, &ValidationError{Name: "Allergic", err: errors.New("ent: missing required field \"Allergic\"")}
	}
	if v, ok := pc.mutation.Allergic(); ok {
		if err := patient.AllergicValidator(v); err != nil {
			return nil, &ValidationError{Name: "Allergic", err: fmt.Errorf("ent: validator failed for field \"Allergic\": %w", err)}
		}
	}
	var (
		err  error
		node *Patient
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PatientCreate) SaveX(ctx context.Context) *Patient {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PatientCreate) sqlSave(ctx context.Context) (*Patient, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PatientCreate) createSpec() (*Patient, *sqlgraph.CreateSpec) {
	var (
		pa    = &Patient{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: patient.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Idcard(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldIdcard,
		})
		pa.Idcard = value
	}
	if value, ok := pc.mutation.PatientName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientName,
		})
		pa.PatientName = value
	}
	if value, ok := pc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldAddress,
		})
		pa.Address = value
	}
	if value, ok := pc.mutation.Congenital(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldCongenital,
		})
		pa.Congenital = value
	}
	if value, ok := pc.mutation.Allergic(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldAllergic,
		})
		pa.Allergic = value
	}
	if nodes := pc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.EmployeeTable,
			Columns: []string{patient.EmployeeColumn},
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
	if nodes := pc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.CategoryTable,
			Columns: []string{patient.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.BloodtypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.BloodtypeTable,
			Columns: []string{patient.BloodtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bloodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.NametitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.NametitleTable,
			Columns: []string{patient.NametitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nametitle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.DiagnosisIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.DiagnosisTable,
			Columns: []string{patient.DiagnosisColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnosis.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pa, _spec
}
