package persistence

// No unit tests because I can't find a way to invoke
// the sqlMock methods that are needed for each case.

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Salam4nder/inventory/internal/entity"
	"github.com/stretchr/testify/assert"

	sqlMock "github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func Test_Create_Success(t *testing.T) {
	driver, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := entity.Item{
		ID:        uuid.New(),
		Name:      "test",
		Unit:      "kg",
		Amount:    1.1,
		ExpiresAt: time.Now().Add(1 * time.Minute),
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO inventory").WithArgs(
		item.Name, item.Unit,
		item.Amount, item.ExpiresAt).WillReturnRows(
		sqlMock.NewRows([]string{"id"}).AddRow(item.ID))
	mock.ExpectCommit()

	_, err = storage.Create(ctx, item)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Create_Rollback_And_Error_On_Timeout(t *testing.T) {
	driver, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := entity.Item{
		ID:        uuid.New(),
		Name:      "test",
		Unit:      "kg",
		Amount:    1.1,
		ExpiresAt: time.Now().Add(6 * time.Minute),
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 500*time.Millisecond)
	defer cancel()

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO inventory").WithArgs(
		item.Name, item.Unit,
		item.Amount, item.ExpiresAt).WillDelayFor(
		1 * time.Second)
	mock.ExpectRollback()

	_, err = storage.Create(ctx, item)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Create_Error_On_BeginTX_Returns_Error(t *testing.T) {
	driver, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := entity.Item{
		ID:        uuid.New(),
		Name:      "test",
		Unit:      "kg",
		Amount:    1.1,
		ExpiresAt: time.Now().Add(6 * time.Minute),
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 500*time.Millisecond)
	defer cancel()

	mock.ExpectBegin().WillReturnError(
		errors.New("begin tx fails"))

	_, err = storage.Create(ctx, item)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Create_Rollback_And_Error_On_Bad_Args(t *testing.T) {
	driver, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := entity.Item{
		ID:        uuid.New(),
		Name:      "bad arg",
		Unit:      "kg",
		Amount:    1.1,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO inventory").WithArgs(
		item.Name, item.Unit,
		item.Amount, item.ExpiresAt).WillReturnError(
		errors.New("bad arg"))
	mock.ExpectRollback()

	_, err = storage.Create(ctx, item)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Create_Commit_Fails_Returns_Error(t *testing.T) {
	driver, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := entity.Item{
		ID:        uuid.New(),
		Name:      "test",
		Unit:      "kg",
		Amount:    1.1,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO inventory").WithArgs(
		item.Name, item.Unit,
		item.Amount, item.ExpiresAt).WillReturnRows(
		sqlMock.NewRows([]string{"id"}).AddRow(item.ID))
	mock.ExpectCommit().WillReturnError(
		errors.New("commit fails"))

	_, err = storage.Create(ctx, item)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Read_Success(t *testing.T) {
	driver, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := entity.Item{
		ID:        uuid.New(),
		Name:      "test",
		Unit:      "kg",
		Amount:    1.1,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectQuery("SELECT FROM inventory").WithArgs(
		item.ID.String()).WillReturnRows(
		sqlMock.NewRows([]string{
			"id", "name", "unit", "amount", "expires_at"}).AddRow(
			item.ID, item.Name, item.Unit, item.Amount, item.ExpiresAt))

	_, err = storage.Read(ctx, item.ID.String())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_Read_Fails_No_Match(t *testing.T) {
	driver, mock, err := sqlMock.New()
	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := entity.Item{
		ID:        uuid.New(),
		Name:      "test",
		Unit:      "kg",
		Amount:    1.1,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectQuery("SELECT FROM inventory").WithArgs(
		item.ID.String()).WillReturnError(
		errors.New("no match"))

	_, err = storage.Read(ctx, item.ID.String())
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_ReadAll_Success(t *testing.T) {
	driver, mock, err := sqlMock.New(
		sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := entity.Item{
		ID:        uuid.New(),
		Name:      "test",
		Unit:      "kg",
		Amount:    1.1,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectQuery("SELECT * FROM inventory").WillReturnRows(
		sqlMock.NewRows([]string{
			"id", "name", "unit", "amount", "expires_at"}).AddRow(
			item.ID, item.Name, item.Unit, item.Amount, item.ExpiresAt))

	_, err = storage.ReadAll(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_ReadAll_Query_Fails_Returns_Error(t *testing.T) {
	driver, mock, err := sqlMock.New(
		sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectQuery("SELECT * FROM inventory").WillReturnError(
		errors.New("query fails"))

	_, err = storage.ReadAll(ctx)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_ReadBy_Success(t *testing.T) {
	driver, mock, err := sqlMock.New(
		sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := entity.Item{
		ID:     uuid.New(),
		Name:   "test",
		Unit:   "kg",
		Amount: 1.1,
	}

	filter := entity.ItemFilter{
		Name:   "test",
		Unit:   "kg",
		Amount: 1.1,
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectQuery(
		"SELECT * FROM inventory WHERE name = $1 AND unit = $2 AND amount = $3").WithArgs(
		filter.Name, filter.Unit, filter.Amount).WillReturnRows(
		sqlMock.NewRows([]string{
			"id", "name", "unit", "amount", "expires_at"}).AddRow(
			item.ID, item.Name, item.Unit, item.Amount, item.ExpiresAt))
	_, err = storage.ReadBy(ctx, filter)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_ReadBy_Fails_With_No_Filter(t *testing.T) {
	driver, _, err := sqlMock.New(
		sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	filter := entity.ItemFilter{}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	_, err = storage.ReadBy(ctx, filter)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_ReadBy_Query_Fails_Returns_Error(t *testing.T) {
	driver, mock, err := sqlMock.New(
		sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	filter := entity.ItemFilter{
		Name:   "test",
		Unit:   "kg",
		Amount: 1.1,
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectBegin()
	mock.ExpectQuery(
		"SELECT * FROM inventory WHERE name = $1 AND unit = $2 AND amount = $3").WithArgs(
		filter.Name, filter.Unit, filter.Amount).WillReturnError(
		errors.New("query failed"))
	mock.ExpectRollback()

	_, err = storage.ReadBy(ctx, filter)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_Update_Success(t *testing.T) {
	driver, mock, err := sqlMock.New(
		sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := &entity.Item{
		ID:        uuid.New(),
		Name:      "test",
		Unit:      "kg",
		Amount:    1.1,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectBegin()
	mock.ExpectExec(
		"UPDATE inventory SET name = $1, unit = $2, amount = $3, expires_at = $4 WHERE id = $5").WithArgs(
		item.Name, item.Unit, item.Amount, item.ExpiresAt, item.ID).WillReturnResult(
		sqlMock.NewResult(1, 1))
	mock.ExpectCommit()

	_, err = storage.Update(ctx, item)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_Update_Fails_With_No_ID(t *testing.T) {
	driver, mock, err := sqlMock.New(
		sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := &entity.Item{
		Name:      "test",
		Unit:      "kg",
		Amount:    1.1,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectBegin()
	mock.ExpectExec(
		"UPDATE inventory SET name = $1, unit = $2, amount = $3, expires_at = $4 WHERE id = $5").WithArgs().WillReturnError(
		errors.New("no id"))
	mock.ExpectRollback()

	_, err = storage.Update(ctx, item)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_Delete_Success(t *testing.T) {
	driver, mock, err := sqlMock.New(
		sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := entity.Item{
		ID:        uuid.New(),
		Name:      "test",
		Unit:      "kg",
		Amount:    1.1,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectBegin()
	mock.ExpectExec(
		"DELETE FROM inventory WHERE id = $1").WithArgs(
		item.ID).WillReturnResult(
		sqlMock.NewResult(1, 1))
	mock.ExpectCommit()

	if err = storage.Delete(ctx, item.ID.String()); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_Delete_Returns_Error_On_Fail(t *testing.T) {
	driver, mock, err := sqlMock.New(
		sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf(
			"unexpected error while creating sqlmock: %s",
			err)
	}
	defer driver.Close()

	storage := Storage{
		DB: driver,
	}

	item := entity.Item{
		ID:        uuid.New(),
		Name:      "test",
		Unit:      "kg",
		Amount:    1.1,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	mock.ExpectBegin()
	mock.ExpectExec(
		"DELETE FROM inventory WHERE id = $1").WithArgs(
		item.ID).WillReturnError(
		errors.New("delete failed"))
	mock.ExpectRollback()

	if err = storage.Delete(ctx, item.ID.String()); err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_filterQueryBuilder(t *testing.T) {
	expiration := time.Now().Add(5 * time.Minute)

	tests := []struct {
		name      string
		input     entity.ItemFilter
		wantQuery string
		wantArgs  []interface{}
	}{
		{
			name:      "empty filter returns default query",
			input:     entity.ItemFilter{},
			wantQuery: "SELECT * FROM inventory",
			wantArgs:  nil,
		},
		{
			name: "filter with name returns query with name",
			input: entity.ItemFilter{
				Name: "test",
			},
			wantQuery: "SELECT * FROM inventory WHERE name = $1",
			wantArgs:  []interface{}{"test"},
		},
		{
			name: "filter with unit returns query with unit",
			input: entity.ItemFilter{
				Unit: "kg",
			},
			wantQuery: "SELECT * FROM inventory WHERE unit = $1",
			wantArgs:  []interface{}{"kg"},
		},
		{
			name: "filter with amount returns query with amount",
			input: entity.ItemFilter{
				Amount: 1.1,
			},
			wantQuery: "SELECT * FROM inventory WHERE amount = $1",
			wantArgs:  []interface{}{1.1},
		},
		{
			name: "filter with expiresat returns query with expiresat",
			input: entity.ItemFilter{
				ExpiresAt: expiration,
			},
			wantQuery: "SELECT * FROM inventory WHERE expires_at = $1",
			wantArgs:  []interface{}{expiration},
		},
		{
			name: "filter with name, amount returns query with name and amount",
			input: entity.ItemFilter{
				Name:   "test",
				Amount: 1.1,
			},
			wantQuery: "SELECT * FROM inventory WHERE name = $1 AND amount = $2",
			wantArgs:  []interface{}{"test", 1.1},
		},
		{
			name: "filter with unit, expiresat returns query with unit and expiresat",
			input: entity.ItemFilter{
				Unit:      "kg",
				ExpiresAt: expiration,
			},
			wantQuery: "SELECT * FROM inventory WHERE unit = $1 AND expires_at = $2",
			wantArgs:  []interface{}{"kg", expiration},
		},
		{
			name: "full filter returns expected query",
			input: entity.ItemFilter{
				Name:      "test",
				Unit:      "kg",
				Amount:    1.1,
				ExpiresAt: expiration,
			},
			wantQuery: "SELECT * FROM inventory WHERE name = $1 AND unit = $2 AND amount = $3 AND expires_at = $4",
			wantArgs: []interface{}{
				"test", "kg", 1.1, expiration},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotQuery, gotArgs := filterQueryBuilder(test.input)

			if test.wantQuery != gotQuery {
				t.Errorf("unexpected query: %s", gotQuery)
			}
			assert.Equal(t, test.wantArgs, gotArgs)
		})
	}
}
