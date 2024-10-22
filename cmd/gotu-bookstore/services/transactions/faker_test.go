package transactions_test

import (
	"gotu-bookstore/cmd/gotu-bookstore/constants"
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
)

func GenerateFakerTransactions(count int) []models.Transactions {
	faker := faker.New()
	results := make([]models.Transactions, count)

	for i := 0; i < count; i++ {
		transactionId := uuid.New()

		numberOfItems := faker.IntBetween(0, 3)
		items := make([]models.TransactionItems, numberOfItems)
		for j := 0; j < numberOfItems; j++ {
			items[j] = models.TransactionItems{
				Id:            uuid.New(),
				TransactionId: transactionId,
				BookId:        uuid.New(),
				Quantity:      faker.Int64Between(0, 3),
				Price:         faker.Float64(2, 10, 20),
				CreatedAt:     time.Now().Add(-time.Duration(i) * 24 * time.Hour),
				UpdatedAt:     time.Now(),
			}
		}

		results[i] = models.Transactions{
			Id:               transactionId,
			UserId:           uuid.New(),
			GrandTotal:       faker.Float64(2, 10, 20),
			Status:           string(constants.TransactionStatusConfirmed),
			InvoiceNumber:    faker.RandomStringWithLength(6),
			CreatedAt:        time.Now().Add(-time.Duration(i) * 24 * time.Hour),
			UpdatedAt:        time.Now(),
			TransactionItems: items,
		}
	}
	return results
}
