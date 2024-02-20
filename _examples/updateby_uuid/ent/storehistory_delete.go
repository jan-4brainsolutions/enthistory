// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/updateby_uuid/ent/predicate"
	"_examples/updateby_uuid/ent/storehistory"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StoreHistoryDelete is the builder for deleting a StoreHistory entity.
type StoreHistoryDelete struct {
	config
	hooks    []Hook
	mutation *StoreHistoryMutation
}

// Where appends a list predicates to the StoreHistoryDelete builder.
func (shd *StoreHistoryDelete) Where(ps ...predicate.StoreHistory) *StoreHistoryDelete {
	shd.mutation.Where(ps...)
	return shd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (shd *StoreHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, shd.sqlExec, shd.mutation, shd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (shd *StoreHistoryDelete) ExecX(ctx context.Context) int {
	n, err := shd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (shd *StoreHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(storehistory.Table, sqlgraph.NewFieldSpec(storehistory.FieldID, field.TypeInt))
	if ps := shd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, shd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	shd.mutation.done = true
	return affected, err
}

// StoreHistoryDeleteOne is the builder for deleting a single StoreHistory entity.
type StoreHistoryDeleteOne struct {
	shd *StoreHistoryDelete
}

// Where appends a list predicates to the StoreHistoryDelete builder.
func (shdo *StoreHistoryDeleteOne) Where(ps ...predicate.StoreHistory) *StoreHistoryDeleteOne {
	shdo.shd.mutation.Where(ps...)
	return shdo
}

// Exec executes the deletion query.
func (shdo *StoreHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := shdo.shd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{storehistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (shdo *StoreHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := shdo.Exec(ctx); err != nil {
		panic(err)
	}
}
