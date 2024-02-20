// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/custompaths/internal/ent/characterhistory"
	"_examples/custompaths/internal/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CharacterHistoryDelete is the builder for deleting a CharacterHistory entity.
type CharacterHistoryDelete struct {
	config
	hooks    []Hook
	mutation *CharacterHistoryMutation
}

// Where appends a list predicates to the CharacterHistoryDelete builder.
func (chd *CharacterHistoryDelete) Where(ps ...predicate.CharacterHistory) *CharacterHistoryDelete {
	chd.mutation.Where(ps...)
	return chd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (chd *CharacterHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, chd.sqlExec, chd.mutation, chd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (chd *CharacterHistoryDelete) ExecX(ctx context.Context) int {
	n, err := chd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (chd *CharacterHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(characterhistory.Table, sqlgraph.NewFieldSpec(characterhistory.FieldID, field.TypeInt))
	if ps := chd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, chd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	chd.mutation.done = true
	return affected, err
}

// CharacterHistoryDeleteOne is the builder for deleting a single CharacterHistory entity.
type CharacterHistoryDeleteOne struct {
	chd *CharacterHistoryDelete
}

// Where appends a list predicates to the CharacterHistoryDelete builder.
func (chdo *CharacterHistoryDeleteOne) Where(ps ...predicate.CharacterHistory) *CharacterHistoryDeleteOne {
	chdo.chd.mutation.Where(ps...)
	return chdo
}

// Exec executes the deletion query.
func (chdo *CharacterHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := chdo.chd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{characterhistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (chdo *CharacterHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := chdo.Exec(ctx); err != nil {
		panic(err)
	}
}
