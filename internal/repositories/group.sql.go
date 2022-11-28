// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: group.sql

package repositories

import (
	"context"
	"time"
)

const createGroup = `-- name: CreateGroup :one
INSERT INTO "group" (
  group_id,
  group_name,
  created_by
) VALUES (
  $1, $2, $3
)
RETURNING group_id, group_name, created_by, created_at, description
`

type CreateGroupParams struct {
	GroupID   string `json:"group_id"`
	GroupName string `json:"group_name"`
	CreatedBy string `json:"created_by"`
}

func (q *Queries) CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error) {
	row := q.db.QueryRowContext(ctx, createGroup, arg.GroupID, arg.GroupName, arg.CreatedBy)
	var i Group
	err := row.Scan(
		&i.GroupID,
		&i.GroupName,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.Description,
	)
	return i, err
}

const deleteGroup = `-- name: DeleteGroup :exec
DELETE FROM "group"
WHERE group_id = $1
`

func (q *Queries) DeleteGroup(ctx context.Context, groupID string) error {
	_, err := q.db.ExecContext(ctx, deleteGroup, groupID)
	return err
}

const getGroup = `-- name: GetGroup :one
SELECT group_id, group_name, created_by, created_at, description
FROM "group"
WHERE group_id = $1
`

func (q *Queries) GetGroup(ctx context.Context, groupID string) (Group, error) {
	row := q.db.QueryRowContext(ctx, getGroup, groupID)
	var i Group
	err := row.Scan(
		&i.GroupID,
		&i.GroupName,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.Description,
	)
	return i, err
}

const listGroupOwned = `-- name: ListGroupOwned :many
SELECT "group".group_id, group_name, created_by, "group".created_at, description, user_id, ug.group_id, role, status, ug.created_at
FROM "group"
JOIN "user_group" ug using (group_id)
WHERE ug.user_id = $1
AND ug.role = 'owner'
AND ug.status = 'joined'
ORDER BY group_id
`

type ListGroupOwnedRow struct {
	GroupID     string    `json:"group_id"`
	GroupName   string    `json:"group_name"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	GroupID_2   string    `json:"group_id_2"`
	Role        string    `json:"role"`
	Status      string    `json:"status"`
	CreatedAt_2 time.Time `json:"created_at_2"`
}

func (q *Queries) ListGroupOwned(ctx context.Context, userID string) ([]ListGroupOwnedRow, error) {
	rows, err := q.db.QueryContext(ctx, listGroupOwned, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListGroupOwnedRow{}
	for rows.Next() {
		var i ListGroupOwnedRow
		if err := rows.Scan(
			&i.GroupID,
			&i.GroupName,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Description,
			&i.UserID,
			&i.GroupID_2,
			&i.Role,
			&i.Status,
			&i.CreatedAt_2,
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
