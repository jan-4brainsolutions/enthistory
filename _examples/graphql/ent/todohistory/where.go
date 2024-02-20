// Code generated by ent, DO NOT EDIT.

package todohistory

import (
	"_examples/graphql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldLTE(FieldID, id))
}

// HistoryTime applies equality check predicate on the "history_time" field. It's identical to HistoryTimeEQ.
func HistoryTime(v time.Time) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// Ref applies equality check predicate on the "ref" field. It's identical to RefEQ.
func Ref(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldRef, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// OtherID applies equality check predicate on the "other_id" field. It's identical to OtherIDEQ.
func OtherID(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldOtherID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldName, v))
}

// HistoryTimeEQ applies the EQ predicate on the "history_time" field.
func HistoryTimeEQ(v time.Time) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// HistoryTimeNEQ applies the NEQ predicate on the "history_time" field.
func HistoryTimeNEQ(v time.Time) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNEQ(FieldHistoryTime, v))
}

// HistoryTimeIn applies the In predicate on the "history_time" field.
func HistoryTimeIn(vs ...time.Time) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldIn(FieldHistoryTime, vs...))
}

// HistoryTimeNotIn applies the NotIn predicate on the "history_time" field.
func HistoryTimeNotIn(vs ...time.Time) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNotIn(FieldHistoryTime, vs...))
}

// HistoryTimeGT applies the GT predicate on the "history_time" field.
func HistoryTimeGT(v time.Time) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldGT(FieldHistoryTime, v))
}

// HistoryTimeGTE applies the GTE predicate on the "history_time" field.
func HistoryTimeGTE(v time.Time) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldGTE(FieldHistoryTime, v))
}

// HistoryTimeLT applies the LT predicate on the "history_time" field.
func HistoryTimeLT(v time.Time) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldLT(FieldHistoryTime, v))
}

// HistoryTimeLTE applies the LTE predicate on the "history_time" field.
func HistoryTimeLTE(v time.Time) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldLTE(FieldHistoryTime, v))
}

// OperationEQ applies the EQ predicate on the "operation" field.
func OperationEQ(v enthistory.OpType) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldOperation, v))
}

// OperationNEQ applies the NEQ predicate on the "operation" field.
func OperationNEQ(v enthistory.OpType) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNEQ(FieldOperation, v))
}

// OperationIn applies the In predicate on the "operation" field.
func OperationIn(vs ...enthistory.OpType) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldIn(FieldOperation, vs...))
}

// OperationNotIn applies the NotIn predicate on the "operation" field.
func OperationNotIn(vs ...enthistory.OpType) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNotIn(FieldOperation, vs...))
}

// RefEQ applies the EQ predicate on the "ref" field.
func RefEQ(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldRef, v))
}

// RefNEQ applies the NEQ predicate on the "ref" field.
func RefNEQ(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNEQ(FieldRef, v))
}

// RefIn applies the In predicate on the "ref" field.
func RefIn(vs ...uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldIn(FieldRef, vs...))
}

// RefNotIn applies the NotIn predicate on the "ref" field.
func RefNotIn(vs ...uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNotIn(FieldRef, vs...))
}

// RefGT applies the GT predicate on the "ref" field.
func RefGT(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldGT(FieldRef, v))
}

// RefGTE applies the GTE predicate on the "ref" field.
func RefGTE(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldGTE(FieldRef, v))
}

// RefLT applies the LT predicate on the "ref" field.
func RefLT(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldLT(FieldRef, v))
}

// RefLTE applies the LTE predicate on the "ref" field.
func RefLTE(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldLTE(FieldRef, v))
}

// RefIsNil applies the IsNil predicate on the "ref" field.
func RefIsNil() predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldIsNull(FieldRef))
}

// RefNotNil applies the NotNil predicate on the "ref" field.
func RefNotNil() predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNotNull(FieldRef))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNotNull(FieldUpdatedBy))
}

// OtherIDEQ applies the EQ predicate on the "other_id" field.
func OtherIDEQ(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldOtherID, v))
}

// OtherIDNEQ applies the NEQ predicate on the "other_id" field.
func OtherIDNEQ(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNEQ(FieldOtherID, v))
}

// OtherIDIn applies the In predicate on the "other_id" field.
func OtherIDIn(vs ...uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldIn(FieldOtherID, vs...))
}

// OtherIDNotIn applies the NotIn predicate on the "other_id" field.
func OtherIDNotIn(vs ...uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNotIn(FieldOtherID, vs...))
}

// OtherIDGT applies the GT predicate on the "other_id" field.
func OtherIDGT(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldGT(FieldOtherID, v))
}

// OtherIDGTE applies the GTE predicate on the "other_id" field.
func OtherIDGTE(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldGTE(FieldOtherID, v))
}

// OtherIDLT applies the LT predicate on the "other_id" field.
func OtherIDLT(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldLT(FieldOtherID, v))
}

// OtherIDLTE applies the LTE predicate on the "other_id" field.
func OtherIDLTE(v uuid.UUID) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldLTE(FieldOtherID, v))
}

// OtherIDIsNil applies the IsNil predicate on the "other_id" field.
func OtherIDIsNil() predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldIsNull(FieldOtherID))
}

// OtherIDNotNil applies the NotNil predicate on the "other_id" field.
func OtherIDNotNil() predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNotNull(FieldOtherID))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.TodoHistory {
	return predicate.TodoHistory(sql.FieldContainsFold(FieldName, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TodoHistory) predicate.TodoHistory {
	return predicate.TodoHistory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TodoHistory) predicate.TodoHistory {
	return predicate.TodoHistory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TodoHistory) predicate.TodoHistory {
	return predicate.TodoHistory(sql.NotPredicates(p))
}
