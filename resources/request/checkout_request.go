package request

type ReqCheckout struct {
	UserID int               `json:"user_id"`
	Order  []ReqCheckoutItem `json:"order"`
}

type ReqCheckoutItem struct {
	SKU string `json:"sku"`
	Qty int    `json:"qty"`
}
