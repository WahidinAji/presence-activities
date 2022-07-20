package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

func (d *UserDeps) LoginRepo(ctx context.Context, in LoginIn) (*LoginOut, error) {
	if err := d.DB.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Connection error: %w", err)
	}

	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	//check user
	var authenticate bool
	query := "select exists(select name from users where email=$1)"
	tx.QueryRow(ctx, query, &in.Email).Scan(&authenticate)
	if !authenticate {
		return nil, fmt.Errorf("user not found")
	}

	//check email and password must be matches
	var password string
	query = "select password from users where email=$1"
	tx.QueryRow(ctx, query, &in.Email).Scan(&password)
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(in.Password))
	if err != nil {
		return nil, fmt.Errorf("Password doesn't matches: %v", err)
	}

	var out LoginOut
	query = "select email from users where email=$1 and password=$2"
	row := tx.QueryRow(ctx, query, &in.Email, password)
	if err := row.Scan(&out.Email); err != nil {
		return nil, fmt.Errorf("Failed to scan email: %v", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return &out, nil
}

func (d *UserDeps) RegisterRepo(ctx context.Context, in RegisIn) (*RegisOut, error) {
	if err := d.DB.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Connection error: %w", err)
	}

	tx, err := d.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	//check email
	var exists bool
	query := "select exists(select email from users where email=$1)"
	row := tx.QueryRow(ctx, query, &in.Email)
	if err := row.Scan(&exists); err != nil {
		return nil, fmt.Errorf("Scan error: %v", err)
	}
	if exists {
		return nil, fmt.Errorf("User already exists")
	}

	query = "INSERT INTO users (name, email, password) values($1,$2,$3)"
	_, err = tx.Exec(ctx, query, &in.Name, &in.Email, &in.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to execute insert query: %v", err)
	}

	var out RegisOut
	out.Email = in.Email

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return &out, nil
}
