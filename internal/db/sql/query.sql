-- name: CreateJob :one
INSERT INTO jobs (jid, data) VALUES ($1, $2)
RETURNING jid, data;

-- name: GetJob :one
SELECT jid, data FROM jobs WHERE jid = $1;

-- name: UpdateJob :exec
UPDATE jobs SET data = $2 WHERE jid = $1;

-- name: DeleteJob :exec
DELETE FROM jobs WHERE jid = $1;

-- name: ListJobs :many
SELECT jid, data FROM jobs;