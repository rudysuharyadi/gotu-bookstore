package dto

/*
{
    "last_id": "uuid",
    "limit": 100,
    "page": 2,
    "size": 1000
}
*/
type PaginationMetadataDTO struct {
	Size int `json:"size"`
	Page int `json:"page"`
	Limit int `json:"limit"`
	LastId string `json:"last_id"`
}
