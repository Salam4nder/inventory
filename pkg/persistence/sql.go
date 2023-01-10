package persistence

import (
	"context"
	"errors"
	"strconv"

	"github.com/Salam4nder/inventory/pkg/entity"
	"github.com/google/uuid"
)

const (
	insertItem = `INSERT INTO inventory (
        name, unit, amount, expires_at) 
        VALUES (
        $1, $2, $3, $4) RETURNING id`

	selectItem = `SELECT FROM inventory WHERE id = $1`

	selectAll = `SELECT * FROM inventory`

	updateItem = `UPDATE inventory SET 
    name = $1, unit = $2, amount = $3, expires_at = $4 WHERE id = $5`

	deleteItem = `DELETE FROM inventory WHERE id = $1`
)

// Context is an alias for context.Context.
type Context context.Context

// Create creates a new item in the database.
func (s *Storage) Create(
	ctx Context, item entity.Item) (uuid.UUID, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback()

	if err := tx.QueryRowContext(
		ctx, insertItem, item.Name, item.Unit,
		item.Amount, item.ExpiresAt).Scan(&item.ID); err != nil {
		return uuid.Nil, err
	}

	if err := tx.Commit(); err != nil {
		return uuid.Nil, err
	}

	return item.ID, nil
}

// Read reads an item from the database based off of an uuid.
func (s *Storage) Read(
	ctx Context, uuid string) (*entity.Item, error) {
	item := &entity.Item{}

	if err := s.DB.QueryRowContext(
		ctx, selectItem, uuid).Scan(
		&item.ID, &item.Name, &item.Unit,
		&item.Amount, &item.ExpiresAt); err != nil {
		return &entity.Item{}, err
	}

	return item, nil
}

// ReadBy reads items fro the database by the given filter.
// It returns an error if the filter is empty.
func (s *Storage) ReadBy(
	ctx Context, filter entity.ItemFilter) ([]*entity.Item, error) {
	items := []*entity.Item{}

	query, args := filterQueryBuilder(filter)
	if len(args) == 0 {
		return items, errors.New("empty filter")
	}

	rows, err := s.DB.QueryContext(
		ctx, query, args...)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		item := entity.Item{}
		if err := rows.Scan(
			&item.ID, &item.Name, &item.Unit,
			&item.Amount, &item.ExpiresAt); err != nil {
			return []*entity.Item{}, err
		}
		items = append(items, &item)
	}

	err = rows.Err()

	return items, err
}

// Update updates an item in the database.
func (s *Storage) Update(
	ctx Context, item *entity.Item) (*entity.Item, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return &entity.Item{}, err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(
		ctx, updateItem, item.Name, item.Unit,
		item.Amount, item.ExpiresAt, item.ID); err != nil {
		return &entity.Item{}, err
	}

	if err := tx.Commit(); err != nil {
		return &entity.Item{}, err
	}

	return item, nil
}

// Delete deletes an item from the database.
func (s *Storage) Delete(
	ctx Context, uuid string) error {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(
		ctx, deleteItem, uuid); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func filterQueryBuilder(filter entity.ItemFilter) (
	query string, args []interface{}) {
	query = "SELECT * FROM inventory WHERE "

	if filter.Name != "" {
		args = append(args, filter.Name)
		query += "name = $" + strconv.Itoa(len(args))
	}

	if filter.Unit != "" {
		if len(args) > 0 {
			query += " AND "
		}
		args = append(args, filter.Unit)
		query += "unit = $" + strconv.Itoa(len(args))
	}

	if filter.Amount != 0.0 {
		if len(args) > 0 {
			query += " AND "
		}
		args = append(args, filter.Amount)
		query += "amount = $" + strconv.Itoa(len(args))
	}

	if !filter.ExpiresAt.IsZero() {
		if len(args) > 0 {
			query += " AND "
		}
		args = append(args, filter.ExpiresAt)
		query += "expires_at = $" + strconv.Itoa(len(args))
	}

	if len(args) == 0 {
		query = selectAll
	}

	return query, args
}
