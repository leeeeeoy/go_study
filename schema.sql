CREATE TABLE public.categories (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	"name" varchar NOT NULL,
	created_at timestamptz NOT NULL DEFAULT now(),
	CONSTRAINT categories_pkey PRIMARY KEY (id)
);

CREATE TABLE public.hashtags (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	value varchar NOT NULL,
	used_count int8 NOT NULL DEFAULT 0,
	created_at timestamptz NOT NULL DEFAULT now(),
	CONSTRAINT hashtags_pkey PRIMARY KEY (id)
);

CREATE TABLE public.report_types (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	description varchar NULL,
	in_active bool NOT NULL DEFAULT false,
	order_num int8 NOT NULL,
	CONSTRAINT report_types_pkey PRIMARY KEY (id)
);

CREATE TABLE public.users (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	email varchar NOT NULL,
	"password" varchar NOT NULL,
	username varchar NOT NULL,
	created_at timestamptz NOT NULL DEFAULT now(),
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE public.topics (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	"name" varchar NOT NULL,
	created_at timestamptz NOT NULL DEFAULT now(),
	category_id uuid NULL,
	CONSTRAINT topics_pkey PRIMARY KEY (id),
	CONSTRAINT topics_categories_topics FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE SET NULL
);

CREATE TABLE public.posts (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	title varchar NOT NULL,
	contents varchar NOT NULL,
	like_count int8 NOT NULL DEFAULT 0,
	comment_count int8 NOT NULL DEFAULT 0,
	view_count int8 NOT NULL DEFAULT 0,
	report_count int8 NOT NULL DEFAULT 0,
	"status" public."status" NOT NULL DEFAULT 'activate'::status,
	private bool NOT NULL DEFAULT false,
	language_type varchar NOT NULL,
	attachments varchar NULL,
	created_at timestamptz NOT NULL DEFAULT now(),
	updated_at timestamptz NOT NULL DEFAULT now(),
	topic_id uuid NULL,
	user_id uuid NULL,
	CONSTRAINT posts_pkey PRIMARY KEY (id),
	CONSTRAINT posts_topics_posts FOREIGN KEY (topic_id) REFERENCES public.topics(id) ON DELETE SET NULL,
	CONSTRAINT posts_users_posts FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL
);

CREATE TABLE public.bookmarks (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	created_at timestamptz NOT NULL DEFAULT now(),
	post_id uuid NULL,
	user_id uuid NULL,
	CONSTRAINT bookmarks_pkey PRIMARY KEY (id),
	CONSTRAINT bookmarks_posts_book_marks FOREIGN KEY (post_id) REFERENCES public.posts(id) ON DELETE SET NULL,
	CONSTRAINT bookmarks_users_book_marks FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL
);

CREATE TABLE public.post_comments (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	contents varchar NOT NULL,
	like_count int8 NOT NULL DEFAULT 0,
	"status" public."status" NOT NULL DEFAULT 'activate'::status,
	report_count int8 NOT NULL DEFAULT 0,
	language_type varchar NOT NULL,
	author_heart bool NOT NULL DEFAULT false,
	created_at timestamptz NOT NULL DEFAULT now(),
	updated_at timestamptz NOT NULL DEFAULT now(),
	post_id uuid NULL,
	user_id uuid NULL,
	CONSTRAINT comments_pkey PRIMARY KEY (id),
	CONSTRAINT comments_postss_comments FOREIGN KEY (post_id) REFERENCES public.posts(id) ON DELETE SET NULL,
	CONSTRAINT comments_users_comments FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL
);

CREATE TABLE public.post_hashtags (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	post_id uuid NULL,
	hashtag_id uuid NULL,
	CONSTRAINT poar_hashtags_pkey PRIMARY KEY (id),
	CONSTRAINT post_hashtags_hashtags_post_hashtag FOREIGN KEY (hashtag_id) REFERENCES public.hashtags(id) ON DELETE SET NULL,
	CONSTRAINT post_hashtags_posts_post_hashtag FOREIGN KEY (post_id) REFERENCES public.posts(id) ON DELETE SET NULL
);

CREATE TABLE public.post_likes (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	created_at timestamptz NOT NULL DEFAULT now(),
	post_id uuid NULL,
	user_id uuid NULL,
	CONSTRAINT post_likes_pkey PRIMARY KEY (id),
	CONSTRAINT post_likes_posts_post_like FOREIGN KEY (post_id) REFERENCES public.posts(id) ON DELETE SET NULL,
	CONSTRAINT post_likes_users_post_like FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL
);

CREATE TABLE public.post_reports (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	"comments" varchar NULL,
	"status" public."status" NOT NULL DEFAULT 'activate'::status,
	created_at timestamptz NOT NULL DEFAULT now(),
	updated_at timestamptz NOT NULL DEFAULT now(),
	post_id uuid NULL,
	report_type_id uuid NULL,
	reporter_id uuid NULL,
	CONSTRAINT post_reports_pkey PRIMARY KEY (id),
	CONSTRAINT post_reports_posts_post_report FOREIGN KEY (post_id) REFERENCES public.posts(id) ON DELETE SET NULL,
	CONSTRAINT post_reports_report_types_post_report FOREIGN KEY (report_type_id) REFERENCES public.report_types(id) ON DELETE SET NULL,
	CONSTRAINT post_reports_users_post_report FOREIGN KEY (reporter_id) REFERENCES public.users(id) ON DELETE SET NULL
);

CREATE TABLE public.comment_likes (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	created_at timestamptz NOT NULL DEFAULT now(),
	comment_id uuid NULL,
	user_id uuid NULL,
	CONSTRAINT comment_likes_pkey PRIMARY KEY (id),
	CONSTRAINT comment_likes_comments_comment_like FOREIGN KEY (comment_id) REFERENCES public.post_comments(id) ON DELETE SET NULL,
	CONSTRAINT comment_likes_users_comment_like FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL
);

CREATE TABLE public.comment_mentions (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	comment_id uuid NULL,
	user_id uuid NULL,
	CONSTRAINT comment_mentions_pkey PRIMARY KEY (id),
	CONSTRAINT comment_mentions_comments_comment_mention FOREIGN KEY (comment_id) REFERENCES public.post_comments(id) ON DELETE SET NULL,
	CONSTRAINT comment_mentions_users_comment_mention FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE SET NULL
);

CREATE TABLE public.comment_reports (
	id uuid NOT NULL DEFAULT uuid_generate_v1(),
	"desc" varchar NULL,
	"status" public."status" NOT NULL DEFAULT 'activate'::status,
	created_at timestamptz NOT NULL DEFAULT now(),
	updated_at timestamptz NOT NULL DEFAULT now(),
	comment_id uuid NULL,
	report_type_id uuid NULL,
	reporter_id uuid NULL,
	CONSTRAINT comment_reports_pkey PRIMARY KEY (id),
	CONSTRAINT comment_reports_comments_comment_report FOREIGN KEY (comment_id) REFERENCES public.post_comments(id) ON DELETE SET NULL,
	CONSTRAINT comment_reports_report_types_comment_report FOREIGN KEY (report_type_id) REFERENCES public.report_types(id) ON DELETE SET NULL,
	CONSTRAINT comment_reports_users_comment_report FOREIGN KEY (reporter_id) REFERENCES public.users(id) ON DELETE SET NULL
);