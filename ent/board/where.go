// Code generated by ent, DO NOT EDIT.

package board

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/leeeeeoy/go_study/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldTitle, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldText, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldUserID, v))
}

// LikeCount applies equality check predicate on the "like_count" field. It's identical to LikeCountEQ.
func LikeCount(v int) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldLikeCount, v))
}

// CommentCount applies equality check predicate on the "comment_count" field. It's identical to CommentCountEQ.
func CommentCount(v int) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldCommentCount, v))
}

// ViewCount applies equality check predicate on the "view_count" field. It's identical to ViewCountEQ.
func ViewCount(v int) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldViewCount, v))
}

// ReportCount applies equality check predicate on the "report_count" field. It's identical to ReportCountEQ.
func ReportCount(v int) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldReportCount, v))
}

// LanguageType applies equality check predicate on the "language_type" field. It's identical to LanguageTypeEQ.
func LanguageType(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldLanguageType, v))
}

// Attachments applies equality check predicate on the "attachments" field. It's identical to AttachmentsEQ.
func Attachments(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldAttachments, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Board {
	return predicate.Board(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Board {
	return predicate.Board(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Board {
	return predicate.Board(sql.FieldContainsFold(FieldTitle, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Board {
	return predicate.Board(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Board {
	return predicate.Board(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Board {
	return predicate.Board(sql.FieldContainsFold(FieldText, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Board {
	return predicate.Board(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Board {
	return predicate.Board(sql.FieldNotNull(FieldUserID))
}

// LikeCountEQ applies the EQ predicate on the "like_count" field.
func LikeCountEQ(v int) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldLikeCount, v))
}

// LikeCountNEQ applies the NEQ predicate on the "like_count" field.
func LikeCountNEQ(v int) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldLikeCount, v))
}

// LikeCountIn applies the In predicate on the "like_count" field.
func LikeCountIn(vs ...int) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldLikeCount, vs...))
}

// LikeCountNotIn applies the NotIn predicate on the "like_count" field.
func LikeCountNotIn(vs ...int) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldLikeCount, vs...))
}

// LikeCountGT applies the GT predicate on the "like_count" field.
func LikeCountGT(v int) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldLikeCount, v))
}

// LikeCountGTE applies the GTE predicate on the "like_count" field.
func LikeCountGTE(v int) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldLikeCount, v))
}

// LikeCountLT applies the LT predicate on the "like_count" field.
func LikeCountLT(v int) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldLikeCount, v))
}

// LikeCountLTE applies the LTE predicate on the "like_count" field.
func LikeCountLTE(v int) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldLikeCount, v))
}

// CommentCountEQ applies the EQ predicate on the "comment_count" field.
func CommentCountEQ(v int) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldCommentCount, v))
}

// CommentCountNEQ applies the NEQ predicate on the "comment_count" field.
func CommentCountNEQ(v int) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldCommentCount, v))
}

// CommentCountIn applies the In predicate on the "comment_count" field.
func CommentCountIn(vs ...int) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldCommentCount, vs...))
}

// CommentCountNotIn applies the NotIn predicate on the "comment_count" field.
func CommentCountNotIn(vs ...int) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldCommentCount, vs...))
}

// CommentCountGT applies the GT predicate on the "comment_count" field.
func CommentCountGT(v int) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldCommentCount, v))
}

// CommentCountGTE applies the GTE predicate on the "comment_count" field.
func CommentCountGTE(v int) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldCommentCount, v))
}

// CommentCountLT applies the LT predicate on the "comment_count" field.
func CommentCountLT(v int) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldCommentCount, v))
}

// CommentCountLTE applies the LTE predicate on the "comment_count" field.
func CommentCountLTE(v int) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldCommentCount, v))
}

// ViewCountEQ applies the EQ predicate on the "view_count" field.
func ViewCountEQ(v int) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldViewCount, v))
}

// ViewCountNEQ applies the NEQ predicate on the "view_count" field.
func ViewCountNEQ(v int) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldViewCount, v))
}

