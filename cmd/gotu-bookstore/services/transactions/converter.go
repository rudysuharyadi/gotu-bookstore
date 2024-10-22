package transactions

import (
	"gotu-bookstore/cmd/gotu-bookstore/contracts/dto"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"gotu-bookstore/pkg/utils"
	"time"
)

func ConvertToTransactionDTO(transaction models.Transactions) (*dto.TransactionDTO, error) {
	transactionItemDTOs, err := ConvertToTransactionItemDTO(transaction.TransactionItems)
	if err != nil {
		return nil, err
	}

	transactionDTO := dto.TransactionDTO{
		Id:            transaction.Id.String(),
		InvoiceNumber: transaction.InvoiceNumber,
		GrandTotal:    utils.FloatToString(transaction.GrandTotal),
		Status:        transaction.Status,
		CreatedAt:     transaction.CreatedAt.Format(time.RFC3339Nano),
		Items:         transactionItemDTOs,
	}

	return &transactionDTO, nil
}

func ConvertToTransactionDTOs(transactions []models.Transactions) ([]dto.TransactionDTO, error) {
	var transactionDTOs = make([]dto.TransactionDTO, 0)
	for _, transaction := range transactions {
		transactionDTO, err := ConvertToTransactionDTO(transaction)
		if err != nil {
			return nil, err
		}

		transactionDTOs = append(transactionDTOs, *transactionDTO)
	}

	return transactionDTOs, nil
}

func ConvertToTransactionItemDTO(transactionItems []models.TransactionItems) ([]dto.TransactionItemDTO, error) {
	var transactionItemDTOs = make([]dto.TransactionItemDTO, 0)
	for _, transactionDetail := range transactionItems {
		transactionItemDTOs = append(transactionItemDTOs, dto.TransactionItemDTO{
			BookId:   transactionDetail.BookId.String(),
			Quantity: transactionDetail.Quantity,
			Price:    utils.FloatToString(transactionDetail.Price),
		})
	}
	return transactionItemDTOs, nil
}
