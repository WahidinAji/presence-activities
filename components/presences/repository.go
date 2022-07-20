package presences

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
)

func (d *PresenceDeps) FindAll(ctx context.Context, in ListIn) ([]PresenceList, error) {
	if err := d.DB.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Connection error: %w", err)
	}
	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	query := "select p.id, u.name, p.status, p.created_at from presences as p left join users as u on u.id = p.user_id where p.user_id = $1"
	rows, err := tx.Query(ctx, query, &in.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to query find all: %v", err)
	}

	var records []PresenceList
	for rows.Next() {
		var record PresenceList
		if err := rows.Scan(&record.ID, &record.Name, &record.Status, &record.PresenceCheck); err != nil {
			return nil, fmt.Errorf("scanning presence list: %v", err)
		}
		records = append(records, record)
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}
	return records, nil
}

func (d *PresenceDeps) PresenceRepo(ctx context.Context, in PresenceIn) (*PresenceOut, error) {
	if err := d.DB.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Connection error: %w", err)
	}

	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	status := strings.Replace(in.Status, "_", " ", 1)
	var exists bool
	query := "select exists(select p.user_id, p.status, u.id from presences as p left join users as u on u.id = p.user_id where DATE(p.created_at) = current_date AND u.id = $1 AND p.status = $2)"
	row := tx.QueryRow(ctx, query, &in.UserId, &in.Status)

	if err := row.Scan(&exists); err != nil {
		return nil, fmt.Errorf("Error scanning check existing row: %v", err)
	}
	if exists {
		return nil, fmt.Errorf(fmt.Sprint("You already ", status))
	}

	query = "insert into presences (user_id, status) values($1,$2)"
	_, err = tx.Exec(ctx, query, &in.UserId, &in.Status)
	if err != nil {
		return nil, fmt.Errorf("error inserting presences: %v", err)
	}

	var out PresenceOut
	query = "select p.status, p.created_at, u.name from presences as p left join users as u on u.id = p.user_id where DATE(p.created_at) = current_date AND u.id = $1 AND p.status = $2"
	row = tx.QueryRow(ctx, query, &in.UserId, &in.Status)
	if err := row.Scan(&out.Status, &out.PresenceCheck, &out.Name); err != nil {
		return nil, fmt.Errorf("error scanning presence out row: %v", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error committing: %v", err)
	}
	return &out, nil
}
