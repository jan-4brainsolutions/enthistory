// Code generated by enthistory, DO NOT EDIT.

package schema

import (
	"_examples/basic/ent/schema/mixins"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

type ResidenceHistory struct {
	ent.Schema
}

func (ResidenceHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Immutable(),
		field.Time("history_time").
			Immutable().
			Default(time.Now),
		field.Enum("operation").
			Immutable().
			GoType(enthistory.OpType("")),
		field.UUID("ref", uuid.UUID{}).
			Optional().
			Immutable(),
		field.String("name").
			Immutable()}
}
func (ResidenceHistory) Edges() []ent.Edge {
	return nil
}
func (ResidenceHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "residence_history"}}
}
func (ResidenceHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{mixins.TimeMixin{}}
}
func (ResidenceHistory) Indexes() []ent.Index {
	return []ent.Index{index.Fields("history_time")}
}
