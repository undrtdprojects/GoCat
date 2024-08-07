package transaction1

import (
	"GoCat/helpers/constant"
	"database/sql"
)

type Repository interface {
	CreateTransaction1Repository(transaction1 Transaction1) (err error)
	GetAllTransaction1Repository() (transaction1s []Transaction1, err error)
	GetTransaction1ByIdRepository(id int) (transaction1 Transaction1, err error)
	DeleteTransaction1Repository(transaction1 Transaction1) (err error)
	UpdateTransaction1Repository(transaction1 Transaction1) (err error)
}

type transaction1Repository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &transaction1Repository{
		db: database,
	}
}

func (r *transaction1Repository) CreateTransaction1Repository(transaction1 Transaction1) (err error) {
	sqlStmt := "INSERT INTO " + constant.Transaction1TableName.String() + "\n" +
		"(transaction_id, menu_id, date_transaction, qty, total_price, payment_id, created_at, created_by, modified_at, modified_by)" + "\n" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	params := []interface{}{
		transaction1.TransactionId,
		transaction1.MenuId,
		transaction1.DateTransaction,
		transaction1.Qty,
		transaction1.TotalPrice,
		transaction1.PaymentId,
		transaction1.CreatedAt,
		transaction1.CreatedBy,
		transaction1.ModifiedAt,
		transaction1.ModifiedBy,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *transaction1Repository) GetAllTransaction1Repository() (transaction1s []Transaction1, err error) {
	sqlStmt := "SELECT id, transaction_id, menu_id, date_transaction, qty, total_price, payment_id, \n" +
		"created_at, created_by, modified_at, modified_by \n" +
		"FROM " + constant.Transaction1TableName.String()

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction1 Transaction1
		if err = rows.Scan(&transaction1.Id, &transaction1.TransactionId, &transaction1.MenuId, &transaction1.DateTransaction,
			&transaction1.Qty, &transaction1.TotalPrice, &transaction1.PaymentId, &transaction1.CreatedAt, &transaction1.CreatedBy,
			&transaction1.ModifiedAt, &transaction1.ModifiedBy); err != nil {
			return nil, err
		}
		transaction1s = append(transaction1s, transaction1)
	}

	return transaction1s, nil
}

func (r *transaction1Repository) GetTransaction1ByIdRepository(id int) (transaction1 Transaction1, err error) {
	sqlStmt := "SELECT id, transaction_id, menu_id, date_transaction, qty, total_price, payment_id, \n" +
		"created_at, created_by, modified_at, modified_by \n" +
		"FROM " + constant.Transaction1TableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		id,
	}

	rows, err := r.db.Query(sqlStmt, params...)
	if err != nil {
		return transaction1, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&transaction1.Id, &transaction1.TransactionId, &transaction1.MenuId, &transaction1.DateTransaction,
			&transaction1.Qty, &transaction1.TotalPrice, &transaction1.PaymentId, &transaction1.CreatedAt, &transaction1.CreatedBy,
			&transaction1.ModifiedAt, &transaction1.ModifiedBy); err != nil {
			return transaction1, err
		}
	}
	return transaction1, nil
}

func (r *transaction1Repository) DeleteTransaction1Repository(transaction1 Transaction1) (err error) {
	sqlStmt := "DELETE FROM " + constant.Transaction1TableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		transaction1.Id,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *transaction1Repository) UpdateTransaction1Repository(transaction1 Transaction1) (err error) {
	sqlStmt := "UPDATE " + constant.Transaction1TableName.String() + "\n" +
		"SET transaction_id = $1, menu_id = $2, date_transaction = $3, qty = $4, total_price = $5, payment_id = $6, \n" +
		"modified_at = $7, modified_by = $8 \n" +
		"WHERE id = $9"

	params := []interface{}{
		transaction1.TransactionId,
		transaction1.MenuId,
		transaction1.DateTransaction,
		transaction1.Qty,
		transaction1.TotalPrice,
		transaction1.PaymentId,
		transaction1.ModifiedAt,
		transaction1.ModifiedBy,
		transaction1.Id,
	}
	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}
