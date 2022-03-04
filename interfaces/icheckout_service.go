package interfaces

import (
	"context"
	"discount-service/resources/request"
)

type ICheckoutService interface {
	ServiceCheckoutOrder(ctx context.Context, req request.ReqCheckout) error
}
