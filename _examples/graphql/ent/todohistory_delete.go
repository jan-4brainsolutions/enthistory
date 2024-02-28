// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/predicate"
	"_examples/graphql/ent/todohistory"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoHistoryDelete is the builder for deleting a TodoHistory entity.
type TodoHistoryDelete struct {
	config
	hooks    []Hook
	mutation *TodoHistoryMutation
}

// Where appends a list predicates to the TodoHistoryDelete builder.
func (thd *TodoHistoryDelete) Where(ps ...predicate.TodoHistory) *TodoHistoryDelete {
	thd.mutation.Where(ps...)
	return thd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (thd *TodoHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, thd.sqlExec, thd.mutation, thd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (thd *TodoHistoryDelete) ExecX(ctx context.Context) int {
	n, err := thd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (thd *TodoHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(todohistory.Table, sqlgraph.NewFieldSpec(todohistory.FieldID, field.TypeUUID))
	if ps := thd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, thd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	thd.mutation.done = true
	return affected, err
}

// TodoHistoryDeleteOne is the builder for deleting a single TodoHistory entity.
type TodoHistoryDeleteOne struct {
	thd *TodoHistoryDelete
}

// Where appends a list predicates to the TodoHistoryDelete builder.
func (thdo *TodoHistoryDeleteOne) Where(ps ...predicate.TodoHistory) *TodoHistoryDeleteOne {
	thdo.thd.mutation.Where(ps...)
	return thdo
}

// Exec executes the deletion query.
func (thdo *TodoHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := thdo.thd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{todohistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (thdo *TodoHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := thdo.Exec(ctx); err != nil {
		panic(err)
	}
}