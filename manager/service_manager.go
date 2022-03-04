package manager

import (
	"discount-service/infra"
	"discount-service/interfaces"
	"discount-service/services"
	"sync"
)

type ServiceManager interface {
	CheckoutService() interfaces.ICheckoutService
}

type serviceManager struct {
	infra infra.Infra
	repo  RepoManager
}

func NewServiceManager(data infra.Infra) ServiceManager {
	return &serviceManager{
		infra: data,
		repo:  NewRepoManager(data),
	}
}

var (
	checkoutServiceOnce sync.Once
	checkoutService     interfaces.ICheckoutService
)

func (c *serviceManager) CheckoutService() interfaces.ICheckoutService {
	checkoutServiceOnce.Do(func() {
		checkoutService = services.NewCheckoutService(
			c.infra.ConfigInfra(),
			c.repo.CheckoutRepo(),
		)
	})

	return checkoutService
}
