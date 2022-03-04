package handlers

import (
	"bytes"
	"discount-service/interfaces"
	mock_interfaces "discount-service/mocks"
	"discount-service/resources/request"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

// TestNewCheckoutHandler ...
func TestNewCheckoutHandler(t *testing.T) {
	type args struct {
		checkoutService interfaces.ICheckoutService
	}
	tests := []struct {
		name string
		args args
		want func() *checkoutHandler
	}{
		{
			name: "success",
			args: args{
				checkoutService: &mock_interfaces.MockICheckoutService{},
			},
			want: func() *checkoutHandler {
				return &checkoutHandler{
					checkoutService: &mock_interfaces.MockICheckoutService{},
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCheckoutHandler(tt.args.checkoutService)
			want := tt.want()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("NewCheckoutHandler() = %v, want %v", got, want)
			}
		})
	}
}

// Test_checkoutHandler_HandlerCheckoutOrder ...
func Test_checkoutHandler_HandlerCheckoutOrder(t *testing.T) {
	url := "/checkout"
	type fields struct {
		checkoutService interfaces.ICheckoutService
	}
	tests := []struct {
		name          string
		fields        fields
		initRouter    func() (*http.Request, *httptest.ResponseRecorder)
		doMockService func(mock *mock_interfaces.MockICheckoutService)
	}{
		{
			name:   "error from request",
			fields: fields{},
			initRouter: func() (*http.Request, *httptest.ResponseRecorder) {
				r, _ := http.NewRequest(http.MethodGet, url, nil)
				r.Method = "GET"
				r.Body = http.NoBody
				w := httptest.NewRecorder()

				return r, w
			},
			doMockService: func(mock *mock_interfaces.MockICheckoutService) {},
		},
		{
			name:   "error from service",
			fields: fields{},
			initRouter: func() (*http.Request, *httptest.ResponseRecorder) {
				objReq := request.ReqCheckout{
					UserID: 190,
				}
				reqByte, _ := json.Marshal(objReq)
				r, _ := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(reqByte))
				w := httptest.NewRecorder()

				return r, w
			},
			doMockService: func(mock *mock_interfaces.MockICheckoutService) {
				mock.EXPECT().ServiceCheckoutOrder(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error"))
			},
		},
		{
			name:   "flow normal",
			fields: fields{},
			initRouter: func() (*http.Request, *httptest.ResponseRecorder) {
				objReq := request.ReqCheckout{
					UserID: 190,
				}
				reqByte, _ := json.Marshal(objReq)
				r, _ := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(reqByte))
				w := httptest.NewRecorder()

				return r, w
			},
			doMockService: func(mock *mock_interfaces.MockICheckoutService) {
				mock.EXPECT().ServiceCheckoutOrder(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &checkoutHandler{
				checkoutService: tt.fields.checkoutService,
			}

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockService := mock_interfaces.NewMockICheckoutService(mockCtrl)
			c.checkoutService = mockService
			tt.doMockService(mockService)

			r, w := tt.initRouter()
			c.HandlerCheckoutOrder(w, r)
		})
	}
}
