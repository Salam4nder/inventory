package persistence

// No unit tests because I can't find a way to invoke
// the sqlMock methods that are needed for each case.

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Salam4nder/inventory/internal/entity"

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

	mock.ExpectQuery(
		"SELECT * FROM inventory WHERE name = $1 AND unit = $2 AND amount = $3").WithArgs(
		filter.Name, filter.Unit, filter.Amount).WillReturnError(
		errors.New("query failed"))
	_, err = storage.ReadBy(ctx, filter)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}
}
