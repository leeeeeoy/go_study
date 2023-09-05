// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BoardsColumns holds the columns for the "boards" table.
	BoardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "text", Type: field.TypeString},
		{Name: "like_count", Type: field.TypeInt},
		{Name: "comment_count", Type: field.TypeInt},
		{Name: "view_count", Type: field.TypeInt},
		{Name: "report_count", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"0", "1"}},
		{Name: "private", Type: field.TypeBool, Default: false},
		{Name: "language_type", Type: field.TypeString},
		{Name: "attachments", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "topic_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// BoardsTable holds the schema information for the "boards" table.
	BoardsTable = &schema.Table{
		Name:       "boards",
		Columns:    BoardsColumns,
		PrimaryKey: []*schema.Column{BoardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "boards_topics_boards",
				Columns:    []*schema.Column{BoardsColumns[13]},
				RefColumns: []*schema.Column{TopicsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "boards_users_boards",
				Columns:    []*schema.Column{BoardsColumns[14]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BoardHashtagsColumns holds the columns for the "board_hashtags" table.
	BoardHashtagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "board_id", Type: field.TypeInt, Nullable: true},
		{Name: "hashtag_id", Type: field.TypeInt, Nullable: true},
	}
	// BoardHashtagsTable holds the schema information for the "board_hashtags" table.
	BoardHashtagsTable = &schema.Table{
		Name:       "board_hashtags",
		Columns:    BoardHashtagsColumns,
		PrimaryKey: []*schema.Column{BoardHashtagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "board_hashtags_boards_board_hashtag",
				Columns:    []*schema.Column{BoardHashtagsColumns[1]},
				RefColumns: []*schema.Column{BoardsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "board_hashtags_hashtags_board_hashtag",
				Columns:    []*schema.Column{BoardHashtagsColumns[2]},
				RefColumns: []*schema.Column{HashtagsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BoardLikesColumns holds the columns for the "board_likes" table.
	BoardLikesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "board_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// BoardLikesTable holds the schema information for the "board_likes" table.
	BoardLikesTable = &schema.Table{
		Name:       "board_likes",
		Columns:    BoardLikesColumns,
		PrimaryKey: []*schema.Column{BoardLikesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "board_likes_boards_board_like",
				Columns:    []*schema.Column{BoardLikesColumns[2]},
				RefColumns: []*schema.Column{BoardsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "board_likes_users_board_like",
				Columns:    []*schema.Column{BoardLikesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BoardReportsColumns holds the columns for the "board_reports" table.
	BoardReportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "comment", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"0", "1"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "board_id", Type: field.TypeInt, Nullable: true},
		{Name: "report_type_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "reporter_id", Type: field.TypeInt, Nullable: true},
	}
	// BoardReportsTable holds the schema information for the "board_reports" table.
	BoardReportsTable = &schema.Table{
		Name:       "board_reports",
		Columns:    BoardReportsColumns,
		PrimaryKey: []*schema.Column{BoardReportsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "board_reports_boards_board_report",
				Columns:    []*schema.Column{BoardReportsColumns[5]},
				RefColumns: []*schema.Column{BoardsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "board_reports_report_types_board_report",
				Columns:    []*schema.Column{BoardReportsColumns[6]},
				RefColumns: []*schema.Column{ReportTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "board_reports_users_board_report",
				Columns:    []*schema.Column{BoardReportsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BookMarksColumns holds the columns for the "book_marks" table.
	BookMarksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "board_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// BookMarksTable holds the schema information for the "book_marks" table.
	BookMarksTable = &schema.Table{
		Name:       "book_marks",
		Columns:    BookMarksColumns,
		PrimaryKey: []*schema.Column{BookMarksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "book_marks_boards_book_marks",
				Columns:    []*schema.Column{BookMarksColumns[2]},
				RefColumns: []*schema.Column{BoardsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "book_marks_users_book_marks",
				Columns:    []*schema.Column{BookMarksColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString},
		{Name: "like_count", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"0", "1"}},
		{Name: "report_count", Type: field.TypeInt},
		{Name: "language_type", Type: field.TypeString},
		{Name: "author_heart", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "board_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_boards_comments",
				Columns:    []*schema.Column{CommentsColumns[9]},
				RefColumns: []*schema.Column{BoardsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommentLikesColumns holds the columns for the "comment_likes" table.
	CommentLikesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "comment_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// CommentLikesTable holds the schema information for the "comment_likes" table.
	CommentLikesTable = &schema.Table{
		Name:       "comment_likes",
		Columns:    CommentLikesColumns,
		PrimaryKey: []*schema.Column{CommentLikesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comment_likes_comments_comment_like",
				Columns:    []*schema.Column{CommentLikesColumns[2]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comment_likes_users_comment_like",
				Columns:    []*schema.Column{CommentLikesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommentMentionsColumns holds the columns for the "comment_mentions" table.
	CommentMentionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "comment_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// CommentMentionsTable holds the schema information for the "comment_mentions" table.
	CommentMentionsTable = &schema.Table{
		Name:       "comment_mentions",
		Columns:    CommentMentionsColumns,
		PrimaryKey: []*schema.Column{CommentMentionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comment_mentions_comments_comment_mention",
				Columns:    []*schema.Column{CommentMentionsColumns[1]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comment_mentions_users_comment_mention",
				Columns:    []*schema.Column{CommentMentionsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommentReportsColumns holds the columns for the "comment_reports" table.
	CommentReportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "desc", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"0", "1"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "comment_id", Type: field.TypeInt, Nullable: true},
		{Name: "report_type_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "reporter_id", Type: field.TypeInt, Nullable: true},
	}
	// CommentReportsTable holds the schema information for the "comment_reports" table.
	CommentReportsTable = &schema.Table{
		Name:       "comment_reports",
		Columns:    CommentReportsColumns,
		PrimaryKey: []*schema.Column{CommentReportsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comment_reports_comments_comment_report",
				Columns:    []*schema.Column{CommentReportsColumns[5]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comment_reports_report_types_comment_report",
				Columns:    []*schema.Column{CommentReportsColumns[6]},
				RefColumns: []*schema.Column{ReportTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comment_reports_users_comment_report",
				Columns:    []*schema.Column{CommentReportsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// HashtagsColumns holds the columns for the "hashtags" table.
	HashtagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeString},
		{Name: "used_count", Type: field.TypeInt, Default: 0},
	}
	// HashtagsTable holds the schema information for the "hashtags" table.
	HashtagsTable = &schema.Table{
		Name:       "hashtags",
		Columns:    HashtagsColumns,
		PrimaryKey: []*schema.Column{HashtagsColumns[0]},
	}
	// ReportTypesColumns holds the columns for the "report_types" table.
	ReportTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "in_active", Type: field.TypeBool, Default: false},
		{Name: "order_num", Type: field.TypeInt},
	}
	// ReportTypesTable holds the schema information for the "report_types" table.
	ReportTypesTable = &schema.Table{
		Name:       "report_types",
		Columns:    ReportTypesColumns,
		PrimaryKey: []*schema.Column{ReportTypesColumns[0]},
	}
	// TopicsColumns holds the columns for the "topics" table.
	TopicsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "category_id", Type: field.TypeInt, Nullable: true},
	}
	// TopicsTable holds the schema information for the "topics" table.
	TopicsTable = &schema.Table{
		Name:       "topics",
		Columns:    TopicsColumns,
		PrimaryKey: []*schema.Column{TopicsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "topics_categories_topics",
				Columns:    []*schema.Column{TopicsColumns[3]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Size: 20},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Size: 10},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BoardsTable,
		BoardHashtagsTable,
		BoardLikesTable,
		BoardReportsTable,
		BookMarksTable,
		CategoriesTable,
		CommentsTable,
		CommentLikesTable,
		CommentMentionsTable,
		CommentReportsTable,
		HashtagsTable,
		ReportTypesTable,
		TopicsTable,
		UsersTable,
	}
)

func init() {
	BoardsTable.ForeignKeys[0].RefTable = TopicsTable
	BoardsTable.ForeignKeys[1].RefTable = UsersTable
	BoardHashtagsTable.ForeignKeys[0].RefTable = BoardsTable
	BoardHashtagsTable.ForeignKeys[1].RefTable = HashtagsTable
	BoardLikesTable.ForeignKeys[0].RefTable = BoardsTable
	BoardLikesTable.ForeignKeys[1].RefTable = UsersTable
	BoardReportsTable.ForeignKeys[0].RefTable = BoardsTable
	BoardReportsTable.ForeignKeys[1].RefTable = ReportTypesTable
	BoardReportsTable.ForeignKeys[2].RefTable = UsersTable
	BookMarksTable.ForeignKeys[0].RefTable = BoardsTable
	BookMarksTable.ForeignKeys[1].RefTable = UsersTable
	CommentsTable.ForeignKeys[0].RefTable = BoardsTable
	CommentsTable.ForeignKeys[1].RefTable = UsersTable
	CommentLikesTable.ForeignKeys[0].RefTable = CommentsTable
	CommentLikesTable.ForeignKeys[1].RefTable = UsersTable
	CommentMentionsTable.ForeignKeys[0].RefTable = CommentsTable
	CommentMentionsTable.ForeignKeys[1].RefTable = UsersTable
	CommentReportsTable.ForeignKeys[0].RefTable = CommentsTable
	CommentReportsTable.ForeignKeys[1].RefTable = ReportTypesTable
	CommentReportsTable.ForeignKeys[2].RefTable = UsersTable
	TopicsTable.ForeignKeys[0].RefTable = CategoriesTable
}
