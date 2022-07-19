package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

func Migrate(ctx context.Context, db *pgx.Conn) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Failed to start transaction db: %v", err))
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
		create table if not exists users(
			id bigserial primary key,
			name varchar(255) not null,
			email varchar(255) not null,
			password text not null,
			created_at timestamp default current_timestamp,
			updated_at timestamp default current_timestamp
		);
	`)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Failed to migrate users table: %v", err))
	}
	log.Println("Successfully migrated users table")

	_, err = tx.Exec(ctx, `
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'presence_status') THEN
				CREATE TYPE presence_status AS ENUM ('check_in','check_out');
			END IF;
			--more types here...
			create table if not exists presences(
				id bigserial primary key,
				user_id bigserial not null,
				status presence_status default 'check_in',
				created_at timestamp default current_timestamp,
				updated_at timestamp default current_timestamp,
				CONSTRAINT fk_user
					FOREIGN key (user_id)
						REFERENCES users(id)
			);
		END$$;
	`)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Failed to migrate presences table: %v", err))
	}
	log.Println("Successfully migrated presences table")

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Failed to commit transaction db: %v", err))
	}

	return nil
}
