-- name: GetTriggersInGroup :many
SELECT 
triggers.uuid, 
triggers."name", 
triggers.description, 
triggers.description, 
triggers."type", 
triggers.condition, 
triggers.command,
elements.uuid as element_uuid,
triggers.active
FROM 
public.triggers
LEFT JOIN public.elements ON elements.id = triggers.element_id,
( SELECT id FROM public.trigger_groups WHERE uuid = @parentuuid::uuid ) trgrp
WHERE triggers.group_id = trgrp.id;

-- name: GetTriggerByUuid :one
SELECT 
triggers.uuid, 
triggers."name", 
triggers.description, 
triggers.description, 
triggers."type", 
triggers.condition, 
triggers.command,
elements.uuid as element_uuid,
triggers.active
FROM 
public.triggers
LEFT JOIN public.elements ON elements.id = triggers.element_id
WHERE triggers.uuid = $1;

-- name: CreateStandaloneTrigger :one
INSERT INTO public.triggers ("name", description, type, condition, command, group_id, active)
SELECT $1, $2, $3, $4, $5, grp.id, $6
FROM public.trigger_groups grp 
WHERE grp.uuid = @groupuuid::uuid
RETURNING uuid;

-- name: CreateTriggerForElement :one
INSERT INTO public.triggers (name, description, type, condition, element_id, group_id, active)
SELECT
$1,
$2,
$3,
$4,
elm.id,
grp.id,
$5
FROM
(SELECT id FROM public.trigger_groups WHERE uuid = @groupuuid::uuid ) grp,
(SELECT id FROM public.elements WHERE uuid = @elmuuid::uuid ) elm
RETURNING uuid;

-- name: UpdateTriggerActive :exec
UPDATE public.triggers
SET active = $1
WHERE uuid = $2;

-- name: DeleteTrigger :exec
DELETE FROM public.triggers
WHERE uuid = @uuid::uuid;