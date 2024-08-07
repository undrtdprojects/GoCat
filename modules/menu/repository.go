package menu

import (
	"GoCat/helpers/constant"
	"database/sql"
	"quiz-3-sanbercode-greg/helpers/constant"
)

type Repository interface {
	CreateMenuRepository(menu Menu) (err error)
	GetAllMenuRepository() (result []Menu, err error)
	GetMenuByIdRepository(id int) (menu Menu, err error)
	DeleteMenuRepository(menu Menu) (err error)
	UpdateMenuRepository(menu Menu) (err error)
}

type menuRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &menuRepository{
		db: database,
	}
}

func (r *menuRepository) CreateMenuRepository(menu Menu) (err error) {
	sqlStmt := "INSERT INTO " + constant.MenuTableName.String() + "\n" +
		"(id, name, price, category_id, created_at, created_by, modified_at, modified_by)" + "\n" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	params := []interface{}{
		menu.Id,
		menu.Name,
		menu.Price,
		menu.CategoryId,
		menu.CreatedAt,
		menu.CreatedBy,
		menu.ModifiedAt,
		menu.ModifiedBy,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *menuRepository) GetAllMenuRepository() (menus []Menu, err error) {
	sqlStmt := "SELECT id, name, price, category_id, created_at, created_by, modified_at, modified_by \n" +
		"FROM " + constant.MenuTableName.String()

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var menu Menu
		if err = rows.Scan(&menu.Id, &menu.Name, &menu.Price, &menu.CategoryId,
			&menu.CreatedAt, &menu.CreatedBy, &menu.ModifiedAt, &menu.ModifiedBy); err != nil {
			return nil, err
		}
		menus = append(menus, menu)
	}

	return menus, nil
}

func (r *menuRepository) GetMenuByIdRepository(id int) (menu Menu, err error) {
	sqlStmt := "SELECT id, name, price, category_id, created_at, created_by, modified_at, modified_by \n" +
		"FROM " + constant.MenuTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		id,
	}

	rows, err := r.db.Query(sqlStmt, params...)
	if err != nil {
		return menu, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&menu.Id, &menu.Name, &menu.Price, &menu.CategoryId,
			&menu.CreatedAt, &menu.CreatedBy, &menu.ModifiedAt, &menu.ModifiedBy); err != nil {
			return nil, err
		}
	}
	return menu, nil
}

func (r *menuRepository) DeleteMenuRepository(menu Menu) (err error) {
	sqlStmt := "DELETE FROM " + constant.MenuTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		menu.Id,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *menuRepository) UpdateMenuRepository(menu Menu) (err error) {
	sqlStmt := "UPDATE " + constant.MenuTableName.String() + "\n" +
		"SET name = $1, price = $2, category_id = $3,  \n" +
		"modified_at = $4, modified_by = $5 " + "\n" +
		"WHERE id = $6"

	params := []interface{}{
		menu.Name,
		menu.Price,
		menu.CategoryId,
		menu.ModifiedAt,
		menu.ModifiedBy,
		menu.Id,
	}
	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}
