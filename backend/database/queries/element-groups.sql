-- name: GetMasterElementGroups :many
SELECT uuid, name, description
FROM public.element_groups
WHERE parent_group_id IS NULL;

-- name: GetChildElementGroupsOf :many
SELECT grp.uuid, grp.name, grp.description
FROM
public.element_groups grp,
(SELECT id FROM public.element_groups WHERE uuid = @parentUuid::uuid) ids
WHERE grp.parent_group_id = ids.id;

-- name: CreateChildElementsGroup :one
INSERT INTO public.element_groups (parent_group_id, "name", description)
SELECT id, @name::text, @description::text
FROM public.element_groups 
WHERE uuid = @parentUuid::uuid
RETURNING uuid;

-- name: CreateMasterElementsGroup :one
INSERT INTO public.element_groups (parent_group_id, "name", description)
VALUES (NULL, @name::text, @description::text)
RETURNING uuid;

-- name: DeleteElementsGroup :exec
DELETE FROM public.element_groups
WHERE uuid = @uuid::uuid;