// ViewCountIn applies the In predicate on the "view_count" field.
func ViewCountIn(vs ...int) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldViewCount, vs...))
}

// ViewCountNotIn applies the NotIn predicate on the "view_count" field.
func ViewCountNotIn(vs ...int) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldViewCount, vs...))
}

// ViewCountGT applies the GT predicate on the "view_count" field.
func ViewCountGT(v int) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldViewCount, v))
}

// ViewCountGTE applies the GTE predicate on the "view_count" field.
func ViewCountGTE(v int) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldViewCount, v))
}

// ViewCountLT applies the LT predicate on the "view_count" field.
func ViewCountLT(v int) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldViewCount, v))
}

// ViewCountLTE applies the LTE predicate on the "view_count" field.
func ViewCountLTE(v int) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldViewCount, v))
}

// ReportCountEQ applies the EQ predicate on the "report_count" field.
func ReportCountEQ(v int) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldReportCount, v))
}

// ReportCountNEQ applies the NEQ predicate on the "report_count" field.
func ReportCountNEQ(v int) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldReportCount, v))
}

// ReportCountIn applies the In predicate on the "report_count" field.
func ReportCountIn(vs ...int) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldReportCount, vs...))
}

// ReportCountNotIn applies the NotIn predicate on the "report_count" field.
func ReportCountNotIn(vs ...int) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldReportCount, vs...))
}

// ReportCountGT applies the GT predicate on the "report_count" field.
func ReportCountGT(v int) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldReportCount, v))
}

// ReportCountGTE applies the GTE predicate on the "report_count" field.
func ReportCountGTE(v int) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldReportCount, v))
}

// ReportCountLT applies the LT predicate on the "report_count" field.
func ReportCountLT(v int) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldReportCount, v))
}

// ReportCountLTE applies the LTE predicate on the "report_count" field.
func ReportCountLTE(v int) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldReportCount, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldStatus, vs...))
}

// LanguageTypeEQ applies the EQ predicate on the "language_type" field.
func LanguageTypeEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldLanguageType, v))
}

// LanguageTypeNEQ applies the NEQ predicate on the "language_type" field.
func LanguageTypeNEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldLanguageType, v))
}

// LanguageTypeIn applies the In predicate on the "language_type" field.
func LanguageTypeIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldLanguageType, vs...))
}

// LanguageTypeNotIn applies the NotIn predicate on the "language_type" field.
func LanguageTypeNotIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldLanguageType, vs...))
}

// LanguageTypeGT applies the GT predicate on the "language_type" field.
func LanguageTypeGT(v string) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldLanguageType, v))
}

// LanguageTypeGTE applies the GTE predicate on the "language_type" field.
func LanguageTypeGTE(v string) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldLanguageType, v))
}

// LanguageTypeLT applies the LT predicate on the "language_type" field.
func LanguageTypeLT(v string) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldLanguageType, v))
}

// LanguageTypeLTE applies the LTE predicate on the "language_type" field.
func LanguageTypeLTE(v string) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldLanguageType, v))
}

// LanguageTypeContains applies the Contains predicate on the "language_type" field.
func LanguageTypeContains(v string) predicate.Board {
	return predicate.Board(sql.FieldContains(FieldLanguageType, v))
}

// LanguageTypeHasPrefix applies the HasPrefix predicate on the "language_type" field.
func LanguageTypeHasPrefix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasPrefix(FieldLanguageType, v))
}

// LanguageTypeHasSuffix applies the HasSuffix predicate on the "language_type" field.
func LanguageTypeHasSuffix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasSuffix(FieldLanguageType, v))
}

// LanguageTypeEqualFold applies the EqualFold predicate on the "language_type" field.
func LanguageTypeEqualFold(v string) predicate.Board {
	return predicate.Board(sql.FieldEqualFold(FieldLanguageType, v))
}

// LanguageTypeContainsFold applies the ContainsFold predicate on the "language_type" field.
func LanguageTypeContainsFold(v string) predicate.Board {
	return predicate.Board(sql.FieldContainsFold(FieldLanguageType, v))
}

