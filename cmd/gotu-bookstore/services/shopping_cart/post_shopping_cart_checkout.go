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

type PostShoppingCartCheckoutService struct {
	shoppingCartRepo ShoppingCartRepoInterface
	userRepo         UserRepoInterface
	transactionRepo  TransactionRepoInterface
	utils.CommonContext
}

func NewPostShoppingCartCheckoutService(
	context utils.CommonContext,
	shoppingCartRepo ShoppingCartRepoInterface,
	userRepo UserRepoInterface,
	transactionRepo TransactionRepoInterface,
) PostShoppingCartCheckoutService {
	return PostShoppingCartCheckoutService{
		shoppingCartRepo: shoppingCartRepo,
		userRepo:         userRepo,
		transactionRepo:  transactionRepo,
		CommonContext:    context,
	}
}

func (s PostShoppingCartCheckoutService) ProcessingPostShoppingCartCheckout() (*contracts.PostShoppingCartCheckoutResponse, error) {
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

	// Get shopping cart data
	cartItems, err := s.shoppingCartRepo.GetByUserId(user.Id.String())
	if err != nil {
		s.LogDebug(err)
		return nil, base_error.NewInternalError(constants.IC0017)
	}

	// Create transaction data
	err = s.checkout(cartItems, *user)
	if err != nil {
		s.LogError(err)
		return nil, base_error.NewInternalError(constants.IC0021)
	}

	// Clear shopping cart
	err = s.shoppingCartRepo.ClearShoppingCart(user.Id.String())
	if err != nil {
		s.LogDebug(err)
		return nil, base_error.NewInternalError(constants.IC0020)
	}

	return nil, nil
}

func (s PostShoppingCartCheckoutService) checkout(cartItems []models.ShoppingCarts, user models.Users) error {
	// Create transaction uuid
	transactionId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	// Calculate grandTotal, generate transactionItems input
	transactionItems := make([]models.TransactionItems, 0)
	grandTotal := 0.0
	for _, cartItem := range cartItems {
		transactionDetailId, err := uuid.NewV7()
		if err != nil {
			return err
		}

		inputItem := models.TransactionItems{
			Id:            transactionDetailId,
			TransactionId: transactionId,
			BookId:        cartItem.BookId,
			Quantity:      cartItem.Quantity,
			Price:         cartItem.Book.Price,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		transactionItems = append(transactionItems, inputItem)
		grandTotal += cartItem.Book.Price
	}

	// generate invoice number
	invoiceNumber, err := s.generateInvoiceNumber()
	if err != nil {
		return err
	}

	// create transactions
	input := models.Transactions{
		Id:               transactionId,
		UserId:           user.Id,
		GrandTotal:       grandTotal,
		Status:           string(constants.TransactionStatusConfirmed),
		InvoiceNumber:    invoiceNumber,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		TransactionItems: transactionItems,
	}
	_, err = s.transactionRepo.Create(input)
	if err != nil {
		return err
	}

	return nil
}

func (s PostShoppingCartCheckoutService) generateInvoiceNumber() (string, error) {
	invoiceCounter, err := s.transactionRepo.GenerateInvoiceCounter()
	if err != nil {
		return "", err
	}
	invoiceNumber := utils.ConvertBase62(invoiceCounter)
	return invoiceNumber, nil
}
