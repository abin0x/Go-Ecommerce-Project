package user

import "ecommerce/domain"

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	// List() ([]*User, error)
	// Delete(productID int) error
	// Update(user User) (*User, error)
}
