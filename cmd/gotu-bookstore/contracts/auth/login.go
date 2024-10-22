package auth

/*
	{
	    "email": "rudy.suharyadi@gmail.com",
	    "password": "password",
		"device_token": "device_token"
	}
*/
type LoginRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DeviceToken string `json:"device_token"`
}

/*
	{
	    "status": "success"
	}
*/
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
