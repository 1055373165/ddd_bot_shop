package postgres

import (
	"context"
	"database/sql"
	"eda-in-go/customers/internal/domain"
	"fmt"

	"github.com/stackus/errors"
)

type CustomerRepository struct {
	tableName string
	db        *sql.DB
}

func NewCustomerRepository(tableName string, db *sql.DB) CustomerRepository {
	return CustomerRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r CustomerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	const query = "INSERT INTO %s (id, name, sms_number, enabled) VALUES ($1, $2, $3, $4)"

	_, err := r.db.ExecContext(ctx, r.table(query), customer.ID, customer.Name, customer.SmsNumber, customer.Enabled)

	return err
}

func (r CustomerRepository) Find(ctx context.Context, customerID string) (*domain.Customer, error) {
	const query = "SELECT name, sms_number, enabled FROM %s WHERE id = $1 LIMIT 1"

	customer := &domain.Customer{
		ID: customerID,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), customerID).Scan(&customer.Name, &customer.SmsNumber, &customer.Enabled)

	return customer, err
}

func (r CustomerRepository) FindAll(ctx context.Context) ([]*domain.Customer, error) {
	const query = "SELECT id, name, sms_number, enabled FROM %s"

	rows, err := r.db.QueryContext(ctx, r.table(query))
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			err = errors.Wrap(err, "closing store rows")
		}
	}(rows)

	customers := []*domain.Customer{}
	for rows.Next() {
		customer := &domain.Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.SmsNumber, &customer.Enabled)
		if err != nil {
			return nil, errors.ErrInternalServerError.Err(err)
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (r CustomerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	const query = "UPDATE %s SET name = $2, sms_number = $3, enabled = $4 WHERE id = $1"

	_, err := r.db.ExecContext(ctx, r.table(query), customer.ID, customer.Name, customer.SmsNumber, customer.Enabled)

	return err
}

func (r CustomerRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
