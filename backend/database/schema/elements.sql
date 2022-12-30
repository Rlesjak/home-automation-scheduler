CREATE TABLE public.elements
(
    id bigint NOT NULL DEFAULT nextval('global_id_sequence'),
    uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    description text,
    type character varying(8) NOT NULL,
    command text,
    group_id bigint NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (group_id)
        REFERENCES public.element_groups (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)
WITH (
    OIDS = FALSE
);

ALTER TABLE IF EXISTS public.elements
    OWNER to postgres;