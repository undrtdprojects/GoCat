package role

import (
	"GoCat/helpers/constant"
	"database/sql"
)

type Repository interface {
	CreateRoleRepository(role Role) (err error)
	GetAllRoleRepository() (result []Role, err error)
	GetRoleByIdRepository(id int) (role Role, err error)
	DeleteRoleRepository(role Role) (err error)
	UpdateRoleRepository(role Role) (err error)
}

type roleRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &roleRepository{
		db: database,
	}
}

func (r *roleRepository) CreateRoleRepository(role Role) (err error) {
	sqlStmt := "INSERT INTO " + constant.RoleTableName.String() + "\n" +
		"(name, created_at, created_by, modified_at, modified_by)" + "\n" +
		"VALUES ($1, $2, $3, $4, $5)"

	params := []interface{}{
		role.Name,
		role.CreatedAt,
		role.CreatedBy,
		role.ModifiedAt,
		role.ModifiedBy,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *roleRepository) GetAllRoleRepository() (roles []Role, err error) {
	sqlStmt := "SELECT id, name, created_at, created_by, modified_at, modified_by \n" +
		"FROM " + constant.RoleTableName.String()

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var role Role
		if err = rows.Scan(&role.Id, &role.Name, &role.CreatedAt, &role.CreatedBy,
			&role.ModifiedAt, &role.ModifiedBy); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (r *roleRepository) GetRoleByIdRepository(id int) (role Role, err error) {
	sqlStmt := "SELECT id, name, created_at, created_by, modified_at, modified_by \n" +
		"FROM " + constant.RoleTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		id,
	}

	rows, err := r.db.Query(sqlStmt, params...)
	if err != nil {
		return role, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&role.Id, &role.Name, &role.CreatedAt, &role.CreatedBy,
			&role.ModifiedAt, &role.ModifiedBy); err != nil {
			return role, err
		}
	}
	return role, nil
}

func (r *roleRepository) DeleteRoleRepository(role Role) (err error) {
	sqlStmt := "DELETE FROM " + constant.RoleTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		role.Id,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *roleRepository) UpdateRoleRepository(role Role) (err error) {
	sqlStmt := "UPDATE " + constant.RoleTableName.String() + "\n" +
		"SET name = $1, modified_at = $2, modified_by = $3 \n" +
		"WHERE id = $4"

	params := []interface{}{
		role.Name,
		role.ModifiedAt,
		role.ModifiedBy,
		role.Id,
	}
	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}
