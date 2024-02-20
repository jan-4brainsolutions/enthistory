// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"_examples/without_updatedby/ent/predicate"
	"_examples/without_updatedby/ent/residencehistory"
)

// ResidenceHistoryUpdate is the builder for updating ResidenceHistory entities.
type ResidenceHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *ResidenceHistoryMutation
}

// Where appends a list predicates to the ResidenceHistoryUpdate builder.
func (rhu *ResidenceHistoryUpdate) Where(ps ...predicate.ResidenceHistory) *ResidenceHistoryUpdate {
	rhu.mutation.Where(ps...)
	return rhu
}

// SetUpdatedAt sets the "updated_at" field.
func (rhu *ResidenceHistoryUpdate) SetUpdatedAt(t time.Time) *ResidenceHistoryUpdate {
	rhu.mutation.SetUpdatedAt(t)
	return rhu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rhu *ResidenceHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *ResidenceHistoryUpdate {
	if t != nil {
		rhu.SetUpdatedAt(*t)
	}
	return rhu
}

// SetName sets the "name" field.
func (rhu *ResidenceHistoryUpdate) SetName(s string) *ResidenceHistoryUpdate {
	rhu.mutation.SetName(s)
	return rhu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rhu *ResidenceHistoryUpdate) SetNillableName(s *string) *ResidenceHistoryUpdate {
	if s != nil {
		rhu.SetName(*s)
	}
	return rhu
}

// Mutation returns the ResidenceHistoryMutation object of the builder.
func (rhu *ResidenceHistoryUpdate) Mutation() *ResidenceHistoryMutation {
	return rhu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rhu *ResidenceHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rhu.sqlSave, rhu.mutation, rhu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rhu *ResidenceHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := rhu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rhu *ResidenceHistoryUpdate) Exec(ctx context.Context) error {
	_, err := rhu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rhu *ResidenceHistoryUpdate) ExecX(ctx context.Context) {
	if err := rhu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rhu *ResidenceHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(residencehistory.Table, residencehistory.Columns, sqlgraph.NewFieldSpec(residencehistory.FieldID, field.TypeInt))
	if ps := rhu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if rhu.mutation.RefCleared() {
		_spec.ClearField(residencehistory.FieldRef, field.TypeUUID)
	}
	if value, ok := rhu.mutation.UpdatedAt(); ok {
		_spec.SetField(residencehistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rhu.mutation.Name(); ok {
		_spec.SetField(residencehistory.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rhu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{residencehistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rhu.mutation.done = true
	return n, nil
}

// ResidenceHistoryUpdateOne is the builder for updating a single ResidenceHistory entity.
type ResidenceHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ResidenceHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (rhuo *ResidenceHistoryUpdateOne) SetUpdatedAt(t time.Time) *ResidenceHistoryUpdateOne {
	rhuo.mutation.SetUpdatedAt(t)
	return rhuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rhuo *ResidenceHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *ResidenceHistoryUpdateOne {
	if t != nil {
		rhuo.SetUpdatedAt(*t)
	}
	return rhuo
}

// SetName sets the "name" field.
func (rhuo *ResidenceHistoryUpdateOne) SetName(s string) *ResidenceHistoryUpdateOne {
	rhuo.mutation.SetName(s)
	return rhuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rhuo *ResidenceHistoryUpdateOne) SetNillableName(s *string) *ResidenceHistoryUpdateOne {
	if s != nil {
		rhuo.SetName(*s)
	}
	return rhuo
}

// Mutation returns the ResidenceHistoryMutation object of the builder.
func (rhuo *ResidenceHistoryUpdateOne) Mutation() *ResidenceHistoryMutation {
	return rhuo.mutation
}

// Where appends a list predicates to the ResidenceHistoryUpdate builder.
func (rhuo *ResidenceHistoryUpdateOne) Where(ps ...predicate.ResidenceHistory) *ResidenceHistoryUpdateOne {
	rhuo.mutation.Where(ps...)
	return rhuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rhuo *ResidenceHistoryUpdateOne) Select(field string, fields ...string) *ResidenceHistoryUpdateOne {
	rhuo.fields = append([]string{field}, fields...)
	return rhuo
}

// Save executes the query and returns the updated ResidenceHistory entity.
func (rhuo *ResidenceHistoryUpdateOne) Save(ctx context.Context) (*ResidenceHistory, error) {
	return withHooks(ctx, rhuo.sqlSave, rhuo.mutation, rhuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rhuo *ResidenceHistoryUpdateOne) SaveX(ctx context.Context) *ResidenceHistory {
	node, err := rhuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rhuo *ResidenceHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := rhuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rhuo *ResidenceHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := rhuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rhuo *ResidenceHistoryUpdateOne) sqlSave(ctx context.Context) (_node *ResidenceHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(residencehistory.Table, residencehistory.Columns, sqlgraph.NewFieldSpec(residencehistory.FieldID, field.TypeInt))
	id, ok := rhuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ResidenceHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rhuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, residencehistory.FieldID)
		for _, f := range fields {
			if !residencehistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != residencehistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rhuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if rhuo.mutation.RefCleared() {
		_spec.ClearField(residencehistory.FieldRef, field.TypeUUID)
	}
	if value, ok := rhuo.mutation.UpdatedAt(); ok {
		_spec.SetField(residencehistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := rhuo.mutation.Name(); ok {
		_spec.SetField(residencehistory.FieldName, field.TypeString, value)
	}
	_node = &ResidenceHistory{config: rhuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rhuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{residencehistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rhuo.mutation.done = true
	return _node, nil
}
