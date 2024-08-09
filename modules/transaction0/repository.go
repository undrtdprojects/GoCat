package transaction0

import (
	"GoCat/helpers/constant"
	"database/sql"
)

type Repository interface {
	CreateTransaction0Repository(transaction0 Transaction0) (err error)
	GetAllTransaction0Repository() (transaction0s []Transaction0, err error)
	GetTransaction0CountRepository() (count int, err error)
	GetTransaction0ByIdRepository(id string) (transaction0 Transaction0, err error)
	DeleteTransaction0Repository(transaction0 Transaction0) (err error)
	UpdateTransaction0Repository(transaction0 Transaction0) (err error)
}

type transaction0Repository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &transaction0Repository{
		db: database,
	}
}

func (r *transaction0Repository) CreateTransaction0Repository(transaction0 Transaction0) (err error) {
	sqlStmt := "INSERT INTO " + constant.Transaction0TableName.String() + "\n" +
		"(id, user_id, payment_id, grand_total_price, created_at, created_by, created_on, modified_at, modified_by, modified_on)" + "\n" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	params := []interface{}{
		transaction0.Id,
		transaction0.UserId,
		transaction0.PaymentId,
		transaction0.GrandTotalPrice,
		transaction0.CreatedAt,
		transaction0.CreatedBy,
		transaction0.CreatedOn,
		transaction0.ModifiedAt,
		transaction0.ModifiedBy,
		transaction0.ModifiedOn,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *transaction0Repository) GetAllTransaction0Repository() (transaction0s []Transaction0, err error) {
	sqlStmt := "SELECT id, user_id, grand_total_price, created_at, created_by, created_on, modified_at, modified_by, modified_on \n" +
		"FROM " + constant.Transaction0TableName.String()

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction0 Transaction0
		if err = rows.Scan(&transaction0.Id, &transaction0.UserId, &transaction0.GrandTotalPrice,
			&transaction0.CreatedAt, &transaction0.CreatedBy, &transaction0.CreatedOn,
			&transaction0.ModifiedAt, &transaction0.ModifiedBy, &transaction0.ModifiedOn); err != nil {
			return nil, err
		}
		transaction0s = append(transaction0s, transaction0)
	}

	return transaction0s, nil
}

func (r *transaction0Repository) GetTransaction0CountRepository() (count int, err error) {
	sqlStmt := "SELECT COUNT(*) FROM " + constant.Transaction0TableName.String()

	err = r.db.QueryRow(sqlStmt).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count + 1, nil
}

func (r *transaction0Repository) GetTransaction0ByIdRepository(id string) (transaction0 Transaction0, err error) {
	sqlStmt := "SELECT id, user_id, grand_total_price, created_at, created_by, created_on, modified_at, modified_by, modified_on \n" +
		"FROM " + constant.Transaction0TableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		id,
	}

	rows, err := r.db.Query(sqlStmt, params...)
	if err != nil {
		return transaction0, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&transaction0.Id, &transaction0.UserId, &transaction0.GrandTotalPrice,
			&transaction0.CreatedAt, &transaction0.CreatedBy, &transaction0.CreatedOn,
			&transaction0.ModifiedAt, &transaction0.ModifiedBy, &transaction0.ModifiedOn); err != nil {
			return transaction0, err
		}
	}
	return transaction0, nil
}

func (r *transaction0Repository) DeleteTransaction0Repository(transaction0 Transaction0) (err error) {
	sqlStmt := "DELETE FROM " + constant.Transaction0TableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		transaction0.Id,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *transaction0Repository) UpdateTransaction0Repository(transaction0 Transaction0) (err error) {
	sqlStmt := "UPDATE " + constant.Transaction0TableName.String() + "\n" +
		"SET user_id = $1, grand_total_price = $2, modified_at = $3, modified_by = $4 \n" +
		"WHERE id = $5"

	params := []interface{}{
		transaction0.UserId,
		transaction0.GrandTotalPrice,
		transaction0.ModifiedAt,
		transaction0.ModifiedBy,
		transaction0.Id,
	}
	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}
