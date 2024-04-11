// Code generated by enthistory, DO NOT EDIT.

package otherschema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

type FriendshipHistory struct {
	ent.Schema
}

func (FriendshipHistory) Fields() []ent.Field {
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
		field.UUID("character_id", uuid.UUID{}),
		field.UUID("friend_id", uuid.UUID{})}
}
func (FriendshipHistory) Edges() []ent.Edge {
	return nil
}
func (FriendshipHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "friendship_history"}, enthistory.Annotations{IsHistory: true, Triggers: []enthistory.OpType{enthistory.OpTypeInsert, enthistory.OpTypeUpdate, enthistory.OpTypeDelete}}}
}
