package main

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(name, email string) (*User, error) {
	result, err := r.db.Exec(
		"INSERT INTO users(name, email) VALUES (?, ?)",
		name, email,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    int(id),
		Name:  name,
		Email: email,
	}, nil
}

func (r *UserRepository) FindByID(id int) (*User, error) {
	row := r.db.QueryRow(
		"SELECT id, name, email FROM users WHERE id = ?",
		id,
	)

	var u User
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
