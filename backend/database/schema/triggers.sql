CREATE TABLE public.triggers
(
    id bigint NOT NULL DEFAULT nextval('global_id_sequence'),
    uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
    name text NOT NULL,
    description text,
    type character varying(8) NOT NULL DEFAULT 'ELEMENT',
    condition text,
    command text,
    element_id bigint,
    group_id bigint NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (group_id)
        REFERENCES public.trigger_groups (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    FOREIGN KEY (element_id)
        REFERENCES public.elements (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)
WITH (
    OIDS = FALSE
);

ALTER TABLE IF EXISTS public.triggers
    OWNER to postgres;