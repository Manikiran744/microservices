package ports

import (
	"context"
	"github.com/Manikiran744/microservices/order/internal/application/core/domain"
)

type PaymentPort interface {
    Charge(context.Context, *domain.Order) error
}
