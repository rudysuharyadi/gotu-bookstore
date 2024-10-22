package shopping_cart

/*

*/
type PostShoppingCartClearRequest struct {
}

/*
{
    "status": "success"
}
*/
type PostShoppingCartClearResponse struct {
	Status string `json:"status"`
}
