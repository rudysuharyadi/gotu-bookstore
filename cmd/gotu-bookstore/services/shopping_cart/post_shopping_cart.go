package shopping_cart

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/shopping_cart"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/pkg/resfmt/base_error"
	"gotu-bookstore/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type PostShoppingCartService struct {
	shoppingCartRepo ShoppingCartRepoInterface
	userRepo         UserRepoInterface
	bookRepo         BookRepoInterface
	utils.CommonContext
}

func NewPostShoppingCartService(
	context utils.CommonContext,
	shoppingCartRepo ShoppingCartRepoInterface,
	userRepo UserRepoInterface,
	bookRepo BookRepoInterface,
) PostShoppingCartService {
	return PostShoppingCartService{
		shoppingCartRepo: shoppingCartRepo,
		userRepo:         userRepo,
		bookRepo:         bookRepo,
		CommonContext:    context,
	}
}

func (s PostShoppingCartService) ProcessingPostShoppingCart(request contracts.PostShoppingCartRequest) (*contracts.ShoppingCartResponse, error) {
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

	// Get book data from request
	book, err := s.bookRepo.GetById(request.BookId)
	if err != nil {
		s.LogDebug(err)
		return nil, base_error.NewInternalError(constants.IC0014)
	}

	// Check existing cartItem
	existingItem, err := s.shoppingCartRepo.GetItemByBookId(book.Id.String(), user.Id.String())
	if err != nil {
		s.LogDebug(err)
		return nil, base_error.NewInternalError(constants.IC0018)
	}
	if existingItem != nil {
		if request.Quantity == 0 {
			// Delete
			err = s.shoppingCartRepo.DeleteByBookId(book.Id.String(), user.Id.String())
			if err != nil {
				s.LogDebug(err)
				return nil, base_error.NewInternalError(constants.IC0022)
			}

		} else {
			// Update quantity
			err = s.shoppingCartRepo.UpdateByBookId(book.Id.String(), int64(request.Quantity), user.Id.String())
			if err != nil {
				s.LogDebug(err)
				return nil, base_error.NewInternalError(constants.IC0019)
			}
		}

	} else {
		if request.Quantity == 0 {
			return nil, base_error.NewBadRequestError(constants.IC0009)
		}

		// Insert
		uuid, err := uuid.NewV7()
		if err != nil {
			s.LogError(err)
			return nil, base_error.NewInternalError(constants.IC0015)
		}

		item := models.ShoppingCarts{
			Id:        uuid,
			UserId:    user.Id,
			BookId:    book.Id,
			Quantity:  int64(request.Quantity),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		_, err = s.shoppingCartRepo.AddItemToCart(item)
		if err != nil {
			s.LogDebug(err)
			return nil, base_error.NewInternalError(constants.IC0016)
		}
	}

	// Get shopping cart data
	cartItems, err := s.shoppingCartRepo.GetByUserId(user.Id.String())
	if err != nil {
		s.LogDebug(err)
		return nil, base_error.NewInternalError(constants.IC0017)
	}

	return ConvertToShoppingCartResponse(cartItems)
}
