// Code generated by sqlc. DO NOT EDIT.
// source: dashboard.sql

package sqlc

import (
	"context"
)

const createDashboard = `-- name: CreateDashboard :one
INSERT INTO dashboard (id, dashboard_name, dashboard_admin)
VALUES ($1, $2, $3)
RETURNING id, dashboard_admin, dashboard_name, logo, created_at, updated_at, status, plan
`

type CreateDashboardParams struct {
	ID             string `json:"id"`
	DashboardName  string `json:"dashboard_name"`
	DashboardAdmin int32  `json:"dashboard_admin"`
}

func (q *Queries) CreateDashboard(ctx context.Context, arg CreateDashboardParams) (Dashboard, error) {
	row := q.db.QueryRowContext(ctx, createDashboard, arg.ID, arg.DashboardName, arg.DashboardAdmin)
	var i Dashboard
	err := row.Scan(
		&i.ID,
		&i.DashboardAdmin,
		&i.DashboardName,
		&i.Logo,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
		&i.Plan,
	)
	return i, err
}

const getAllDashboards = `-- name: GetAllDashboards :many
SELECT id, dashboard_admin, dashboard_name, logo, created_at, updated_at, status, plan FROM dashboard
`

func (q *Queries) GetAllDashboards(ctx context.Context) ([]Dashboard, error) {
	rows, err := q.db.QueryContext(ctx, getAllDashboards)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Dashboard
	for rows.Next() {
		var i Dashboard
		if err := rows.Scan(
			&i.ID,
			&i.DashboardAdmin,
			&i.DashboardName,
			&i.Logo,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Status,
			&i.Plan,
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

const updateDashboardName = `-- name: UpdateDashboardName :one
UPDATE dashboard
SET dashboard_name = $3
WHERE id = $1 AND dashboard_admin = $2
RETURNING id, dashboard_admin, dashboard_name, logo, created_at, updated_at, status, plan
`

type UpdateDashboardNameParams struct {
	ID             string `json:"id"`
	DashboardAdmin int32  `json:"dashboard_admin"`
	DashboardName  string `json:"dashboard_name"`
}

func (q *Queries) UpdateDashboardName(ctx context.Context, arg UpdateDashboardNameParams) (Dashboard, error) {
	row := q.db.QueryRowContext(ctx, updateDashboardName, arg.ID, arg.DashboardAdmin, arg.DashboardName)
	var i Dashboard
	err := row.Scan(
		&i.ID,
		&i.DashboardAdmin,
		&i.DashboardName,
		&i.Logo,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Status,
		&i.Plan,
	)
	return i, err
}