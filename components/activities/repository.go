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

func (d *ActivityDeps) FindByDate(ctx context.Context, in QueryDateIn) ([]QueryDateOut, error) {
	if err := d.DB.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Connection error: %w", err)
	}

	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	query := "select a.title, a.description, a.status from activities where a.user_id = $1 AND DATE(a.created_at) => $2 and DATE(a.created_at) <= $3"
	rows, err := tx.Query(ctx, query, &in.UserId, &in.DateFrom, &in.DateTo)
	if err != nil {
		return nil, fmt.Errorf("error querying database : %v", err)
	}
	var records []QueryDateOut
	for rows.Next() {
		var record QueryDateOut
		if err := rows.Scan(&record.Title, &record.Description, &record.Status); err != nil {
			return nil, fmt.Errorf("error scanning database : %v", err)
		}
		records = append(records, record)
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}
	return records, nil
	// query = "select p.status, p.created_at, u.name from presences as p left join users as u on u.id = p.user_id where DATE(p.created_at) = current_date AND u.id = $1 AND p.status = $2"
}
