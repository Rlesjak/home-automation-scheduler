-- name: GetMasterTriggerGroups :many
SELECT uuid, name, description
FROM public.trigger_groups
WHERE parent_group_id IS NULL;

-- name: GetChildTriggerGroupsOf :many
SELECT grp.uuid, grp.name, grp.description
FROM
public.trigger_groups grp,
(SELECT id FROM public.trigger_groups WHERE uuid = @parentUuid::uuid) ids
WHERE grp.parent_group_id = ids.id;

-- name: CreateChildTriggersGroup :one
INSERT INTO public.trigger_groups (parent_group_id, "name", description)
SELECT id, @name::text, @description::text
FROM public.trigger_groups 
WHERE uuid = @parentUuid::uuid
RETURNING uuid;

-- name: CreateMasterTriggersGroup :one
INSERT INTO public.trigger_groups (parent_group_id, "name", description)
VALUES (NULL, @name::text, @description::text)
RETURNING uuid;

-- name: DeleteTriggersGroup :exec
DELETE FROM public.trigger_groups
WHERE uuid = @uuid::uuid;