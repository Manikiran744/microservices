package ports

import (
	"context"

	"github.com/Manikiran744/microservices/order/internal/application/core/domain"
)
type DBPort interface {
    Get(context.Context, int64) (domain.Order, error)
    Save(context.Context, *domain.Order) error
}