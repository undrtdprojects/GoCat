package categories

import (
	"GoCat/helpers/constant"
	"database/sql"
)

type Repository interface {
	CreateCategoriesRepository(category Categories) (err error)
	GetAllCategoriesRepository() (result []Categories, err error)
	GetCategoriesByIdRepository(id string) (category Categories, err error)
	DeleteCategoriesRepository(category Categories) (err error)
	UpdateCategoriesRepository(category Categories) (err error)
}

type categoriesRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &categoriesRepository{
		db: database,
	}
}

func (r *categoriesRepository) CreateCategoriesRepository(category Categories) (err error) {
	sqlStmt := "INSERT INTO " + constant.CategoriesTableName.String() + "\n" +
		" (id, name, created_at, created_by, created_on, modified_at, modified_by, modified_on)" + "\n" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	params := []interface{}{
		category.Id,
		category.Name,
		category.CreatedAt,
		category.CreatedBy,
		category.CreatedOn,
		category.ModifiedAt,
		category.ModifiedBy,
		category.ModifiedOn,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *categoriesRepository) GetAllCategoriesRepository() (categories []Categories, err error) {
	sqlStmt := "SELECT id, name, created_at, created_by, created_on, modified_at, modified_by, modified_on " + "\n" +
		"FROM " + constant.CategoriesTableName.String()

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category Categories
		if err = rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.CreatedOn,
			&category.ModifiedAt, &category.ModifiedBy, &category.ModifiedOn); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *categoriesRepository) GetCategoriesByIdRepository(id string) (category Categories, err error) {
	sqlStmt := "SELECT id, name, created_at, created_by, created_on, modified_at, modified_by, modified_on" + "\n" +
		"FROM " + constant.CategoriesTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		id,
	}

	rows, err := r.db.Query(sqlStmt, params...)
	if err != nil {
		return category, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.CreatedOn,
			&category.ModifiedAt, &category.ModifiedBy, &category.ModifiedOn); err != nil {
			return category, err
		}
	}
	return category, nil
}

func (r *categoriesRepository) DeleteCategoriesRepository(category Categories) (err error) {
	sqlStmt := "DELETE FROM " + constant.CategoriesTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		category.Id,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *categoriesRepository) UpdateCategoriesRepository(category Categories) (err error) {
	sqlStmt := "UPDATE " + constant.CategoriesTableName.String() + "\n" +
		"SET name = $1, modified_at = $2, modified_by = $3 " + "\n" +
		"WHERE id = $4"

	params := []interface{}{
		category.Name,
		category.ModifiedAt,
		category.ModifiedBy,
		category.Id,
	}
	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}
