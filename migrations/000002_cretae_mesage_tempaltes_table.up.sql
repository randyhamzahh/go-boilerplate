CREATE TABLE public.message_templates (
	id bigserial NOT NULL,
	name text NULL,
	message text NULL,
	user_id bigint NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT message_templates_pkey PRIMARY KEY (id),
    CONSTRAINT message_templates_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id)
);
