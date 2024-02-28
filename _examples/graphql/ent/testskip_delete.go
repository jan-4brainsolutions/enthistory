// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/predicate"
	"_examples/graphql/ent/testskip"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestSkipDelete is the builder for deleting a TestSkip entity.
type TestSkipDelete struct {
	config
	hooks    []Hook
	mutation *TestSkipMutation
}

// Where appends a list predicates to the TestSkipDelete builder.
func (tsd *TestSkipDelete) Where(ps ...predicate.TestSkip) *TestSkipDelete {
	tsd.mutation.Where(ps...)
	return tsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tsd *TestSkipDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tsd.sqlExec, tsd.mutation, tsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tsd *TestSkipDelete) ExecX(ctx context.Context) int {
	n, err := tsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tsd *TestSkipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(testskip.Table, sqlgraph.NewFieldSpec(testskip.FieldID, field.TypeUUID))
	if ps := tsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tsd.mutation.done = true
	return affected, err
}

// TestSkipDeleteOne is the builder for deleting a single TestSkip entity.
type TestSkipDeleteOne struct {
	tsd *TestSkipDelete
}

// Where appends a list predicates to the TestSkipDelete builder.
func (tsdo *TestSkipDeleteOne) Where(ps ...predicate.TestSkip) *TestSkipDeleteOne {
	tsdo.tsd.mutation.Where(ps...)
	return tsdo
}

// Exec executes the deletion query.
func (tsdo *TestSkipDeleteOne) Exec(ctx context.Context) error {
	n, err := tsdo.tsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{testskip.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tsdo *TestSkipDeleteOne) ExecX(ctx context.Context) {
	if err := tsdo.Exec(ctx); err != nil {
		panic(err)
	}
}