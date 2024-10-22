package transactions

import "gotu-bookstore/cmd/gotu-bookstore/contracts/dto"

type GetTransactionDetailsRequest struct {
	TransactionId string `uri:"transaction_id"`
}

type GetTransactionDetailsResponse struct {
	dto.TransactionDTO
}
