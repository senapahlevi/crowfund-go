package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//handleer;: input dari user akan dimapping ke struct input
// service: struct Input akan dimapping ke user Struct
//service: dependencies dari handler
//repoository :userstruct akan di save ke db
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailableInput(input CheckEmailInput) (bool, error)
}

type service struct {
	repository Repository //return user dan error
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}
	user.PasswordHash = string(PasswordHash) //convert ke string karena massih []byte\
	user.Role = "user"
	newUser, err := s.repository.Save(user)

	if err != nil {
		// response := helper.APIResponse("Account failed register", http.StatusBadRequest, "failed", nil)
		return newUser, err
	}
	return newUser, nil
}
func (s *service) Login(input LoginInput) (User, error) {
	user := User{}

	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("email user not found ")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) //dibandingkan db dan inputan user password sekarang di form

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailableInput(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return false, nil //nilai defaul

}
