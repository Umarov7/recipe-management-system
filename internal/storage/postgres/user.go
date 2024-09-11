package postgres

import (
	"context"
	"database/sql"
	"recipe-management/internal/models"
	"recipe-management/internal/storage"

	"github.com/pkg/errors"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) storage.IUserStorage {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *models.NewUser) (string, error) {
	query := `
	insert into users
		(username, email, password)
	values
		($1, $2, $3)
	returning id
	`

	var id string
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		return "", errors.Wrap(err, "failed to add user")
	}

	return id, nil
}

func (r *UserRepo) Read(ctx context.Context, id string) (*models.User, error) {
	query := `
	select
		username, email, role, created_at, updated_at
	from
		users
	where
		deleted_at is null and id = $1
	`

	user := models.User{ID: id}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(err, "user not found")
		}
		return nil, errors.Wrap(err, "failed to read user")
	}

	return &user, nil
}

func (r *UserRepo) Update(ctx context.Context, user *models.NewUserData) error {
	query := `
	update
		users
	set
		username = $1, email = $2, updated_at = now()
	where
		deleted_at is null and id = $3
	`

	res, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.ID)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *UserRepo) Delete(ctx context.Context, id string) error {
	query := `
	update
		users
	set
		deleted_at = now()
	where
		deleted_at is null and id = $1
	`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
