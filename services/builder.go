package services

import (
	"discount-service/models"
	"discount-service/resources/request"
	"fmt"
)

func builderCheckoutRequest(req request.ReqCheckout) ([]string, map[string]int) {
	var mapItems = make(map[string]int)
	listSKU := []string{}

	for _, v := range req.Order {
		val, ok := mapItems[v.SKU]
		if ok {
			mapItems[v.SKU] = val + v.Qty
		} else {
			listSKU = append(listSKU, v.SKU)
			mapItems[v.SKU] = v.Qty
		}
	}

	return listSKU, mapItems
}

func checkSKUNotFoundAndQty(mapItem map[string]int, data []models.Items) (map[string]models.Items, error) {
	skuStok := []string{}
	mapListItem := map[string]models.Items{}

	if len(data) != len(mapItem) {
		return nil, fmt.Errorf("please check your item request")
	}

	for _, v := range data {
		val, ok := mapItem[v.SKU]
		if ok {
			// check stok
			if v.Qty < val {
				skuStok = append(skuStok, v.SKU)
			}
		}

		mapListItem[v.SKU] = models.Items{
			ID:    v.ID,
			SKU:   v.SKU,
			Name:  v.Name,
			Qty:   val,
			Price: v.Price,
		}
	}

	if len(skuStok) > 0 {
		return nil, fmt.Errorf("please check stok sku items : %v", skuStok)
	}

	return mapListItem, nil
}
