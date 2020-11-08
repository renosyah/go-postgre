package model

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type (
	User struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Email string    `json:"email"`
	}

	UserResponse struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Email string    `json:"email"`
	}

	AllUser struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (u *User) Response() UserResponse {
	return UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func (u *User) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "user" (name,email) VALUES ($1,$2) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.Name, u.Email).Scan(&u.ID)
	if err != nil {
		return u.ID, err
	}

	return u.ID, nil
}

func (u *User) All(ctx context.Context, db *sql.DB, param AllUser) ([]User, error) {
	all := []User{}

	query := `SELECT id,name,email FROM "user" WHERE %s LIKE $1 ORDER BY %s %s OFFSET $2 LIMIT $3`
	rows, err := db.QueryContext(ctx, fmt.Sprintf(query, param.SearchBy, param.OrderBy, param.OrderDir), "%"+param.SearchValue+"%", param.Offset, param.Limit)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := User{}
		err = rows.Scan(
			&one.ID, &one.Name, &one.Email,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (u *User) One(ctx context.Context, db *sql.DB) (User, error) {
	one := User{}

	query := `SELECT id,name,email FROM "user" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.ID).Scan(
		&one.ID, &one.Name, &one.Email,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (u *User) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "user" SET name = $1, email = $2 WHERE id = $3 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.Name, u.Email, u.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (u *User) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "user" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), u.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
