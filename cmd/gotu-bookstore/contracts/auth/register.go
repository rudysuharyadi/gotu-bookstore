package auth

/*
	{
	    "confirm_password": "",
	    "email": "rudy.suharyadi@gmail.com",
	    "name": "Rudy Suharyadi",
	    "password": ""
	}
*/
type RegisterRequest struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Name            string `json:"name"`
	Email           string `json:"email"`
}

/*
 */
type RegisterResponse struct {
}
