package services

import (
	"context"
	"discount-service/interfaces"
	"discount-service/models"
	"discount-service/resources/request"
	"fmt"
	"math"
	"strings"

	"github.com/rs/zerolog/log"
)

type checkoutService struct {
	config       *models.Config
	checkoutRepo interfaces.ICheckoutRepo
}

// NewCheckoutService ...
func NewCheckoutService(
	config *models.Config,
	checkoutRepo interfaces.ICheckoutRepo,
) interfaces.ICheckoutService {
	return &checkoutService{
		config:       config,
		checkoutRepo: checkoutRepo,
	}
}

// ServiceCheckoutOrder ...
func (cs *checkoutService) ServiceCheckoutOrder(ctx context.Context, req request.ReqCheckout) error {

	// mapping requestcheckout
	listSKU, mapItems := builderCheckoutRequest(req)

	if len(listSKU) == 0 {
		return fmt.Errorf("please check your request")
	}

	// get from DB for checking
	data, err := cs.checkoutRepo.RepoGetItemsList(ctx, listSKU)
	if err != nil {
		log.Error().Err(err).Msg("checkoutRepo RepoGetItemsList error")
		return fmt.Errorf("please try again")
	}

	if len(data) == 0 {
		return fmt.Errorf("all sku not found %v", listSKU)
	}

	// check any sku not found and check QTY
	mapListItems, err := checkSKUNotFoundAndQty(mapItems, data)
	if err != nil {
		log.Error().Err(err).Msg("checkSKUNotFoundAndQty error")
		return err
	}

	// check promotion_items
	respPromo, err := cs.checkoutRepo.RepoPromotionActiveList(ctx)
	if err != nil {
		log.Error().Err(err).Msg("checkoutRepo RepoPromotionActiveList error")
		return fmt.Errorf("please try again")
	}

	err = cs.checkoutRepo.BeginTrx(ctx)
	if err != nil {
		log.Error().Err(err).Msg("checkoutRepo BeginTrx error")
		return fmt.Errorf("error transaction")
	}

	var dataOrder []models.Orders
	if len(respPromo) > 0 {
		// process promo
		for _, v := range respPromo {
			order := models.Orders{}
			// check item first exist or not
			if val, ok := mapListItems[v.ItemSKU]; ok {
				// process promo check minQTY
				var discount float64
				if v.MinQty <= val.Qty {
					// category free item
					if v.FreeItem != nil && *v.FreeItem != "" {

						listItemID := strings.Split(*v.FreeItem, ",")
						data, err := cs.checkoutRepo.RepoGetItemsList(ctx, listItemID)
						if err != nil {
							cs.checkoutRepo.RoolBackTrx(ctx)
							log.Error().Err(err).Msg("checkoutRepo RepoGetItemsList error")
							return fmt.Errorf("please try again")
						}

						for _, v2 := range data {
							order = models.Orders{
								UserID:    req.UserID,
								ItemID:    int(v2.ID),
								ItemSku:   v2.SKU,
								ItemName:  v2.Name,
								ItemPrice: v2.Price,
								ItemQty:   val.Qty,
								Total:     math.Round(v2.Price * float64(val.Qty)),
							}
							dataOrder = append(dataOrder, order)
						}
					}

					// category discount
					if *v.Discount != 0 {
						discount = *v.Discount
					}
				}

				total := ((val.Price * float64(val.Qty)) - (val.Price*float64(val.Qty))*(discount/100))
				order = models.Orders{
					UserID:    req.UserID,
					ItemID:    int(val.ID),
					ItemSku:   val.SKU,
					ItemName:  val.Name,
					ItemPrice: val.Price,
					ItemQty:   val.Qty,
					Discount:  discount,
					Total:     math.Round(total*100) / 100,
				}

				dataOrder = append(dataOrder, order)
			}
		}

		// insert to order
		err = cs.checkoutRepo.RepoInsertOrder(ctx, dataOrder)
		if err != nil {
			cs.checkoutRepo.RoolBackTrx(ctx)
			log.Error().Err(err).Msg("checkoutRepo RepoInsertOrder error")
			return fmt.Errorf("please try again")
		}
	}

	// update stok
	for _, v := range dataOrder {
		err = cs.checkoutRepo.RepoUpdateItems(ctx, v)
		if err != nil {
			cs.checkoutRepo.RoolBackTrx(ctx)
			log.Error().Err(err).Msg("checkoutRepo RepoUpdateItems error")
			return fmt.Errorf("please try again")
		}
	}

	cs.checkoutRepo.CommitTrx(ctx)
	return nil
}
