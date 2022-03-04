package services

import (
	"context"
	"discount-service/interfaces"
	mock_interfaces "discount-service/mocks"
	"discount-service/models"
	"discount-service/resources/request"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

// TestNewCheckoutService ...
func TestNewCheckoutService(t *testing.T) {
	type args struct {
		config       *models.Config
		checkoutRepo interfaces.ICheckoutRepo
	}
	tests := []struct {
		name string
		args args
		want func() interfaces.ICheckoutService
	}{
		{
			name: "flow normal",
			args: args{
				config:       &models.Config{},
				checkoutRepo: &mock_interfaces.MockICheckoutRepo{},
			},
			want: func() interfaces.ICheckoutService {
				return &checkoutService{
					config:       &models.Config{},
					checkoutRepo: &mock_interfaces.MockICheckoutRepo{},
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCheckoutService(tt.args.config, tt.args.checkoutRepo)
			want := tt.want()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("NewCheckoutService() = %v, want %v", got, want)
			}
		})
	}
}

// Test_checkoutService_ServiceCheckoutOrder ...
func Test_checkoutService_ServiceCheckoutOrder(t *testing.T) {
	stringTesting := "testing"
	floatTesting := float64(190)
	type fields struct {
		config       *models.Config
		checkoutRepo interfaces.ICheckoutRepo
	}
	type args struct {
		ctx context.Context
		req request.ReqCheckout
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		doMockRepo func(mock *mock_interfaces.MockICheckoutRepo)
		wantErr    error
	}{
		{
			name: "error list SKU",
			args: args{
				ctx: context.Background(),
				req: request.ReqCheckout{
					UserID: 190,
					Order:  []request.ReqCheckoutItem{},
				},
			},
			doMockRepo: func(mock *mock_interfaces.MockICheckoutRepo) {},
			wantErr:    fmt.Errorf("please check your request"),
		},
		{
			name: "error repo RepoGetItemsList",
			args: args{
				ctx: context.Background(),
				req: request.ReqCheckout{
					UserID: 190,
					Order: []request.ReqCheckoutItem{
						{
							SKU: "testing",
							Qty: 1,
						},
					},
				},
			},
			doMockRepo: func(mock *mock_interfaces.MockICheckoutRepo) {
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{}, errors.New("error"))
			},
			wantErr: fmt.Errorf("please try again"),
		},
		{
			name: "error data == 0 from repo RepoGetItemsList",
			args: args{
				ctx: context.Background(),
				req: request.ReqCheckout{
					UserID: 190,
					Order: []request.ReqCheckoutItem{
						{
							SKU: "testing",
							Qty: 1,
						},
					},
				},
			},
			doMockRepo: func(mock *mock_interfaces.MockICheckoutRepo) {
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{}, nil)
			},
			wantErr: fmt.Errorf("all sku not found %v", "[testing]"),
		},
		{
			name: "error checkSKUNotFoundAndQty",
			args: args{
				ctx: context.Background(),
				req: request.ReqCheckout{
					UserID: 190,
					Order: []request.ReqCheckoutItem{
						{
							SKU: "testing",
							Qty: 1,
						},
						{
							SKU: "testing-2",
							Qty: 1,
						},
					},
				},
			},
			doMockRepo: func(mock *mock_interfaces.MockICheckoutRepo) {
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{
					{
						SKU:      "testing",
						ID:       1,
						Name:     "macbook",
						Price:    160,
						Currency: "USD",
						Qty:      1,
					},
				}, nil)
			},
			wantErr: fmt.Errorf("please check your item request"),
		},
		{
			name: "error checkSKUNotFoundAndQtykedua",
			args: args{
				ctx: context.Background(),
				req: request.ReqCheckout{
					UserID: 190,
					Order: []request.ReqCheckoutItem{
						{
							SKU: "testing",
							Qty: 7,
						},
					},
				},
			},
			doMockRepo: func(mock *mock_interfaces.MockICheckoutRepo) {
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{
					{
						SKU:      "testing",
						ID:       1,
						Name:     "macbook",
						Price:    160,
						Currency: "USD",
						Qty:      5,
					},
				}, nil)
			},
			wantErr: fmt.Errorf("please check stok sku items : [testing]"),
		},
		{
			name: "error RepoPromotionActiveList",
			args: args{
				ctx: context.Background(),
				req: request.ReqCheckout{
					UserID: 190,
					Order: []request.ReqCheckoutItem{
						{
							SKU: "testing",
							Qty: 1,
						},
					},
				},
			},
			doMockRepo: func(mock *mock_interfaces.MockICheckoutRepo) {
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{
					{
						SKU:      "testing",
						ID:       1,
						Name:     "macbook",
						Price:    160,
						Currency: "USD",
						Qty:      1,
					},
				}, nil)
				mock.EXPECT().RepoPromotionActiveList(gomock.Any()).Return([]models.PromotionItems{}, errors.New("error"))
			},
			wantErr: fmt.Errorf("please try again"),
		},
		{
			name: "error BeginTrx",
			args: args{
				ctx: context.Background(),
				req: request.ReqCheckout{
					UserID: 190,
					Order: []request.ReqCheckoutItem{
						{
							SKU: "testing",
							Qty: 1,
						},
					},
				},
			},
			doMockRepo: func(mock *mock_interfaces.MockICheckoutRepo) {
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{
					{
						SKU:      "testing",
						ID:       1,
						Name:     "macbook",
						Price:    160,
						Currency: "USD",
						Qty:      1,
					},
				}, nil)
				mock.EXPECT().RepoPromotionActiveList(gomock.Any()).Return([]models.PromotionItems{}, nil)
				mock.EXPECT().BeginTrx(gomock.Any()).Return(errors.New("error"))
			},
			wantErr: fmt.Errorf("error transaction"),
		},
		{
			name: "error RepoGetItemsList second",
			args: args{
				ctx: context.Background(),
				req: request.ReqCheckout{
					UserID: 190,
					Order: []request.ReqCheckoutItem{
						{
							SKU: "testing",
							Qty: 1,
						},
					},
				},
			},
			doMockRepo: func(mock *mock_interfaces.MockICheckoutRepo) {
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{
					{
						SKU:      "testing",
						ID:       1,
						Name:     "macbook",
						Price:    160,
						Currency: "USD",
						Qty:      1,
					},
				}, nil)
				mock.EXPECT().RepoPromotionActiveList(gomock.Any()).Return([]models.PromotionItems{
					{
						ItemID:   190,
						ItemSKU:  "testing",
						MinQty:   1,
						FreeItem: &stringTesting,
					},
				}, nil)
				mock.EXPECT().BeginTrx(gomock.Any()).Return(nil)
				mock.EXPECT().RoolBackTrx(gomock.Any()).Return(nil)
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{}, errors.New("error"))
			},
			wantErr: fmt.Errorf("please try again"),
		},
		{
			name: "error RepoInsertOrder",
			args: args{
				ctx: context.Background(),
				req: request.ReqCheckout{
					UserID: 190,
					Order: []request.ReqCheckoutItem{
						{
							SKU: "testing",
							Qty: 1,
						},
					},
				},
			},
			doMockRepo: func(mock *mock_interfaces.MockICheckoutRepo) {
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{
					{
						SKU:      "testing",
						ID:       1,
						Name:     "macbook",
						Price:    160,
						Currency: "USD",
						Qty:      1,
					},
				}, nil)
				mock.EXPECT().RepoPromotionActiveList(gomock.Any()).Return([]models.PromotionItems{
					{
						ItemID:   190,
						ItemSKU:  "testing",
						MinQty:   1,
						FreeItem: &stringTesting,
						Discount: &floatTesting,
					},
				}, nil)
				mock.EXPECT().BeginTrx(gomock.Any()).Return(nil)
				mock.EXPECT().RoolBackTrx(gomock.Any()).Return(nil)
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{
					{
						SKU:      "testing",
						Name:     "macbook",
						Price:    160,
						Currency: "USD",
						Qty:      1,
					},
				}, nil)
				mock.EXPECT().RepoInsertOrder(gomock.Any(), gomock.Any()).Return(errors.New("error"))

			},
			wantErr: fmt.Errorf("please try again"),
		},
		{
			name: "error RepoUpdateItems",
			args: args{
				ctx: context.Background(),
				req: request.ReqCheckout{
					UserID: 190,
					Order: []request.ReqCheckoutItem{
						{
							SKU: "testing",
							Qty: 1,
						},
					},
				},
			},
			doMockRepo: func(mock *mock_interfaces.MockICheckoutRepo) {
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{
					{
						SKU:      "testing",
						ID:       1,
						Name:     "macbook",
						Price:    160,
						Currency: "USD",
						Qty:      1,
					},
				}, nil)
				mock.EXPECT().RepoPromotionActiveList(gomock.Any()).Return([]models.PromotionItems{
					{
						ItemID:   190,
						ItemSKU:  "testing",
						MinQty:   1,
						FreeItem: &stringTesting,
						Discount: &floatTesting,
					},
				}, nil)
				mock.EXPECT().BeginTrx(gomock.Any()).Return(nil)
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{
					{
						SKU:      "testing",
						Name:     "macbook",
						Price:    160,
						Currency: "USD",
						Qty:      1,
					},
				}, nil)
				mock.EXPECT().RepoInsertOrder(gomock.Any(), gomock.Any()).Return(nil)
				mock.EXPECT().RepoUpdateItems(gomock.Any(), gomock.Any()).Return(errors.New("error"))
				mock.EXPECT().RoolBackTrx(gomock.Any()).Return(nil)
			},
			wantErr: fmt.Errorf("please try again"),
		},
		{
			name: "flow normal",
			args: args{
				ctx: context.Background(),
				req: request.ReqCheckout{
					UserID: 190,
					Order: []request.ReqCheckoutItem{
						{
							SKU: "testing",
							Qty: 1,
						},
					},
				},
			},
			doMockRepo: func(mock *mock_interfaces.MockICheckoutRepo) {
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{
					{
						SKU:      "testing",
						ID:       1,
						Name:     "macbook",
						Price:    160,
						Currency: "USD",
						Qty:      1,
					},
				}, nil)
				mock.EXPECT().RepoPromotionActiveList(gomock.Any()).Return([]models.PromotionItems{
					{
						ItemID:   190,
						ItemSKU:  "testing",
						MinQty:   1,
						FreeItem: &stringTesting,
						Discount: &floatTesting,
					},
				}, nil)
				mock.EXPECT().BeginTrx(gomock.Any()).Return(nil)
				mock.EXPECT().RepoGetItemsList(gomock.Any(), gomock.Any()).Return([]models.Items{
					{
						SKU:      "testing",
						Name:     "macbook",
						Price:    160,
						Currency: "USD",
						Qty:      1,
					},
				}, nil)
				mock.EXPECT().RepoInsertOrder(gomock.Any(), gomock.Any()).Return(nil)
				mock.EXPECT().RepoUpdateItems(gomock.Any(), gomock.Any()).Return(nil)
				mock.EXPECT().RepoUpdateItems(gomock.Any(), gomock.Any()).Return(nil)
				mock.EXPECT().CommitTrx(gomock.Any()).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &checkoutService{
				config:       tt.fields.config,
				checkoutRepo: tt.fields.checkoutRepo,
			}

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := mock_interfaces.NewMockICheckoutRepo(mockCtrl)
			tt.doMockRepo(mockRepo)
			cs.checkoutRepo = mockRepo

			err := cs.ServiceCheckoutOrder(tt.args.ctx, tt.args.req)
			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("checkoutService.ServiceCheckoutOrder() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
