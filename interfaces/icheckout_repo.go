package interfaces

import (
	"context"
	"discount-service/models"
)

type ICheckoutRepo interface {
	// db
	BeginTrx(ctx context.Context) (err error)
	RoolBackTrx(ctx context.Context) (err error)
	CommitTrx(ctx context.Context) (err error)

	RepoGetItemsList(ctx context.Context, sku []string) ([]models.Items, error)
	RepoPromotionActiveList(ctx context.Context) ([]models.PromotionItems, error)
	RepoInsertOrder(ctx context.Context, data []models.Orders) error
	RepoUpdateItems(ctx context.Context, data models.Orders) error
}
