package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"eda-in-go/payments/internal/application"
	"eda-in-go/payments/internal/models"

	"github.com/stackus/errors"
)

type PaymentRepository struct {
	tableName string
	db        *sql.DB
}

var _ application.PaymentRepository = (*PaymentRepository)(nil)

func NewPaymentRepository(tableName string, db *sql.DB) PaymentRepository {
	return PaymentRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r PaymentRepository) Save(ctx context.Context, payment *models.Payment) error {
	const query = "INSERT INTO %s (id, customer_id, amount) VALUES ($1, $2, $3)"

	_, err := r.db.ExecContext(ctx, r.table(query), payment.ID, payment.CustomerID, payment.Amount)

	return err
}

func (r PaymentRepository) Find(ctx context.Context, paymentID string) (*models.Payment, error) {
	const query = "SELECT customer_id, amount FROM %s WHERE id = $1 LIMIT 1"

	payment := &models.Payment{
		ID: paymentID,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), paymentID).Scan(&payment.CustomerID, &payment.Amount)

	return payment, err
}

func (r PaymentRepository) FindAll(ctx context.Context) ([]*models.Payment, error) {
	const query = "SELECT id, customer_id, amount FROM %s"

	rows, err := r.db.QueryContext(ctx, r.table(query))
	if err != nil {
		return nil, errors.Wrap(err, "query payment")
	}
	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			err = errors.Wrap(err, "closing payment")
		}
	}(rows)

	payments := []*models.Payment{}
	for rows.Next() {
		payment := &models.Payment{}
		err = rows.Scan(&payment.ID, &payment.CustomerID, &payment.Amount)
		if err != nil {
			return nil, errors.Wrap(err, "scanning payment")
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func (r PaymentRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
