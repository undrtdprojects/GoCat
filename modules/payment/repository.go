package payment

import (
	"GoCat/helpers/constant"
	"database/sql"
)

type Repository interface {
	CreatePaymentRepository(payment Payment) (err error)
	GetAllPaymentRepository() (result []Payment, err error)
	GetPaymentByIdRepository(id int) (payment Payment, err error)
	DeletePaymentRepository(payment Payment) (err error)
	UpdatePaymentRepository(payment Payment) (err error)
}

type paymentRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &paymentRepository{
		db: database,
	}
}

func (r *paymentRepository) CreatePaymentRepository(payment Payment) (err error) {
	sqlStmt := "INSERT INTO " + constant.PaymentTableName.String() + "\n" +
		"(name, created_at, created_by, created_on, modified_at, modified_by, modified_on)" + "\n" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)"

	params := []interface{}{
		payment.Name,
		payment.CreatedAt,
		payment.CreatedBy,
		payment.CreatedOn,
		payment.ModifiedAt,
		payment.ModifiedBy,
		payment.ModifiedOn,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *paymentRepository) GetAllPaymentRepository() (payments []Payment, err error) {
	sqlStmt := "SELECT id, name, created_at, created_by, created_on, modified_at, modified_by, modified_on \n" +
		"FROM " + constant.PaymentTableName.String()

	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var payment Payment
		if err = rows.Scan(&payment.Id, &payment.Name, &payment.CreatedAt, &payment.CreatedBy, &payment.CreatedOn,
			&payment.ModifiedAt, &payment.ModifiedBy, &payment.ModifiedOn); err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func (r *paymentRepository) GetPaymentByIdRepository(id int) (payment Payment, err error) {
	sqlStmt := "SELECT id, name, created_at, created_by, created_on, modified_at, modified_by, modified_on \n" +
		"FROM " + constant.PaymentTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		id,
	}

	rows, err := r.db.Query(sqlStmt, params...)
	if err != nil {
		return payment, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&payment.Id, &payment.Name, &payment.CreatedAt, &payment.CreatedBy, &payment.CreatedOn,
			&payment.ModifiedAt, &payment.ModifiedBy, &payment.ModifiedOn); err != nil {
			return payment, err
		}
	}
	return payment, nil
}

func (r *paymentRepository) DeletePaymentRepository(payment Payment) (err error) {
	sqlStmt := "DELETE FROM " + constant.PaymentTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		payment.Id,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

func (r *paymentRepository) UpdatePaymentRepository(payment Payment) (err error) {
	sqlStmt := "UPDATE " + constant.PaymentTableName.String() + "\n" +
		"SET name = $1, modified_at = $2, modified_by = $3 " + "\n" +
		"WHERE id = $4"

	params := []interface{}{
		payment.Name,
		payment.ModifiedAt,
		payment.ModifiedBy,
		payment.Id,
	}
	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}
