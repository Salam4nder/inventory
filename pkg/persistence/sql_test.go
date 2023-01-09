package persistence

// No unit tests because I can't find a way to invoke
// the sqlMock methods that are needed for each case.

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Salam4nder/inventory/pkg/entity"

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

func Test_Create_Error_On_Begin_TX(t *testing.T) {
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
