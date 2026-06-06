package repositories

import (
	"bookmark-api/internal/db"
	"bookmark-api/internal/models"
	"time"
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

func (r *UserRepository) CreatePasswordResetToken(
	userID int64,
	token string,
	expiresAt time.Time,
) error {
	query := `
	INSERT INTO password_reset_tokens(user_id,token,expires_at)
	VALUES($1,$2,$3)
	`

	_, err := db.Conn.Exec(
		query,
		userID,
		token,
		expiresAt,
	)

	return err
}

func (r *UserRepository) FindUserByResetToken(
	token string,
) (*models.User, error) {
	query := `
	SELECT u.id,u.name,u.email,u.password,u.created_at
	FROM users u
	JOIN password_reset_tokens prt ON u.id=prt.user_id
	WHERE prt.token=$1 AND prt.expires_at > NOW()
	`

	var user models.User

	err := db.Conn.QueryRow(query, token).Scan(
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

func (r *UserRepository) UpdatePassword(
	userID int64,
	hashedPassword string,
) error {
	query := `
	UPDATE users
	SET password=$1
	WHERE id=$2
	`

	_, err := db.Conn.Exec(query, hashedPassword, userID)
	return err
}

func (r *UserRepository) InvalidateResetToken(
	token string,
) error {
	query := `
	DELETE FROM password_reset_tokens
	WHERE token=$1
	`

	_, err := db.Conn.Exec(query, token)
	return err
}
