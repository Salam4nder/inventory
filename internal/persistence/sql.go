package persistence

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/google/uuid"
)

// Create creates a new item in the database.
func (s *SQLDatabase) Create(
	ctx context.Context, item Item) (uuid.UUID, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback()

	query := `INSERT INTO inventory (
        name, unit, amount, expires_at)
        VALUES (
        $1, $2, $3, $4) RETURNING id`

	if err := tx.QueryRowContext(
		ctx,
		query,
		item.Name,
		item.Unit,
		item.Amount,
		item.ExpiresAt).
		Scan(&item.ID); err != nil {
		return uuid.Nil, err
	}

	if err := tx.Commit(); err != nil {
		return uuid.Nil, err
	}

	return item.ID, nil
}

// Read reads an item from the database based off of an uuid.
func (s *SQLDatabase) Read(
	ctx context.Context, uuid string) (Item, error) {
	var item Item

	query := `SELECT * FROM inventory WHERE id = $1`

	if err := s.DB.QueryRowContext(
		ctx,
		query,
		uuid).Scan(
		&item.ID,
		&item.Name,
		&item.Unit,
		&item.Amount,
		&item.ExpiresAt); err != nil {
		if err == sql.ErrNoRows {
			return Item{}, ErrNotFound
		}
		return Item{}, err
	}

	return item, nil
}

// ReadAll reads all items from the database.
func (s *SQLDatabase) ReadAll(
	ctx context.Context) ([]Item, error) {
	query := `SELECT * FROM inventory`

	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item

	for rows.Next() {
		var item Item

		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Unit,
			&item.Amount,
			&item.ExpiresAt); err != nil {
			if err == sql.ErrNoRows {
				return []Item{}, ErrNotFound
			}

			return []Item{}, err
		}

		items = append(items, item)
	}

	if len(items) < 1 {
		return []Item{}, nil
	}

	return items, nil
}

// ReadBy reads items fro the database by the given filter.
// It returns an error if the filter is empty.
func (s *SQLDatabase) ReadBy(
	ctx context.Context, filter ItemFilter) (
	[]Item, error) {
	var items []Item

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
		var item Item
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Unit,
			&item.Amount,
			&item.ExpiresAt); err != nil {
			if err == sql.ErrNoRows {
				return []Item{}, ErrNotFound
			}

			return []Item{}, err
		}
		items = append(items, item)
	}

	err = rows.Err()

	return items, err
}

// Update updates an item in the database.
func (s *SQLDatabase) Update(
	ctx context.Context, item Item) (Item, error) {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return Item{}, err
	}
	defer tx.Rollback()

	query := `UPDATE inventory SET
        name = $1, unit = $2, amount = $3, expires_at = $4
        WHERE id = $5`

	if _, err := tx.ExecContext(
		ctx,
		query,
		item.Name,
		item.Unit,
		item.Amount,
		item.ExpiresAt,
		item.ID); err != nil {
		if err == sql.ErrNoRows {
			return Item{}, ErrNotFound
		}

		return Item{}, err
	}

	if err := tx.Commit(); err != nil {
		return Item{}, err
	}

	return item, nil
}

// Delete deletes an item from the database.
func (s *SQLDatabase) Delete(
	ctx context.Context, uuid string) error {
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `DELETE FROM inventory WHERE id = $1`

	if _, err := tx.ExecContext(
		ctx,
		query,
		uuid); err != nil {
		if err == sql.ErrNoRows {
			return ErrNotFound
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func filterQueryBuilder(filter ItemFilter) (
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
		query = ""
	}

	return query, args
}
