package activities

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func (d *ActivityDeps) CreateRepo(ctx context.Context, in ActivityIn) (*ActivityOut, error) {
	if err := d.DB.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Connection error: %w", err)
	}

	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	query := "insert into activities (user_id, title, description, status) values($1,$2,$3,$4)"
	_, err = tx.Exec(ctx, query, &in.UserId, &in.Title, &in.Description, &in.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to execute insert query: %v", err)
	}

	var out ActivityOut
	out.Title = in.Title
	out.Description = in.Description
	out.Status = in.Status

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return &out, nil
}

func (d *ActivityDeps) FindAll(ctx context.Context, in ListIn) ([]ActivitiesList, error) {
	if err := d.DB.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Connection error: %w", err)
	}

	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	query := "select a.id, a.title, a.description, a.status, u.name from activities as a left join users as u on a.user_id = u.id where a.user_id =$1"
	rows, err := tx.Query(ctx, query, &in.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to query find all: %v", err)
	}

	var records []ActivitiesList
	for rows.Next() {
		var record ActivitiesList
		if err := rows.Scan(&record.ID, &record.Title, &record.Description, &record.Status, &record.Name); err != nil {
			return nil, fmt.Errorf("scanning presence list: %v", err)
		}
		records = append(records, record)
	}
	
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}
	return records, nil

}
