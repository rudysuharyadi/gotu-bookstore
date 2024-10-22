package shopping_cart_test

import (
	"gotu-bookstore/cmd/gotu-bookstore/models"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
)

func GenerateFakerShoppingCart(count int) []models.ShoppingCarts {
	faker := faker.New()
	results := make([]models.ShoppingCarts, count)
	for i := 0; i < count; i++ {
		bookId := uuid.New()
		results[i] = models.ShoppingCarts{
			Id:        uuid.New(),
			UserId:    uuid.New(),
			BookId:    bookId,
			Quantity:  faker.Int64Between(1, 3),
			CreatedAt: time.Now().Add(-time.Duration(i) * 24 * time.Hour),
			UpdatedAt: time.Now(),
			Book: &models.Books{
				Id:          bookId,
				Author:      faker.Person().Name(),
				Title:       faker.Lorem().Sentence(3),
				Description: faker.Lorem().Paragraph(2),
				Category:    faker.Lorem().Word(),
				Publisher:   faker.Company().Name(),
				Price:       faker.Float64(2, 1, 99),
				Isbn:        faker.Lorem().Word(),
				Language:    faker.Language().Language(),
				PublishDate: faker.Time().TimeBetween(time.Now().AddDate(-5, 0, 0), time.Now()),
				ImageUrl:    faker.Internet().URL(),
				Page:        int64(faker.IntBetween(100, 1000)),
				Rating:      faker.Float64(1, 1.0, 5.0),
				CreatedAt:   time.Now().Add(-time.Duration(i) * 24 * time.Hour),
				UpdatedAt:   time.Now(),
			},
		}
	}
	return results
}

func GenerateFakeBook(count int) []models.Books {
	faker := faker.New()
	books := make([]models.Books, count)

	for i := 0; i < count; i++ {
		books[i] = models.Books{
			Id:          uuid.New(),
			Author:      faker.Person().Name(),
			Title:       faker.Lorem().Sentence(3),
			Description: faker.Lorem().Paragraph(2),
			Category:    faker.Lorem().Word(),
			Publisher:   faker.Company().Name(),
			Price:       faker.Float64(2, 1, 99),
			Isbn:        faker.Lorem().Word(),
			Language:    faker.Language().Language(),
			PublishDate: faker.Time().TimeBetween(time.Now().AddDate(-5, 0, 0), time.Now()),
			ImageUrl:    faker.Internet().URL(),
			Page:        int64(faker.IntBetween(100, 1000)),
			Rating:      faker.Float64(1, 1.0, 5.0),
			CreatedAt:   time.Now().Add(-time.Duration(i) * 24 * time.Hour),
			UpdatedAt:   time.Now(),
		}
	}
	return books
}
