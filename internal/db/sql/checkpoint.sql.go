// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: checkpoint.sql

package sql

import (
	"context"
	"strings"
	"time"
)

const createCheckpoint = `-- name: CreateCheckpoint :exec
INSERT INTO checkpoints (ID, JID, Path, Time, Size) VALUES (?, ?, ?, ?, ?)
`

type CreateCheckpointParams struct {
	ID   string
	Jid  string
	Path string
	Time time.Time
	Size int64
}

func (q *Queries) CreateCheckpoint(ctx context.Context, arg CreateCheckpointParams) error {
	_, err := q.db.ExecContext(ctx, createCheckpoint,
		arg.ID,
		arg.Jid,
		arg.Path,
		arg.Time,
		arg.Size,
	)
	return err
}

const deleteCheckpoint = `-- name: DeleteCheckpoint :exec
DELETE FROM checkpoints WHERE ID = ?
`

func (q *Queries) DeleteCheckpoint(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCheckpoint, id)
	return err
}

const listCheckpoints = `-- name: ListCheckpoints :many
SELECT ID, JID, Path, Time, Size FROM checkpoints ORDER BY Time DESC
`

func (q *Queries) ListCheckpoints(ctx context.Context) ([]Checkpoint, error) {
	rows, err := q.db.QueryContext(ctx, listCheckpoints)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Checkpoint
	for rows.Next() {
		var i Checkpoint
		if err := rows.Scan(
			&i.ID,
			&i.Jid,
			&i.Path,
			&i.Time,
			&i.Size,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCheckpointsByIDs = `-- name: ListCheckpointsByIDs :many
SELECT ID, JID, Path, Time, Size FROM checkpoints WHERE ID in (/*SLICE:ids*/?)
ORDER BY Time DESC
`

func (q *Queries) ListCheckpointsByIDs(ctx context.Context, ids []string) ([]Checkpoint, error) {
	query := listCheckpointsByIDs
	var queryParams []interface{}
	if len(ids) > 0 {
		for _, v := range ids {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:ids*/?", strings.Repeat(",?", len(ids))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:ids*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Checkpoint
	for rows.Next() {
		var i Checkpoint
		if err := rows.Scan(
			&i.ID,
			&i.Jid,
			&i.Path,
			&i.Time,
			&i.Size,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCheckpointsByJIDs = `-- name: ListCheckpointsByJIDs :many
SELECT ID, JID, Path, Time, Size FROM checkpoints WHERE JID in (/*SLICE:jids*/?)
ORDER BY Time DESC
`

func (q *Queries) ListCheckpointsByJIDs(ctx context.Context, jids []string) ([]Checkpoint, error) {
	query := listCheckpointsByJIDs
	var queryParams []interface{}
	if len(jids) > 0 {
		for _, v := range jids {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:jids*/?", strings.Repeat(",?", len(jids))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:jids*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Checkpoint
	for rows.Next() {
		var i Checkpoint
		if err := rows.Scan(
			&i.ID,
			&i.Jid,
			&i.Path,
			&i.Time,
			&i.Size,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCheckpoint = `-- name: UpdateCheckpoint :exec
UPDATE checkpoints SET
    JID = ?,
    Path = ?,
    Time = ?,
    Size = ?
WHERE ID = ?
`

type UpdateCheckpointParams struct {
	Jid  string
	Path string
	Time time.Time
	Size int64
	ID   string
}

func (q *Queries) UpdateCheckpoint(ctx context.Context, arg UpdateCheckpointParams) error {
	_, err := q.db.ExecContext(ctx, updateCheckpoint,
		arg.Jid,
		arg.Path,
		arg.Time,
		arg.Size,
		arg.ID,
	)
	return err
}
