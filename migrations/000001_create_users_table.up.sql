CREATE TABLE public.users (
	id bigserial NOT NULL,
	name text NULL,
	email text NULL,
	"password" text NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
