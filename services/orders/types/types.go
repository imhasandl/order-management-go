package types

import (
	"context"

	"github.com/imhasandl/grpc-go-project/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}