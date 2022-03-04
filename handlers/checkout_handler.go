package handlers

import (
	"discount-service/interfaces"
	"discount-service/resources/request"
	"discount-service/resources/response"
	"encoding/json"
	"net/http"
)

type checkoutHandler struct {
	checkoutService interfaces.ICheckoutService
}

func NewCheckoutHandler(
	checkoutService interfaces.ICheckoutService,
) *checkoutHandler {
	return &checkoutHandler{
		checkoutService: checkoutService,
	}
}

func (c *checkoutHandler) HandlerCheckoutOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req request.ReqCheckout

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		response.BadRequest(w, "please check yout request")
		return
	}

	err = c.checkoutService.ServiceCheckoutOrder(ctx, req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.Write(w, http.StatusOK, nil)
}
