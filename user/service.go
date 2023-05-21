package user

import (
	"golang.org/x/crypto/bcrypt"
)

//handleer;: input dari user akan dimapping ke struct input
// service: struct Input akan dimapping ke user Struct
//service: dependencies dari handler
//repoository :userstruct akan di save ke db
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
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
		return newUser, err
	}
	return newUser, nil
}
