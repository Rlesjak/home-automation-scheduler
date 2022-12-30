CREATE TABLE public.element_groups
(
    id bigint NOT NULL DEFAULT nextval('global_id_sequence'),
    uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
    parent_group_id bigint,
    name text NOT NULL,
    description text,
    PRIMARY KEY (id),
    FOREIGN KEY (parent_group_id) REFERENCES public.element_groups(id) ON DELETE CASCADE
)
WITH (
    OIDS = FALSE
);

ALTER TABLE IF EXISTS public.element_groups
    OWNER to postgres;