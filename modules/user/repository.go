package user

import (
	"database/sql"
	"quiz-3-sanbercode-greg/helpers/constant"
)

type Repository interface {
	Login(user LoginRequest) (result User, err error)
	SignUp(user User) (err error)
	Update(user User) (err error)
	Delete(user User) (err error)
	GetList() (user []User, err error)
	GetUserByUsername(username string) (user User, err error)
}
type userRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &userRepository{
		db: database,
	}
}

func (r *userRepository) Login(user LoginRequest) (result User, err error) {
	sqlStmt := "SELECT id, password FROM " + constant.UsersTableName.String() + " WHERE username = $1"

	params := []interface{}{
		user.Username,
	}

	err = r.db.QueryRow(sqlStmt, params...).Scan(&result.Id, &result.Password)
	if err != nil && err != sql.ErrNoRows {
		return result, err
	}

	return result, nil
}

// query untuk insert data user ke database (sign up)
func (r *userRepository) SignUp(user User) (err error) {
	sqlStmt := "INSERT INTO " + constant.UsersTableName.String() +
		" (username, password, created_at, created_by, modified_at, modified_by) VALUES ($1, $2, $3, $4, $5, $6)"

	params := []interface{}{
		user.Username,
		user.Password,
		user.CreatedAt,
		user.CreatedBy,
		user.ModifiedAt,
		user.ModifiedBy,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *userRepository) Update(user User) (err error) {
	sqlStmt := "UPDATE " + constant.UsersTableName.String() + " SET password = $1 WHERE username = $2"

	params := []interface{}{
		user.Password,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *userRepository) Delete(user User) (err error) {
	sqlStmt := "DELETE FROM " + constant.UsersTableName.String() + " WHERE username = $1"

	params := []interface{}{
		user.Username,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *userRepository) GetList() (users []User, err error) {
	sqlStmt := "SELECT username, password FROM " + constant.UsersTableName.String()

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.Username, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUserByUsername(username string) (user User, err error) {
	sqlStmt := "SELECT username, password FROM " + constant.UsersTableName.String() + " WHERE username = $1"

	params := []interface{}{
		username,
	}
	rows, err := r.db.Query(sqlStmt, params...)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&user.Username, &user.Password); err != nil {
			return user, err
		}
	}
	return user, nil
}