// AttachmentsEQ applies the EQ predicate on the "attachments" field.
func AttachmentsEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldAttachments, v))
}

// AttachmentsNEQ applies the NEQ predicate on the "attachments" field.
func AttachmentsNEQ(v string) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldAttachments, v))
}

// AttachmentsIn applies the In predicate on the "attachments" field.
func AttachmentsIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldAttachments, vs...))
}

// AttachmentsNotIn applies the NotIn predicate on the "attachments" field.
func AttachmentsNotIn(vs ...string) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldAttachments, vs...))
}

// AttachmentsGT applies the GT predicate on the "attachments" field.
func AttachmentsGT(v string) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldAttachments, v))
}

// AttachmentsGTE applies the GTE predicate on the "attachments" field.
func AttachmentsGTE(v string) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldAttachments, v))
}

// AttachmentsLT applies the LT predicate on the "attachments" field.
func AttachmentsLT(v string) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldAttachments, v))
}

// AttachmentsLTE applies the LTE predicate on the "attachments" field.
func AttachmentsLTE(v string) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldAttachments, v))
}

// AttachmentsContains applies the Contains predicate on the "attachments" field.
func AttachmentsContains(v string) predicate.Board {
	return predicate.Board(sql.FieldContains(FieldAttachments, v))
}

// AttachmentsHasPrefix applies the HasPrefix predicate on the "attachments" field.
func AttachmentsHasPrefix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasPrefix(FieldAttachments, v))
}

// AttachmentsHasSuffix applies the HasSuffix predicate on the "attachments" field.
func AttachmentsHasSuffix(v string) predicate.Board {
	return predicate.Board(sql.FieldHasSuffix(FieldAttachments, v))
}

// AttachmentsIsNil applies the IsNil predicate on the "attachments" field.
func AttachmentsIsNil() predicate.Board {
	return predicate.Board(sql.FieldIsNull(FieldAttachments))
}

// AttachmentsNotNil applies the NotNil predicate on the "attachments" field.
func AttachmentsNotNil() predicate.Board {
	return predicate.Board(sql.FieldNotNull(FieldAttachments))
}

// AttachmentsEqualFold applies the EqualFold predicate on the "attachments" field.
func AttachmentsEqualFold(v string) predicate.Board {
	return predicate.Board(sql.FieldEqualFold(FieldAttachments, v))
}

// AttachmentsContainsFold applies the ContainsFold predicate on the "attachments" field.
func AttachmentsContainsFold(v string) predicate.Board {
	return predicate.Board(sql.FieldContainsFold(FieldAttachments, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Board {
	return predicate.Board(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Board {
	return predicate.Board(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Board {
	return predicate.Board(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		step := newCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBoardLike applies the HasEdge predicate on the "board_like" edge.
func HasBoardLike() predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BoardLikeTable, BoardLikeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBoardLikeWith applies the HasEdge predicate on the "board_like" edge with a given conditions (other predicates).
func HasBoardLikeWith(preds ...predicate.BoardLike) predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		step := newBoardLikeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBoardHashtag applies the HasEdge predicate on the "board_hashtag" edge.
func HasBoardHashtag() predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BoardHashtagTable, BoardHashtagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBoardHashtagWith applies the HasEdge predicate on the "board_hashtag" edge with a given conditions (other predicates).
func HasBoardHashtagWith(preds ...predicate.BoardHashtag) predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		step := newBoardHashtagStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBoardReport applies the HasEdge predicate on the "board_report" edge.
func HasBoardReport() predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BoardReportTable, BoardReportColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBoardReportWith applies the HasEdge predicate on the "board_report" edge with a given conditions (other predicates).
func HasBoardReportWith(preds ...predicate.BoardReport) predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		step := newBoardReportStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Board) predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Board) predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Board) predicate.Board {
	return predicate.Board(func(s *sql.Selector) {
		p(s.Not())
	})
}
