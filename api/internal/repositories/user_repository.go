package repositories

import (
	"bookmark-api/internal/db"
	"bookmark-api/internal/models"
)

type UserRepository struct{}

func (r *UserRepository) Create(
	name string,
	email string,
	hashedPassword string,
) error {
	query := `
	INSERT INTO users(name,email,password)
	VALUES($1,$2,$3)
	`

	_, err := db.Conn.Exec(
		query,
		name,
		email,
		hashedPassword,
	)

	return err
}

func (r *UserRepository) FindByEmail(
	email string,
) (*models.User, error) {
	query := `
	SELECT id,name,email,password,created_at
	FROM users
	WHERE email=$1
	`

	var user models.User

	err := db.Conn.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
