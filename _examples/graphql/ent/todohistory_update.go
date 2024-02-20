// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/predicate"
	"_examples/graphql/ent/todohistory"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TodoHistoryUpdate is the builder for updating TodoHistory entities.
type TodoHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *TodoHistoryMutation
}

// Where appends a list predicates to the TodoHistoryUpdate builder.
func (thu *TodoHistoryUpdate) Where(ps ...predicate.TodoHistory) *TodoHistoryUpdate {
	thu.mutation.Where(ps...)
	return thu
}

// SetOtherID sets the "other_id" field.
func (thu *TodoHistoryUpdate) SetOtherID(u uuid.UUID) *TodoHistoryUpdate {
	thu.mutation.SetOtherID(u)
	return thu
}

// SetNillableOtherID sets the "other_id" field if the given value is not nil.
func (thu *TodoHistoryUpdate) SetNillableOtherID(u *uuid.UUID) *TodoHistoryUpdate {
	if u != nil {
		thu.SetOtherID(*u)
	}
	return thu
}

// ClearOtherID clears the value of the "other_id" field.
func (thu *TodoHistoryUpdate) ClearOtherID() *TodoHistoryUpdate {
	thu.mutation.ClearOtherID()
	return thu
}

// SetName sets the "name" field.
func (thu *TodoHistoryUpdate) SetName(s string) *TodoHistoryUpdate {
	thu.mutation.SetName(s)
	return thu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (thu *TodoHistoryUpdate) SetNillableName(s *string) *TodoHistoryUpdate {
	if s != nil {
		thu.SetName(*s)
	}
	return thu
}

// Mutation returns the TodoHistoryMutation object of the builder.
func (thu *TodoHistoryUpdate) Mutation() *TodoHistoryMutation {
	return thu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (thu *TodoHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, thu.sqlSave, thu.mutation, thu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (thu *TodoHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := thu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (thu *TodoHistoryUpdate) Exec(ctx context.Context) error {
	_, err := thu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thu *TodoHistoryUpdate) ExecX(ctx context.Context) {
	if err := thu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (thu *TodoHistoryUpdate) check() error {
	if v, ok := thu.mutation.Name(); ok {
		if err := todohistory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TodoHistory.name": %w`, err)}
		}
	}
	return nil
}

func (thu *TodoHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := thu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(todohistory.Table, todohistory.Columns, sqlgraph.NewFieldSpec(todohistory.FieldID, field.TypeUUID))
	if ps := thu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if thu.mutation.RefCleared() {
		_spec.ClearField(todohistory.FieldRef, field.TypeUUID)
	}
	if thu.mutation.UpdatedByCleared() {
		_spec.ClearField(todohistory.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := thu.mutation.OtherID(); ok {
		_spec.SetField(todohistory.FieldOtherID, field.TypeUUID, value)
	}
	if thu.mutation.OtherIDCleared() {
		_spec.ClearField(todohistory.FieldOtherID, field.TypeUUID)
	}
	if value, ok := thu.mutation.Name(); ok {
		_spec.SetField(todohistory.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, thu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todohistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	thu.mutation.done = true
	return n, nil
}

// TodoHistoryUpdateOne is the builder for updating a single TodoHistory entity.
type TodoHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoHistoryMutation
}

// SetOtherID sets the "other_id" field.
func (thuo *TodoHistoryUpdateOne) SetOtherID(u uuid.UUID) *TodoHistoryUpdateOne {
	thuo.mutation.SetOtherID(u)
	return thuo
}

// SetNillableOtherID sets the "other_id" field if the given value is not nil.
func (thuo *TodoHistoryUpdateOne) SetNillableOtherID(u *uuid.UUID) *TodoHistoryUpdateOne {
	if u != nil {
		thuo.SetOtherID(*u)
	}
	return thuo
}

// ClearOtherID clears the value of the "other_id" field.
func (thuo *TodoHistoryUpdateOne) ClearOtherID() *TodoHistoryUpdateOne {
	thuo.mutation.ClearOtherID()
	return thuo
}

// SetName sets the "name" field.
func (thuo *TodoHistoryUpdateOne) SetName(s string) *TodoHistoryUpdateOne {
	thuo.mutation.SetName(s)
	return thuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (thuo *TodoHistoryUpdateOne) SetNillableName(s *string) *TodoHistoryUpdateOne {
	if s != nil {
		thuo.SetName(*s)
	}
	return thuo
}

// Mutation returns the TodoHistoryMutation object of the builder.
func (thuo *TodoHistoryUpdateOne) Mutation() *TodoHistoryMutation {
	return thuo.mutation
}

// Where appends a list predicates to the TodoHistoryUpdate builder.
func (thuo *TodoHistoryUpdateOne) Where(ps ...predicate.TodoHistory) *TodoHistoryUpdateOne {
	thuo.mutation.Where(ps...)
	return thuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (thuo *TodoHistoryUpdateOne) Select(field string, fields ...string) *TodoHistoryUpdateOne {
	thuo.fields = append([]string{field}, fields...)
	return thuo
}

// Save executes the query and returns the updated TodoHistory entity.
func (thuo *TodoHistoryUpdateOne) Save(ctx context.Context) (*TodoHistory, error) {
	return withHooks(ctx, thuo.sqlSave, thuo.mutation, thuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (thuo *TodoHistoryUpdateOne) SaveX(ctx context.Context) *TodoHistory {
	node, err := thuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (thuo *TodoHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := thuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thuo *TodoHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := thuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (thuo *TodoHistoryUpdateOne) check() error {
	if v, ok := thuo.mutation.Name(); ok {
		if err := todohistory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TodoHistory.name": %w`, err)}
		}
	}
	return nil
}

func (thuo *TodoHistoryUpdateOne) sqlSave(ctx context.Context) (_node *TodoHistory, err error) {
	if err := thuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(todohistory.Table, todohistory.Columns, sqlgraph.NewFieldSpec(todohistory.FieldID, field.TypeUUID))
	id, ok := thuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TodoHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := thuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todohistory.FieldID)
		for _, f := range fields {
			if !todohistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todohistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := thuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if thuo.mutation.RefCleared() {
		_spec.ClearField(todohistory.FieldRef, field.TypeUUID)
	}
	if thuo.mutation.UpdatedByCleared() {
		_spec.ClearField(todohistory.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := thuo.mutation.OtherID(); ok {
		_spec.SetField(todohistory.FieldOtherID, field.TypeUUID, value)
	}
	if thuo.mutation.OtherIDCleared() {
		_spec.ClearField(todohistory.FieldOtherID, field.TypeUUID)
	}
	if value, ok := thuo.mutation.Name(); ok {
		_spec.SetField(todohistory.FieldName, field.TypeString, value)
	}
	_node = &TodoHistory{config: thuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, thuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todohistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	thuo.mutation.done = true
	return _node, nil
}
