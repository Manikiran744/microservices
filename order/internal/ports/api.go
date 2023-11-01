package ports

import (
	"context"
	"github.com/Manikiran744/microservices/order/internal/application/core/domain"
)

type APIPort interface {
    PlaceOrder(context.Context, domain.Order) (domain.Order, error)
}
