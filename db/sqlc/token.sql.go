// Code generated by sqlc. DO NOT EDIT.
// source: token.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const blacklistToken = `-- name: BlacklistToken :exec
INSERT INTO blacklisted_access_tokens (
  token_id
) VALUES (
  $1
)
`

func (q *Queries) BlacklistToken(ctx context.Context, tokenID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, blacklistToken, tokenID)
	return err
}

const listBlacklistedAcessTokens = `-- name: ListBlacklistedAcessTokens :many
SELECT token_id FROM blacklisted_access_tokens
`

func (q *Queries) ListBlacklistedAcessTokens(ctx context.Context) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, listBlacklistedAcessTokens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []uuid.UUID
	for rows.Next() {
		var token_id uuid.UUID
		if err := rows.Scan(&token_id); err != nil {
			return nil, err
		}
		items = append(items, token_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}