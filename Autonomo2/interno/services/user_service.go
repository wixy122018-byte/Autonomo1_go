
package repositories

import (
	"errors"
	"go-code-gestion/interno/modelos"
)

type UserRepository interface {
	Create(user *modelos.User) error
	GetByEmail(email string) (*modelos.User, error)
	GetByID(id uint) (*modelos.User, error)
	GetAll() ([]modelos.User, error)
}

type UserRepositoryImpl struct {
	users []modelos.User
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		users: []modelos.User{},
	}
}

func (r *UserRepositoryImpl) Create(user *modelos.User) error {

	for _, u := range r.users {
		if u.Email == user.Email {
			return errors.New("el correo ya existe")
		}
	}

	r.users = append(r.users, *user)

	return nil
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*modelos.User, error) {

	for _, user := range r.users {

		if user.Email == email {
			return &user, nil
		}
	}

	return nil, errors.New("usuario no encontrado")
}

func (r *UserRepositoryImpl) GetByID(id uint) (*modelos.User, error) {

	for _, user := range r.users {

		if user.ID == id {
			return &user, nil
		}
	}

	return nil, errors.New("usuario no encontrado")
}

func (r *UserRepositoryImpl) GetAll() ([]modelos.User, error) {
	return r.users, nil
}
