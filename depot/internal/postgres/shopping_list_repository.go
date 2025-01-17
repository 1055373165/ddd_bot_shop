package postgres

import (
	"context"
	"database/sql"
	"eda-in-go/depot/internal/domain"
	"encoding/json"
	"fmt"

	"github.com/stackus/errors"
)

type ShoppingListRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.ShoppingListRepository = (*ShoppingListRepository)(nil)

func NewShoppingListRepository(tableName string, db *sql.DB) ShoppingListRepository {
	return ShoppingListRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r ShoppingListRepository) Find(ctx context.Context, id string) (*domain.ShoppingList, error) {
	const query = "SELECT order_id, stops, assigned_bot_id, status FROM %s WHERE id = $1 LIMIT 1"

	shoppingList := &domain.ShoppingList{
		ID: id,
	}
	var stops []byte
	var status string

	err := r.db.QueryRowContext(ctx, r.table(query), id).Scan(&shoppingList.OrderID, &stops, &shoppingList.AssignedBotID, &status)
	if err != nil {
		return nil, err
	}

	shoppingList.Status = domain.ToShoppingListStatus(status)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	err = json.Unmarshal(stops, &shoppingList.Stops)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	return shoppingList, nil
}

func (r ShoppingListRepository) FindAll(ctx context.Context) ([]*domain.ShoppingList, error) {
	const query = "SELECT id, order_id, stops, assigned_bot_id, status FROM %s"
	rows, err := r.db.QueryContext(ctx, r.table(query))
	if err != nil {
		return nil, err
	}

	shoppingLists := []*domain.ShoppingList{}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			err = errors.Wrap(err, "close rows")
		}
	}(rows)

	for rows.Next() {
		shoppingList := &domain.ShoppingList{}
		var stops []byte
		var status string
		err := rows.Scan(&shoppingList.ID, &shoppingList.OrderID, &stops, &shoppingList.AssignedBotID, &status)
		if err != nil {
			return nil, errors.Wrap(err, "scan rows")
		}
		shoppingList.Status = domain.ToShoppingListStatus(status)
		err = json.Unmarshal(stops, &shoppingList.Stops)
		if err != nil {
			return nil, errors.ErrInternalServerError.Err(err)
		}
		shoppingLists = append(shoppingLists, shoppingList)
	}

	return shoppingLists, nil
}

func (r ShoppingListRepository) FindByOrderID(ctx context.Context, orderID string) (*domain.ShoppingList, error) {
	const query = "SELECT id, stops, assigned_bot_id, status FROM %s WHERE order_id = $1 LIMIT 1"

	shoppingList := &domain.ShoppingList{
		OrderID: orderID,
	}
	var stops []byte
	var status string

	err := r.db.QueryRowContext(ctx, r.table(query), orderID).Scan(&shoppingList.ID, &stops, &shoppingList.AssignedBotID, &status)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	shoppingList.Status = domain.ToShoppingListStatus(status)

	err = json.Unmarshal(stops, &shoppingList.Stops)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	return shoppingList, nil
}

func (r ShoppingListRepository) Save(ctx context.Context, list *domain.ShoppingList) error {
	const query = "INSERT INTO %s (id, order_id, stops, assigned_bot_id, status) VALUES ($1, $2, $3, $4, $5)"

	stops, err := json.Marshal(list.Stops)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query), list.ID, list.OrderID, stops, list.AssignedBotID, list.Status.String())

	return errors.ErrInternalServerError.Err(err)
}

func (r ShoppingListRepository) Update(ctx context.Context, list *domain.ShoppingList) error {
	const query = "UPDATE %s SET stops = $2, assigned_bot_id = $3, status = $4 WHERE id = $1"

	stops, err := json.Marshal(list.Stops)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query), list.ID, stops, list.AssignedBotID, list.Status.String())

	return errors.ErrInternalServerError.Err(err)
}

func (r ShoppingListRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
