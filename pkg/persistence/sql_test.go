package persistence

import (
	"context"
	"testing"
	"time"

	"github.com/Salam4nder/inventory/pkg/entity"

	sqlMock "github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func TestCreateSuccess(t *testing.T) {
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
