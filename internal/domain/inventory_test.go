package domain

import (
    "testing"
    "time"
    "context"

    "github.com/Salam4nder/inventory/internal/mock"
    "github.com/Salam4nder/inventory/internal/entity"

    "github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
    tests := []struct {
        name string
        input []entity.Item
        want []entity.Item
        ctxTimeout time.Duration
        wantErr bool
    } {
        {
            name: "test1",
        },
    }

    for _,test := range tests {
        t.Run(test.name, func(t *testing.T) {
             mock := mock.NewPersistence()

            service := New(mock)

            ctx, cancel := context.WithTimeout(
                context.Background(), test.ctxTimeout)
            defer cancel()

            service.Create(ctx,test.input)
        })
    }
}
