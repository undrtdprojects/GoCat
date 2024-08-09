package user

import (
	"GoCat/helpers/constant"
	"database/sql"
	"fmt"
)

type Repository interface {
	Login(user LoginRequest) (result User, err error)
	SignUp(user User) (err error)
	Update(user User) (err error)
	ChangePassword(user User) (err error)
	Delete(user User) (err error)
	GetList() (user []User, err error)
	GetUserByUsername(username string) (user User, err error)
	GetUserById(id int) (user User, err error)
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
	sqlStmt := "SELECT id, username, password, role_id FROM " + constant.UsersTableName.String() + " WHERE username = $1"

	params := []interface{}{
		user.Username,
	}

	err = r.db.QueryRow(sqlStmt, params...).Scan(&result.Id, &result.Username, &result.Password, &result.RoleId)
	if err != nil && err != sql.ErrNoRows {
		return result, err
	}

	return result, nil
}

func (r *userRepository) SignUp(user User) (err error) {
	sqlStmt := "INSERT INTO " + constant.UsersTableName.String() +
		" (username, password, role_id, created_at, created_by, created_on, modified_at, modified_by, modified_on) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	params := []interface{}{
		user.Username,
		user.Password,
		user.RoleId,
		user.CreatedAt,
		user.CreatedBy,
		user.CreatedOn,
		user.ModifiedAt,
		user.ModifiedBy,
		user.ModifiedOn,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *userRepository) Update(user User) (err error) {
	sqlStmt := "UPDATE " + constant.UsersTableName.String() +
		"SET username = $1, password = $2, role_id = $3, modified_at = $4, modified_by = $5, modified_on = $6 WHERE id = $7"

	params := []interface{}{
		user.Username,
		user.Password,
		user.RoleId,
		user.ModifiedAt,
		user.ModifiedBy,
		user.ModifiedOn,
		user.Id,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *userRepository) ChangePassword(user User) (err error) {
	sqlStmt := "UPDATE " + constant.UsersTableName.String() +
		"SET password = $1, modified_at = $2, modified_by = $3, modified_on = $4 WHERE id = $5"

	params := []interface{}{
		user.Password,
		user.ModifiedAt,
		user.ModifiedBy,
		user.ModifiedOn,
		user.Id,
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
	sqlStmt := "SELECT id, username, password, role_id, created_at, created_by, created_on, modified_at, modified_by, modified_on FROM " + constant.UsersTableName.String()

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.RoleId, &user.CreatedAt, &user.CreatedBy, &user.CreatedOn,
			&user.ModifiedAt, &user.ModifiedBy, &user.ModifiedOn); err != nil {
			return nil, err
		}
		fmt.Println("rows:", user)

		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUserByUsername(username string) (user User, err error) {
	sqlStmt := "SELECT id, username, password, role_id FROM " + constant.UsersTableName.String() + " WHERE username = $1"

	params := []interface{}{
		username,
	}
	rows, err := r.db.Query(sqlStmt, params...)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.RoleId); err != nil {
			return user, err
		}
	}
	return user, nil
}

func (r *userRepository) GetUserById(id int) (user User, err error) {
	sqlStmt := "SELECT username, password, role_id FROM " + constant.UsersTableName.String() + " WHERE id = $1"

	params := []interface{}{
		id,
	}
	rows, err := r.db.Query(sqlStmt, params...)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&user.Username, &user.Password, &user.RoleId); err != nil {
			return user, err
		}
	}
	return user, nil
}
