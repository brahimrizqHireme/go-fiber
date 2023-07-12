package inputs

import "errors"

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	isActive bool   `json:"isActive"`
}

func (i *CreateUserInput) Validate() error {
	if i.Name == "" {
		return errors.New("name is required")
	}
	if i.Email == "" {
		return errors.New("email is required")
	}
	if i.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
