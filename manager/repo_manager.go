package manager

import (
	"discount-service/infra"
	"discount-service/interfaces"
	"discount-service/repo"
	"sync"
)

type RepoManager interface {
	CheckoutRepo() interfaces.ICheckoutRepo
}

type repoManager struct {
	infra infra.Infra
}

func NewRepoManager(infra infra.Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}

var (
	checkoutRepoOnce sync.Once
	checkoutRepo     interfaces.ICheckoutRepo
)

// PackageRedeemHistory Get PackageRedeemHistory repository.
func (c *repoManager) CheckoutRepo() interfaces.ICheckoutRepo {
	checkoutRepoOnce.Do(func() {
		checkoutRepo = repo.NewCheckoutRepo(c.infra.ConnectionDB())
	})

	return checkoutRepo
}
