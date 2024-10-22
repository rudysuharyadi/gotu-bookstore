package shopping_cart

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/shopping_cart"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
)

type PostShoppingCartClearService struct {
	shoppingCartRepo ShoppingCartRepoInterface
	userRepo         UserRepoInterface
	utils.CommonContext
}

func NewPostShoppingCartClearService(
	context utils.CommonContext,
	shoppingCartRepo ShoppingCartRepoInterface,
	userRepo UserRepoInterface,
) PostShoppingCartClearService {
	return PostShoppingCartClearService{
		shoppingCartRepo: shoppingCartRepo,
		userRepo:         userRepo,
		CommonContext:    context,
	}
}

func (s PostShoppingCartClearService) ProcessingPostShoppingCartClear() (*contracts.PostShoppingCartClearResponse, error) {
	// Get session data
	session, err := s.GetSession()
	if err != nil {
		s.LogError(err)
		return nil, base_error.NewUnauthorizedError(constants.IC0006)
	}

	// Get user By ID from session
	user, err := s.userRepo.GetById(session.Id)
	if err != nil {
		s.LogDebug(err)
		return nil, base_error.NewInternalError(constants.IC0011)
	}
	if user == nil {
		return nil, base_error.NewUnauthorizedError(constants.IC0006)
	}

	// Clear shopping cart
	err = s.shoppingCartRepo.ClearShoppingCart(user.Id.String())
	if err != nil {
		s.LogDebug(err)
		return nil, base_error.NewInternalError(constants.IC0020)
	}

	return nil, nil
}
