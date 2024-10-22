package auth

import (
	"regexp"

	"gotu-bookstore/cmd/gotu-bookstore/constants"
	contracts "gotu-bookstore/cmd/gotu-bookstore/contracts/auth"
	"gotu-bookstore/pkg/resfmt/base_error"
)

type UsersValidator struct {
}

func NewUsersValidator() UsersValidator {
	return UsersValidator{}
}

func (s UsersValidator) Validate(request contracts.RegisterRequest) []error {
	responses := make([]error, 0)
	err := s.ValidatePassword(request.Password, request.ConfirmPassword)
	if err != nil {
		responses = append(responses, err)
	}

	err = s.ValidateEmail(request.Email)
	if err != nil {
		responses = append(responses, err)
	}

	err = s.ValidateName(request.Name)
	if err != nil {
		responses = append(responses, err)
	}

	return responses
}

func (s UsersValidator) ValidateEmail(email string) error {
	regex := regexp.MustCompile(constants.EmailRegex)
	if !regex.MatchString(email) {
		return base_error.NewSubError("email", "Email is not valid")
	}

	return nil
}

func (s UsersValidator) ValidateName(name string) error {
	if len(name) == 0 {
		return base_error.NewSubError("name", "Name cannot be empty")
	}

	return nil
}

func (s UsersValidator) ValidatePassword(password, confirmPassword string) error {
	if len(password) == 0 {
		return base_error.NewSubError("password", "Password cannot be empty")
	}

	if password != confirmPassword {
		return base_error.NewSubError("confirm_password", "Confirm password does not match")
	}

	return nil
}
